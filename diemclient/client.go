// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemclient

import (
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/diem/client-sdk-go/jsonrpc"
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

// StaleResponseError is error for the case server response latest ledger state is older than
// client knows
type StaleResponseError struct {
	Client LedgerState
	Server LedgerState
}

// Error implements error interface
func (e *StaleResponseError) Error() string {
	return fmt.Sprintf("stale response error: expected server response ledger %v >= %v", e.Server, e.Client)
}

// InvalidTransactionError is error for get a transaction with unexpected details (e.g. vm status is failure)
type InvalidTransactionError struct {
	Transaction Transaction
	Msg         string
}

// Error implements error interface
func (e *InvalidTransactionError) Error() string {
	return e.Msg
}

// Client is Diem client implements high level APIs
type Client interface {
	GetCurrencies() ([]*CurrencyInfo, error)
	GetMetadata() (*Metadata, error)
	GetMetadataByVersion(uint64) (*Metadata, error)
	GetAccount(diemtypes.AccountAddress) (*Account, error)
	GetAccountTransaction(diemtypes.AccountAddress, uint64, bool) (*Transaction, error)
	GetAccountTransactions(diemtypes.AccountAddress, uint64, uint64, bool) ([]*Transaction, error)
	GetTransactions(uint64, uint64, bool) ([]*Transaction, error)
	GetEvents(string, uint64, uint64) ([]*Event, error)
	Submit(signedTxnHex string) error
	SubmitTransaction(txn *diemtypes.SignedTransaction) error

	WaitForTransaction(
		address diemtypes.AccountAddress,
		seq uint64,
		hash string,
		expirationTimeSec uint64,
		timeout time.Duration,
	) (*Transaction, error)
	WaitForTransaction2(
		txn *diemtypes.SignedTransaction,
		timeout time.Duration,
	) (*Transaction, error)
	WaitForTransaction3(
		signedTxnHex string,
		timeout time.Duration,
	) (*Transaction, error)

	LastResponseLedgerState() LedgerState
	UpdateLastResponseLedgerState(state LedgerState) error
	WithRetryOptions(opts ...retry.Option) Client
}

// New creates a `DiemClient` connect to given server URL.
// It creates default jsonrpc client `http.Transport` config, if you need to customize
// `http.Transport` config (for better connection pool production usage), call `NewWithJsonRpcClient` with
// `jsonrpc.NewClientWithTransport(url, <your http.Transport>)`
func New(chainID byte, url string) Client {
	return NewWithJsonRpcClient(chainID, jsonrpc.NewClient(url))
}

// NewWithJsonRpcClient creates a `DiemClient` with given `jsonrpc.Client`
func NewWithJsonRpcClient(chainID byte, rpc jsonrpc.Client) Client {
	return &client{chainID: chainID, rpc: rpc, retryOpts: []retry.Option{retry.LastErrorOnly(true)}}
}

// LedgerState represents response DiemLedgerTimestampusec & DiemLedgerVersion
type LedgerState struct {
	TimestampUsec uint64
	Version       uint64
}

type client struct {
	chainID   byte
	rpc       jsonrpc.Client
	mux       sync.RWMutex
	last      LedgerState
	retryOpts []retry.Option
}

// WithRetryOptions appends given retry options
func (c *client) WithRetryOptions(opts ...retry.Option) Client {
	c.retryOpts = append(c.retryOpts, opts...)
	return c
}

// LastResponseLedgerState returns last recorded response ledger state
func (c *client) LastResponseLedgerState() LedgerState {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.last
}

// UpdateLastResponseLedgerState updates LastResponseLedgerState
func (c *client) UpdateLastResponseLedgerState(state LedgerState) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	var last = c.last
	if last.Version == state.Version && last.TimestampUsec == state.TimestampUsec {
		return nil
	}
	if last.Version > state.Version || last.TimestampUsec > state.TimestampUsec {
		return &StaleResponseError{Client: last, Server: state}
	}

	c.last = state
	return nil
}

// WaitForTransaction3 waits for given `SignedTransaction` hex string
func (c *client) WaitForTransaction3(signedTxnHex string, timeout time.Duration) (*Transaction, error) {
	bytes, err := hex.DecodeString(signedTxnHex)
	if err != nil {
		return nil, err
	}
	txn, err := diemtypes.BcsDeserializeSignedTransaction(bytes)
	if err != nil {
		return nil, fmt.Errorf("Deserialize given hex string as SignedTransaction BCS failed: %v", err.Error())
	}
	return c.WaitForTransaction2(&txn, timeout)
}

// WaitForTransaction2 waits for given `SignedTransaction`
func (c *client) WaitForTransaction2(txn *diemtypes.SignedTransaction, timeout time.Duration) (*Transaction, error) {
	return c.WaitForTransaction(
		txn.RawTxn.Sender,
		txn.RawTxn.SequenceNumber,
		txn.TransactionHash(),
		txn.RawTxn.ExpirationTimestampSecs,
		timeout,
	)
}

// WaitForTransaction waits for given (address, sequence number, hash) transaction.
func (c *client) WaitForTransaction(address diemtypes.AccountAddress, seq uint64, hash string, expirationTimeSec uint64, timeout time.Duration) (*Transaction, error) {
	step := time.Millisecond * 500
	start := time.Now()
	for {
		if time.Since(start) > timeout {
			break
		}
		txn, err := c.GetAccountTransaction(address, seq, true)
		if _, ok := err.(*StaleResponseError); ok {
			continue
		}
		if err != nil {
			return nil, err
		}
		if txn != nil {
			if txn.Hash != hash {
				return nil, &InvalidTransactionError{
					Transaction: *txn,
					Msg: fmt.Sprintf(
						"transaction hash does not match, given %#v, but got %#v",
						hash, txn.Hash),
				}
			}
			if txn.VmStatus.Type != VmStatusExecuted {
				return nil, &InvalidTransactionError{
					Transaction: *txn,
					Msg: fmt.Sprintf(
						"transaction execution failed: %v", txn.VmStatus),
				}

			}
			return txn, nil
		}
		if expirationTimeSec*1_000_000 <= c.LastResponseLedgerState().TimestampUsec {
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

func (c *client) GetAccount(address diemtypes.AccountAddress) (*Account, error) {
	var ret Account
	ok, err := c.call(GetAccount, &ret, address.Hex())
	if !ok {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetAccountTransaction(address diemtypes.AccountAddress, sequenceNum uint64, includeEvent bool) (*Transaction, error) {
	var ret Transaction
	ok, err := c.call(GetAccountTransaction, &ret, address.Hex(), sequenceNum, includeEvent)
	if !ok {
		return nil, err
	}
	return &ret, nil
}

func (c *client) GetAccountTransactions(address diemtypes.AccountAddress, start uint64, limit uint64, includeEvent bool) ([]*Transaction, error) {
	var ret []*Transaction
	_, err := c.call(GetAccountTransactions, &ret, address.Hex(), start, limit, includeEvent)
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

// Submit hex-encoded signed transaction bytes to mempool.
// This function ignores StaleResponseError and does not retry on any errors.
func (c *client) Submit(data string) error {
	ok, err := c.callWithoutRetry(Submit, nil, data)
	if !ok {
		if _, ok := err.(*StaleResponseError); ok {
			return nil
		}
		return err
	}
	return nil
}

func (c *client) SubmitTransaction(txn *diemtypes.SignedTransaction) error {
	return c.Submit(diemtypes.ToHex(txn))
}

func (c *client) call(method jsonrpc.Method, ret interface{}, params ...jsonrpc.Param) (ok bool, err error) {
	err = retry.Do(
		func() error {
			ok, err = c.callWithoutRetry(method, ret, params...)
			return err
		},
		c.retryOpts...,
	)
	return ok, err
}

func (c *client) callWithoutRetry(method jsonrpc.Method, ret interface{}, params ...jsonrpc.Param) (bool, error) {
	req := jsonrpc.NewRequest(method, params...)
	resps, err := c.rpc.Call(req)
	if err != nil {
		return false, err
	}
	resp := resps[req.ID]

	if err = c.validateChainID(byte(resp.DiemChainID)); err != nil {
		return false, err
	}
	err = c.UpdateLastResponseLedgerState(LedgerState{
		TimestampUsec: resp.DiemLedgerTimestampusec,
		Version:       resp.DiemLedgerVersion,
	})
	if err != nil {
		return false, err
	}

	if resp.Error != nil {
		return false, resp.Error
	}
	return resp.UnmarshalResult(ret)
}

func (c *client) validateChainID(chainID byte) error {
	if c.chainID != chainID {
		return fmt.Errorf("chain id mismatch error: expected server response chain id == %d, but got %d", c.chainID, chainID)
	}
	return nil
}
