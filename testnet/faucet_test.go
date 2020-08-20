package testnet_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {
	keys := librakeys.MustGenKeys()
	seq := testnet.MustMint(keys.AuthKey.Hex(), 1000, "LBR")
	assert.True(t, seq > 0)
}

func TestMustMintPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
		assert.Fail(t, "should panic, but not")
	}()

	testnet.MustMint("invalid", 1000, "HELLO")
}
