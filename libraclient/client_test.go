package libraclient_test

import (
	"encoding/hex"
	"encoding/json"
	"testing"
	"time"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/jsonrpc/jsonrpctest"
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/librakeys"
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
			name:     "wait for transaction: timeout",
			response: jsonrpc.Response{},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					hex.EncodeToString(account.AccountAddress),
					0,
					"sig",
					time.Second*1,
				)
				require.EqualError(t, err, "transaction not found within timeout period: 1s")
				assert.Nil(t, ret)
			},
		},
		{
			name: "wait for transaction: signature mismatch",
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
					hex.EncodeToString(account.AccountAddress),
					0,
					"mismatched sig",
					time.Second*5,
				)
				assert.EqualError(t, err, "found transaction, but signature does not match")
				assert.NotNil(t, ret)
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
    "vm_status": { "type": "move_abort", "abort_code": 5, "location":  "00000000000000000000000000000001::LibraAccount"}
}`)),
			},
			call: func(t *testing.T, client libraclient.Client) {
				account := librakeys.MustGenKeys()
				ret, err := client.WaitForTransaction(
					hex.EncodeToString(account.AccountAddress),
					0,
					"a181a036ba68fcd25a7ba9f3895caf720af7aee4bf86c4d798050a1101e75f71ccd891158c8fa0bf349bbb66fb0ba50b29b6fb29822dc04071aff831735e6402",
					time.Second*5,
				)
				assert.EqualError(t, err, "transaction execution failed: map[abort_code:5 location:00000000000000000000000000000001::LibraAccount type:move_abort]")
				assert.NotNil(t, ret)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := libraclient.NewWithJsonRpcClient(&jsonrpctest.Stub{
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
