package testnet_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {
	keys := libraclient.MustGenKeys()
	seq := testnet.MustMint(keys.AuthKey.ToString(), 1000, "LBR")
	assert.True(t, seq > 0)
}
