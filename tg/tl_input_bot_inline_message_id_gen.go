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

// InputBotInlineMessageID represents TL type `inputBotInlineMessageID#890c3d89`.
type InputBotInlineMessageID struct {
	// DCID field of InputBotInlineMessageID.
	DCID int
	// ID field of InputBotInlineMessageID.
	ID int64
	// AccessHash field of InputBotInlineMessageID.
	AccessHash int64
}

// InputBotInlineMessageIDTypeID is TL type id of InputBotInlineMessageID.
const InputBotInlineMessageIDTypeID = 0x890c3d89

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageID#890c3d89 as nil")
	}
	b.PutID(InputBotInlineMessageIDTypeID)
	b.PutInt(i.DCID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageID#890c3d89 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field dc_id: %w", err)
		}
		i.DCID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageID#890c3d89: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputBotInlineMessageID.
var (
	_ bin.Encoder = &InputBotInlineMessageID{}
	_ bin.Decoder = &InputBotInlineMessageID{}
)
