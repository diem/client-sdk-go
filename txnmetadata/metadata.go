// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package txnmetadata

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/libratypes"
)

// NewTravelRuleMetadata creates metadata and signature message for given
// offChainReferenceID.
// This is used for peer to peer transfer between 2 custodial accounts.
func NewTravelRuleMetadata(
	offChainReferenceID string,
	senderAccountAddress libratypes.AccountAddress,
	amount uint64,
) ([]byte, []byte) {
	metadata := libratypes.Metadata__TravelRuleMetadata{
		Value: &libratypes.TravelRuleMetadata__TravelRuleMetadataVersion0{
			Value: libratypes.TravelRuleMetadataV0{
				OffChainReferenceId: &offChainReferenceID,
			},
		},
	}

	// receiver_lcs_data = lcs(metadata, sender_address, amount, "@@$$LIBRA_ATTEST$$@@" /*ASCII-encoded string*/);
	s := new(lcs.Serializer)
	metadata.Serialize(s)
	senderAccountAddress.Serialize(s)
	s.SerializeU64(amount)
	sigMsg := append(s.GetBytes(), []byte("@@$$LIBRA_ATTEST$$@@")...)

	return libratypes.ToLCS(&metadata), sigMsg
}

// NewGeneralMetadataToSubAddress creates metadata for creating peer to peer
// transaction script with ToSubaddress
// This is used for peer to peer transfer from non-custodial account to custodial account.
func NewGeneralMetadataToSubAddress(toSubAddress libratypes.SubAddress) []byte {
	to := toSubAddress[:]
	return newGeneralMetadata(nil, &to)
}

// NewGeneralMetadataFromSubAddress creates metadata for creating peer to peer
// transaction script with FromSubaddress
// This is used for peer to peer transfer from custodial account to non-custodial account.
func NewGeneralMetadataFromSubAddress(fromSubAddress libratypes.SubAddress) []byte {
	from := fromSubAddress[:]
	return newGeneralMetadata(&from, nil)
}

// NewGeneralMetadataWithFromToSubaddresses creates metadata for creating peer to peer
// transaction script with fromSubaddress and toSubaddress.
// Use this function to create metadata with from and to subaddresses for peer to peer transfer
// from custodial account to custodial account under travel rule threshold.
func NewGeneralMetadataWithFromToSubaddresses(fromSubAddress libratypes.SubAddress, toSubAddress libratypes.SubAddress) []byte {
	from := fromSubAddress[:]
	to := toSubAddress[:]
	return newGeneralMetadata(&from, &to)
}

// newGeneralMetadata is internal methods for constructing with *[]byte as from and to subaddress type
func newGeneralMetadata(fromSubAddress *[]byte, toSubAddress *[]byte) []byte {
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				FromSubaddress: fromSubAddress,
				ToSubaddress:   toSubAddress,
			},
		},
	}
	return libratypes.ToLCS(&metadata)
}

// FindRefundReferenceEventFromTransaction looks for receivedpayment type event in the
// given transaction and event receiver is given receiver account address.
func FindRefundReferenceEventFromTransaction(txn *libraclient.Transaction, receiver libratypes.AccountAddress) *libraclient.Event {
	if txn == nil {
		return nil
	}
	address := receiver.Hex()
	for i, event := range txn.Events {
		if event.Data.Type == "receivedpayment" &&
			event.Data.Receiver == address {
			return &txn.Events[i]
		}
	}
	return nil
}

// NewNonCustodyToCustodyRefundMetadataFromEvent creates GeneralMetadata for refund
// given event.
// The given event must be custody to non-custody payment event with
// `libratypes.Metadata__GeneralMetadata` includes `FromSubaddress` as metadata.
func NewNonCustodyToCustodyRefundMetadataFromEvent(event *libraclient.Event) ([]byte, error) {
	if event == nil {
		return nil, errors.New("must provide refund reference event")
	}
	bytes, err := hex.DecodeString(event.Data.Metadata)
	if err != nil {
		return nil, fmt.Errorf("decode event metadata failed: %v", err.Error())
	}
	metadata, err := libratypes.DeserializeMetadata(lcs.NewDeserializer(bytes))
	if err != nil {
		return nil, err
	}
	gm, ok := metadata.(*libratypes.Metadata__GeneralMetadata)
	if !ok {
		return nil, fmt.Errorf("unexpected metadata: %v", metadata)
	}
	gmv0, ok := gm.Value.(*libratypes.GeneralMetadata__GeneralMetadataVersion0)
	if !ok {
		return nil, fmt.Errorf("unexpected metadata: %v", metadata)
	}

	if gmv0.Value.FromSubaddress == nil {
		return nil, errors.New("event metadata FromSubaddress is not set")
	}
	subaddress, err := libratypes.MakeSubAddressFromBytes(*gmv0.Value.FromSubaddress)
	if err != nil {
		return nil, err
	}
	return NewNonCustodyToCustodyRefundMetadata(subaddress, event.SequenceNumber), nil
}

// NewNonCustodyToCustodyRefundMetadata creates metadata
// for creating refund peer to peer transaction script.
// Only required for refund a transaction that is transfer from custodial account to
// non-custodial account.
func NewNonCustodyToCustodyRefundMetadata(
	toSubAddress libratypes.SubAddress,
	referencedEventSequenceNumber uint64,
) []byte {
	to := toSubAddress[:]
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				ToSubaddress:    &to,
				ReferencedEvent: &referencedEventSequenceNumber,
			},
		},
	}
	return libratypes.ToLCS(&metadata)
}
