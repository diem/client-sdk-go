// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys

import (
	"github.com/libra/libra-client-sdk-go/libratypes"
)

// Sign transaction
func (keys *Keys) Sign(
	sequenceNum uint64, script libratypes.Script,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) *libratypes.SignedTransaction {
	rawTxn := libratypes.RawTransaction{
		Sender:                  keys.AccountAddress,
		SequenceNumber:          sequenceNum,
		Payload:                 &libratypes.TransactionPayload__Script{script},
		MaxGasAmount:            maxGasAmmount,
		GasUnitPrice:            gasUnitPrice,
		GasCurrencyCode:         gasCurrencyCode,
		ExpirationTimestampSecs: expirationTimeSec,
		ChainId:                 libratypes.ChainId{chainID},
	}

	signingMsg := append(libratypes.HashPrefix("RawTransaction"), libratypes.ToLCS(&rawTxn)...)
	signature := keys.PrivateKey.Sign(signingMsg)

	return &libratypes.SignedTransaction{
		RawTxn:        rawTxn,
		Authenticator: keys.PublicKey.NewAuthenticator(signature),
	}
}