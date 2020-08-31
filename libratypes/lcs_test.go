// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libratypes_test

import (
	"errors"
	"testing"

	"github.com/libra/libra-client-sdk-go/libratypes"
	"github.com/stretchr/testify/assert"
)

func TestToLCS(t *testing.T) {
	address, _ := libratypes.MakeAccountAddress("f72589b71ff4f8d139674a3f7369c69b")
	bytes := libratypes.ToLCS(&address)
	assert.NotEmpty(t, bytes)
}

func TestToLCSPanicsForSerializationError(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	libratypes.ToLCS(new(lcsError))
}

type lcsError struct {
}

func (l *lcsError) LcsSerialize() ([]byte, error) {
	return nil, errors.New("unexpected")
}
