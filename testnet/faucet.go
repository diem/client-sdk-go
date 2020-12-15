// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
)

// GenAccount generate account with single keys
func GenAccount() *diemkeys.Keys {
	keys := diemkeys.MustGenKeys()
	MustMint(keys.AuthKey().Hex(), 1000000, "XUS")
	return keys
}

// GenMultiSigAccount generate account with multi sig keys
func GenMultiSigAccount() *diemkeys.Keys {
	keys := diemkeys.MustGenMultiSigKeys()
	MustMint(keys.AuthKey().Hex(), 2000000, "XUS")
	return keys
}

// MustMint mints coins with retry, and panics if all retries failed.
// This func also wait for next account seq.
func MustMint(authKey string, amount uint64, currencyCode string) {
	retry := 5
	var err error
	var txns []diemtypes.SignedTransaction
	for i := 0; i < retry; i++ {
		if txns, err = Mint(authKey, amount, currencyCode); err == nil {
			if err = waitForTransactionsExecuted(txns); err == nil {
				return
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	panic(fmt.Sprintf("mint coins failed with retry: %s", err))
}

// Mint mints coints once without retry
func Mint(authKey string, amount uint64, currencyCode string) ([]diemtypes.SignedTransaction, error) {
	url := fmt.Sprintf("%v?amount=%d&auth_key=%s&currency_code=%s&return_txns=true", FaucetURL, amount, authKey, currencyCode)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Non 200 response: %s", string(body))
	}

	return deserializeMintTransactions(body)
}

func waitForTransactionsExecuted(txns []diemtypes.SignedTransaction) error {
	for i := range txns {
		_, err := Client.WaitForTransaction2(&txns[i], time.Second*30)
		if err != nil {
			return err
		}
	}
	return nil
}

func deserializeMintTransactions(body []byte) ([]diemtypes.SignedTransaction, error) {
	bytes, err := hex.DecodeString(string(body))
	if err != nil {
		return nil, fmt.Errorf("decode mint transactions hex string failed: %v", err)
	}
	deserializer := bcs.NewDeserializer(bytes)
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, fmt.Errorf("deserialize mint transactions length failed: %v", err)
	}
	ret := make([]diemtypes.SignedTransaction, length)
	for i := range ret {
		val, err := diemtypes.DeserializeSignedTransaction(deserializer)
		if err != nil {
			return nil, fmt.Errorf("deserialize %v mint transaction failed: %v", i, err)
		}
		ret[i] = val
	}

	return ret, nil
}
