// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package librastd

import (
	"github.com/libra/libra-client-sdk-go/libratypes"
)

// CurrencyCode converts given currency code string into Move TypeTag that is required by
// move script argument.
func CurrencyCode(code string) libratypes.TypeTag {
	return &libratypes.TypeTag__Struct{
		Value: libratypes.StructTag{
			Address: libratypes.AccountAddress{
				Value: codeAddress,
			},
			Module:     libratypes.Identifier{Value: code},
			Name:       libratypes.Identifier{Value: code},
			TypeParams: []libratypes.TypeTag{},
		},
	}
}
