// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemkeys_test

import (
	"encoding/hex"
	"testing"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEd25519PublicKey(t *testing.T) {
	keyHex := "447fc3be296803c2303951c7816624c7566730a5cc6860a4a1bd3c04731569f5"
	publicKey, _ := diemkeys.NewEd25519PublicKeyFromString(keyHex)
	t.Run("Hex", func(t *testing.T) {
		assert.Equal(t, keyHex, publicKey.Hex())
	})
	t.Run("IsMulti", func(t *testing.T) {
		assert.False(t, publicKey.IsMulti())
	})
	t.Run("Bytes", func(t *testing.T) {
		bytes, _ := hex.DecodeString(keyHex)
		assert.Equal(t, bytes, publicKey.Bytes())
	})
}

func TestNewEd25519PublicKeyFromStringError(t *testing.T) {
	_, err := diemkeys.NewEd25519PublicKeyFromString("invalid")
	assert.Error(t, err)
}

func TestEd25519PrivateKey(t *testing.T) {
	keyHex := "b38318e91089220c144854881c48b88975c25d6395ac3aeeb21a287bcfa1ebe9fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c63"
	key, err := diemkeys.NewEd25519PrivateKeyFromString(keyHex)
	require.NoError(t, err)

	t.Run("sign", func(t *testing.T) {
		expectedSig := "46a8d7cb7ba2fe5703b18d72cbbd6f3e19d3d05793a5870b5d22cac191ad757286c9222ed82a21ff3d2ef02bd2f08380607417e21403da44318ecb39a12f2904"
		assert.Equal(t, expectedSig, hex.EncodeToString(key.Sign([]byte("test"))))
	})
	t.Run("hex", func(t *testing.T) {
		assert.Equal(t, keyHex, key.Hex())
	})
}

func TestNewEd25519PrivateKeyFromStringError(t *testing.T) {
	_, err := diemkeys.NewEd25519PrivateKeyFromString("invalid")
	assert.Error(t, err)
}
