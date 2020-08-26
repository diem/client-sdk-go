// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"github.com/libra/libra-client-sdk-go/libraclient"
	"github.com/libra/libra-client-sdk-go/stdlib"
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
	LBR   = stdlib.CurrencyCode("LBR")
	Coin1 = stdlib.CurrencyCode("Coin1")
	Coin2 = stdlib.CurrencyCode("Coin2")
)
