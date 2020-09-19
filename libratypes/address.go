// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes

import (
	"encoding/hex"
	"fmt"
)

const AccountAddressLength = 16

// MakeAccountAddress creates account address from given hex string,
// it returns error if given string is not hex-encoded or decoded bytes length
// does not meet requirement (16 bytes).
func MakeAccountAddress(address string) (AccountAddress, error) {
	bytes, err := hex.DecodeString(address)
	if err != nil {
		return AccountAddress{}, err
	}
	return MakeAccountAddressFromBytes(bytes)
}

// MakeAccountAddressFromBytes creates account address from given bytes, it returns
// error if given bytes length does not meet requirement (16 bytes).
func MakeAccountAddressFromBytes(bytes []byte) (AccountAddress, error) {
	var ret AccountAddress
	if len(bytes) != AccountAddressLength {
		return ret, fmt.Errorf(
			"invalid account address bytes length: %v", len(bytes))
	}
	copy(ret[:], bytes)
	return ret, nil
}

// MustMakeAccountAddress panics if parse given string address failed
func MustMakeAccountAddress(address string) AccountAddress {
	ret, err := MakeAccountAddress(address)
	if err != nil {
		panic(err)
	}
	return ret
}

// Hex returns hex-encoded string of the address
func (a AccountAddress) Hex() string {
	return hex.EncodeToString(a[:])
}
