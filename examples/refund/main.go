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
	metadata, err := txnmetadata.DeserializeMetadata(event)
	if err != nil {
		panic(err)
	}
	var refundMetadata []byte
	switch v := metadata.(type) {
	case *libratypes.Metadata__GeneralMetadata:
		refundMetadata, err = txnmetadata.NewRefundMetadataFromEventMetadata(event.SequenceNumber, v)
		if err != nil {
			panic(err)
		}
	case *libratypes.Metadata__TravelRuleMetadata:
		// If original peer to peer transaction script contains travel rule metadata,
		// refund should be same process.
		// It requires communication through off-chain API first and then create peer to
		// peer transaction script with travel rule metadata and recipient signature.
		// Please see https://github.com/libra/libra-client-sdk-go/blob/master/examples/p2p-transfers/main.go
		// for custodial account to custodial account over threshold example.
		//
		// Here as we expect GeneralMetadata, so we panic for simplicity.
		panic("unexpected event metadata")
	default:
		// Nil or other type case, no refund metadata required.
		//
		// Here as we expect GeneralMetadata, so we panic for simplicity.
		panic("unexpected event metadata")
	}

	exampleutils.SubmitAndWait(
		"refund transaction",
		receiver,
		stdlib.EncodePeerToPeerWithMetadataScript(
			libratypes.Currency(event.Data.Amount.Currency),
			sender.AccountAddress(),
			event.Data.Amount.Amount,
			refundMetadata,
			nil, // no metadata signature for GeneralMetadata
		),
	)
	exampleutils.PrintAccountsBalances("after refund", sender, receiver)
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
