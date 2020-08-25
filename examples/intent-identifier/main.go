// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"gopkg.in/yaml.v3"
)

func main() {
	merchant := librakeys.MustGenKeys()
	currency := "LBR"
	amount := uint64(5000)
	account := libraid.NewAccount(libraid.TestnetPrefix, merchant.AccountAddress, nil)
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
