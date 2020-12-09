// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

// This file implemenets Diem Account Identifier proposal
// https://github.com/diem/lip/blob/master/lips/lip-5.md

package diemid

import (
	"errors"

	"github.com/diem/client-sdk-go/diemid/bech32"
	"github.com/diem/client-sdk-go/diemtypes"
)

const (
	// MainnetPrefix is mainnet account identifier prefix
	MainnetPrefix NetworkPrefix = "lbr"
	// TestnetPrefix is testnet account identifier prefix
	TestnetPrefix NetworkPrefix = "tlb"

	// version 1
	V1 byte = 1

	// account address length
	AccountAddressLength = diemtypes.AccountAddressLength
	SubAddressLength     = diemtypes.SubAddressLength
)

// NetworkPrefix is account identifier prefix type
type NetworkPrefix string

// Account captures all parts of account identifier
type Account struct {
	Prefix         NetworkPrefix
	Version        byte
	AccountAddress diemtypes.AccountAddress
	SubAddress     diemtypes.SubAddress
}

// NewAccount create new Account with version set to v1.
// Set subAddress == nil for no sub-address case.
func NewAccount(prefix NetworkPrefix, accountAddress diemtypes.AccountAddress, subAddress diemtypes.SubAddress) *Account {
	return &Account{
		Prefix:         prefix,
		Version:        V1,
		AccountAddress: accountAddress,
		SubAddress:     subAddress,
	}
}

// EncodeAccount creates account v1 encode string
// Set subAddress == nil for no sub-address case.
func EncodeAccount(prefix NetworkPrefix, accountAddress diemtypes.AccountAddress, subAddress diemtypes.SubAddress) (string, error) {
	return NewAccount(prefix, accountAddress, subAddress).Encode()
}

// DecodeToAccount decode given encoded account identifier string to `Account`.
// Given NetworkPrefix is used to validate account identifier network prefix, and returns error
// if the network prefix mismatched.
func DecodeToAccount(prefix NetworkPrefix, encodedAccountIdentifier string) (*Account, error) {
	version, data, err := bech32.SegwitAddrDecode(string(prefix), encodedAccountIdentifier)
	if err != nil {
		return nil, err
	}
	if len(data) != AccountAddressLength+SubAddressLength {
		return nil, errors.New("invalid account identifier, account address and sub-address length does not match")
	}

	address, _ := diemtypes.MakeAccountAddressFromBytes(
		ints2bytes(data[:AccountAddressLength]))
	subAddress, _ := diemtypes.MakeSubAddressFromBytes(
		ints2bytes(data[AccountAddressLength:]))
	return &Account{
		Prefix:         prefix,
		Version:        byte(version),
		AccountAddress: address,
		SubAddress:     subAddress,
	}, nil
}

// Encode encodes Account into SegwitAddr string
func (ai *Account) Encode() (string, error) {
	if len(ai.SubAddress) != SubAddressLength {
		return "", errors.New("invalid sub address")
	}
	data := make([]byte, 0, AccountAddressLength+SubAddressLength)
	data = append(data, ai.AccountAddress[:]...)
	data = append(data, ai.SubAddress[:]...)

	return bech32.SegwitAddrEncode(string(ai.Prefix), int(ai.Version), bytes2ints(data))
}

func bytes2ints(bs []byte) []int {
	ret := make([]int, len(bs))
	for i, b := range bs {
		ret[i] = int(b)
	}
	return ret
}

func ints2bytes(is []int) []byte {
	ret := make([]byte, len(is))
	for i, b := range is {
		ret[i] = byte(b)
	}
	return ret
}
