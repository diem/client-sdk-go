// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package exampleutils

import (
	"fmt"
	"time"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/librasigner"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/testnet"
)

// Client should be singleton instance for your application.
// It is initialized with connection pool, see it's document
// for how to configure.
var Client = libraclient.New(testnet.ChainID, testnet.URL)

// SubmitAndWait creates transaction for given script, then submit and wait for
// the transaction executed.
// Title is passed in for output with transaction version.
// To keep logic simple:
//   - this function simply panic when got unexpected error.
//   - assume sender account has "LBR" currency and use it as gas currency
//   - always use 0 gasUnitPrice
// This function returns back executed transaction version.
func SubmitAndWait(title string, sender *librakeys.Keys, script libratypes.Script) uint64 {
	fmt.Println(title)
	address := sender.AccountAddress()
Retry:
	account, err := Client.GetAccount(address)
	if err != nil {
		if _, ok := err.(*libraclient.StaleResponseError); ok {
			// retry to hit another server if got stale response
			goto Retry
		}
		panic(err)
	}
	sequenceNum := account.SequenceNumber
	// it is recommended to set short expiration time for peer to peer transaction,
	// as Libra blockchain transaction execution is fast.
	expirationDuration := 30 * time.Second
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := librasigner.Sign(
		sender,
		address,
		sequenceNum,
		script,
		1_000_000, 0, "LBR",
		expiration,
		testnet.ChainID,
	)
	err = Client.SubmitTransaction(txn)
	if err != nil {
		if _, ok := err.(*libraclient.StaleResponseError); !ok {
			panic(err)
		} else {
			// ignore *libraclient.StaleResponseError as we know
			// submit probably succeed even hit a stale server
		}
	}
	transaction, err := Client.WaitForTransaction2(txn, expirationDuration)
	if err != nil {
		// WaitForTransaction retried for *libraclient.StaleResponseError
		// already, hence here we panic if got error (including timeout error)
		panic(err)
	}
	fmt.Printf("=> version: %v, status: %v\n",
		transaction.Version, transaction.VmStatus.Type)
	return transaction.Version
}
