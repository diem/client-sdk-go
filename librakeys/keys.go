// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/libra/libra-client-sdk-go/libraid"
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

// PublicKey is Libra account public key
type PublicKey interface {
	NewAuthenticator(sig []byte) libratypes.TransactionAuthenticator
	NewAuthKey() AuthKey
	Hex() string
}

// PrivateKey is Libra account private key
type PrivateKey interface {
	Sign(msg []byte) []byte
	Hex() string
}

// Keys holds Libra local account keys
type Keys struct {
	PublicKey      PublicKey
	PrivateKey     PrivateKey
	AuthKey        AuthKey
	AccountAddress libratypes.AccountAddress
}

// MustGenKeys generates local account keys, panics if got error
func MustGenKeys() *Keys {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	pk := NewPublicKey(publicKey)
	authKey := pk.NewAuthKey()
	return &Keys{
		pk,
		NewPrivateKey(privateKey),
		authKey,
		authKey.AccountAddress(),
	}
}

// MustNewKeysFromPublicAndPrivateKeyHexStrings creates `*Keys` from given public and private keys
// it panics if given string is not valid hex-encoded bytes.
func MustNewKeysFromPublicAndPrivateKeyHexStrings(publicKey string, privateKey string) *Keys {
	pk := MustNewPublicKeyFromString(publicKey)
	authKey := pk.NewAuthKey()
	return &Keys{
		pk,
		MustNewPrivateKeyFromString(privateKey),
		authKey,
		authKey.AccountAddress(),
	}
}

// NewPrivateKey from single `ed25519.PrivateKey`
func NewPrivateKey(key ed25519.PrivateKey) PrivateKey {
	return &singlePrivateKey{key}
}

// NewPublicKey from single `ed25519.PublicKey`
func NewPublicKey(key ed25519.PublicKey) PublicKey {
	return &singlePublicKey{key}
}

// NewPrivateKeyFromString creates `PrivateKey` from given hex-encoded key string
func NewPrivateKeyFromString(key string) (PrivateKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return &singlePrivateKey{bytes}, nil
}

// NewPublicKeyFromString creates `PublicKey` from given hex-encoded key string
func NewPublicKeyFromString(key string) (PublicKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return &singlePublicKey{bytes}, nil
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

// MustNewPrivateKeyFromString creates `PrivateKey` from given hex-encoded key string
// or panic
func MustNewPrivateKeyFromString(key string) PrivateKey {
	ret, err := NewPrivateKeyFromString(key)
	if err != nil {
		panic(err)
	}
	return ret
}

// MustNewPublicKeyFromString creates `PublicKey` from given hex-encoded key string
// or panic
func MustNewPublicKeyFromString(key string) PublicKey {
	ret, err := NewPublicKeyFromString(key)
	if err != nil {
		panic(err)
	}
	return ret
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
func (k AuthKey) AccountAddress() libratypes.AccountAddress {
	return libratypes.AccountAddress{k[len(k)-libraid.AccountAddressLength:]}
}

// Hex returns hex encoded string for the AuthKey
func (k AuthKey) Hex() string {
	return hex.EncodeToString(k)
}

type singlePublicKey struct {
	pk ed25519.PublicKey
}

func (k *singlePublicKey) NewAuthenticator(signature []byte) libratypes.TransactionAuthenticator {
	return &libratypes.TransactionAuthenticator__Ed25519{
		PublicKey: libratypes.Ed25519PublicKey{[]byte(k.pk)},
		Signature: libratypes.Ed25519Signature{signature},
	}
}

func (k *singlePublicKey) NewAuthKey() AuthKey {
	hash := sha3.New256()
	hash.Write([]byte(k.pk))
	hash.Write([]byte{byte(Ed25519Key)})
	return AuthKey(hash.Sum(nil))
}

func (k *singlePublicKey) Hex() string {
	return hex.EncodeToString(k.pk)
}

type singlePrivateKey struct {
	pk ed25519.PrivateKey
}

func (k *singlePrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(k.pk, msg)
}

func (k *singlePrivateKey) Hex() string {
	return hex.EncodeToString(k.pk)
}
