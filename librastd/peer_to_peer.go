package librastd

import (
	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"
)

// EncodePeerToPeerScriptWithMetadata
func EncodePeerToPeerScriptWithMetadata(address []byte, currencyCode string, amount uint64, metadata []byte, metadataSignature []byte) []byte {
	s := new(lcs.Serializer)
	// Script
	s.SerializeVariantIndex(1)
	// code
	s.SerializeBytes([]byte{
		161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 2, 2, 2, 4, 3, 6, 16, 4, 22, 2, 5, 24, 29, 7, 53, 97, 8,
		150, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 2, 3, 1, 1, 0, 4, 1, 3, 0, 1, 5, 1, 6, 12,
		1, 8, 0, 5, 6, 8, 0, 5, 3, 10, 2, 10, 2, 0, 5, 6, 12, 5, 3, 10, 2, 10, 2, 1, 9, 0, 12, 76, 105,
		98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97,
		112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116,
		104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 8, 112, 97, 121, 95,
		102, 114, 111, 109, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114,
		97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 4, 1, 12, 11, 0, 17, 0, 12, 5, 14, 5, 10, 1, 10, 2, 11, 3, 11, 4, 56, 0, 11,
		5, 17, 2, 2,
	})
	serializeCurrencyCode(s, currencyCode)
	// arguments
	s.SerializeLen(4)
	// address
	s.SerializeVariantIndex(3)
	for _, b := range address {
		s.SerializeU8(b)
	}
	// ammount
	s.SerializeVariantIndex(1)
	s.SerializeU64(amount)
	// metadata, u8 vector
	s.SerializeVariantIndex(4)
	s.SerializeBytes(metadata)
	// metadata sig, u8 vector
	s.SerializeVariantIndex(4)
	s.SerializeBytes(metadataSignature)

	return s.GetBytes()
}
