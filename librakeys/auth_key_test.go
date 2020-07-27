// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/stretchr/testify/assert"
)

func TestAuthKey(t *testing.T) {
	key := librakeys.MustNewAuthKeyFromString(
		"459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d")
	t.Run("hex", func(t *testing.T) {
		assert.Equal(t, "459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d", key.Hex())
	})
	t.Run("prefix", func(t *testing.T) {
		assert.Equal(t, "459c77a38803bd53f3adee52703810e3", hex.EncodeToString(key.Prefix()))
	})
}

func TestNewAuthKey(t *testing.T) {
	keyHex := "447fc3be296803c2303951c7816624c7566730a5cc6860a4a1bd3c04731569f5"
	publicKey, _ := librakeys.NewEd25519PublicKeyFromString(keyHex)
	authKey := librakeys.NewAuthKey(publicKey)
	assert.Equal(t,
		"459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d",
		authKey.Hex())
	assert.Equal(t,
		"a74fd7c46952c497e75afb0a7932586d",
		authKey.AccountAddress().Hex())
	assert.Equal(t,
		"459c77a38803bd53f3adee52703810e3",
		hex.EncodeToString(authKey.Prefix()))
}

func TestNewAuthKeyFromMultiSigPublicKey(t *testing.T) {
	keys := []string{
		"20fdbac9b10b7587bba7b5bc163bce69e796d71e4ed44c10fcb4488689f7a144",
		"75e4174dd58822548086f17b037cecb0ee86516b7d13400a80c856b4bdaf7fe1",
		"631c1541f3a4bf44d4d897061564aa8495d766f6191a3ff61562003f184b8c65",
	}

	publicKeys := make([]ed25519.PublicKey, len(keys))
	for i, k := range keys {
		bytes, _ := hex.DecodeString(k)
		publicKeys[i] = ed25519.PublicKey(bytes)
	}

	pk := librakeys.NewMultiEd25519PublicKey(publicKeys, byte(2))
	authKey := librakeys.NewAuthKey(pk)
	assert.Equal(t,
		"4b09784ca88af28c16b6e8cf24c36c45c8b5290aa97c1d392679f636790fa5de",
		authKey.Hex())
	assert.Equal(t,
		"c8b5290aa97c1d392679f636790fa5de",
		authKey.AccountAddress().Hex())
	assert.Equal(t,
		"4b09784ca88af28c16b6e8cf24c36c45",
		hex.EncodeToString(authKey.Prefix()))
}

func TestMustNewAuthKeyFromStringPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
		assert.Fail(t, "should panic, but not")
	}()
	librakeys.MustNewAuthKeyFromString("invalid")
}
