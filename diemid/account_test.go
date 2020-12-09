// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemid_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemid"
	"github.com/diem/client-sdk-go/diemid/bech32"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeAccountIdentifier(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	subAddress, _ := diemtypes.MakeSubAddress("cf64428bdeb62af2")

	ret, err := diemid.EncodeAccount(diemid.MainnetPrefix, address, subAddress)
	require.NoError(t, err)
	assert.Equal(t, "lbr1p7ujcndcl7nudzwt8fglhx6wxn08kgs5tm6mz4usw5p72t", string(ret))

	id, err := diemid.DecodeToAccount(diemid.MainnetPrefix, ret)
	require.NoError(t, err)
	assert.Equal(t, "f72589b71ff4f8d139674a3f7369c69b", id.AccountAddress.Hex())
	assert.Equal(t, "cf64428bdeb62af2", id.SubAddress.Hex())
	assert.Equal(t, byte(1), id.Version)
	assert.Equal(t, diemid.MainnetPrefix, id.Prefix)
}

func TestEncodeDecodeAccountIdentifierWithoutSubAddress(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")

	ret, err := diemid.EncodeAccount(diemid.MainnetPrefix, address, diemtypes.EmptySubAddress)
	require.NoError(t, err)
	assert.Equal(t, "lbr1p7ujcndcl7nudzwt8fglhx6wxnvqqqqqqqqqqqqqflf8ma", string(ret))

	id, err := diemid.DecodeToAccount(diemid.MainnetPrefix, ret)
	require.NoError(t, err)
	assert.Equal(t, "f72589b71ff4f8d139674a3f7369c69b", id.AccountAddress.Hex())
	assert.Equal(t, "0000000000000000", id.SubAddress.Hex())
	assert.Equal(t, byte(1), id.Version)
	assert.Equal(t, diemid.MainnetPrefix, id.Prefix)
}

func TestDecodeInvalidAccountIdentifierString(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	subAddress, _ := diemtypes.MakeSubAddress("cf64428bdeb62af2")
	t.Run("invalid checksum", func(t *testing.T) {
		ret, err := diemid.EncodeAccount(diemid.MainnetPrefix, address, subAddress)
		require.NoError(t, err)
		id, err := diemid.DecodeToAccount(diemid.MainnetPrefix, ret[:len(ret)-1])
		require.Error(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Error(), "invalid checksum")
	})
	t.Run("invalid account address length", func(t *testing.T) {
		data := make([]int, diemtypes.AccountAddressLength)
		for i, b := range address {
			data[i] = int(b)
		}

		encoded, err := bech32.SegwitAddrEncode(string(diemid.MainnetPrefix), 1, data)
		require.NoError(t, err)

		id, err := diemid.DecodeToAccount(diemid.MainnetPrefix, encoded)
		require.Error(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Error(), "invalid account identifier")
	})
}
