// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes

import (
	"encoding/hex"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
)

// LCSable interface for `ToLCS`
type LCSable interface {
	Serialize(serializer serde.Serializer) error
}

// ToLCS serialize given `LCSable` into LCS bytes
func ToLCS(t LCSable) []byte {
	s := lcs.NewSerializer()
	t.Serialize(s)
	return s.GetBytes()
}

// ToLCS convert `RawTransaction` into LCS bytes
func (t *RawTransaction) ToLCS() []byte {
	return ToLCS(t)
}

// ToLCS convert `SignedTransaction` into LCS bytes
func (t *SignedTransaction) ToLCS() []byte {
	return ToLCS(t)
}

// ToHex convert `SignedTransaction` into hex-encoded LCS bytes
func (t *SignedTransaction) ToHex() string {
	return hex.EncodeToString(ToLCS(t))
}
