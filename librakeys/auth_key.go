// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"encoding/hex"

	"github.com/libra/libra-client-sdk-go/libratypes"
	"golang.org/x/crypto/sha3"
)

type KeyScheme byte

const (
	Ed25519Key      KeyScheme = 0
	MultiEd25519Key KeyScheme = 1
)

// AuthKey is Libra account authentication key
type AuthKey []byte

// NewAuthKeyFromString creates AuthKey from given hex-encoded key string.
// Returns error if given string is not hex encoded.
func NewAuthKeyFromString(key string) (AuthKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return AuthKey(bytes), nil
}

// MustNewAuthKeyFromString parses given key or panic
func MustNewAuthKeyFromString(key string) AuthKey {
	ret, err := NewAuthKeyFromString(key)
	if err != nil {
		panic(err)
	}
	return ret
}

// NewAuthKey creates AuthKey PublicKey
func NewAuthKey(publicKey PublicKey) AuthKey {
	hash := sha3.New256()
	hash.Write(publicKey.Bytes())
	if publicKey.IsMulti() {
		hash.Write([]byte{byte(MultiEd25519Key)})
	} else {
		hash.Write([]byte{byte(Ed25519Key)})
	}
	return AuthKey(hash.Sum(nil))
}

// Hex returns hex encoded string for the AuthKey
func (k AuthKey) Hex() string {
	return hex.EncodeToString(k)
}

// Prefix returns AuthKey's first 16 bytes
func (k AuthKey) Prefix() []uint8 {
	return k[:libratypes.AccountAddressLength]
}

// AccountAddress return account address from auth key
func (k AuthKey) AccountAddress() libratypes.AccountAddress {
	ret, _ := libratypes.MakeAccountAddressFromBytes(
		k[len(k)-libratypes.AccountAddressLength:])
	return ret
}
