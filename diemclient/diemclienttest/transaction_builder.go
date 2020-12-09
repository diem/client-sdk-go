// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemclienttest

import "github.com/diem/client-sdk-go/diemclient"

type transactionBuilderPart func(*diemclient.Transaction)
type TransactionBuilder struct {
	parts []transactionBuilderPart
}

func (b *TransactionBuilder) Build() *diemclient.Transaction {
	t := new(diemclient.Transaction)
	for _, part := range b.parts {
		part(t)
	}
	return t
}

func (b TransactionBuilder) Events(events ...*EventBuilder) *TransactionBuilder {
	return b.append(func(t *diemclient.Transaction) {
		for _, event := range events {
			t.Events = append(t.Events, event.Build())
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
