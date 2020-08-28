// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package exampleutils

import (
	"fmt"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
)

// PrintAccountsBalances prints sender & receiver's account balances
func PrintAccountsBalances(title string, sender, receiver *librakeys.Keys) {
	fmt.Printf("\n> %v\n", title)
	PrintAccountBalances("sender", sender)
	PrintAccountBalances("receiver", receiver)
}

// PrintAccountBalances prints given account balances
func PrintAccountBalances(name string, account *librakeys.Keys) {
RetryGetAccount:
	ret, err := Client.GetAccount(account.AccountAddress().Hex())
	if _, ok := err.(*libraclient.StaleResponseError); ok {
		// retry to hit another server if got stale response
		goto RetryGetAccount
	}
	fmt.Println(name)
	for _, b := range ret.Balances {
		fmt.Printf(" - %v: %v\n", b.Currency, b.Amount)
	}
}
