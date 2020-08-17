package librastd

import "github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/lcs"

func serializeCurrencyCode(s *lcs.Serializer, currencyCode string) {
	// []TypeTag size
	s.SerializeLen(1)
	// TypeTag Struct
	s.SerializeVariantIndex(7)
	// address
	for _, b := range codeAddress {
		s.SerializeU8(b)
	}
	// module
	s.SerializeStr(currencyCode)
	s.SerializeStr(currencyCode)
	// type params
	s.SerializeLen(0)
}
