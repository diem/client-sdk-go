// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemtypes"
)

const (
	URL            = "http://dev.testnet.diem.com/v1"
	FaucetURL      = "http://dev.testnet.diem.com/mint"
	ChainID   byte = 3 // temp to 3 devnet before we upgrade testnet to new version
)

var (
	// DDAccountAddress is testnet default dd account address
	DDAccountAddress = diemtypes.MustMakeAccountAddress("000000000000000000000000000000DD")
	// Client is testnet client
	Client = diemclient.New(ChainID, URL)
)

// Currencies
var (
	XDX = diemtypes.Currency("XDX")
	XUS = diemtypes.Currency("XUS")
)
