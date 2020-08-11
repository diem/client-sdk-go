package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Method is type of JSON-RPC method
type Method string

// Param is type of JSON-RPC params list
type Param interface{}

// Request is type of JSON-RPC request struct
type Request struct {
	JsonRpc string  `json:"jsonrpc"`
	Method  Method  `json:"method"`
	Params  []Param `json:"params"`
	ID      int     `json:"id"`
}

// Response is type of JSON-RPC response struct
type Response struct {
	JsonRpc                  string           `json:"jsonrpc"`
	ID                       *int             `json:"id"`
	Result                   *json.RawMessage `json:"result"`
	Error                    *ResponseError   `json:"error"`
	LibraChainID             uint64           `json:"libra_chain_id"`
	LibraLedgerTimestampusec uint64           `json:"libra_ledger_timestampusec"`
	LibraLedgerVersion       uint64           `json:"libra_ledger_version"`
}

// ResponseResult is type for serializing JSON-RPC response result
type ResponseResult interface{}

// ResponseError is type of JSON-RPC response error, it implements error interface.
type ResponseError struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error implements error interface, returns `Code` + `Message`
func (e *ResponseError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

// NewClient creates a new JSON-RPC Client
func NewClient(url string) Client {
	return &client{url: url, http: &http.Client{Transport: &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}}}
}

// Client is interface of the JSON-RPC client
type Client interface {
	Call(Method, ResponseResult, ...Param) (*Response, *Error)
}

type client struct {
	url  string
	http *http.Client
}

// Call implements Client interface
func (c *client) Call(method Method, respResult ResponseResult, params ...Param) (*Response, *Error) {
	if params == nil {
		params = make([]Param, 0)
	}
	request := Request{JsonRpc: "2.0", Method: method, Params: params, ID: 1}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, NewError(SerializeRequestJsonError, err)
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return nil, NewError(HttpCallError, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(ReadHttpResponseBodyError, err)
	}

	var jsonRpcResponse Response
	if err = json.Unmarshal(body, &jsonRpcResponse); err != nil {
		return nil, NewError(ParseResponseJsonError, err)
	}
	if jsonRpcResponse.JsonRpc != "2.0" {
		return nil, NewError(InvalidJsonRpcResponseError,
			fmt.Errorf("unexpected jsonrpc version: %s", jsonRpcResponse.JsonRpc))
	}
	if jsonRpcResponse.Result != nil {
		if err := json.Unmarshal(*jsonRpcResponse.Result, respResult); err != nil {
			return nil, NewError(ParseResponseResultJsonError, err)
		}
	}
	return &jsonRpcResponse, nil
}
