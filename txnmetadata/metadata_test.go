// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package txnmetadata_test

import (
	"encoding/hex"
	"testing"

	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemclient/diemclienttest"
	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/diem/client-sdk-go/txnmetadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTravelRuleMetadata(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	metadata, sigMsg := txnmetadata.NewTravelRuleMetadata(
		"off chain reference id",
		address,
		1000)
	assert.Equal(t, "020001166f666620636861696e207265666572656e6365206964", hex.EncodeToString(metadata))
	assert.Equal(t, "020001166f666620636861696e207265666572656e6365206964f72589b71ff4f8d139674a3f7369c69be803000000000000404024244449454d5f41545445535424244040", hex.EncodeToString(sigMsg))
}

func TestNewGeneralMetadataToSubAddress(t *testing.T) {
	subAddress, _ := diemtypes.MakeSubAddress("8f8b82153010a1bd")
	ret := txnmetadata.NewGeneralMetadataToSubAddress(subAddress)
	assert.Equal(t, "010001088f8b82153010a1bd0000", hex.EncodeToString(ret))
}

func TestNewGeneralMetadataFromSubAddress(t *testing.T) {
	subAddress, _ := diemtypes.MakeSubAddress("8f8b82153010a1bd")
	ret := txnmetadata.NewGeneralMetadataFromSubAddress(subAddress)
	assert.Equal(t, "01000001088f8b82153010a1bd00", hex.EncodeToString(ret))
}

func TestNewGeneralMetadataWithFromToSubaddresses(t *testing.T) {
	subAddress1, _ := diemtypes.MakeSubAddress("8f8b82153010a1bd")
	subAddress2, _ := diemtypes.MakeSubAddress("111111153010a111")
	ret := txnmetadata.NewGeneralMetadataWithFromToSubAddresses(subAddress1, subAddress2)
	assert.Equal(t, "01000108111111153010a11101088f8b82153010a1bd00", hex.EncodeToString(ret))
}

func TestFindRefundReferenceEventFromTransaction(t *testing.T) {
	receiver := diemkeys.MustGenKeys()

	t.Run("return nil for given transaction is nil", func(t *testing.T) {
		ret := txnmetadata.FindRefundReferenceEventFromTransaction(nil, receiver.AccountAddress())
		assert.Nil(t, ret)
	})
	t.Run("return event that is receivedpayment type and receiver account address", func(t *testing.T) {
		txn := diemclienttest.TransactionBuilder{}.Events(
			diemclienttest.EventBuilder{}.
				Type("unknowntype").
				Receiver(receiver.AccountAddress().Hex()),
			diemclienttest.EventBuilder{}.
				Type("receivedpayment").
				Receiver("unknwon address"),
			diemclienttest.EventBuilder{}.
				Type("receivedpayment").
				Receiver(receiver.AccountAddress().Hex()),
		).Build()
		ret := txnmetadata.FindRefundReferenceEventFromTransaction(txn, receiver.AccountAddress())
		require.NotNil(t, ret)
		assert.Equal(t, "receivedpayment", ret.Data.Type)
		assert.Equal(t, receiver.AccountAddress().Hex(), ret.Data.Receiver)
	})
	t.Run("return nil event if not found", func(t *testing.T) {
		txn := diemclienttest.TransactionBuilder{}.Events(
			diemclienttest.EventBuilder{}.
				Type("unknowntype").
				Receiver(receiver.AccountAddress().Hex()),
			diemclienttest.EventBuilder{}.
				Type("receivedpayment").
				Receiver("unknwon address"),
		).Build()
		ret := txnmetadata.FindRefundReferenceEventFromTransaction(txn, receiver.AccountAddress())
		require.Nil(t, ret)
	})
}

func TestNewRefundMetadataFromEvent(t *testing.T) {
	referencedEventSeqNum := uint64(123)

	cases := []struct {
		name             string
		event            *diemclienttest.EventBuilder
		expectedErrorMsg string
		expected         *diemtypes.Metadata__GeneralMetadata
	}{
		{
			name: "return event metadata with referenced event: include both from & to subaddress",
			event: diemclienttest.EventBuilder{}.
				Metadata(
					hex.EncodeToString(txnmetadata.NewGeneralMetadataWithFromToSubAddresses(
						diemtypes.SubAddress{1, 2, 3, 4, 5, 6, 7, 8},
						diemtypes.SubAddress{8, 7, 6, 5, 4, 3, 2, 1},
					))).
				SequenceNumber(referencedEventSeqNum),
			expected: &diemtypes.Metadata__GeneralMetadata{
				Value: &diemtypes.GeneralMetadata__GeneralMetadataVersion0{
					Value: diemtypes.GeneralMetadataV0{
						FromSubaddress:  &[]byte{8, 7, 6, 5, 4, 3, 2, 1},
						ToSubaddress:    &[]byte{1, 2, 3, 4, 5, 6, 7, 8},
						ReferencedEvent: &referencedEventSeqNum,
					},
				},
			},
		},
		{
			name: "return event metadata with referenced event: only has to subaddress",
			event: diemclienttest.EventBuilder{}.
				Metadata(
					hex.EncodeToString(txnmetadata.NewGeneralMetadataFromSubAddress(
						diemtypes.SubAddress{1, 2, 3, 4, 5, 6, 7, 8},
					))).
				SequenceNumber(referencedEventSeqNum),
			expected: &diemtypes.Metadata__GeneralMetadata{
				Value: &diemtypes.GeneralMetadata__GeneralMetadataVersion0{
					Value: diemtypes.GeneralMetadataV0{
						ToSubaddress:    &[]byte{1, 2, 3, 4, 5, 6, 7, 8},
						ReferencedEvent: &referencedEventSeqNum,
					},
				},
			},
		},
		{
			name: "return event metadata with referenced event: only has from subaddress",
			event: diemclienttest.EventBuilder{}.
				Metadata(
					hex.EncodeToString(txnmetadata.NewGeneralMetadataToSubAddress(
						diemtypes.SubAddress{1, 2, 3, 4, 5, 6, 7, 8},
					))).
				SequenceNumber(referencedEventSeqNum),
			expected: &diemtypes.Metadata__GeneralMetadata{
				Value: &diemtypes.GeneralMetadata__GeneralMetadataVersion0{
					Value: diemtypes.GeneralMetadataV0{
						FromSubaddress:  &[]byte{1, 2, 3, 4, 5, 6, 7, 8},
						ReferencedEvent: &referencedEventSeqNum,
					},
				},
			},
		},
		{
			name:             "event is nil",
			event:            nil,
			expectedErrorMsg: "must provide refund reference event",
		},
		{
			name:             "event metadata is not hex-encoded string",
			event:            diemclienttest.EventBuilder{}.Metadata("lj;lafda"),
			expectedErrorMsg: "decode event metadata failed: encoding/hex: invalid byte: U+006C 'l'",
		},
		{
			name:             "event metadata is not valid LCS bytes",
			event:            diemclienttest.EventBuilder{}.Metadata("1112233333"),
			expectedErrorMsg: "can't deserialize metadata: Unknown variant index for Metadata: 17",
		},
		{
			name: "return nil without error if event metadata is empty",
			event: diemclienttest.EventBuilder{}.
				Metadata("").
				SequenceNumber(referencedEventSeqNum),
			expected: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var event *diemclient.Event
			if tc.event != nil {
				event = tc.event.Build()
			}
			ret, err := newRefundMetadataFromEvent(event)
			if tc.expectedErrorMsg != "" {
				assert.EqualError(t, err, tc.expectedErrorMsg)
			} else if tc.expected != nil {
				require.NoError(t, err)
				ret, err := diemtypes.LcsDeserializeMetadata(ret)
				require.NoError(t, err)
				assert.EqualValues(t, tc.expected, ret)
			} else {
				assert.Nil(t, ret)
				assert.Nil(t, err)
			}
		})
	}
}

func newRefundMetadataFromEvent(event *diemclient.Event) ([]byte, error) {
	md, err := txnmetadata.DeserializeMetadata(event)
	if err != nil {
		return nil, err
	}
	if md == nil {
		return nil, nil
	}
	return txnmetadata.NewRefundMetadataFromEventMetadata(event.SequenceNumber,
		md.(*diemtypes.Metadata__GeneralMetadata))
}
