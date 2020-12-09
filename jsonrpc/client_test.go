// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package jsonrpc_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/diem/client-sdk-go/jsonrpc"
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
			serve:  `{"jsonrpc": "2.0", "result": {"code": 1, "msg": "hello"}, "id": 1}`,
			expect: response_result(),
		},
		{
			name:   "response result == null",
			method: "get_code",
			serve:  `{"jsonrpc": "2.0", "result": null, "id": 1}`,
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
			serve:  `{"jsonrpc": "2.0", "result": {"code": 1, "msg": "hello"}, "id": 1}`,
			expect: response_result(),
		},
		{
			name:   "success with result and diem extension fields",
			method: "get_code",
			params: []jsonrpc.Param{"hello", 1},
			serve: `{
  "jsonrpc": "2.0",
  "result": {"code": 1, "msg": "hello"},
  "diem_chain_id": 2,
  "diem_ledger_timestampusec": 3,
  "diem_ledger_version": 4,
  "id": 1
}`,
			expect: list(response_result(), diem_extension()),
		},
		{
			name:   "success with error and diem extension fields",
			method: "get_code",
			params: []jsonrpc.Param{"hello", 1},
			serve: `{
  "jsonrpc": "2.0",
  "error": {"code": 32000, "message": "hello world", "data": {"foo": "bar"}},
  "diem_chain_id": 2,
  "diem_ledger_timestampusec": 3,
  "diem_ledger_version": 4,
  "id": 1
}`,
			expect: list(
				response_error(32000, "hello world", map[string]interface{}{"foo": "bar"}),
				diem_extension(),
			),
		},
		{
			name:   "invalid json response",
			method: "get_code",
			serve:  `{ ... }`,
			expect: expectError(jsonrpc.ParseResponseJsonError),
		},
		{
			name:   "invalid jsonrpc response: jsonrpc version is not 2.0",
			method: "get_code",
			serve:  `{}`,
			expect: expectError(jsonrpc.InvalidJsonRpcResponseError),
		},
		{
			name:   "invalid jsonrpc response: invalid result json",
			method: "get_code",
			serve:  `{"jsonrpc": "2.0", "result": { ... }, "id": 1}`,
			expect: expectError(jsonrpc.ParseResponseJsonError),
		},
		{
			name:   "jsonrpc response type mismatch",
			method: "get_another_code",
			serve:  `{"jsonrpc": "2.0", "result": {"code": "hello", "msg": "hello"}, "id": 1}`,
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
			expect: expectError(jsonrpc.HttpCallError),
		},
		{
			name:   "serialize request error",
			method: "get_code",
			params: []jsonrpc.Param{func() {}},
			expect: expectError(jsonrpc.SerializeRequestJsonError),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := jsonrpc.NewRequest(tc.method, tc.params...)
			server := serve(t, tc.serve, request)
			defer server.Close()
			if tc.url == "" {
				tc.url = server.URL
			}
			client := jsonrpc.NewClient(tc.url)
			resp, err := client.Call(request)
			jerr, _ := err.(*jsonrpc.Error)
			require.True(t, resp == nil || len(resp) == 1)
			tc.expect(t, resp[1], jerr)
		})
	}
}

func TestBatchRequests(t *testing.T) {
	cases := []struct {
		name     string
		requests []*jsonrpc.Request
		serve    string
		expect   func(*testing.T, map[jsonrpc.RequestID]*jsonrpc.Response, error)
	}{
		{
			name: "success",
			requests: []*jsonrpc.Request{
				jsonrpc.NewRequestWithID(1, "get_code"),
				jsonrpc.NewRequestWithID(2, "get_code"),
			},
			serve: `[
  {"jsonrpc": "2.0", "result": {"code": 2, "msg": "hello"}, "id": 2},
  {"jsonrpc": "2.0", "result": {"code": 1, "msg": "world"}, "id": 1}
]`,
			expect: func(t *testing.T, resp map[jsonrpc.RequestID]*jsonrpc.Response, err error) {
				require.NoError(t, err)
				require.Len(t, resp, 2)
				require.NotNil(t, resp[1])
				require.NotNil(t, resp[2])
				var ret result
				ok, _ := resp[1].UnmarshalResult(&ret)
				require.True(t, ok)
				assert.Equal(t, uint64(1), ret.Code)
				assert.Equal(t, "world", ret.Msg)

				ok, _ = resp[2].UnmarshalResult(&ret)
				require.True(t, ok)
				assert.Equal(t, uint64(2), ret.Code)
				assert.Equal(t, "hello", ret.Msg)
			},
		},
		{
			name: "success with one error",
			requests: []*jsonrpc.Request{
				jsonrpc.NewRequestWithID(1, "get_code"),
				jsonrpc.NewRequestWithID(2, "get_code"),
			},
			serve: `[
  {"jsonrpc": "2.0", "result": {"code": 2, "msg": "hello"}, "id": 2},
  {"jsonrpc": "2.0", "error": {"code": 32000, "message": "hello world", "data": null}, "id": 1}
]`,
			expect: func(t *testing.T, resp map[jsonrpc.RequestID]*jsonrpc.Response, err error) {
				require.NoError(t, err)
				require.Len(t, resp, 2)
				require.NotNil(t, resp[1])
				require.NotNil(t, resp[2])
				var ret result
				ok, _ := resp[1].UnmarshalResult(&ret)
				require.False(t, ok)
				assert.Error(t, resp[1].Error)

				ok, _ = resp[2].UnmarshalResult(&ret)
				require.True(t, ok)
				assert.Equal(t, uint64(2), ret.Code)
				assert.Equal(t, "hello", ret.Msg)
			},
		},
		{
			name: "serialize request error",
			requests: []*jsonrpc.Request{
				jsonrpc.NewRequestWithID(1, "get_code"),
				jsonrpc.NewRequestWithID(2, "get_code", func() {}),
			},
			expect: func(t *testing.T, resp map[jsonrpc.RequestID]*jsonrpc.Response, err error) {
				require.Nil(t, resp)
				require.Error(t, err)
				expectError(jsonrpc.SerializeRequestJsonError)(t, nil, err.(*jsonrpc.Error))
			},
		},
		{
			name: "invalid response: missing response",
			requests: []*jsonrpc.Request{
				jsonrpc.NewRequestWithID(1, "get_code"),
				jsonrpc.NewRequestWithID(2, "get_code"),
			},
			serve: `[
  {"jsonrpc": "2.0", "result": {"code": 2, "msg": "hello"}, "id": 2}
]`,
			expect: func(t *testing.T, resp map[jsonrpc.RequestID]*jsonrpc.Response, err error) {
				require.Error(t, err)
				assert.Equal(t, jsonrpc.InvalidJsonRpcResponseError,
					err.(*jsonrpc.Error).ErrorType)
				assert.Len(t, resp, 1)
			},
		},
		{
			name: "invalid response: invalid json",
			requests: []*jsonrpc.Request{
				jsonrpc.NewRequestWithID(1, "get_code"),
				jsonrpc.NewRequestWithID(2, "get_code"),
			},
			serve: `{"jsonrpc": "2.0", "result": {"code": 2, "msg": "hello"}, "id": 2}`,
			expect: func(t *testing.T, resp map[jsonrpc.RequestID]*jsonrpc.Response, err error) {
				require.Error(t, err)
				assert.Equal(t, jsonrpc.ParseResponseJsonError,
					err.(*jsonrpc.Error).ErrorType)
				assert.Nil(t, resp)
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := serve(t, tc.serve, tc.requests...)
			defer server.Close()
			client := jsonrpc.NewClient(server.URL)
			resps, err := client.Call(tc.requests...)
			tc.expect(t, resps, err)
		})
	}
}

func TestNoRequestsCall(t *testing.T) {
	client := jsonrpc.NewClient("url")
	resps, err := client.Call()
	assert.Error(t, err)
	assert.Nil(t, resps)
}

func TestHandleNon200Response(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()
	client := jsonrpc.NewClient(server.URL)
	resps, err := client.Call(jsonrpc.NewRequest("hello"))
	require.Error(t, err)
	assert.Equal(t, jsonrpc.HttpCallError, err.(*jsonrpc.Error).ErrorType)
	assert.Nil(t, resps)
}

func serve(t *testing.T, content string, expectedReqs ...*jsonrpc.Request) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		if len(expectedReqs) == 1 {
			req := make(map[string]interface{})
			err = json.Unmarshal(body, &req)
			require.NoError(t, err)
			expectSame(t, expectedReqs[0], req)
		} else {
			var req []map[string]interface{}
			err = json.Unmarshal(body, &req)
			require.NoError(t, err)
			assert.Equal(t, len(expectedReqs), len(req))
			for i := range req {
				expectSame(t, expectedReqs[i], req[i])
			}
		}

		fmt.Fprintln(w, content)
	}))
}

func expectSame(t *testing.T, expected *jsonrpc.Request, req map[string]interface{}) {
	assert.Equal(t, "2.0", req["jsonrpc"])
	assert.Equal(t, string(expected.Method), req["method"])

	reqParams := req["params"].([]interface{})
	require.NotNil(t, reqParams)
	assert.Len(t, reqParams, len(expected.Params))
	for i, expected := range expected.Params {
		assert.EqualValues(t, expected, reqParams[i])
	}
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

func diem_extension() expectation {
	return func(t *testing.T, resp *jsonrpc.Response, err *jsonrpc.Error) {
		require.NotNil(t, resp)
		assert.Equal(t, uint64(2), resp.DiemChainID)
		assert.Equal(t, uint64(3), resp.DiemLedgerTimestampusec)
		assert.Equal(t, uint64(4), resp.DiemLedgerVersion)
	}
}

func expectError(errorType jsonrpc.ErrorType) expectation {
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
