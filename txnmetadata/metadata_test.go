// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package txnmetadata_test

import (
	"encoding/hex"
	"testing"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/txnmetadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTravelRuleMetadata(t *testing.T) {
	address, _ := libratypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	metadata, sigMsg := txnmetadata.NewTravelRuleMetadata(
		"off chain reference id",
		address,
		1000)
	assert.Equal(t, "020001166f666620636861696e207265666572656e6365206964", hex.EncodeToString(metadata))
	assert.Equal(t, "020001166f666620636861696e207265666572656e6365206964f72589b71ff4f8d139674a3f7369c69be803000000000000404024244c494252415f41545445535424244040", hex.EncodeToString(sigMsg))
}

func TestNewGeneralMetadataToSubAddress(t *testing.T) {
	subAddress, _ := libratypes.MakeSubAddress("8f8b82153010a1bd")
	ret := txnmetadata.NewGeneralMetadataToSubAddress(subAddress)
	assert.Equal(t, "010001088f8b82153010a1bd0000", hex.EncodeToString(ret))
}

func TestNewGeneralMetadataFromSubAddress(t *testing.T) {
	subAddress, _ := libratypes.MakeSubAddress("8f8b82153010a1bd")
	ret := txnmetadata.NewGeneralMetadataFromSubAddress(subAddress)
	assert.Equal(t, "01000001088f8b82153010a1bd00", hex.EncodeToString(ret))
}

func TestNewGeneralMetadataWithFromToSubaddresses(t *testing.T) {
	subAddress1, _ := libratypes.MakeSubAddress("8f8b82153010a1bd")
	subAddress2, _ := libratypes.MakeSubAddress("111111153010a111")
	ret := txnmetadata.NewGeneralMetadataWithFromToSubaddresses(subAddress1, subAddress2)
	assert.Equal(t, "01000108111111153010a11101088f8b82153010a1bd00", hex.EncodeToString(ret))
}

func TestNewNonCustodyToCustodyRefundMetadataFromEvent(t *testing.T) {
	subAddress, _ := libratypes.MakeSubAddress("8f8b82153010a1bd")
	bytes := txnmetadata.NewGeneralMetadataFromSubAddress(subAddress)
	ret, err := txnmetadata.NewNonCustodyToCustodyRefundMetadataFromEvent(&libraclient.Event{
		Data: libraclient.EventData{
			Metadata: hex.EncodeToString(bytes),
		},
		SequenceNumber: 123,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, ret)
}
