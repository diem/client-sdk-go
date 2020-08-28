// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/ed25519"

	"github.com/libra/libra-client-sdk-go/examples/exampleutils"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/libra/libra-client-sdk-go/txnmetadata"
)

const currency = "LBR"

func main() {
	nonCustodialAccount := testnet.GenAccount()
	nonCustodialAccount2 := testnet.GenAccount()

	custodialAccountParentVasp, custodialAccountChildVasp, custodialAccountSubAddress := createCustodialAccount()
	amount := uint64(10000)

	exampleutils.SubmitAndWait(
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

	exampleutils.SubmitAndWait(
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

	exampleutils.SubmitAndWait(
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
	exampleutils.SubmitAndWait(
		"setup parent vasp compliance key, testnet defaults it to fake key.",
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

	exampleutils.SubmitAndWait(
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

func createCustodialAccount() (*librakeys.Keys, *librakeys.Keys, libratypes.SubAddress) {
	parentVASP := testnet.GenAccount()
	childVASPAccount := librakeys.MustGenKeys()
	exampleutils.SubmitAndWait(
		"create child vasp for custodial account",
		parentVASP,
		stdlib.EncodeCreateChildVaspAccountScript(
			testnet.LBR,
			childVASPAccount.AccountAddress(),
			childVASPAccount.AuthKey().Prefix(),
			false,
			uint64(100000),
		),
	)
	custodialAccountSubAddress := libratypes.MustGenSubAddress()
	return parentVASP, childVASPAccount, custodialAccountSubAddress
}
