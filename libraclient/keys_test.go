package libraclient_test

import (
	"encoding/hex"
	"testing"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAuthKeyFromPublicKey(t *testing.T) {
	bytes, err := hex.DecodeString("447fc3be296803c2303951c7816624c7566730a5cc6860a4a1bd3c04731569f5")
	require.NoError(t, err)
	publicKey := libraclient.PublicKey(bytes)
	authKey := libraclient.NewAuthKey(publicKey, libraclient.Ed25519Key)
	assert.Equal(t, "459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d",
		hex.EncodeToString(authKey))
}

func TestAccountAddressFromAuthKey(t *testing.T) {
	key := libraclient.MustNewAuthKeyFromString(
		"459c77a38803bd53f3adee52703810e3a74fd7c46952c497e75afb0a7932586d")
	assert.Equal(t, "a74fd7c46952c497e75afb0a7932586d", hex.EncodeToString(key.AccountAddress()))
}

func TestMustGenKeys(t *testing.T) {
	keys := libraclient.MustGenKeys()
	assert.NotEmpty(t, keys.PublicKey)
	assert.NotEmpty(t, keys.PrivateKey)
	assert.NotEmpty(t, keys.AuthKey)
	assert.NotEmpty(t, keys.AccountAddress)
}
