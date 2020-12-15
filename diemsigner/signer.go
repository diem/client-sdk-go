// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemsigner

import (
	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemtypes"
)

// Sign transaction
func Sign(
	keys *diemkeys.Keys,
	accountAddress diemtypes.AccountAddress,
	sequenceNum uint64, script diemtypes.Script,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) *diemtypes.SignedTransaction {
	rawTxn, signingMsg := NewRawTransactionAndSigningMsg(
		accountAddress,
		sequenceNum, script,
		maxGasAmmount, gasUnitPrice, gasCurrencyCode,
		expirationTimeSec,
		chainID)

	signature := keys.PrivateKey.Sign(signingMsg)
	return NewSignedTransaction(keys.PublicKey, rawTxn, signature)
}

// NewRawTransactionAndSigningMsg creates raw transaction and signing message
func NewRawTransactionAndSigningMsg(
	accountAddress diemtypes.AccountAddress,
	sequenceNum uint64, script diemtypes.Script,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) (*diemtypes.RawTransaction, []byte) {
	rawTxn := diemtypes.RawTransaction{
		Sender:                  accountAddress,
		SequenceNumber:          sequenceNum,
		Payload:                 &diemtypes.TransactionPayload__Script{script},
		MaxGasAmount:            maxGasAmmount,
		GasUnitPrice:            gasUnitPrice,
		GasCurrencyCode:         gasCurrencyCode,
		ExpirationTimestampSecs: expirationTimeSec,
		ChainId:                 diemtypes.ChainId(chainID),
	}

	signingMsg := append(diemtypes.HashPrefix("RawTransaction"), diemtypes.ToBCS(&rawTxn)...)
	return &rawTxn, signingMsg
}

// NewSignedTransaction creates new `SignedTransaction`
func NewSignedTransaction(publicKey diemkeys.PublicKey, rawTxn *diemtypes.RawTransaction, signature []byte) *diemtypes.SignedTransaction {
	var auth diemtypes.TransactionAuthenticator
	if publicKey.IsMulti() {
		auth = &diemtypes.TransactionAuthenticator__MultiEd25519{
			PublicKey: diemtypes.MultiEd25519PublicKey(publicKey.Bytes()),
			Signature: diemtypes.MultiEd25519Signature(signature),
		}
	} else {
		auth = &diemtypes.TransactionAuthenticator__Ed25519{
			PublicKey: diemtypes.Ed25519PublicKey(publicKey.Bytes()),
			Signature: diemtypes.Ed25519Signature(signature),
		}
	}
	return &diemtypes.SignedTransaction{
		RawTxn:        *rawTxn,
		Authenticator: auth,
	}
}
