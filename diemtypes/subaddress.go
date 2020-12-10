// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes

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
type SubAddress [SubAddressLength]uint8

// EmptySubAddress represents empty sub-address, used for creating account identifier without sub-address
var EmptySubAddress SubAddress

// MakeSubAddress creates SubAddress from given hex-encoded bytes string
// SubAddress should be 8 bytes.
func MakeSubAddress(str string) (SubAddress, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return EmptySubAddress, err
	}
	return MakeSubAddressFromBytes(bytes)
}

// MakeSubAddressFromBytes from given bytes
func MakeSubAddressFromBytes(bytes []byte) (SubAddress, error) {
	if len(bytes) != SubAddressLength {
		return EmptySubAddress, fmt.Errorf(
			"SubAddress should be %v uint8, but given %v",
			SubAddressLength,
			len(bytes),
		)
	}
	var ret SubAddress
	copy(ret[:], bytes)
	return ret, nil
}

// GenSubAddress generates a random subaddress.
func GenSubAddress() (SubAddress, error) {
	bytes := make([]byte, SubAddressLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return EmptySubAddress, err
	}
	return MakeSubAddressFromBytes(bytes)
}

// MustGenSubAddress calls `GenSubAddress` and panics if got error
func MustGenSubAddress() SubAddress {
	ret, err := GenSubAddress()
	if err != nil {
		panic(err)
	}
	return ret
}

// Hex returns hex-encoded address string
func (a SubAddress) Hex() string {
	return hex.EncodeToString(a[:])
}
