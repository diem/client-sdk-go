// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes

import (
	"encoding/hex"
	"fmt"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/serde"
)

const (
	// AccountAddressLength is valid account address length
	AccountAddressLength = 16
)

// Serializable interface for `ToLCS`
type Serializable interface {
	Serialize(serializer serde.Serializer) error
}

// ToLCS seralize given `Serializable` into LCS bytes
func ToLCS(t Serializable) []byte {
	s := new(lcs.Serializer)
	t.Serialize(s)
	return s.GetBytes()
}

// Hex returns signed transaction hex encoded string
func (t *SignedTransaction) Hex() string {
	return hex.EncodeToString(ToLCS(t))
}

// HexSignature returns transaction signature hex encoded string
func (t *SignedTransaction) HexSignature() string {
	switch t.Authenticator.(type) {
	case *TransactionAuthenticator__Ed25519:
		sig := t.Authenticator.(*TransactionAuthenticator__Ed25519).Signature
		return hex.EncodeToString(sig.Value)
	}
	panic("t.Authenticator type not found")
}

// Hex returns hex encoded string for the AccountAddress
func (a AccountAddress) Hex() string {
	return hex.EncodeToString(a.Value)
}

// NewAccountAddressFromHex creates account address from given hex string
func NewAccountAddressFromHex(address string) (*AccountAddress, error) {
	bytes, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}
	if len(bytes) != AccountAddressLength {
		return nil, fmt.Errorf(
			"Account address should be 16 bytes, but given %d bytes", len(bytes))
	}
	return &AccountAddress{bytes}, nil
}

// MustNewAccountAddressFromHex creates account address or panic for invalid address hex string
func MustNewAccountAddressFromHex(address string) *AccountAddress {
	ret, err := NewAccountAddressFromHex(address)
	if err != nil {
		panic(err)
	}
	return ret
}
