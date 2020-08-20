package librasigner

import (
	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"golang.org/x/crypto/sha3"
)

// Sign transaction
func Sign(
	account *librakeys.Keys, sequenceNum uint64, transactionPayload libratypes.TransactionPayload,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) *libratypes.SignedTransaction {
	rawTxn := libratypes.RawTransaction{
		Sender:                  libratypes.AccountAddress{account.AccountAddress},
		SequenceNumber:          sequenceNum,
		Payload:                 transactionPayload,
		MaxGasAmount:            maxGasAmmount,
		GasUnitPrice:            gasUnitPrice,
		GasCurrencyCode:         gasCurrencyCode,
		ExpirationTimestampSecs: expirationTimeSec,
		ChainId:                 libratypes.ChainId{chainID},
	}

	hash := sha3.New256()
	hash.Write([]byte("LIBRA::RawTransaction"))
	rawTransactionPrefix := hash.Sum(nil)
	signingMsg := append(rawTransactionPrefix, libratypes.ToLCS(rawTxn)...)
	signature := account.PrivateKey.Sign(signingMsg)

	return &libratypes.SignedTransaction{
		RawTxn:        rawTxn,
		Authenticator: account.PublicKey.NewAuthenticator(signature),
	}
}
