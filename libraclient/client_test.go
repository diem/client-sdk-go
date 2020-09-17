// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/jsonrpc/jsonrpctest"
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitForTransaction(t *testing.T) {
	cases := []struct {
		name     string
		response jsonrpc.Response
		call     func(t *testing.T, client libraclient.Client)
	}{
		{
			name: "wait for transaction: success",
			response: jsonrpc.Response{
				Result: toPtr(json.RawMessage(`{
    "events": [],
    "gas_used": 175,
    "hash": "0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
    "transaction": {
      "chain_id": 2,
      "expiration_timestamp_secs": 100000000000,
      "sequence_number": 0,
      "signature": "a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
      "type": "user"
    },
    "version": 106548,
    "vm_status": { "type": "executed" }
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
					uint64(time.Now().Add(2*time.Second).Unix()),
					time.Second*1,
				)
				require.NoError(t, err)
				assert.NotNil(t, ret)
			},
		},
		{
			name: "wait for transaction: timeout when server response stale version",
			response: jsonrpc.Response{
				LibraLedgerVersion:       10,
				LibraLedgerTimestampusec: 1597722856123456,
				Result: toPtr(json.RawMessage(`{
    "events": [],
    "gas_used": 175,
    "hash": "0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
    "transaction": {
      "chain_id": 2,
      "expiration_timestamp_secs": 100000000000,
      "sequence_number": 0,
      "signature": "a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
      "type": "user"
    },
    "version": 106548,
    "vm_status": { "type": "executed" }
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				client.UpdateLastResponseLedgerState(libraclient.LedgerState{
					Version:       11,
					TimestampUsec: 1597722856123457,
				})
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
					uint64(time.Now().Add(2*time.Second).Unix()),
					time.Second*1,
				)
				require.EqualError(t, err, "transaction not found within timeout period: 1s")
				assert.Nil(t, ret)
			},
		},
		{
			name:     "wait for transaction: timeout",
			response: jsonrpc.Response{},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"invalid-hash",
					uint64(time.Now().Add(2*time.Second).Unix()),
					time.Second*1,
				)
				require.EqualError(t, err, "transaction not found within timeout period: 1s")
				assert.Nil(t, ret)
			},
		},
		{
			name: "wait for transaction: hash mismatch",
			response: jsonrpc.Response{
				Result: toPtr(json.RawMessage(`{
    "events": [],
    "gas_used": 175,
    "hash": "0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
    "transaction": {
      "chain_id": 2,
      "expiration_timestamp_secs": 100000000000,
      "sequence_number": 0,
      "signature": "a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
      "type": "user"
    },
    "version": 106548,
    "vm_status": { "type": "executed" }
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"mismatched hash",
					uint64(time.Now().Add(time.Second).Unix()),
					time.Second*5,
				)
				assert.EqualError(t, err, "found transaction, but hash does not match, given \"mismatched hash\", but got \"0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e\"")
				assert.Nil(t, ret)
			},
		},
		{
			name: "wait for transaction: execution failed",
			response: jsonrpc.Response{
				Result: toPtr(json.RawMessage(`{
    "events": [],
    "gas_used": 175,
    "hash": "0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
    "transaction": {
      "chain_id": 2,
      "expiration_timestamp_secs": 100000000000,
      "sequence_number": 0,
      "signature": "a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
      "type": "user"
    },
    "version": 106548,
    "vm_status": { "type": "move_abort", "abort_code": 5, "location": "00000000000000000000000000000001::LibraAccount"}
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
					uint64(time.Now().Add(time.Second).Unix()),
					time.Second*5,
				)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "transaction execution failed")
				assert.Nil(t, ret)
			},
		},
		{
			name: "wait for transaction: expired",
			response: jsonrpc.Response{
				Result:                   nil,
				LibraLedgerTimestampusec: 1597722856123456,
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"0fa27a781a9086e80a870851ea4f1b14090fb8b5bd9933e27447ab806443e08e",
					uint64(1597722856),
					time.Second*5,
				)
				assert.EqualError(t, err, "transaction expired")
				assert.Nil(t, ret)
			},
		},
		{
			name: "wait for transaction: if onchain time is exactly same with expiration time",
			response: jsonrpc.Response{
				Result:                   nil,
				LibraLedgerTimestampusec: 1597722856000000,
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					account.AccountAddress().Hex(),
					0,
					"a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
					uint64(1597722856),
					time.Second*5,
				)
				assert.EqualError(t, err, "transaction expired")
				assert.Nil(t, ret)
			},
		},
		{
			name:     "wait for transaction3 invalid hex string",
			response: jsonrpc.Response{},
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.WaitForTransaction3("invalid", time.Second)
				require.EqualError(t, err, "encoding/hex: invalid byte: U+0069 'i'")
				assert.Nil(t, ret)
			},
		},
		{
			name:     "wait for transaction3: not a signed transaction lcs",
			response: jsonrpc.Response{},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction3(
					account.AccountAddress().Hex(),
					time.Second)
				require.EqualError(t, err, "Deserialize given hex string as SignedTransaction LCS failed: EOF")
				assert.Nil(t, ret)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.NewWithJsonRpcClient(testnet.ChainID, &jsonrpctest.Stub{
				Responses: map[jsonrpc.RequestID]jsonrpc.Response{
					1: tc.response,
				},
			})
			tc.call(t, client)
		})
	}
}

func TestHandleStaleResponse(t *testing.T) {
	cases := []struct {
		name     string
		response jsonrpc.Response
		call     func(t *testing.T, client libraclient.Client)
	}{
		{
			name: "return error if server response version is older",
			response: jsonrpc.Response{
				LibraLedgerVersion:       9,
				LibraLedgerTimestampusec: 1597722856123456,
				Result: toPtr(json.RawMessage(`{
    "timestamp": 1597722856123456,
    "version": 9,
    "chain_id": 2
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				client.UpdateLastResponseLedgerState(libraclient.LedgerState{
					Version:       10,
					TimestampUsec: 1597722856123477,
				})
				ret, err := client.GetMetadata()
				assert.EqualError(t, err, "stale response error: expected server response ledger {1597722856123456 9} >= {1597722856123477 10}")
				assert.Nil(t, ret)
			},
		},
		{
			name: "return error if server response timestamp is older",
			response: jsonrpc.Response{
				LibraLedgerVersion:       10,
				LibraLedgerTimestampusec: 1597722856123456,
				Result: toPtr(json.RawMessage(`{
    "timestamp": 1597722856123456,
    "version": 10,
    "chain_id": 2
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				client.UpdateLastResponseLedgerState(libraclient.LedgerState{
					Version:       10,
					TimestampUsec: 1597722856123477,
				})
				ret, err := client.GetMetadata()
				assert.EqualError(t, err, "stale response error: expected server response ledger {1597722856123456 10} >= {1597722856123477 10}")
				assert.Nil(t, ret)
			},
		},
		{
			name: "update last response state if server response version & timestamp is new",
			response: jsonrpc.Response{
				LibraLedgerVersion:       11,
				LibraLedgerTimestampusec: 1597722856123488,
				Result: toPtr(json.RawMessage(`{
    "timestamp": 1597722856123488,
    "version": 11,
    "chain_id": 2
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				client.UpdateLastResponseLedgerState(libraclient.LedgerState{
					Version:       10,
					TimestampUsec: 1597722856123477,
				})
				_, err := client.GetMetadata()
				assert.NoError(t, err)
				last := client.LastResponseLedgerState()
				assert.Equal(t, uint64(11), last.Version)
				assert.Equal(t, uint64(1597722856123488), last.TimestampUsec)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.NewWithJsonRpcClient(testnet.ChainID, &jsonrpctest.Stub{
				Responses: map[jsonrpc.RequestID]jsonrpc.Response{
					1: tc.response,
				},
			})
			tc.call(t, client)
		})
	}
}

func TestValidateChainID(t *testing.T) {
	cases := []struct {
		name     string
		response jsonrpc.Response
		call     func(t *testing.T, client libraclient.Client)
	}{
		{
			name: "return error if server response chain id mismatched",
			response: jsonrpc.Response{
				LibraChainID: 9,
				Result: toPtr(json.RawMessage(`{
    "timestamp": 1597722856123456,
    "version": 9,
    "chain_id": 9
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				ret, err := client.GetMetadata()
				assert.EqualError(t, err, "chain id mismatch error: expected server response chain id == 2, but got 9")
				assert.Nil(t, ret)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.NewWithJsonRpcClient(testnet.ChainID, &jsonrpctest.Stub{
				Responses: map[jsonrpc.RequestID]jsonrpc.Response{
					1: tc.response,
				},
			})
			tc.call(t, client)
		})
	}
}

func toPtr(msg json.RawMessage) *json.RawMessage {
	return &msg
}
