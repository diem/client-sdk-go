// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/ed25519"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/diem/client-sdk-go/examples/exampleutils"
	"github.com/diem/client-sdk-go/stdlib"
	"github.com/diem/client-sdk-go/testnet"
	"github.com/diem/client-sdk-go/txnmetadata"
)

const currency = "XUS"

func main() {
	nonCustodialAccount := testnet.GenAccount()
	nonCustodialAccount2 := testnet.GenAccount()

	custodialAccountParentVasp, custodialAccountChildVasp, custodialAccountSubAddress := createCustodialAccount()
	amount := uint64(10000)

	// Non custodial to non custodial has no requirement on metadata
	exampleutils.PrintAccountsBalances("before transfer", nonCustodialAccount, nonCustodialAccount2)
	exampleutils.SubmitAndWait(
		"non custodial account to non custodial account transaction",
		nonCustodialAccount,
		stdlib.EncodePeerToPeerWithMetadataScript(
			diemtypes.Currency(currency),
			nonCustodialAccount2.AccountAddress(),
			amount,
			nil,
			nil,
		),
	)
	exampleutils.PrintAccountsBalances("after transfer", nonCustodialAccount, nonCustodialAccount2)

	// Non custodial account to custodial account requires target custodial account subaddress,
	// hence we need construct a general metadata includes to_subaddress
	exampleutils.PrintAccountsBalances("before transfer", nonCustodialAccount, custodialAccountChildVasp)
	exampleutils.SubmitAndWait(
		"non custodial account to custodial account transaction",
		nonCustodialAccount,
		stdlib.EncodePeerToPeerWithMetadataScript(
			diemtypes.Currency(currency),
			custodialAccountChildVasp.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataToSubAddress(custodialAccountSubAddress),
			nil, // no metadata signature for GeneralMetadata
		),
	)
	exampleutils.PrintAccountsBalances("after transfer", nonCustodialAccount, custodialAccountChildVasp)

	// Custodial account to non-custodial account requires sender's custodial account subaddress,
	// hence we need construct a general metadata includes from_subaddress
	exampleutils.PrintAccountsBalances("before transfer", custodialAccountChildVasp, nonCustodialAccount)
	exampleutils.SubmitAndWait(
		"custodial account to non custodial account transaction",
		custodialAccountChildVasp,
		stdlib.EncodePeerToPeerWithMetadataScript(
			diemtypes.Currency(currency),
			nonCustodialAccount.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataFromSubAddress(custodialAccountSubAddress),
			nil, // no metadata signature for GeneralMetadata
		),
	)
	exampleutils.PrintAccountsBalances("after transfer", custodialAccountChildVasp, nonCustodialAccount)

	// Custodial account to custodial account transaction has 2 cases

	// setup sender custodial account
	_, senderCustodialAccountChildVasp, senderCustodialAccountSubAddress := createCustodialAccount()

	// Case 1: For transactions under the travel rule threshold, transaction metadata inclusive of both to_subaddress and from_subaddress should be composed.
	exampleutils.PrintAccountsBalances("before transfer",
		senderCustodialAccountChildVasp, custodialAccountChildVasp)
	exampleutils.SubmitAndWait(
		"custodial account to custodial account transaction under travel rule threshold",
		senderCustodialAccountChildVasp,
		stdlib.EncodePeerToPeerWithMetadataScript(
			diemtypes.Currency(currency),
			custodialAccountChildVasp.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataWithFromToSubAddresses(
				senderCustodialAccountSubAddress,
				custodialAccountSubAddress,
			),
			nil, // no metadata signature for GeneralMetadata
		),
	)
	exampleutils.PrintAccountsBalances("after transfer",
		senderCustodialAccountChildVasp, custodialAccountChildVasp)

	// Case 2: For transactions over the travel rule limit, custodial to custodial transactions must exchange travel rule compliance data off-chain

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

	// sender & receiver communicate by off chain APIs
	offChainReferenceId := "32323abc"

	// metadata and signature message
	metadata, sigMsg := txnmetadata.NewTravelRuleMetadata(
		offChainReferenceId,
		senderCustodialAccountChildVasp.AccountAddress(),
		amount,
	)

	// receiver_signature is passed to the sender via the off-chain APIs as per
	// https://github.com/diem/lip/blob/master/lips/lip-1.mdx#recipient-signature
	recipientSignature := ed25519.Sign(compliancePrivateKey, sigMsg)

	exampleutils.PrintAccountsBalances("before transfer",
		senderCustodialAccountChildVasp, custodialAccountChildVasp)
	exampleutils.SubmitAndWait(
		"custodial account to custodial account transaction",
		senderCustodialAccountChildVasp,
		stdlib.EncodePeerToPeerWithMetadataScript(
			diemtypes.Currency(currency),
			custodialAccountChildVasp.AccountAddress(), //receiverAccountAddress,
			amount,
			metadata,
			recipientSignature,
		),
	)
	exampleutils.PrintAccountsBalances("after transfer",
		senderCustodialAccountChildVasp, custodialAccountChildVasp)
}

func createCustodialAccount() (*diemkeys.Keys, *diemkeys.Keys, diemtypes.SubAddress) {
	parentVASP := testnet.GenAccount()
	childVASPAccount := diemkeys.MustGenKeys()
	exampleutils.SubmitAndWait(
		"create child vasp for custodial account",
		parentVASP,
		stdlib.EncodeCreateChildVaspAccountScript(
			testnet.XUS,
			childVASPAccount.AccountAddress(),
			childVASPAccount.AuthKey().Prefix(),
			false,
			uint64(100000),
		),
	)
	custodialAccountSubAddress := diemtypes.MustGenSubAddress()
	return parentVASP, childVASPAccount, custodialAccountSubAddress
}
