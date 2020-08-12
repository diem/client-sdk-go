package libraclient

import (
	"crypto/ed25519"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type KeyScheme byte

const (
	Ed25519Key      KeyScheme = 0
	MultiEd25519Key KeyScheme = 1
)

// AccountAddress represents Libra account address
type AccountAddress []byte

// AccountAddress bytes length
const AccountAddressLength = 16

// AuthKey is Libra account authentication key
type AuthKey []byte

// PublicKey is Libra account public key
type PublicKey ed25519.PublicKey

// PrivateKey is Libra account private key
type PrivateKey ed25519.PrivateKey

// Keys holds Libra local account keys
type Keys struct {
	PublicKey      PublicKey
	PrivateKey     PrivateKey
	AuthKey        AuthKey
	AccountAddress AccountAddress
}

// MustGenKeys generates local account keys, panics if got error
func MustGenKeys() *Keys {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	authKey := NewAuthKey(PublicKey(publicKey), Ed25519Key)
	return &Keys{
		PublicKey(publicKey),
		PrivateKey(privateKey),
		authKey,
		authKey.AccountAddress(),
	}
}

// NewAuthKey return auth key from public key
func NewAuthKey(publicKey PublicKey, scheme KeyScheme) AuthKey {
	hash := sha3.New256()
	hash.Write(publicKey)
	hash.Write([]byte{byte(scheme)})
	return AuthKey(hash.Sum(nil))
}

// NewAuthKeyFromString creates AuthKey from given hex-encoded key string.
// Returns error if given string is not hex encoded.
func NewAuthKeyFromString(key string) (AuthKey, error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return AuthKey(bytes), nil
}

// MustNewAuthKeyFromString parses given key or panic
func MustNewAuthKeyFromString(key string) AuthKey {
	ret, err := NewAuthKeyFromString(key)
	if err != nil {
		panic(err)
	}
	return ret
}

// AccountAddress return account address from auth key
func (k AuthKey) AccountAddress() AccountAddress {
	return AccountAddress(k[len(k)-AccountAddressLength:])
}
