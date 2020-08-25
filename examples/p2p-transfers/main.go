// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
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
		newNonCustodialToNonCustodialTransactionScript(
			nonCustodialAccount2.AccountAddress,
			amount,
		),
	)

	newTransactionSubmitAndWait(
		"non custodial account to custodial account transaction",
		nonCustodialAccount,
		newNonCustodialToCustodialTransactionScript(
			custodialAccountChildVasp.AccountAddress,
			custodialAccountSubAddress,
			amount,
		),
	)

	newTransactionSubmitAndWait(
		"custodial account to non custodial account transaction",
		custodialAccountChildVasp,
		newCustodialToNonCustodialTransactionScript(
			custodialAccountSubAddress,
			nonCustodialAccount.AccountAddress,
			amount,
		),
	)

	compliancePublicKey, compliancePrivateKey, _ := ed25519.GenerateKey(nil)
	newTransactionSubmitAndWait(
		"testnet created parent vasp has a fake compliance key, need rotate first",
		custodialAccountParentVasp,
		stdlib.EncodeRotateDualAttestationInfoScript(
			[]byte("http://helloworld.com"),
			[]byte(compliancePublicKey),
		),
	)

	_, senderCustodialAccountChildVasp, _ := createCustodialAccount()
	newTransactionSubmitAndWait(
		"custodial account to custodial account transaction",
		senderCustodialAccountChildVasp,
		newCustodialToCustodialTransactionScript(
			senderCustodialAccountChildVasp.AccountAddress,
			custodialAccountChildVasp.AccountAddress,
			compliancePrivateKey,
			amount,
		),
	)
}

func newNonCustodialToNonCustodialTransactionScript(
	receiverAccountAddress libratypes.AccountAddress,
	amount uint64,
) libratypes.Script {
	return stdlib.EncodePeerToPeerWithMetadataScript(
		stdlib.CurrencyCode(currency),
		receiverAccountAddress,
		amount,
		nil,
		nil,
	)
}

func newCustodialToCustodialTransactionScript(
	senderAccountAddress libratypes.AccountAddress,
	receiverAccountAddress libratypes.AccountAddress,
	compliancePrivateKey ed25519.PrivateKey,
	amount uint64,
) libratypes.Script {
	// from off chain APIs
	offChainReferenceId := "32323abc"
	metadata := libratypes.Metadata__TravelRuleMetadata{
		Value: &libratypes.TravelRuleMetadata__TravelRuleMetadataVersion0{
			Value: libratypes.TravelRuleMetadataV0{
				OffChainReferenceId: &offChainReferenceId,
			},
		},
	}

	// receiver_signature is passed to the sender via the off-chain APIs as per
	// https://github.com/libra/lip/blob/master/lips/lip-1.mdx#recipient-signature
	// receiver_lcs_data = lcs(metadata, sender_address, amount, "@@$$LIBRA_ATTEST$$@@" /*ASCII-encoded string*/);

	s := new(lcs.Serializer)
	metadata.Serialize(s)
	senderAccountAddress.Serialize(s)
	s.SerializeU64(amount)
	msg := append(s.GetBytes(), []byte("@@$$LIBRA_ATTEST$$@@")...)
	recipientSignature := ed25519.Sign(compliancePrivateKey, msg)

	// sender constructs transaction script
	return stdlib.EncodePeerToPeerWithMetadataScript(
		stdlib.CurrencyCode(currency),
		receiverAccountAddress,
		amount,
		libratypes.ToLCS(&metadata),
		recipientSignature,
	)
}

func newNonCustodialToCustodialTransactionScript(
	toCustodialAddress libratypes.AccountAddress,
	toCustodialSubAddress []byte,
	amount uint64,
) libratypes.Script {
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				ToSubaddress: &toCustodialSubAddress,
			},
		},
	}
	return stdlib.EncodePeerToPeerWithMetadataScript(
		stdlib.CurrencyCode(currency),
		toCustodialAddress,
		amount,
		libratypes.ToLCS(&metadata),
		nil, // no metadata signature for GeneralMetadata
	)
}

func newCustodialToNonCustodialTransactionScript(
	fromCustodialSubAddress []byte,
	toNonCustodialAddress libratypes.AccountAddress,
	amount uint64,
) libratypes.Script {
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				FromSubaddress: &fromCustodialSubAddress,
			},
		},
	}
	return stdlib.EncodePeerToPeerWithMetadataScript(
		stdlib.CurrencyCode(currency),
		toNonCustodialAddress,
		amount,
		libratypes.ToLCS(&metadata),
		nil, // no metadata signature for GeneralMetadata
	)
}

func newTransactionSubmitAndWait(title string, sender *librakeys.Keys, script libratypes.Script) {
	account, err := testnet.Client.GetAccount(sender.AccountAddress.Hex())
	if err != nil {
		panic(err)
	}
	sequenceNum := account.SequenceNumber
	expirationDuration := 30 * time.Second
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := sender.Sign(
		sequenceNum,
		script,
		100000, 0, currency,
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
	fmt.Printf("\n====== %v ======\n", title)
	yaml, _ := yaml.Marshal(transaction)
	fmt.Println(string(yaml))
}

func createNonCustodialAccount() *librakeys.Keys {
	keys := librakeys.MustGenKeys()
	testnet.MustMint(keys.AuthKey.Hex(), 1000000, currency)
	return keys
}

func createCustodialAccount() (*librakeys.Keys, *librakeys.Keys, libraid.SubAddress) {
	parentVASP := createNonCustodialAccount()
	childVASPAccount := librakeys.MustGenKeys()
	script := stdlib.EncodeCreateChildVaspAccountScript(
		testnet.LBR,
		childVASPAccount.AccountAddress,
		childVASPAccount.AuthKey.Prefix(),
		false,
		uint64(100000),
	)
	sequenceNum := uint64(0) // we just generated new parentVASP, hence it is 0
	expirationDuration := time.Second * 30
	expiration := uint64(time.Now().Add(expirationDuration).Unix())
	txn := parentVASP.Sign(
		sequenceNum,
		script,
		10000, 0, currency,
		expiration,
		testnet.ChainID,
	)
	err := testnet.Client.Submit(txn.Hex())
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
