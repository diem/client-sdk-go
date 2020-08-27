// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/stretchr/testify/assert"
)

func TestMustGenKeys(t *testing.T) {
	keys := librakeys.MustGenKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey())
	assert.NotEmpty(t, keys.AccountAddress().Hex())
}

func TestMustGenMultiSigKeys(t *testing.T) {
	keys := librakeys.MustGenMultiSigKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey())
	assert.NotEmpty(t, keys.AccountAddress().Hex())

	for i := 0; i < 1000; i++ {
		keys2 := librakeys.MustGenMultiSigKeys()
		assert.NotEqual(t, keys.PrivateKey, keys2.PrivateKey)
	}
}
