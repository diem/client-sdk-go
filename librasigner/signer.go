package librasigner

import (
	"encoding/hex"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/libra/libra-client-sdk-go/libraid"
	"github.com/libra/libra-client-sdk-go/librakeys"
	"golang.org/x/crypto/sha3"
)

// RawTransaction captures raw transaction data
type RawTransaction struct {
	Address            libraid.AccountAddress
	SequenceNum        uint64
	TransactionPayload []byte
	MaxGasAmount       uint64
	GasUnitPrice       uint64
	GasCurrencyCode    string
	ExpirationTimeSec  uint64
	ChainID            byte
}

func (t *RawTransaction) SigningMessage() []byte {
	hash := sha3.New256()
	hash.Write([]byte("LIBRA::RawTransaction"))
	rawTransactionPrefix := hash.Sum(nil)
	return append(rawTransactionPrefix, t.ToLCS()...)
}

func (t *RawTransaction) ToLCS() []byte {
	s := new(lcs.Serializer)
	for _, b := range t.Address {
		s.SerializeU8(b)
	}
	s.SerializeU64(t.SequenceNum)
	// Script
	for _, b := range t.TransactionPayload {
		s.SerializeU8(b)
	}
	s.SerializeU64(t.MaxGasAmount)
	s.SerializeU64(t.GasUnitPrice)
	s.SerializeStr(t.GasCurrencyCode)
	s.SerializeU64(t.ExpirationTimeSec)
	s.SerializeU8(t.ChainID)

	return s.GetBytes()
}

type SignedTransaction struct {
	Raw       RawTransaction
	Signature []byte
	PublicKey librakeys.PublicKey
}

func (t *SignedTransaction) ToLCS() []byte {
	s := new(lcs.Serializer)
	for _, b := range t.Raw.ToLCS() {
		s.SerializeU8(b)
	}

	s.SerializeVariantIndex(uint32(t.PublicKey.KeyScheme()))
	s.SerializeBytes(t.PublicKey.ToBytes())
	s.SerializeBytes(t.Signature)
	return s.GetBytes()
}

// HexSignature returns signature of raw transaction hex encoded string
func (t *SignedTransaction) HexSignature() string {
	return hex.EncodeToString(t.Signature)
}

// HexSignedTransaction returns signed transaction hex encoded string
func (t *SignedTransaction) HexSignedTransaction() string {
	return hex.EncodeToString(t.ToLCS())
}

// Sign transaction
func Sign(
	account *librakeys.Keys, sequenceNum uint64, transactionPayload []byte,
	maxGasAmmount uint64, gasUnitPrice uint64, gasCurrencyCode string,
	expirationTimeSec uint64,
	chainID byte,
) *SignedTransaction {
	rawTxn := RawTransaction{
		account.AccountAddress, sequenceNum, transactionPayload,
		maxGasAmmount, gasUnitPrice, gasCurrencyCode,
		expirationTimeSec, chainID,
	}
	signature := account.PrivateKey.Sign(rawTxn.SigningMessage())
	return &SignedTransaction{
		Raw:       rawTxn,
		Signature: signature,
		PublicKey: account.PublicKey,
	}
}
