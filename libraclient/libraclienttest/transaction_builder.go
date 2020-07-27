// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclienttest

import "github.com/libra/libra-client-sdk-go/libraclient"

type transactionBuilderPart func(*libraclient.Transaction)
type TransactionBuilder struct {
	parts []transactionBuilderPart
}

func (b *TransactionBuilder) Build() *libraclient.Transaction {
	t := new(libraclient.Transaction)
	for _, part := range b.parts {
		part(t)
	}
	return t
}

func (b TransactionBuilder) Events(events ...*EventBuilder) *TransactionBuilder {
	return b.append(func(t *libraclient.Transaction) {
		for _, event := range events {
			t.Events = append(t.Events, *event.Build())
		}
	})
}

func (b *TransactionBuilder) append(parts ...transactionBuilderPart) *TransactionBuilder {
	newParts := make([]transactionBuilderPart, len(b.parts)+len(parts))
	copy(newParts, b.parts)
	copy(newParts[len(b.parts):], parts)
	b.parts = newParts
	return b
}
