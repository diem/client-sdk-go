// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemid_test

import (
	"fmt"
	"testing"

	"github.com/diem/client-sdk-go/diemid"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeIntent(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	subAddress, _ := diemtypes.MakeSubAddress("cf64428bdeb62af2")
	account := diemid.NewAccount(diemid.MainnetPrefix, address, subAddress)
	accountEncode, _ := account.Encode()

	t.Run("without params", func(t *testing.T) {
		intent := diemid.Intent{Account: *account}
		intentEncode, err := intent.Encode()
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("diem://%s", accountEncode), intentEncode)

		ret, err := diemid.DecodeToIntent(diemid.MainnetPrefix, intentEncode)
		require.NoError(t, err)
		require.NotNil(t, ret)
		assert.Equal(t, intent, *ret)
	})

	t.Run("with params", func(t *testing.T) {
		amount := uint64(123)
		intent := diemid.Intent{
			Account: *account,
			Params: diemid.Params{
				Currency: "XUS",
				Amount:   &amount,
			},
		}
		intentEncode, err := intent.Encode()
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("diem://%s?am=123&c=XUS", accountEncode), intentEncode)

		ret, err := diemid.DecodeToIntent(diemid.MainnetPrefix, intentEncode)
		require.NoError(t, err)
		require.NotNil(t, ret)
		assert.Equal(t, intent, *ret)
	})
}

func TestDecodeIntentErrors(t *testing.T) {
	t.Run("invalid url", func(t *testing.T) {
		ret, err := diemid.DecodeToIntent(diemid.MainnetPrefix, "s/s/###...")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid intent identifier")
	})
	t.Run("invalid scheme", func(t *testing.T) {
		ret, err := diemid.DecodeToIntent(diemid.MainnetPrefix, "http://account")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid intent scheme")
	})
	t.Run("invalid account identifier", func(t *testing.T) {
		ret, err := diemid.DecodeToIntent(diemid.MainnetPrefix, "diem://accountid")
		require.Error(t, err)
		require.Nil(t, ret)
		assert.Contains(t, err.Error(), "invalid account identifier")
	})
}

func TestEncodeIntentErrors(t *testing.T) {
	t.Run("invalid account identifier", func(t *testing.T) {
		intent := diemid.Intent{Account: diemid.Account{}}
		ret, err := intent.Encode()
		require.Error(t, err)
		require.Empty(t, ret)
		assert.Contains(t, err.Error(), "encode account identifier failed")
	})
}
