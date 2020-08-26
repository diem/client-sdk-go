// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"crypto/ed25519"
	"math/rand"
	"time"

	"github.com/libra/libra-client-sdk-go/libratypes"
)

// PublicKey is Libra account public key
type PublicKey interface {
	NewAuthenticator(sig []byte) libratypes.TransactionAuthenticator
	NewAuthKey() AuthKey
	Hex() string
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
	AccountAddress libratypes.AccountAddress
}

// NewKeysFromPublicAndPrivateKeys creates new `Keys` from given public key and private key
func NewKeysFromPublicAndPrivateKeys(publicKey PublicKey, privateKey PrivateKey) *Keys {
	authKey := publicKey.NewAuthKey()
	return &Keys{
		PublicKey:      publicKey,
		PrivateKey:     privateKey,
		AuthKey:        authKey,
		AccountAddress: *authKey.AccountAddress(),
	}
}

// MustGenKeys generates local account keys, panics if got error
func MustGenKeys() *Keys {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	return NewKeysFromPublicAndPrivateKeys(
		NewEd25519PublicKey(publicKey), NewEd25519PrivateKey(privateKey))
}

// MustGenMultiSigKeys generates `*Keys`, mostly for testing purpose.
// It panics if got error while generating key
func MustGenMultiSigKeys() *Keys {
	rand.Seed(time.Now().UnixNano())
	numOfKeys := 1 + rand.Intn(MaxNumOfKeys)
	publicKeys := make([]ed25519.PublicKey, numOfKeys)
	privateKeys := make([]ed25519.PrivateKey, numOfKeys)
	var err error
	for i := 0; i < numOfKeys; i++ {
		publicKeys[i], privateKeys[i], err = ed25519.GenerateKey(nil)
		if err != nil {
			panic(err)
		}
	}
	threshold := 1 + rand.Intn(numOfKeys)
	return NewKeysFromPublicAndPrivateKeys(
		NewMultiEd25519PublicKey(publicKeys, byte(threshold)),
		NewMultiEd25519PrivateKey(privateKeys, byte(threshold)),
	)
}
