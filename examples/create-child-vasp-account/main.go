// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"time"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"gopkg.in/yaml.v3"
)

func main() {
	parentVASP := librakeys.MustGenKeys()
	testnet.MustMint(parentVASP.AuthKey.Hex(), 1000000, "LBR")

	account, err := testnet.Client.GetAccount(parentVASP.AccountAddress.Hex())
	if err != nil {
		panic(err)
	}
	print("Parent VASP account", account)

	childVASPAccount := librakeys.MustGenKeys()

	script := stdlib.EncodeCreateChildVaspAccountScript(
		testnet.LBR,
		childVASPAccount.AccountAddress,
		childVASPAccount.AuthKey.Prefix(),
		false,
		uint64(1000),
	)

	sequenceNum := uint64(0) // we just generated new parentVASP, hence it is 0
	expirationDuration := time.Second * 30
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := parentVASP.Sign(
		sequenceNum,
		script,
		10000, 0, "LBR",
		expiration,
		testnet.ChainID,
	)
	err = testnet.Client.Submit(txn.Hex())
	if err != nil {
		panic(err)
	}

	transaction, err := testnet.Client.WaitForTransaction2(txn, 5*time.Second)
	if err != nil {
		panic(err)
	}
	print("Create child VASP account transaction", transaction)

	child, err := testnet.Client.GetAccount(childVASPAccount.AccountAddress.Hex())
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
