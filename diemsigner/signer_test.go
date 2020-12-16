// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemsigner_test

import (
	"testing"

	"github.com/diem/client-sdk-go/diemkeys"
	"github.com/diem/client-sdk-go/diemsigner"
	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/diem/client-sdk-go/stdlib"
	"github.com/diem/client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	var maxGasAmount uint64 = 1000000
	var gasUnitPrice uint64 = 0
	var seq uint64 = 42
	var expiration uint64 = 1593189628
	var amount uint64 = 100
	var currencyCode = "XDX"

	sender := newKeysFromHexKeys("fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c63", "b38318e91089220c144854881c48b88975c25d6395ac3aeeb21a287bcfa1ebe9fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c63")
	receiver := newKeysFromHexKeys(
		"a761194c93feb3983e6fffb0af9ccc02bc91fe21e1a9c38b24e03dabc40105ed",
		"6762610fdb4bc8acee054bf11870277c63386d64a22ae67a90936e74cb6c4ccba761194c93feb3983e6fffb0af9ccc02bc91fe21e1a9c38b24e03dabc40105ed",
	)

	script := stdlib.EncodePeerToPeerWithMetadataScript(
		diemtypes.Currency(currencyCode),
		receiver.AccountAddress(),
		amount, []byte{}, []byte{})

	txn := diemsigner.Sign(
		sender,
		sender.AccountAddress(),
		seq,
		script,
		maxGasAmount, gasUnitPrice, currencyCode,
		expiration,
		testnet.ChainID,
	)
	expected := "e6866fc23780715681be9febd4f771f72a0000000000000001e001a11ceb0b010000000701000202020403061004160205181d0735600895011000000001010000020001000003020301010004010300010501060c0108000506080005030a020a020005060c05030a020a020109000b4469656d4163636f756e741257697468647261774361706162696c6974791b657874726163745f77697468647261775f6361706162696c697479087061795f66726f6d1b726573746f72655f77697468647261775f6361706162696c69747900000000000000000000000000000001010104010c0b0011000c050e050a010a020b030b0438000b051102020107000000000000000000000000000000010358445803584458000403b4b71dbdfaa82e63855337e615889c970164000000000000000400040040420f0000000000000000000000000003584458fc24f65e00000000020020fc4ea02dc1e42b332ac221d716ece959d5b1fc86c156fa4a5d8b77b3886c3c6340833bb10a6b7a45c327426d0f6f20fe140f8641840d7a20cd22ed711ebca0daa4fe9d8d557d1836517435abc21e5d2e423b5d4e331e3f74aafd2c8eeaccbe470e"
	assert.Equal(t, expected, diemtypes.ToHex(txn))
}

func newKeysFromHexKeys(publicKeyHex string, privateKeyHex string) *diemkeys.Keys {
	publicKey, err := diemkeys.NewEd25519PublicKeyFromString(publicKeyHex)
	if err != nil {
		panic(err)
	}
	privateKey, err := diemkeys.NewEd25519PrivateKeyFromString(privateKeyHex)
	if err != nil {
		panic(err)
	}
	return diemkeys.NewKeysFromPublicAndPrivateKeys(publicKey, privateKey)
}
