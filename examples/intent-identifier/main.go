// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/diem/client-sdk-go/diemid"
	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
	"gopkg.in/yaml.v3"
)

func main() {
	merchant := diemkeys.MustGenKeys()
	address := merchant.AccountAddress()
	currency := "XUS"
	amount := uint64(5000)
	account := diemid.NewAccount(
		diemid.TestnetPrefix, address, diemtypes.EmptySubAddress)
	intent := diemid.Intent{
		Account: *account,
		Params: diemid.Params{
			Currency: currency,
			Amount:   &amount,
		},
	}
	encodedIntent, err := intent.Encode()
	if err != nil {
		panic(err)
	}

	fmt.Printf("==== encoded intent identifier ====\n%v\n", encodedIntent)

	decodedIntent, err := diemid.DecodeToIntent(diemid.TestnetPrefix, encodedIntent)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\n==== decoded intent ====")
	yaml, _ := yaml.Marshal(decodedIntent)
	fmt.Println(string(yaml))
}
