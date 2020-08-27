// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"time"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/librasigner"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"gopkg.in/yaml.v3"
)

func main() {
	parentVASP := testnet.GenAccount()
	parentVASPAddress := parentVASP.AccountAddress()
	account, err := testnet.Client.GetAccount(parentVASPAddress.Hex())
	if err != nil {
		panic(err)
	}
	print("Parent VASP account", account)

	childVASPAccount := librakeys.MustGenKeys()
	childVASPAddress := childVASPAccount.AccountAddress()
	childAuthKey := childVASPAccount.AuthKey()

	script := stdlib.EncodeCreateChildVaspAccountScript(
		testnet.LBR,
		childVASPAddress,
		childAuthKey.Prefix(),
		false,
		uint64(1000),
	)

	sequenceNum := uint64(0) // we just generated new parentVASP, hence it is 0
	expirationDuration := time.Second * 30
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := librasigner.Sign(
		parentVASP,
		parentVASPAddress,
		sequenceNum,
		script,
		10000, 0, "LBR",
		expiration,
		testnet.ChainID,
	)
	err = testnet.Client.SubmitTransaction(txn)
	if err != nil {
		panic(err)
	}

	transaction, err := testnet.Client.WaitForTransaction2(txn, 5*time.Second)
	if err != nil {
		panic(err)
	}
	print("Create child VASP account transaction", transaction)

	child, err := testnet.Client.GetAccount(childVASPAddress.Hex())
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
