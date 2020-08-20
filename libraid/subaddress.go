// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	// SubAddressLength is valid sub-address length
	SubAddressLength = 8
)

// SubAddress represents sub-address bytes
type SubAddress []byte

// GenSubAddress generates a random subaddress.
func GenSubAddress() (SubAddress, error) {
	bytes := make([]byte, SubAddressLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return SubAddress(bytes), nil
}

// MustGenSubAddress calls `GenSubAddress` and panics if got error
func MustGenSubAddress() SubAddress {
	ret, err := GenSubAddress()
	if err != nil {
		panic(err)
	}
	return ret
}

// NewSubAddressFromHex creates SubAddress from given hex-encoded bytes string
// SubAddress should be 8 bytes.
func NewSubAddressFromHex(str string) (SubAddress, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	if len(bytes) != SubAddressLength {
		return nil, fmt.Errorf(
			"Subaddress should be 8 bytes, but given %d bytes", len(bytes))
	}

	return SubAddress(bytes), nil
}

// MustNewSubAddressFromHex creates SubAddress or panic if given hex-encoded subaddress
// is invalid
func MustNewSubAddressFromHex(str string) SubAddress {
	ret, err := NewSubAddressFromHex(str)
	if err != nil {
		panic(err)
	}
	return ret
}

// Hex returns hex-encoded address string
func (a SubAddress) Hex() string {
	return hex.EncodeToString(a)
}
