// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountAddress(t *testing.T) {
	t.Run("MakeAccountAddress", func(t *testing.T) {
		keys := diemkeys.MustGenKeys()
		address := keys.AccountAddress().Hex()
		accountAddress, err := diemtypes.MakeAccountAddress(address)
		require.NoError(t, err)
		assert.Equal(t, address, accountAddress.Hex())
	})

	t.Run("MakeAccountAddress: invalid hex-encoded string", func(t *testing.T) {
		_, err := diemtypes.MakeAccountAddress("xx")
		assert.EqualError(t, err, "encoding/hex: invalid byte: U+0078 'x'")
	})

	t.Run("MakeAccountAddress: invalid bytes length", func(t *testing.T) {
		_, err := diemtypes.MakeAccountAddress("22")
		assert.EqualError(t, err, "invalid account address bytes length: 1")
	})
}
