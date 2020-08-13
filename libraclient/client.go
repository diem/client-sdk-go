// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient

import (
	"github.com/libra/libra-client-sdk-go/jsonrpc"
)

const (
	TESTNET_URL = "https://client.testnet.libra.org/v1"
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
	return &client{rpc}
}

type client struct {
	rpc jsonrpc.Client
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
	return resp.UnmarshalResult(ret)
}
