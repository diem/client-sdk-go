// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package jsonrpctest

import (
	"time"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
	"github.com/libra/libra-client-sdk-go/testnet"
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
		if resp.LibraChainID == 0 {
			resp.LibraChainID = uint64(testnet.ChainID)
		}
		if resp.LibraLedgerTimestampusec == 0 {
			resp.LibraLedgerTimestampusec = uint64(time.Now().Unix() * 1000000)
		}
		if resp.LibraLedgerVersion == 0 {
			resp.LibraLedgerVersion = 100
		}

		ret[req.ID] = &resp
	}
	return ret, nil
}
