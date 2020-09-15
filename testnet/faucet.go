// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
)

// GenAccount generate account with single keys
func GenAccount() *librakeys.Keys {
	keys := librakeys.MustGenKeys()
	MustMint(keys.AuthKey().Hex(), 1000000, "LBR")
	return keys
}

// GenMultiSigAccount generate account with multi sig keys
func GenMultiSigAccount() *librakeys.Keys {
	keys := librakeys.MustGenMultiSigKeys()
	MustMint(keys.AuthKey().Hex(), 2000000, "LBR")
	return keys
}

// MustMint mints coins with retry, and panics if all retries failed.
// This func also wait for next account seq.
func MustMint(authKey string, amount uint64, currencyCode string) {
	retry := 5
	var err error
	var accountSeq int
	for i := 0; i < retry; i++ {
		accountSeq, err = Mint(authKey, amount, currencyCode)
		if err != nil {
			time.Sleep(1100 * time.Millisecond)
			continue
		}
		waitAccountSequence(accountSeq)
		return
	}
	panic(fmt.Sprintf("mint coins failed with retry: %s", err))
}

// Mint mints coints once without retry
func Mint(authKey string, amount uint64, currencyCode string) (int, error) {
	url := fmt.Sprintf("%v?amount=%d&auth_key=%s&currency_code=%s", FaucetURL, amount, authKey, currencyCode)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Non 200 response: %s", string(body))
	}
	return strconv.Atoi(string(body))
}

func waitAccountSequence(seq int) {
	for i := 0; i < 100; i++ {
		account, err := Client.GetAccount(DDAccountAddress)
		if _, ok := err.(*libraclient.StaleResponseError); ok {
			continue
		}
		if err != nil {
			panic(fmt.Errorf("get account failed: %v", err))
		}
		if account.SequenceNumber >= uint64(seq) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
	panic("waiting for mint timeout")
}
