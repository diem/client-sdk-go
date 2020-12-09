// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenSubAddress(t *testing.T) {
	address, err := diemtypes.GenSubAddress()
	assert.NoError(t, err)
	assert.Len(t, address, 8)
	for i := 0; i < 10000; i++ {
		require.NotEqual(t, address, diemtypes.MustGenSubAddress())
	}
}

func TestNewSubAddressErrorsForInvalidSubAddress(t *testing.T) {
	t.Run("invalid hex-encoded string", func(t *testing.T) {
		_, err := diemtypes.MakeSubAddress("invalid")
		assert.Error(t, err)
	})
	t.Run("invalid bytes length: too long", func(t *testing.T) {
		_, err := diemtypes.MakeSubAddress("f72589b71ff4f8d139674a3f7369c69b")
		assert.Error(t, err)
	})

	t.Run("invalid bytes length: too short", func(t *testing.T) {
		_, err := diemtypes.MakeSubAddress("f72589b")
		assert.Error(t, err)
	})
}

func TestMakeSubAddress(t *testing.T) {
	address, _ := diemtypes.GenSubAddress()
	newSubAddress, err := diemtypes.MakeSubAddress(address.Hex())
	require.NoError(t, err)
	assert.EqualValues(t, address, newSubAddress)
}
