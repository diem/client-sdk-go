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

// UnmarshalResult unmarshals result json into given struct.
// Returns true, nil for success unmarshal, otherwise first bool
// will always be false.
// Returns false, nil if `Result` is nil.
func (r *Response) UnmarshalResult(result interface{}) (bool, error) {
	if r.Result == nil {
		return false, nil
	}

	if err := json.Unmarshal(*r.Result, result); err != nil {
		return false, newError(ParseResponseResultJsonError, err)
	}
	return true, nil
}

// Client is interface of the JSON-RPC client
type Client interface {
	Call(Method, ...Param) (*Response, error)
}

// NewClient creates a new JSON-RPC Client
func NewClient(url string) Client {
	return &client{url: url, http: &http.Client{Transport: &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}}}
}

type client struct {
	url  string
	http *http.Client
}

// Call implements Client interface
func (c *client) Call(method Method, params ...Param) (*Response, error) {
	if params == nil {
		params = make([]Param, 0)
	}
	request := Request{JsonRpc: "2.0", Method: method, Params: params, ID: 1}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, newError(SerializeRequestJsonError, err)
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return nil, newError(HttpCallError, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, newError(ReadHttpResponseBodyError, err)
	}

	var jsonRpcResponse Response
	if err = json.Unmarshal(body, &jsonRpcResponse); err != nil {
		return nil, newError(ParseResponseJsonError, err)
	}
	if jsonRpcResponse.JsonRpc != "2.0" {
		return nil, newError(InvalidJsonRpcResponseError,
			fmt.Errorf("unexpected jsonrpc version: %s", jsonRpcResponse.JsonRpc))
	}
	return &jsonRpcResponse, nil
}
