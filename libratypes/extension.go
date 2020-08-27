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

// SignatureHex returns transaction signature hex encoded string
func (t *SignedTransaction) SignatureHex() string {
	switch auth := t.Authenticator.(type) {
	case *TransactionAuthenticator__Ed25519:
		return hex.EncodeToString(auth.Signature)
	case *TransactionAuthenticator__MultiEd25519:
		return hex.EncodeToString(auth.Signature)
	default:
		panic(fmt.Sprintf("unknown Authenticator type: %v", auth))
	}
}

// Hex returns hex encoded string for the AccountAddress
func (a *AccountAddress) Hex() string {
	return hex.EncodeToString(a[:])
}

// NewAccountAddressFromHex creates account address from given hex string
func NewAccountAddressFromHex(address string) (*AccountAddress, error) {
	bytes, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}
	return NewAccountAddressFromBytes(bytes)
}

// MustNewAccountAddressFromHex creates account address or panic for invalid address hex string
func MustNewAccountAddressFromHex(address string) *AccountAddress {
	ret, err := NewAccountAddressFromHex(address)
	if err != nil {
		panic(err)
	}
	return ret
}

// NewAccountAddressFromBytes creates `AccountAddress` from given bytes,
// returns error if given bytes length != 16
func NewAccountAddressFromBytes(bytes []byte) (*AccountAddress, error) {
	if len(bytes) != AccountAddressLength {
		return nil, fmt.Errorf(
			"Account address should be 16 bytes, but given %d bytes", len(bytes))
	}
	address := &AccountAddress{}
	copy(address[:], bytes)
	return address, nil
}

// MustNewAccountAddressFromBytes panics if given bytes length != 16
func MustNewAccountAddressFromBytes(bytes []byte) *AccountAddress {
	ret, err := NewAccountAddressFromBytes(bytes)
	if err != nil {
		panic(err)
	}
	return ret
}
