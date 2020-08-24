// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// TransactionHashHex returns hash hex string of a transaction for a user signed transaction
func TransactionHashHex(txn *SignedTransaction) string {
	return hex.EncodeToString(Hash(
		HashPrefix("Transaction"),
		ToLCS(&Transaction__UserTransaction{*txn}),
	))
}

// HashPrefix returns Libra hashing prefix by given type name
func HashPrefix(name string) []byte {
	return Hash([]byte("LIBRA::"), []byte(name))
}

// Hash returns sha3 256 hash bytes for given bytes
func Hash(prefix []byte, bytes []byte) []byte {
	sha256 := sha3.New256()
	sha256.Write(prefix)
	sha256.Write(bytes)
	return sha256.Sum(nil)
}
