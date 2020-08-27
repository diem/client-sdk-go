// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient_test

import (
	"testing"
	"time"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librasigner"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	cases := []struct {
		name string
		call func(t *testing.T, client libraclient.Client)
	}{
		{
			name: "get currencies",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetCurrencies()
				require.Nil(t, err)
				assert.NotEmpty(t, ret)
				assert.Len(t, ret, 3)
			},
		},
		{
			name: "get currencies error",
			call: func(t *testing.T, client libraclient.Client) {
				client = libraclient.New(testnet.ChainID, "invalid")
				ret, err := client.GetCurrencies()
				require.Error(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get metadata",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetMetadata()
				require.Nil(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "get metadata error",
			call: func(t *testing.T, client libraclient.Client) {
				client = libraclient.New(testnet.ChainID, "invalid")
				ret, err := client.GetMetadata()
				require.Error(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get metadata by version",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetMetadataByVersion(1)
				require.Nil(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "get metadata by version error",
			call: func(t *testing.T, client libraclient.Client) {
				client = libraclient.New(testnet.ChainID, "invalid")
				ret, err := client.GetMetadataByVersion(1)
				require.Error(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get account",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccount("0000000000000000000000000A550C18")
				require.Nil(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "get account not found",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccount("10000000010000000000000010000C18")
				require.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get account error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccount("10000000010000000000000010000C1K")
				require.Error(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get account transaction",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccountTransaction(
					"000000000000000000000000000000DD", 0, true)
				require.Nil(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "get account transaction not found",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccountTransaction(
					"10000000010000000000000010000C18", 10000000, true)
				require.Nil(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get account transaction error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccountTransaction(
					"10000000010000000000000010000C1K", 10000000, true)
				require.Error(t, err)
				assert.Nil(t, ret)
			},
		},
		{
			name: "get account transactions",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccountTransactions(
					"000000000000000000000000000000DD", 0, 10, true)
				require.Nil(t, err)
				assert.NotEmpty(t, ret)
			},
		},
		{
			name: "get account transactions error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetAccountTransactions(
					"1668f6be25668c1a17cd8caf6b8d2f2K", 0, 10, true)
				require.Error(t, err)
				assert.Empty(t, ret)
			},
		},
		{
			name: "get transactions",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetTransactions(0, 10, true)
				require.Nil(t, err)
				assert.NotEmpty(t, ret)
				assert.Len(t, ret, 10)
			},
		},
		{
			name: "get transactions error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetTransactions(0, 10000000, true)
				require.Error(t, err)
				assert.Empty(t, ret)
			},
		},
		{
			name: "get events",
			call: func(t *testing.T, client libraclient.Client) {
				account, err := client.GetAccount("000000000000000000000000000000DD")
				require.NoError(t, err)

				ret, err := client.GetEvents(account.SentEventsKey, 2, 5)
				require.Nil(t, err)
				assert.NotEmpty(t, ret)
				assert.Len(t, ret, 5)
			},
		},
		{
			name: "get events error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetEvents(
					"00000000000000001668f6be25668c1a17cd8caf6b8d2f2K", 2, 15)
				require.Error(t, err)
				assert.Empty(t, ret)
			},
		},
		{
			name: "submit data",
			call: func(t *testing.T, client libraclient.Client) {
				err := client.Submit("1668f6be25668c1a17cd8caf6b8d2f25")
				require.Error(t, err)
				jrpcErr, ok := err.(*jsonrpc.ResponseError)
				require.True(t, ok)
				require.Equal(t, "Invalid param data(params[0]): should be hex-encoded string of LCS serialized Libra SignedTransaction type", jrpcErr.Message)
			},
		},
		{
			name: "submit: transaction hex-encoded bytes",
			call: func(t *testing.T, client libraclient.Client) {
				var currencyCode = "LBR"
				var sequenceNum uint64 = 0
				var amount uint64 = 10
				account1 := testnet.GenAccount()
				account2 := testnet.GenAccount()
				script := stdlib.EncodePeerToPeerWithMetadataScript(
					libratypes.Currency(currencyCode),
					account2.AccountAddress(),
					amount, nil, nil)

				txn := librasigner.Sign(
					account1,
					account1.AccountAddress(),
					sequenceNum,
					script,
					10000, 0, currencyCode,
					uint64(time.Now().Add(time.Second*30).Unix()),
					testnet.ChainID,
				)
				err := client.Submit(txn.ToHex())
				require.NoError(t, err)

				ret, err := client.WaitForTransaction3(
					txn.ToHex(), time.Second*5)
				require.NoError(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "submit transaction: with multi-signatures",
			call: func(t *testing.T, client libraclient.Client) {
				var currencyCode = "LBR"
				var sequenceNum uint64 = 0
				var amount uint64 = 10
				account1 := testnet.GenMultiSigAccount()
				address1 := account1.AccountAddress()

				account2 := testnet.GenAccount()
				script := stdlib.EncodePeerToPeerWithMetadataScript(
					libratypes.Currency(currencyCode),
					account2.AccountAddress(),
					amount, nil, nil)

				txn := librasigner.Sign(
					account1,
					address1,
					sequenceNum,
					script,
					10000, 0, currencyCode,
					uint64(time.Now().Add(time.Second*30).Unix()),
					testnet.ChainID,
				)
				err := client.SubmitTransaction(txn)
				require.NoError(t, err)
				ret, err := client.WaitForTransaction2(txn, time.Second*5)
				require.NoError(t, err)
				assert.NotNil(t, ret)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.New(testnet.ChainID, testnet.URL)
			tc.call(t, client)
		})
	}
}
