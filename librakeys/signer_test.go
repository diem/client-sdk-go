// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librakeys_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/stdlib"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	var maxGasAmount uint64 = 1000000
	var gasUnitPrice uint64 = 0
	var seq uint64 = 42
	var expiration uint64 = 1593189628
	var amount uint64 = 100
	var currencyCode = "LBR"

	sender := librakeys.MustNewKeysFromPublicAndPrivateKeyHexStrings(
		"fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c63",
		"b38318e91089220c144854881c48b88975c25d6395ac3aeeb21a287bcfa1ebe9fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c63",
	)
	receiver := librakeys.MustNewKeysFromPublicAndPrivateKeyHexStrings(
		"a761194c93feb3983e6fffb0af9ccc02bc91fe21e1a9c38b24e03dabc40105ed",
		"6762610fdb4bc8acee054bf11870277c63386d64a22ae67a90936e74cb6c4ccba761194c93feb3983e6fffb0af9ccc02bc91fe21e1a9c38b24e03dabc40105ed",
	)

	script := stdlib.EncodePeerToPeerWithMetadataScript(
		stdlib.CurrencyCode(currencyCode),
		receiver.AccountAddress,
		amount, []byte{}, []byte{})

	txn := sender.Sign(
		seq,
		script,
		maxGasAmount, gasUnitPrice, currencyCode,
		expiration,
		testnet.ChainID,
	)
	expected := "e6866fc23780715681be9febd4f771f72a0000000000000001e101a11ceb0b010000000701000202020403061004160205181d0735610896011000000001010000020001000003020301010004010300010501060c0108000506080005030a020a020005060c05030a020a020109000c4c696272614163636f756e741257697468647261774361706162696c6974791b657874726163745f77697468647261775f6361706162696c697479087061795f66726f6d1b726573746f72655f77697468647261775f6361706162696c69747900000000000000000000000000000001010104010c0b0011000c050e050a010a020b030b0438000b05110202010700000000000000000000000000000001034c4252034c4252000403b4b71dbdfaa82e63855337e615889c970164000000000000000400040040420f00000000000000000000000000034c4252fc24f65e00000000020020fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c6340559e5f164fd58fd49947f9ce7b1a9ae2ff8cb799e2a4acec84a140ac2862d9eda66e6d882deb26051522cc85e15623e91b4c1dbc0a7237093d24053e37799e0a"
	assert.Equal(t, expected, txn.Hex())
}
