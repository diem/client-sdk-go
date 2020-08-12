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

// New creates a `LibraClient` connect to given server URL
func New(url string) Client {
	return &client{jsonrpc.NewClient(url)}
}

type client struct {
	rpc jsonrpc.Client
}

// GetCurrencies calls to "get_currencies" method
func (c *client) GetCurrencies() ([]*CurrencyInfo, error) {
	var ret []*CurrencyInfo
	err := c.call(GetCurrencies, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *client) GetMetadata() (*Metadata, error) {
	var ret Metadata
	err := c.call(GetMetadata, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetMetadataByVersion(version uint64) (*Metadata, error) {
	var ret Metadata
	err := c.call(GetMetadata, &ret, version)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *client) GetAccount(address Address) (*Account, error) {
	var ret Account
	err := c.call(GetAccount, &ret, address)
	if err != nil {
		return nil, err
	}
	if ret.AuthenticationKey == "" {
		return nil, nil
	}

	return &ret, nil
}

func (c *client) GetAccountTransaction(address Address, sequenceNum uint64, includeEvent bool) (*Transaction, error) {
	var ret Transaction
	err := c.call(GetAccountTransaction, &ret, address, sequenceNum, includeEvent)
	if err != nil {
		return nil, err
	}
	if ret.Hash == "" {
		return nil, nil
	}
	return &ret, nil
}

func (c *client) GetAccountTransactions(address Address, start uint64, limit uint64, includeEvent bool) ([]*Transaction, error) {
	var ret []*Transaction
	err := c.call(GetAccountTransactions, &ret, address, start, limit, includeEvent)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *client) GetTransactions(startVersion uint64, limit uint64, includeEvent bool) ([]*Transaction, error) {
	var ret []*Transaction
	err := c.call(GetTransactions, &ret, startVersion, limit, includeEvent)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *client) GetEvents(key string, start uint64, limit uint64) ([]*Event, error) {
	var ret []*Event
	err := c.call(GetEvents, &ret, key, start, limit)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *client) Submit(data string) error {
	err := c.call(GetAccountTransaction, nil, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) call(method jsonrpc.Method, ret interface{}, params ...jsonrpc.Param) error {
	resp, err := c.rpc.Call(method, ret, params...)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
