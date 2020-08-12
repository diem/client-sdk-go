package libraclient_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/libraclient"

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
			name: "get metadata",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetMetadata()
				require.Nil(t, err)
				assert.NotNil(t, ret)
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
					"1668f6be25668c1a17cd8caf6b8d2f25", 0, true)
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
					"1668f6be25668c1a17cd8caf6b8d2f25", 0, 10, true)
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
				ret, err := client.GetEvents(
					"00000000000000001668f6be25668c1a17cd8caf6b8d2f25", 0, 15)
				require.Nil(t, err)
				assert.NotEmpty(t, ret)
				assert.Len(t, ret, 15)
			},
		},
		{
			name: "get events error",
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetEvents(
					"00000000000000001668f6be25668c1a17cd8caf6b8d2f2K", 0, 15)
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
				require.Equal(t, "Invalid params", jrpcErr.Message)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.New(libraclient.TESTNET_URL)
			tc.call(t, client)
		})
	}
}