// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package jsonrpc

import "fmt"

// ErrorType is type of the `Error`, which maybe caused by different underlining errors
type ErrorType string

const (
	SerializeRequestJsonError    ErrorType = "serialize request json failed"
	HttpCallError                ErrorType = "http call failed"
	ReadHttpResponseBodyError    ErrorType = "read http response body failed"
	ParseResponseJsonError       ErrorType = "parse response json failed"
	ParseResponseResultJsonError ErrorType = "parse response result json failed"
	InvalidJsonRpcResponseError  ErrorType = "invalid JSON-RPC response: missing result / error field"
)

// Error is a wrap of a type and underlining `Cause` error
type Error struct {
	ErrorType ErrorType
	Cause     error
}

// newError creates new `Error` by gien type and cause
func newError(t ErrorType, cause error) *Error {
	return &Error{t, cause}
}

// Error returns `ErrorType` + `Cause#Error()` as message
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorType, e.Cause.Error())
}
