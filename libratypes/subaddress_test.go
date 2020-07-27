// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenSubAddress(t *testing.T) {
	address, err := libratypes.GenSubAddress()
	assert.NoError(t, err)
	assert.Len(t, address, 8)
	for i := 0; i < 10000; i++ {
		require.NotEqual(t, address, libratypes.MustGenSubAddress())
	}
}

func TestNewSubAddressErrorsForInvalidSubAddress(t *testing.T) {
	t.Run("invalid hex-encoded string", func(t *testing.T) {
		_, err := libratypes.MakeSubAddress("invalid")
		assert.Error(t, err)
	})
	t.Run("invalid bytes length: too long", func(t *testing.T) {
		_, err := libratypes.MakeSubAddress("f72589b71ff4f8d139674a3f7369c69b")
		assert.Error(t, err)
	})

	t.Run("invalid bytes length: too short", func(t *testing.T) {
		_, err := libratypes.MakeSubAddress("f72589b")
		assert.Error(t, err)
	})
}

func TestMakeSubAddress(t *testing.T) {
	address, _ := libratypes.GenSubAddress()
	newSubAddress, err := libratypes.MakeSubAddress(address.Hex())
	require.NoError(t, err)
	assert.EqualValues(t, address, newSubAddress)
}
