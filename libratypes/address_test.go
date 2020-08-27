package libratypes_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountAddress(t *testing.T) {
	t.Run("MakeAccountAddress", func(t *testing.T) {
		keys := librakeys.MustGenKeys()
		address := keys.AccountAddress().Hex()
		accountAddress, err := libratypes.MakeAccountAddress(address)
		require.NoError(t, err)
		assert.Equal(t, address, accountAddress.Hex())
	})

	t.Run("MakeAccountAddress: invalid hex-encoded string", func(t *testing.T) {
		_, err := libratypes.MakeAccountAddress("xx")
		assert.EqualError(t, err, "encoding/hex: invalid byte: U+0078 'x'")
	})

	t.Run("MakeAccountAddress: invalid bytes length", func(t *testing.T) {
		_, err := libratypes.MakeAccountAddress("22")
		assert.EqualError(t, err, "invalid account address bytes length: 1")
	})
}
