// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/librasigner"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/libra/libra-client-sdk-go/txnmetadata"
	"gopkg.in/yaml.v3"
)

const currency = "LBR"

func main() {
	nonCustodialAccount := createNonCustodialAccount()
	nonCustodialAccount2 := createNonCustodialAccount()

	custodialAccountParentVasp, custodialAccountChildVasp, custodialAccountSubAddress := createCustodialAccount()
	amount := uint64(10000)

	newTransactionSubmitAndWait(
		"non custodial account to non custodial account transaction",
		nonCustodialAccount,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(currency),
			nonCustodialAccount2.AccountAddress(),
			amount,
			nil,
			nil,
		),
	)

	newTransactionSubmitAndWait(
		"non custodial account to custodial account transaction",
		nonCustodialAccount,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(currency),
			custodialAccountChildVasp.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataToSubAddress(custodialAccountSubAddress),
			nil, // no metadata signature for GeneralMetadata
		),
	)

	newTransactionSubmitAndWait(
		"custodial account to non custodial account transaction",
		custodialAccountChildVasp,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(currency),
			nonCustodialAccount.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataFromSubAddress(custodialAccountSubAddress),
			nil, // no metadata signature for GeneralMetadata
		),
	)

	// custodial account to custodial account transaction

	// setup receiver compliance public & private keys
	compliancePublicKey, compliancePrivateKey, _ := ed25519.GenerateKey(nil)
	newTransactionSubmitAndWait(
		"testnet created parent vasp has a fake compliance key, need rotate first",
		custodialAccountParentVasp,
		stdlib.EncodeRotateDualAttestationInfoScript(
			[]byte("http://helloworld.com"),
			[]byte(compliancePublicKey),
		),
	)

	// setup sender account
	_, senderCustodialAccountChildVasp, _ := createCustodialAccount()

	// sender & receiver communicate by off chain APIs
	offChainReferenceId := "32323abc"

	// metadata and signature message
	metadata, sigMsg := txnmetadata.NewTravelRuleMetadata(
		offChainReferenceId,
		senderCustodialAccountChildVasp.AccountAddress(),
		amount,
	)

	// receiver_signature is passed to the sender via the off-chain APIs as per
	// https://github.com/libra/lip/blob/master/lips/lip-1.mdx#recipient-signature
	recipientSignature := ed25519.Sign(compliancePrivateKey, sigMsg)

	newTransactionSubmitAndWait(
		"custodial account to custodial account transaction",
		senderCustodialAccountChildVasp,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(currency),
			custodialAccountChildVasp.AccountAddress(), //receiverAccountAddress,
			amount,
			metadata,
			recipientSignature,
		),
	)
}

func newTransactionSubmitAndWait(title string, sender *librakeys.Keys, script libratypes.Script) {
	address := sender.AccountAddress()
	account, err := testnet.Client.GetAccount(address.Hex())
	if err != nil {
		panic(err)
	}
	sequenceNum := account.SequenceNumber
	expirationDuration := 30 * time.Second
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := librasigner.Sign(
		sender,
		address,
		sequenceNum,
		script,
		100000, 0, currency,
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
	fmt.Printf("\n====== %v ======\n", title)
	yaml, _ := yaml.Marshal(transaction)
	fmt.Println(string(yaml))
}

func createNonCustodialAccount() *librakeys.Keys {
	return testnet.GenAccount()
}

func createCustodialAccount() (*librakeys.Keys, *librakeys.Keys, libraid.SubAddress) {
	parentVASP := createNonCustodialAccount()
	childVASPAccount := librakeys.MustGenKeys()
	script := stdlib.EncodeCreateChildVaspAccountScript(
		testnet.LBR,
		childVASPAccount.AccountAddress(),
		childVASPAccount.AuthKey().Prefix(),
		false,
		uint64(100000),
	)
	sequenceNum := uint64(0) // we just generated new parentVASP, hence it is 0
	expirationDuration := time.Second * 30
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := librasigner.Sign(
		parentVASP,
		parentVASP.AccountAddress(),
		sequenceNum,
		script,
		10000, 0, currency,
		expiration,
		testnet.ChainID,
	)
	err := testnet.Client.SubmitTransaction(txn)
	if err != nil {
		panic(err)
	}

	_, err = testnet.Client.WaitForTransaction2(txn, 5*time.Second)
	if err != nil {
		panic(err)
	}
	custodialAccountSubAddress := libraid.MustGenSubAddress()
	return parentVASP, childVASPAccount, custodialAccountSubAddress
}
