// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/examples/exampleutils"
	"github.com/diem/client-sdk-go/stdlib"
	"github.com/diem/client-sdk-go/testnet"
	"gopkg.in/yaml.v3"
)

func main() {
	parentVASP := testnet.GenAccount()
	parentVASPAddress := parentVASP.AccountAddress()
	account, err := exampleutils.Client.GetAccount(parentVASPAddress)
	if err != nil {
		panic(err)
	}
	print("Parent VASP account", account)

	childVASPAccount := diemkeys.MustGenKeys()
	childVASPAddress := childVASPAccount.AccountAddress()
	childAuthKey := childVASPAccount.AuthKey()

	exampleutils.SubmitAndWait(
		"create child vasp account transaction",
		parentVASP,
		stdlib.EncodeCreateChildVaspAccountScript(
			testnet.XUS,
			childVASPAddress,
			childAuthKey.Prefix(),
			false,
			uint64(1000),
		),
	)

	child, err := exampleutils.Client.GetAccount(childVASPAddress)
	if err != nil {
		panic(err)
	}
	print("Child VASP account", child)
}

func print(title string, obj interface{}) {
	fmt.Printf("====== %v ======\n", title)
	yaml, _ := yaml.Marshal(obj)
	fmt.Println(string(yaml))
}
