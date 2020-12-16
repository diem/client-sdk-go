// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemtypes"
)

const (
	URL            = "http://testnet.diem.com/v1"
	FaucetURL      = "http://testnet.diem.com/mint"
	ChainID   byte = 2
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
