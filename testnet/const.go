// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/libratypes"
)

const (
	URL                  = "https://testnet.libra.org/v1"
	FaucetURL            = "https://testnet.libra.org/mint"
	ChainID         byte = 2
	DDAcountAddress      = "000000000000000000000000000000DD"
)

// Client is testnet client
var Client = libraclient.New(ChainID, URL)

// Currencies
var (
	LBR   = libratypes.Currency("LBR")
	Coin1 = libratypes.Currency("Coin1")
	Coin2 = libratypes.Currency("Coin2")
)
