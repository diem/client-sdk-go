// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librajsonrpctypes_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/librajsonrpctypes"
	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseServerResponseJson(t *testing.T) {
	cases := []struct {
		name      string
		unmarshal func(*testing.T, *jsonrpc.Response) interface{}
	}{
		{
			name: "get-account-transactions-with-events.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result []*librajsonrpctypes.Transaction
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-account-with-child-vasp-role.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Account
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-account-with-designated-dealer-role.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Account
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-account-with-parent-vasp-role.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Account
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-account-with-unknown-role.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Account
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-currencies.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result []*librajsonrpctypes.CurrencyInfo
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-latest-metadata.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Metadata
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-received-payment-events.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result []*librajsonrpctypes.Event
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-sent-payment-events.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result []*librajsonrpctypes.Event
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-transactions-with-events.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result []*librajsonrpctypes.Transaction
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
		{
			name: "get-account-transaction-with-events.json",
			unmarshal: func(t *testing.T, resp *jsonrpc.Response) interface{} {
				var result librajsonrpctypes.Transaction
				ok, err := resp.UnmarshalResult(&result)
				assert.True(t, ok)
				require.NoError(t, err)
				return &result
			},
		},
	}
	wd, _ := os.Getwd()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bytes, err := ioutil.ReadFile(wd + "/testdata/" + tc.name)
			var data Data
			err = json.Unmarshal(bytes, &data)
			require.NoError(t, err)
			require.NotNil(t, data.Response)

			result := tc.unmarshal(t, data.Response)

			resultBytes, err := json.Marshal(&result)
			require.NoError(t, err)
			expectedBytes, err := json.Marshal(data.Response.Result)
			require.NoError(t, err)
			opt := jsondiff.DefaultConsoleOptions()
			diff, desc := jsondiff.Compare(expectedBytes, resultBytes, &opt)
			assert.Equal(t, jsondiff.FullMatch, diff, desc)
		})
	}
}

type Data struct {
	Name     string            `json:"name"`
	Request  *jsonrpc.Request  `json:"request"`
	Response *jsonrpc.Response `json:"response"`
}
