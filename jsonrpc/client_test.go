package jsonrpc_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type expectation func(*testing.T, *jsonrpc.Response, *jsonrpc.Error)

func TestCall(t *testing.T) {
	cases := []struct {
		name   string
		method jsonrpc.Method
		params []jsonrpc.Param
		url    string
		serve  string
		expect expectation
	}{
		{
			name:   "success",
			method: "get_code",
			serve:  `{"jsonrpc": "2.0", "result": {"code": 1, "msg": "hello"}}`,
			expect: response_result(),
		},
		{
			name:   "response result == null",
			method: "get_code",
			serve:  `{"jsonrpc": "2.0", "result": null}`,
			expect: func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
				assert.Nil(t, err)
				require.NotNil(t, resp)
				assert.Nil(t, resp.Error)
				assert.Nil(t, resp.Result)

				var ret result
				ok, unmarshalErr := resp.UnmarshalResult(&ret)
				assert.False(t, ok)
				assert.NoError(t, unmarshalErr)
				assert.Equal(t, result{}, ret)
			},
		},
		{
			name:   "success with params",
			method: "get_code",
			params: []jsonrpc.Param{"hello", 1},
			serve:  `{"jsonrpc": "2.0", "result": {"code": 1, "msg": "hello"}}`,
			expect: response_result(),
		},
		{
			name:   "success with result and libra extension fields",
			method: "get_code",
			params: []jsonrpc.Param{"hello", 1},
			serve: `{
  "jsonrpc": "2.0",
  "result": {"code": 1, "msg": "hello"},
  "libra_chain_id": 2,
  "libra_ledger_timestampusec": 3,
  "libra_ledger_version": 4
}`,
			expect: list(response_result(), libra_extension()),
		},
		{
			name:   "success with error and libra extension fields",
			method: "get_code",
			params: []jsonrpc.Param{"hello", 1},
			serve: `{
  "jsonrpc": "2.0",
  "error": {"code": 32000, "message": "hello world", "data": {"foo": "bar"}},
  "libra_chain_id": 2,
  "libra_ledger_timestampusec": 3,
  "libra_ledger_version": 4
}`,
			expect: list(
				response_error(32000, "hello world", map[string]interface{}{"foo": "bar"}),
				libra_extension(),
			),
		},
		{
			name:   "invalid json response",
			method: "get_code",
			serve:  `{ ... }`,
			expect: error(jsonrpc.ParseResponseJsonError),
		},
		{
			name:   "invalid jsonrpc response: jsonrpc version is not 2.0",
			method: "get_code",
			serve:  `{}`,
			expect: error(jsonrpc.InvalidJsonRpcResponseError),
		},
		{
			name:   "invalid jsonrpc response: invalid result json",
			method: "get_code",
			serve:  `{"jsonrpc": "2.0", "result": { ... }}`,
			expect: error(jsonrpc.ParseResponseJsonError),
		},
		{
			name:   "jsonrpc response type mismatch",
			method: "get_another_code",
			serve:  `{"jsonrpc": "2.0", "result": {"code": "hello", "msg": "hello"}}`,
			expect: func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
				assert.Nil(t, err)
				require.NotNil(t, resp)
				assert.Nil(t, resp.Error)
				assert.NotNil(t, resp.Result)

				ok, unmarshalErr := resp.UnmarshalResult(new(result))
				assert.False(t, ok)
				assert.Error(t, unmarshalErr)
				assert.Equal(t, jsonrpc.ParseResponseResultJsonError,
					unmarshalErr.(*jsonrpc.Error).ErrorType)
			},
		},
		{
			name:   "http call error",
			method: "get_code",
			url:    "invalid",
			expect: error(jsonrpc.HttpCallError),
		},
		{
			name:   "serialize request error",
			method: "get_code",
			params: []jsonrpc.Param{func() {}},
			expect: error(jsonrpc.SerializeRequestJsonError),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := serve(t, tc.serve, tc.method, tc.params)
			defer server.Close()
			if tc.url == "" {
				tc.url = server.URL
			}
			client := jsonrpc.NewClient(tc.url)
			resp, err := client.Call(tc.method, tc.params...)
			jerr, _ := err.(*jsonrpc.Error)
			tc.expect(t, resp, jerr)
		})
	}
}

func serve(t *testing.T, content string, method jsonrpc.Method, params []jsonrpc.Param) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		req := make(map[string]interface{})
		err = json.Unmarshal(body, &req)
		require.NoError(t, err)

		assert.Equal(t, "2.0", req["jsonrpc"])
		assert.Equal(t, string(method), req["method"])

		reqParams := req["params"].([]interface{})
		require.NotNil(t, reqParams)
		assert.Len(t, reqParams, len(params))
		for i, expected := range params {
			assert.EqualValues(t, expected, reqParams[i])
		}

		fmt.Fprintln(w, content)
	}))
}

func response_result() expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		assert.Nil(t, err)
		require.NotNil(t, resp)
		assert.Nil(t, resp.Error)
		assert.NotNil(t, resp.Result)

		var ret result
		ok, unmarshalErr := resp.UnmarshalResult(&ret)
		assert.True(t, ok)
		assert.NoError(t, unmarshalErr)
		assert.Equal(t, uint64(1), ret.Code)
		assert.Equal(t, "hello", ret.Msg)
	}
}

func response_error(code int, msg string, data interface{}) expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		assert.Nil(t, err)
		require.NotNil(t, resp)
		assert.NotNil(t, resp.Error)
		assert.Nil(t, resp.Result)

		assert.Contains(t, resp.Error.Error(), strconv.Itoa(code))
		assert.Contains(t, resp.Error.Error(), msg)

		assert.Equal(t, int32(code), resp.Error.Code)
		assert.Equal(t, msg, resp.Error.Message)
		assert.EqualValues(t, data, resp.Error.Data)
	}
}

func libra_extension() expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		require.NotNil(t, resp)
		assert.Equal(t, uint64(2), resp.LibraChainID)
		assert.Equal(t, uint64(3), resp.LibraLedgerTimestampusec)
		assert.Equal(t, uint64(4), resp.LibraLedgerVersion)
	}
}

func error(errorType jsonrpc.ErrorType) expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		assert.Nil(t, resp)
		require.Error(t, err)
		require.NotNil(t, err)
		assert.Equal(t, errorType, err.ErrorType)
		assert.Contains(t, err.Error(), errorType)
	}
}

func list(exps ...expectation) expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		for _, exp := range exps {
			exp(t, resp, err)
		}
	}
}

type result struct {
	Code uint64 `json:"code"`
	Msg  string `json:"msg"`
}
