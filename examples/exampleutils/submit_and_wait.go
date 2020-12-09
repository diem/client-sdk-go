// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package exampleutils

import (
	"fmt"
	"time"

	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemsigner"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/diem/client-sdk-go/testnet"
)

// Client should be singleton instance for your application.
// It is initialized with connection pool, see it's document
// for how to configure.
var Client = diemclient.New(testnet.ChainID, testnet.URL)

// SubmitAndWait creates transaction for given script, then submit and wait for
// the transaction executed.
// Title is passed in for output with transaction version.
// To keep logic simple:
//   - this function simply panic when got unexpected error.
//   - assume sender account has "XUS" currency and use it as gas currency
//   - always use 0 gasUnitPrice
// This function returns back executed transaction version.
func SubmitAndWait(title string, sender *diemkeys.Keys, script diemtypes.Script) uint64 {
	fmt.Println(title)
	address := sender.AccountAddress()
Retry:
	account, err := Client.GetAccount(address)
	if err != nil {
		if _, ok := err.(*diemclient.StaleResponseError); ok {
			// retry to hit another server if got stale response
			goto Retry
		}
		panic(err)
	}
	sequenceNum := account.SequenceNumber
	// it is recommended to set short expiration time for peer to peer transaction,
	// as Diem blockchain transaction execution is fast.
	expirationDuration := 30 * time.Second
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := diemsigner.Sign(
		sender,
		address,
		sequenceNum,
		script,
		1_000_000, 0, "XUS",
		expiration,
		testnet.ChainID,
	)
	err = Client.SubmitTransaction(txn)
	if err != nil {
		if _, ok := err.(*diemclient.StaleResponseError); !ok {
			panic(err)
		} else {
			// ignore *diemclient.StaleResponseError as we know
			// submit probably succeed even hit a stale server
		}
	}
	transaction, err := Client.WaitForTransaction2(txn, expirationDuration)
	if err != nil {
		// WaitForTransaction retried for *diemclient.StaleResponseError
		// already, hence here we panic if got error (including timeout error)
		panic(err)
	}
	fmt.Printf("=> version: %v, status: %v\n",
		transaction.Version, transaction.VmStatus.Type)
	return transaction.Version
}
