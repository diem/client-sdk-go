// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/libratypes"
)

const (
	URL            = "https://testnet.libra.org/v1"
	FaucetURL      = "https://testnet.libra.org/mint"
	ChainID   byte = 2
)

var (
	// DDAccountAddress is testnet default dd account address
	DDAccountAddress = libratypes.MustMakeAccountAddress("000000000000000000000000000000DD")
	// Client is testnet client
	Client = libraclient.New(ChainID, URL)
)

// Currencies
var (
	LBR   = libratypes.Currency("LBR")
	Coin1 = libratypes.Currency("Coin1")
)
