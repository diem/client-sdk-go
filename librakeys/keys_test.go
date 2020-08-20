// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"encoding/hex"
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAuthKeyFromPublicKey(t *testing.T) {
	bytes, err := hex.DecodeString("447fc3be296803c2303951c7816624c7566730a5cc6860a4a1bd3c04731569f5")
	require.NoError(t, err)
	publicKey := librakeys.NewPublicKey(bytes)
	authKey := publicKey.NewAuthKey()
	assert.Equal(t, "459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d",
		authKey.Hex())
}

func TestAccountAddressFromAuthKey(t *testing.T) {
	key := librakeys.MustNewAuthKeyFromString(
		"459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d")
	assert.Equal(t, "a74fd7c46952c497e75afb0a7932586d", key.AccountAddress().Hex())
}

func TestMustGenKeys(t *testing.T) {
	keys := librakeys.MustGenKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey)
	assert.NotEmpty(t, keys.AccountAddress)
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

func TestMustNewKeysFromPublicAndPrivateKeyHexStrings(t *testing.T) {
	sender := librakeys.MustNewKeysFromPublicAndPrivateKeyHexStrings(
		"f549a91fb9989883fb4d38b463308f3ea82074fb39ea74dae61f62e11bf55d25",
		"76e3de861d516283dc285e12ddadc95245a9e98f351c910b0ad722f790bac273",
	)
	assert.Equal(t, "f549a91fb9989883fb4d38b463308f3ea82074fb39ea74dae61f62e11bf55d25",
		sender.PublicKey.Hex())
	assert.Equal(t, "76e3de861d516283dc285e12ddadc95245a9e98f351c910b0ad722f790bac273",
		sender.PrivateKey.Hex())
	assert.Equal(t, "1668f6be25668c1a17cd8caf6b8d2f25",
		sender.AccountAddress.Hex())
	assert.Equal(t, "d939b0214b484bf4d71d08d0247b755a1668f6be25668c1a17cd8caf6b8d2f25",
		sender.AuthKey.Hex())
}

func TestMustNewKeysFromPublicAndPrivateKeyHexStringsPanic(t *testing.T) {
	t.Run("invalid public key", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.MustNewKeysFromPublicAndPrivateKeyHexStrings(
			"invalid",
			"76e3de861d516283dc285e12ddadc95245a9e98f351c910b0ad722f790bac273",
		)
	})
	t.Run("invalid private key", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				return
			}
			assert.Fail(t, "should panic, but not")
		}()
		librakeys.MustNewKeysFromPublicAndPrivateKeyHexStrings(
			"f549a91fb9989883fb4d38b463308f3ea82074fb39ea74dae61f62e11bf55d25",
			"invalid",
		)
	})
}
