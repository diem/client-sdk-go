// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes

import (
	"encoding/hex"
	"fmt"
)

// LCSable interface for `ToLCS`
type LCSable interface {
	LcsSerialize() ([]byte, error)
}

// ToLCS serialize given `LCSable` into LCS bytes.
// It panics if lcs serialization failed.
func ToLCS(t LCSable) []byte {
	ret, err := t.LcsSerialize()
	if err != nil {
		panic(fmt.Sprintf("lcs serialize failed: %v", err.Error()))
	}
	return ret
}

// ToHex serialize given `LCSable` into LCS bytes and then return hex-encoded string
// It panics if lcs serialization failed.
func ToHex(t LCSable) string {
	return hex.EncodeToString(ToLCS(t))
}
