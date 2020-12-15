// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemtypes_test

import (
	"errors"
	"testing"

	"github.com/diem/client-sdk-go/diemtypes"
	"github.com/stretchr/testify/assert"
)

func TestToBCS(t *testing.T) {
	address, _ := diemtypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	bytes := diemtypes.ToBCS(&address)
	assert.NotEmpty(t, bytes)
}

func TestToBCSPanicsForSerializationError(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	diemtypes.ToBCS(new(bcsError))
}

type bcsError struct {
}

func (l *bcsError) BcsSerialize() ([]byte, error) {
	return nil, errors.New("unexpected")
}
