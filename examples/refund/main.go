package main

import (
	"fmt"

	"github.com/libra/libra-client-sdk-go/examples/exampleutils"
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/libra/libra-client-sdk-go/txnmetadata"
)

const currency = "LBR"

func main() {
	sender, senderUserSubAddress := createCustodialAccount()
	receiver := testnet.GenAccount()
	amount := uint64(1000)

	exampleutils.PrintAccountsBalances("before transafer", sender, receiver)
	txnVersion := exampleutils.SubmitAndWait(
		"p2p transfer",
		sender,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(currency),
			receiver.AccountAddress(),
			amount,
			txnmetadata.NewGeneralMetadataFromSubAddress(senderUserSubAddress),
			nil, // no metadata signature for GeneralMetadata
		),
	)

	exampleutils.PrintAccountsBalances("after transfer, before refund", sender, receiver)
	// refund start
RetryGetTransactions:
	// find transaction back with events info
	txns, err := exampleutils.Client.GetTransactions(txnVersion, 1, true)
	if err != nil {
		if _, ok := err.(*libraclient.StaleResponseError); ok {
			// retry to hit another server if got stale response
			goto RetryGetTransactions
		}
		panic(err)
	}
	if len(txns) != 1 {
		panic(fmt.Sprintf("found transactions %v", len(txns)))
	}
	event := txnmetadata.FindRefundReferenceEventFromTransaction(
		txns[0], receiver.AccountAddress())
	if event == nil {
		panic("could not find refund reference event from transaction")
	}
	metadata, err := txnmetadata.NewNonCustodyToCustodyRefundMetadataFromEvent(event)
	if err != nil {
		panic(err)
	}
	exampleutils.SubmitAndWait(
		"refund transaction",
		receiver,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(event.Data.Amount.Currency),
			sender.AccountAddress(),
			event.Data.Amount.Amount,
			metadata,
			nil, // no metadata signature for GeneralMetadata
		),
	)
	exampleutils.PrintAccountsBalances("after transfer", sender, receiver)
}

func createCustodialAccount() (*librakeys.Keys, libratypes.SubAddress) {
	parentVASP := testnet.GenAccount()
	childVASPAccount := librakeys.MustGenKeys()
	script := stdlib.EncodeCreateChildVaspAccountScript(
		testnet.LBR,
		childVASPAccount.AccountAddress(),
		childVASPAccount.AuthKey().Prefix(),
		false,
		uint64(100000),
	)
	exampleutils.SubmitAndWait("create custodial child vasp account",
		parentVASP, script)
	return childVASPAccount, libratypes.MustGenSubAddress()
}
