// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesAffectedMessages represents TL type `messages.affectedMessages#84d19185`.
type MessagesAffectedMessages struct {
	// Pts field of MessagesAffectedMessages.
	Pts int
	// PtsCount field of MessagesAffectedMessages.
	PtsCount int
}

// MessagesAffectedMessagesTypeID is TL type id of MessagesAffectedMessages.
const MessagesAffectedMessagesTypeID = 0x84d19185

// Encode implements bin.Encoder.
func (a *MessagesAffectedMessages) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedMessages#84d19185 as nil")
	}
	b.PutID(MessagesAffectedMessagesTypeID)
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedMessages) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedMessages#84d19185 to nil")
	}
	if err := b.ConsumeID(MessagesAffectedMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAffectedMessages.
var (
	_ bin.Encoder = &MessagesAffectedMessages{}
	_ bin.Decoder = &MessagesAffectedMessages{}
)
