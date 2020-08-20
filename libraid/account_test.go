// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraid_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/sipa/bech32/ref/go/src/bech32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenSubAddress(t *testing.T) {
	address, err := libraid.GenSubAddress()
	assert.NoError(t, err)
	assert.Len(t, address, 8)
	for i := 0; i < 10000; i++ {
		require.NotEqual(t, address, libraid.MustGenSubAddress())
	}
}

func TestEncodeDecodeAccountIdentifier(t *testing.T) {
	address := *libratypes.MustNewAccountAddressFromHex("f72589b71ff4f8d139674a3f7369c69b")
	subAddress := libraid.MustNewSubAddressFromHex("cf64428bdeb62af2")

	ret, err := libraid.EncodeAccount(libraid.MainnetPrefix, address, subAddress)
	require.NoError(t, err)
	assert.Equal(t, "lbr1p7ujcndcl7nudzwt8fglhx6wxn08kgs5tm6mz4usw5p72t", string(ret))

	id, err := libraid.DecodeToAccount(libraid.MainnetPrefix, ret)
	require.NoError(t, err)
	assert.Equal(t, "f72589b71ff4f8d139674a3f7369c69b", id.AccountAddress.Hex())
	assert.Equal(t, "cf64428bdeb62af2", id.SubAddress.Hex())
	assert.Equal(t, byte(1), id.Version)
	assert.Equal(t, libraid.MainnetPrefix, id.Prefix)
}

func TestEncodeDecodeAccountIdentifierWithoutSubAddress(t *testing.T) {
	address := *libratypes.MustNewAccountAddressFromHex("f72589b71ff4f8d139674a3f7369c69b")

	ret, err := libraid.EncodeAccount(libraid.MainnetPrefix, address, nil)
	require.NoError(t, err)
	assert.Equal(t, "lbr1p7ujcndcl7nudzwt8fglhx6wxnvqqqqqqqqqqqqqflf8ma", string(ret))

	id, err := libraid.DecodeToAccount(libraid.MainnetPrefix, ret)
	require.NoError(t, err)
	assert.Equal(t, "f72589b71ff4f8d139674a3f7369c69b", id.AccountAddress.Hex())
	assert.Equal(t, "0000000000000000", id.SubAddress.Hex())
	assert.Equal(t, byte(1), id.Version)
	assert.Equal(t, libraid.MainnetPrefix, id.Prefix)
}

func TestEncodeShouldValidateAddressLen(t *testing.T) {
	t.Run("invalid account address", func(t *testing.T) {
		address := libratypes.AccountAddress{}
		subAddress := libraid.MustNewSubAddressFromHex("cf64428bdeb62af2")

		ret, err := libraid.EncodeAccount(libraid.MainnetPrefix, address, subAddress)
		require.Error(t, err)
		assert.Equal(t, "", ret)
	})
	t.Run("invalid sub account address", func(t *testing.T) {
		address := *libratypes.MustNewAccountAddressFromHex("f72589b71ff4f8d139674a3f7369c69b")
		subAddress := libraid.SubAddress([]byte{1})

		ret, err := libraid.EncodeAccount(libraid.MainnetPrefix, address, subAddress)
		require.Error(t, err)
		assert.Equal(t, "", ret)
	})
}

func TestDecodeInvalidAccountIdentifierString(t *testing.T) {
	address := *libratypes.MustNewAccountAddressFromHex("f72589b71ff4f8d139674a3f7369c69b")
	subAddress := libraid.MustNewSubAddressFromHex("cf64428bdeb62af2")
	t.Run("invalid checksum", func(t *testing.T) {
		ret, err := libraid.EncodeAccount(libraid.MainnetPrefix, address, subAddress)
		require.NoError(t, err)
		id, err := libraid.DecodeToAccount(libraid.MainnetPrefix, ret[:len(ret)-1])
		require.Error(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Error(), "invalid checksum")
	})
	t.Run("invalid account address length", func(t *testing.T) {
		data := make([]int, libratypes.AccountAddressLength)
		for i, b := range address.Value {
			data[i] = int(b)
		}

		encoded, err := bech32.SegwitAddrEncode(string(libraid.MainnetPrefix), 1, data)
		require.NoError(t, err)

		id, err := libraid.DecodeToAccount(libraid.MainnetPrefix, encoded)
		require.Error(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Error(), "invalid account identifier")
	})
}

func TestMustNewSubAddressFromHexPanicsForInvalidSubAddress(t *testing.T) {
	t.Run("invalid hex-encoded string", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		libraid.MustNewSubAddressFromHex("invalid")
	})
	t.Run("invalid bytes length: too long", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		libraid.MustNewSubAddressFromHex("f72589b71ff4f8d139674a3f7369c69b")
	})

	t.Run("invalid bytes length: too short", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		libraid.MustNewSubAddressFromHex("f72589b")
	})
}
