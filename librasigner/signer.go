// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librasigner

import (
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
)

// Sign transaction
func Sign(
	keys *librakeys.Keys,
	accountAddress libratypes.AccountAddress,
	sequenceNum uint64, script libratypes.Script,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) *libratypes.SignedTransaction {
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
	accountAddress libratypes.AccountAddress,
	sequenceNum uint64, script libratypes.Script,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) (*libratypes.RawTransaction, []byte) {
	rawTxn := libratypes.RawTransaction{
		Sender:                  accountAddress,
		SequenceNumber:          sequenceNum,
		Payload:                 &libratypes.TransactionPayload__Script{script},
		MaxGasAmount:            maxGasAmmount,
		GasUnitPrice:            gasUnitPrice,
		GasCurrencyCode:         gasCurrencyCode,
		ExpirationTimestampSecs: expirationTimeSec,
		ChainId:                 libratypes.ChainId(chainID),
	}

	signingMsg := append(libratypes.HashPrefix("RawTransaction"), libratypes.ToLCS(&rawTxn)...)
	return &rawTxn, signingMsg
}

// NewSignedTransaction creates new `SignedTransaction`
func NewSignedTransaction(publicKey librakeys.PublicKey, rawTxn *libratypes.RawTransaction, signature []byte) *libratypes.SignedTransaction {
	var auth libratypes.TransactionAuthenticator
	if publicKey.IsMulti() {
		auth = &libratypes.TransactionAuthenticator__MultiEd25519{
			PublicKey: libratypes.MultiEd25519PublicKey(publicKey.Bytes()),
			Signature: libratypes.MultiEd25519Signature(signature),
		}
	} else {
		auth = &libratypes.TransactionAuthenticator__Ed25519{
			PublicKey: libratypes.Ed25519PublicKey(publicKey.Bytes()),
			Signature: libratypes.Ed25519Signature(signature),
		}
	}
	return &libratypes.SignedTransaction{
		RawTxn:        *rawTxn,
		Authenticator: auth,
	}
}
