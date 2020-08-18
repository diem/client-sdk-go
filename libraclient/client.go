// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/libra/libra-client-sdk-go/jsonrpc"
)

// List of supported methods
const (
	GetCurrencies          jsonrpc.Method = "get_currencies"
	GetMetadata            jsonrpc.Method = "get_metadata"
	GetAccount             jsonrpc.Method = "get_account"
	GetAccountTransaction  jsonrpc.Method = "get_account_transaction"
	GetAccountTransactions jsonrpc.Method = "get_account_transactions"
	GetTransactions        jsonrpc.Method = "get_transactions"
	GetEvents              jsonrpc.Method = "get_events"
	Submit                 jsonrpc.Method = "submit"

	VmStatusExecuted = "executed"
)

// Client is Libra client implements high level APIs
type Client interface {
	GetCurrencies() ([]*CurrencyInfo, error)
	GetMetadata() (*Metadata, error)
	GetMetadataByVersion(uint64) (*Metadata, error)
	GetAccount(Address) (*Account, error)
	GetAccountTransaction(Address, uint64, bool) (*Transaction, error)
	GetAccountTransactions(Address, uint64, uint64, bool) ([]*Transaction, error)
	GetTransactions(uint64, uint64, bool) ([]*Transaction, error)
	GetEvents(string, uint64, uint64) ([]*Event, error)
	Submit(string) error
	WaitForTransaction(
		address Address,
		seq uint64,
		signature string,
		expirationTimeSec uint64,
		timeout time.Duration,
	) (*Transaction, error)
	LastResponseLedgerState() LedgerState
}

// New creates a `LibraClient` connect to given server URL.
// It creates default jsonrpc client `http.Transport` config, if you need to customize
// `http.Transport` config (for better connection pool production usage), call `NewWithJsonRpcClient` with
// `jsonrpc.NewClientWithTransport(url, <your http.Transport>)`
func New(url string) Client {
	return NewWithJsonRpcClient(jsonrpc.NewClient(url))
}

// NewWithJsonRpcClient creates a `LibraClient` with given `jsonrpc.Client`
func NewWithJsonRpcClient(rpc jsonrpc.Client) Client {
	return &client{rpc: rpc}
}

// LedgerState represents response LibraLedgerTimestampusec & LibraLedgerVersion
type LedgerState struct {
	TimestampUsec uint64
	Version       uint64
}

type client struct {
	rpc  jsonrpc.Client
	mux  sync.RWMutex
	last LedgerState
}

// LastResponseLedgerState returns last recorded response ledger state
func (c *client) LastResponseLedgerState() LedgerState {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.last
}

// UpdateLastResponseLedgerState updates LastResponseLedgerState
func (c *client) UpdateLastResponseLedgerState(state LedgerState) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.last = state
}

// WaitForTransaction waits for given (address, sequence number, signature) transaction.
func (c *client) WaitForTransaction(address Address, seq uint64, signature string, expirationTimeSec uint64, timeout time.Duration) (*Transaction, error) {
	step := time.Millisecond * 500
	for i := time.Duration(0); i < timeout; i += step {
		txn, err := c.GetAccountTransaction(address, seq, true)
		if err != nil {
			return nil, err
		}
		if txn != nil {
			if txn.Transaction.Signature != signature {
				return txn, errors.New("found transaction, but signature does not match")
			}
			if txn.VmStatus != VmStatusExecuted {
				return txn, fmt.Errorf("transaction execution failed: %v", txn.VmStatus)

			}
			return txn, nil
		}
		if expirationTimeSec*1000000 < c.LastResponseLedgerState().TimestampUsec {
			return nil, errors.New("transaction expired")
		}
		time.Sleep(step)
	}
	return nil, fmt.Errorf("transaction not found within timeout period: %v", timeout)
}

// GetCurrencies calls to "get_currencies" method
func (c *client) GetCurrencies() ([]*CurrencyInfo, error) {
	var ret []*CurrencyInfo
	ok, err := c.call(GetCurrencies, &ret)
	if !ok {
		return nil, err
	}

	return ret, nil
}

func (c *client) GetMetadata() (*Metadata, error) {
	var ret Metadata
	ok, err := c.call(GetMetadata, &ret)
	if !ok {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetMetadataByVersion(version uint64) (*Metadata, error) {
	var ret Metadata
	ok, err := c.call(GetMetadata, &ret, version)
	if !ok {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetAccount(address Address) (*Account, error) {
	var ret Account
	ok, err := c.call(GetAccount, &ret, address)
	if !ok {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetAccountTransaction(address Address, sequenceNum uint64, includeEvent bool) (*Transaction, error) {
	var ret Transaction
	ok, err := c.call(GetAccountTransaction, &ret, address, sequenceNum, includeEvent)
	if !ok {
		return nil, err
	}
	return &ret, nil
}

func (c *client) GetAccountTransactions(address Address, start uint64, limit uint64, includeEvent bool) ([]*Transaction, error) {
	var ret []*Transaction
	_, err := c.call(GetAccountTransactions, &ret, address, start, limit, includeEvent)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *client) GetTransactions(startVersion uint64, limit uint64, includeEvent bool) ([]*Transaction, error) {
	var ret []*Transaction
	ok, err := c.call(GetTransactions, &ret, startVersion, limit, includeEvent)
	if !ok {
		return nil, err
	}
	return ret, nil
}

func (c *client) GetEvents(key string, start uint64, limit uint64) ([]*Event, error) {
	var ret []*Event
	ok, err := c.call(GetEvents, &ret, key, start, limit)
	if !ok {
		return nil, err
	}
	return ret, nil
}

func (c *client) Submit(data string) error {
	ok, err := c.call(Submit, nil, data)
	if !ok {
		return err
	}
	return nil
}

func (c *client) call(method jsonrpc.Method, ret interface{}, params ...jsonrpc.Param) (bool, error) {
	req := jsonrpc.NewRequest(method, params...)
	resps, err := c.rpc.Call(req)
	if err != nil {
		return false, err
	}
	resp := resps[req.ID]
	if resp.Error != nil {
		return false, resp.Error
	}
	c.UpdateLastResponseLedgerState(LedgerState{
		TimestampUsec: resp.LibraLedgerTimestampusec,
		Version:       resp.LibraLedgerVersion,
	})
	return resp.UnmarshalResult(ret)
}
