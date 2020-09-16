// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclienttest

import "github.com/libra/libra-client-sdk-go/libraclient"

type eventBuilderPart func(*libraclient.Event)
type EventBuilder struct {
	parts []eventBuilderPart
}

func (b EventBuilder) Build() *libraclient.Event {
	t := new(libraclient.Event)
	for _, part := range b.parts {
		part(t)
	}
	return t
}

func (b EventBuilder) Type(t string) *EventBuilder {
	return b.append(func(e *libraclient.Event) {
		if e.Data == nil {
			e.Data = new(libraclient.EventData)
		}
		e.Data.Type = t
	})
}

func (b EventBuilder) SequenceNumber(n uint64) *EventBuilder {
	return b.append(func(e *libraclient.Event) {
		e.SequenceNumber = n
	})
}

func (b EventBuilder) Receiver(receiver string) *EventBuilder {
	return b.append(func(e *libraclient.Event) {
		if e.Data == nil {
			e.Data = new(libraclient.EventData)
		}
		e.Data.Receiver = receiver
	})
}

func (b EventBuilder) Metadata(metadata string) *EventBuilder {
	return b.append(func(e *libraclient.Event) {
		if e.Data == nil {
			e.Data = new(libraclient.EventData)
		}
		e.Data.Metadata = metadata
	})
}

func (b *EventBuilder) append(parts ...eventBuilderPart) *EventBuilder {
	newParts := make([]eventBuilderPart, len(b.parts)+len(parts))
	copy(newParts, b.parts)
	copy(newParts[len(b.parts):], parts)
	b.parts = newParts
	return b
}
