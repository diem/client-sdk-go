// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package exampleutils

import (
	"fmt"

	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemkeys"
)

// PrintAccountsBalances prints sender & receiver's account balances
func PrintAccountsBalances(title string, sender, receiver *diemkeys.Keys) {
	fmt.Printf("\n> %v\n", title)
	PrintAccountBalances("sender", sender)
	PrintAccountBalances("receiver", receiver)
}

// PrintAccountBalances prints given account balances
func PrintAccountBalances(name string, account *diemkeys.Keys) {
RetryGetAccount:
	ret, err := Client.GetAccount(account.AccountAddress())
	if _, ok := err.(*diemclient.StaleResponseError); ok {
		// retry to hit another server if got stale response
		goto RetryGetAccount
	}
	fmt.Println(name)
	for _, b := range ret.Balances {
		fmt.Printf(" - %v: %v\n", b.Currency, b.Amount)
	}
}
