// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemkeys_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/stretchr/testify/assert"
)

func TestMustGenKeys(t *testing.T) {
	keys := diemkeys.MustGenKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey())
	assert.NotEmpty(t, keys.AccountAddress().Hex())
}

func TestMustGenMultiSigKeys(t *testing.T) {
	keys := diemkeys.MustGenMultiSigKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey())
	assert.NotEmpty(t, keys.AccountAddress().Hex())

	for i := 0; i < 1000; i++ {
		keys2 := diemkeys.MustGenMultiSigKeys()
		assert.NotEqual(t, keys.PrivateKey, keys2.PrivateKey)
	}
}
