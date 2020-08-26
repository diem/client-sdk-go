// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package txnmetadata

import (
	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/libra/libra-client-sdk-go/libraid"
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
func NewGeneralMetadataToSubAddress(toSubAddress libraid.SubAddress) []byte {
	to := []byte(toSubAddress)
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				ToSubaddress: &to,
			},
		},
	}
	return libratypes.ToLCS(&metadata)
}

// NewGeneralMetadataFromSubAddress creates metadata for creating peer to peer
// transaction script with FromSubaddress
// This is used for peer to peer transfer from custodial account to non-custodial account.
func NewGeneralMetadataFromSubAddress(fromSubAddress libraid.SubAddress) []byte {
	from := []byte(fromSubAddress)
	metadata := libratypes.Metadata__GeneralMetadata{
		Value: &libratypes.GeneralMetadata__GeneralMetadataVersion0{
			Value: libratypes.GeneralMetadataV0{
				FromSubaddress: &from,
			},
		},
	}
	return libratypes.ToLCS(&metadata)
}
