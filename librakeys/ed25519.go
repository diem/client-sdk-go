// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/libra/libra-client-sdk-go/libratypes"
)

// Ed25519PublicKey implements `PublicKey` interface for ed25519 public key
type Ed25519PublicKey struct {
	pk ed25519.PublicKey
}

// Ed25519PrivateKey implements `PrivateKey` interface for ed25519 private key
type Ed25519PrivateKey struct {
	pk ed25519.PrivateKey
}

// NewEd25519PublicKey creates `Ed25519PublicKey`
func NewEd25519PublicKey(key ed25519.PublicKey) *Ed25519PublicKey {
	return &Ed25519PublicKey{key}
}

// NewEd25519PrivateKey creates `Ed25519PrivateKey`
func NewEd25519PrivateKey(key ed25519.PrivateKey) *Ed25519PrivateKey {
	return &Ed25519PrivateKey{key}
}

// NewEd25519PublicKeyFromString creates `*Ed25519PublicKey` from given hex-encoded
// `ed25519.PublicKey` string
func NewEd25519PublicKeyFromString(key string) (*Ed25519PublicKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return &Ed25519PublicKey{bytes}, nil
}

// NewEd25519PrivateKeyFromString creates `*Ed25519PrivateKey` from given hex-encoded
// `ed25519.PrivateKey` string
func NewEd25519PrivateKeyFromString(key string) (*Ed25519PrivateKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return NewEd25519PrivateKey(ed25519.PrivateKey(bytes)), nil
}

// NewAuthenticator returns `libratypes.TransactionAuthenticator` with given signature bytes and public key
func (k *Ed25519PublicKey) NewAuthenticator(signature []byte) libratypes.TransactionAuthenticator {
	return &libratypes.TransactionAuthenticator__Ed25519{
		PublicKey: libratypes.Ed25519PublicKey([]byte(k.pk)),
		Signature: libratypes.Ed25519Signature(signature),
	}
}

// NewAuthKey creates `AuthKey` from given public key and Ed25519Key scheme
func (k *Ed25519PublicKey) NewAuthKey() AuthKey {
	return newAuthKeyFromPublicKeyAndScheme(k.pk, Ed25519Key)
}

// Hex returns hex string of the public key
func (k *Ed25519PublicKey) Hex() string {
	return hex.EncodeToString(k.pk)
}

// Sign signs given message bytes by private key
func (k *Ed25519PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(k.pk, msg)
}

// Hex returns hex string of private key, used for testing
func (k *Ed25519PrivateKey) Hex() string {
	return hex.EncodeToString(k.pk)
}
