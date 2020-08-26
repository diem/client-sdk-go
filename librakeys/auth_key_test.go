// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"encoding/hex"
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/stretchr/testify/assert"
)

func TestAuthKey(t *testing.T) {
	key := librakeys.MustNewAuthKeyFromString(
		"459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d")
	t.Run("account address", func(t *testing.T) {
		assert.Equal(t, "a74fd7c46952c497e75afb0a7932586d", key.AccountAddress().Hex())
	})
	t.Run("hex", func(t *testing.T) {
		assert.Equal(t, "459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d", key.Hex())
	})
	t.Run("prefix", func(t *testing.T) {
		assert.Equal(t, "459c77a38803bd53f3adee52703810e3", hex.EncodeToString(key.Prefix()))
	})
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
