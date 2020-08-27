// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package stdlib

import (
	"github.com/libra/libra-client-sdk-go/libratypes"
)

var codeAddress = [16]uint8{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
}

// CurrencyCode converts given currency code string into Move TypeTag that is required by
// move script argument.
func CurrencyCode(code string) libratypes.TypeTag {
	return &libratypes.TypeTag__Struct{
		Value: libratypes.StructTag{
			Address:    codeAddress,
			Module:     libratypes.Identifier(code),
			Name:       libratypes.Identifier(code),
			TypeParams: []libratypes.TypeTag{},
		},
	}
}
