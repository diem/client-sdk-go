// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/libra/libra-client-sdk-go/libraid"
	"golang.org/x/crypto/sha3"
)

type KeyScheme byte

const (
	Ed25519Key      KeyScheme = 0
	MultiEd25519Key KeyScheme = 1
)

// AuthKey is Libra account authentication key
type AuthKey []byte

// PublicKey is Libra account public key
type PublicKey interface {
	KeyScheme() KeyScheme
	ToBytes() []byte
}

// PrivateKey is Libra account private key
type PrivateKey interface {
	Sign(msg []byte) []byte
}

// Keys holds Libra local account keys
type Keys struct {
	PublicKey      PublicKey
	PrivateKey     PrivateKey
	AuthKey        AuthKey
	AccountAddress libraid.AccountAddress
}

// MustGenKeys generates local account keys, panics if got error
func MustGenKeys() *Keys {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	pk := NewPublicKey(publicKey)
	authKey := NewAuthKey(pk)
	return &Keys{
		pk,
		NewPrivateKey(privateKey),
		authKey,
		authKey.AccountAddress(),
	}
}

func NewPrivateKey(key ed25519.PrivateKey) PrivateKey {
	return &singlePrivateKey{key}
}

// NewPublicKey from single `ed25519.PublicKey`
func NewPublicKey(key ed25519.PublicKey) PublicKey {
	return &singlePublicKey{key}
}

// NewAuthKey return auth key from public key
func NewAuthKey(publicKey PublicKey) AuthKey {
	hash := sha3.New256()
	hash.Write(publicKey.ToBytes())
	hash.Write([]byte{byte(Ed25519Key)})
	return AuthKey(hash.Sum(nil))
}

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

// AccountAddress return account address from auth key
func (k AuthKey) AccountAddress() libraid.AccountAddress {
	return libraid.AccountAddress(k[len(k)-libraid.AccountAddressLength:])
}

// ToString returns hex encoded string for the AuthKey
func (k AuthKey) ToString() string {
	return hex.EncodeToString(k)
}

type singlePublicKey struct {
	pk ed25519.PublicKey
}

func (k *singlePublicKey) KeyScheme() KeyScheme {
	return Ed25519Key
}

func (k *singlePublicKey) ToBytes() []byte {
	return []byte(k.pk)
}

type singlePrivateKey struct {
	pk ed25519.PrivateKey
}

func (k *singlePrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(k.pk, msg)
}
