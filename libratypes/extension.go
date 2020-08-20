package libratypes

import (
	"encoding/hex"

	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/serde"
)

// Serializable interface for `ToLCS`
type Serializable interface {
	Serialize(serializer serde.Serializer) error
}

// ToLCS seralize given `Serializable` into LCS bytes
func ToLCS(t Serializable) []byte {
	s := new(lcs.Serializer)
	t.Serialize(s)
	return s.GetBytes()
}

// Hex returns signed transaction hex encoded string
func (t *SignedTransaction) Hex() string {
	return hex.EncodeToString(ToLCS(t))
}

// HexSignature returns transaction signature hex encoded string
func (t *SignedTransaction) HexSignature() string {
	switch t.Authenticator.(type) {
	case *TransactionAuthenticator__Ed25519:
		sig := t.Authenticator.(*TransactionAuthenticator__Ed25519).Signature
		return hex.EncodeToString(sig.Value)
	}
	panic("t.Authenticator type not found")
}
