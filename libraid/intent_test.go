// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraid_test

import (
	"fmt"
	"testing"

	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeIntent(t *testing.T) {
	address, _ := libratypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	subAddress, _ := libratypes.MakeSubAddress("cf64428bdeb62af2")
	account := libraid.NewAccount(libraid.MainnetPrefix, address, subAddress)
	accountEncode, _ := account.Encode()

	t.Run("without params", func(t *testing.T) {
		intent := libraid.Intent{Account: *account}
		intentEncode, err := intent.Encode()
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("libra://%s", accountEncode), intentEncode)

		ret, err := libraid.DecodeToIntent(libraid.MainnetPrefix, intentEncode)
		require.NoError(t, err)
		require.NotNil(t, ret)
		assert.Equal(t, intent, *ret)
	})

	t.Run("with params", func(t *testing.T) {
		amount := uint64(123)
		intent := libraid.Intent{
			Account: *account,
			Params: libraid.Params{
				Currency: "Coin1",
				Amount:   &amount,
			},
		}
		intentEncode, err := intent.Encode()
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("libra://%s?am=123&c=Coin1", accountEncode), intentEncode)

		ret, err := libraid.DecodeToIntent(libraid.MainnetPrefix, intentEncode)
		require.NoError(t, err)
		require.NotNil(t, ret)
		assert.Equal(t, intent, *ret)
	})
}

func TestDecodeIntentErrors(t *testing.T) {
	t.Run("invalid url", func(t *testing.T) {
		ret, err := libraid.DecodeToIntent(libraid.MainnetPrefix, "s/s/###...")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid intent identifier")
	})
	t.Run("invalid scheme", func(t *testing.T) {
		ret, err := libraid.DecodeToIntent(libraid.MainnetPrefix, "http://account")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid intent scheme")
	})
	t.Run("invalid account identifier", func(t *testing.T) {
		ret, err := libraid.DecodeToIntent(libraid.MainnetPrefix, "libra://accountid")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid account identifier")
	})
}

func TestEncodeIntentErrors(t *testing.T) {
	t.Run("invalid account identifier", func(t *testing.T) {
		intent := libraid.Intent{Account: libraid.Account{}}
		ret, err := intent.Encode()
		require.Error(t, err)
		require.Empty(t, ret)
		assert.Contains(t, err.Error(), "encode account identifier failed")
	})
}
