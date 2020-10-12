// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"gopkg.in/yaml.v3"
)

func main() {
	merchant := librakeys.MustGenKeys()
	address := merchant.AccountAddress()
	currency := "Coin1"
	amount := uint64(5000)
	account := libraid.NewAccount(
		libraid.TestnetPrefix, address, libratypes.EmptySubAddress)
	intent := libraid.Intent{
		Account: *account,
		Params: libraid.Params{
			Currency: currency,
			Amount:   &amount,
		},
	}
	encodedIntent, err := intent.Encode()
	if err != nil {
		panic(err)
	}

	fmt.Printf("==== encoded intent identifier ====\n%v\n", encodedIntent)

	decodedIntent, err := libraid.DecodeToIntent(libraid.TestnetPrefix, encodedIntent)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\n==== decoded intent ====")
	yaml, _ := yaml.Marshal(decodedIntent)
	fmt.Println(string(yaml))
}
