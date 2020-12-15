// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// HashPrefix returns Diem hashing prefix by given type name
func HashPrefix(name string) []byte {
	return Hash([]byte("DIEM::"), []byte(name))
}

// Hash returns sha3 256 hash bytes for given bytes
func Hash(prefix []byte, bytes []byte) []byte {
	sha256 := sha3.New256()
	sha256.Write(prefix)
	sha256.Write(bytes)
	return sha256.Sum(nil)
}

// TransactionHash returns hex-encoded hash string of the
// transaction that `SignedTransaction` may executed.
func (t *SignedTransaction) TransactionHash() string {
	return hex.EncodeToString(Hash(
		HashPrefix("Transaction"),
		ToBCS(&Transaction__UserTransaction{*t}),
	))
}
