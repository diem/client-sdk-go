// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/libra/libra-client-sdk-go/libratypes"
)

const (
	// BitmapNumOfBytes defines length of bitmap appended to multi-signed signature
	BitmapNumOfBytes = 4

	// MaxNumOfKeys defines max number of keys a multi sig keys can have
	MaxNumOfKeys = 32
)

// MultiEd25519PublicKey implements `PublicKey` interface with multi ed25519 sig support
type MultiEd25519PublicKey struct {
	keys      []ed25519.PublicKey
	threshold byte
}

// MultiEd25519PrivateKey implements `PrivateKey` interface with multi ed25519 sig support
type MultiEd25519PrivateKey struct {
	keys      []ed25519.PrivateKey
	threshold byte
}

// NewMultiEd25519PublicKey creates new `MultiEd25519PublicKey` as `PublicKey`
// with given keys and threshold
func NewMultiEd25519PublicKey(keys []ed25519.PublicKey, threshold byte) PublicKey {
	validate(len(keys), threshold)
	return &MultiEd25519PublicKey{keys, threshold}
}

// NewMultiEd25519PrivateKey creates new `MultiEd25519PrivateKey` as `PrivateKey`
// with given keys and threshold
func NewMultiEd25519PrivateKey(keys []ed25519.PrivateKey, threshold byte) PrivateKey {
	validate(len(keys), threshold)
	return &MultiEd25519PrivateKey{keys, threshold}
}

func validate(keysLen int, threshold byte) {
	if keysLen == 0 {
		panic("should at least have 1 key")
	}
	if int(threshold) > keysLen {
		panic("threshold should be less or equal to len(keys)")
	}
	if keysLen > MaxNumOfKeys {
		panic("len(keys) is more than max num of keys")
	}
}

func (k *MultiEd25519PublicKey) NewAuthenticator(signature []byte) libratypes.TransactionAuthenticator {
	return &libratypes.TransactionAuthenticator__MultiEd25519{
		PublicKey: libratypes.MultiEd25519PublicKey{k.ToBytes()},
		Signature: libratypes.MultiEd25519Signature{signature},
	}
}

func (k *MultiEd25519PublicKey) NewAuthKey() AuthKey {
	return NewAuthKeyFromPublicKeyAndScheme(k.ToBytes(), MultiEd25519Key)
}

func (k *MultiEd25519PublicKey) Hex() string {
	return hex.EncodeToString(k.ToBytes())
}

func (k *MultiEd25519PublicKey) ToBytes() []byte {
	var ret []byte
	for _, key := range k.keys {
		ret = append(ret, key...)
	}
	return append(ret, k.threshold)
}

func (k *MultiEd25519PrivateKey) Sign(msg []byte) []byte {
	var bitmap [BitmapNumOfBytes]byte
	var ret []byte
	for i, key := range k.keys[:k.threshold] {
		bitmapSetBit(&bitmap, byte(i))
		ret = append(ret, ed25519.Sign(key, msg)...)
	}
	return append(ret, bitmap[:]...)
}

func bitmapSetBit(input *[BitmapNumOfBytes]byte, index byte) {
	bucket := index / 8
	// It's always invoked with index < 32, thus there is no need to check range.
	pos := index - (bucket * 8)
	input[bucket] |= uint8(128 >> pos)
}
