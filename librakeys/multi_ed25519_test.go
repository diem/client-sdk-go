// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/stretchr/testify/assert"
)

func TestMultiEd25519PublicKey(t *testing.T) {
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

	t.Run("hex", func(t *testing.T) {
		expectedBytes := "20fdbac9b10b7587bba7b5bc163bce69e796d71e4ed44c10fcb4488689f7a14475e4174dd58822548086f17b037cecb0ee86516b7d13400a80c856b4bdaf7fe1631c1541f3a4bf44d4d897061564aa8495d766f6191a3ff61562003f184b8c6502"
		assert.Equal(t, expectedBytes, pk.Hex())
	})

	t.Run("is multi", func(t *testing.T) {
		assert.True(t, pk.IsMulti())
	})

}

func TestMultiEd25519PrivateKey(t *testing.T) {
	keys := []string{
		"76b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770dc7",
		"da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586",
		"9f07e7be5551387a98ba977c732d080dcb0f29a048e3656912c6533e32ee7aed",
	}
	privateKeys := make([]ed25519.PrivateKey, len(keys))
	for i, k := range keys {
		bytes, _ := hex.DecodeString(k)
		privateKeys[i] = ed25519.PrivateKey(ed25519.NewKeyFromSeed(bytes))
	}

	pk := librakeys.NewMultiEd25519PrivateKey(privateKeys, byte(2))

	t.Run("sign", func(t *testing.T) {
		multiSig := pk.Sign([]byte("test"))

		expectedSig := "8401951ea9303fe7c0245a2a4c159b3f641e4623e15091a6e0557eb26144cac9ebe3ab4338b9bc7f7d54e78e9c50f3de10bf43199956f5ed0fbcd3c54a081c43b407e86b66d43e1c70a69e819247c9df579dca7d6927569a89a1863f74af7d8ebe07947425ddf6d0155b8a193c8e859a8b3f7f85191b4c613718d40d0fd3e09ca400c0000000"
		assert.Equal(t, expectedSig, lcsBytes(multiSig))
	})
}

func TestNewMultiEd25519PrivateKeyErrors(t *testing.T) {
	t.Run("empty keys", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.NewMultiEd25519PrivateKey(nil, 0)
	})
	t.Run("threshold > len(keys)", func(t *testing.T) {
		_, privateKey, _ := ed25519.GenerateKey(nil)

		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.NewMultiEd25519PrivateKey([]ed25519.PrivateKey{privateKey}, 2)
	})
	t.Run("len(keys) > max num of keys", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		keys := make([]ed25519.PrivateKey, librakeys.MaxNumOfKeys+1)
		librakeys.NewMultiEd25519PrivateKey(keys, 2)
	})
}

func TestNewMultiEd25519PublicKeyErrors(t *testing.T) {
	t.Run("empty keys", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.NewMultiEd25519PublicKey(nil, 0)
	})
	t.Run("threshold > len(keys)", func(t *testing.T) {
		publicKey, _, _ := ed25519.GenerateKey(nil)

		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.NewMultiEd25519PublicKey([]ed25519.PublicKey{publicKey}, 2)
	})
	t.Run("len(keys) > max num of keys", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		keys := make([]ed25519.PublicKey, librakeys.MaxNumOfKeys+1)
		librakeys.NewMultiEd25519PublicKey(keys, 2)
	})
}

func lcsBytes(bytes []byte) string {
	s := new(lcs.Serializer)
	s.SerializeBytes(bytes)
	return hex.EncodeToString(s.GetBytes())
}
