package jsonrpc

import (
	"encoding/json"
	"fmt"
)

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
	ID                       *RequestID       `json:"id"`
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

// Validate validates response data with JSON-RPC 2.0 spec
func (r *Response) Validate() error {
	if r.JsonRpc != "2.0" {
		return newError(InvalidJsonRpcResponseError,
			fmt.Errorf("unexpected jsonrpc version: %s", r.JsonRpc))
	}
	return nil
}
