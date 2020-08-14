package jsonrpc

import "encoding/json"

const (
	jsonRpcVersion = "2.0"
)

// Method is type of JSON-RPC method
type Method string

// Param is type of JSON-RPC params list
type Param interface{}

// RequestID is type for request ID.
type RequestID uint

// Request is type of JSON-RPC request struct
type Request struct {
	Method  Method    `json:"method"`
	Params  []Param   `json:"params"`
	ID      RequestID `json:"id"`
	JsonRpc string    `json:"jsonrpc"`
}

// NewRequest creates `Request` with default `JsonRpc` = "2.0" and ID = 1
func NewRequest(m Method, params ...Param) *Request {
	return NewRequestWithID(1, m, params...)
}

// NewRequestWithID creates `Request` with default `JsonRpc` = "2.0"
func NewRequestWithID(id RequestID, m Method, params ...Param) *Request {
	if params == nil {
		params = []Param{}
	}
	return &Request{m, params, id, jsonRpcVersion}
}

// ToString returns json format of the request
func (r *Request) ToString() string {
	b, _ := json.MarshalIndent(r, "", "\t")
	return string(b)
}
