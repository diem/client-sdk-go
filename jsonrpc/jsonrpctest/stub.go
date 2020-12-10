// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

// Provides a simple json-rpc client stub for testing client without connecting to remote server.
package jsonrpctest

import (
	"time"

	"github.com/diem/client-sdk-go/jsonrpc"
	"github.com/diem/client-sdk-go/testnet"
)

type Stub struct {
	Responses map[jsonrpc.RequestID]jsonrpc.Response
}

func (s *Stub) Call(requests ...*jsonrpc.Request) (map[jsonrpc.RequestID]*jsonrpc.Response, error) {
	ret := make(map[jsonrpc.RequestID]*jsonrpc.Response)
	for _, req := range requests {
		resp := s.Responses[req.ID]
		resp.JsonRpc = req.JsonRpc
		resp.ID = &req.ID
		if resp.DiemChainID == 0 {
			resp.DiemChainID = uint64(testnet.ChainID)
		}
		if resp.DiemLedgerTimestampusec == 0 {
			resp.DiemLedgerTimestampusec = uint64(time.Now().Unix() * 1000000)
		}
		if resp.DiemLedgerVersion == 0 {
			resp.DiemLedgerVersion = 100
		}

		ret[req.ID] = &resp
	}
	return ret, nil
}
