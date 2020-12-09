// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMint(t *testing.T) {
	keys := diemkeys.MustGenKeys()
	testnet.MustMint(keys.AuthKey().Hex(), 1000, "XUS")
	account, err := testnet.Client.GetAccount(keys.AccountAddress())
	require.NoError(t, err)

	assert.Len(t, account.Balances, 1)
	assert.Equal(t, "XUS", account.Balances[0].Currency)
	assert.Equal(t, uint64(1000), account.Balances[0].Amount)
}

func TestMustMintPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
		assert.Fail(t, "should panic, but not")
	}()

	testnet.MustMint("invalid", 1000, "HELLO")
}
