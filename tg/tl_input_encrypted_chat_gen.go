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

// InputEncryptedChat represents TL type `inputEncryptedChat#f141b5e1`.
type InputEncryptedChat struct {
	// ChatID field of InputEncryptedChat.
	ChatID int
	// AccessHash field of InputEncryptedChat.
	AccessHash int64
}

// InputEncryptedChatTypeID is TL type id of InputEncryptedChat.
const InputEncryptedChatTypeID = 0xf141b5e1

// Encode implements bin.Encoder.
func (i *InputEncryptedChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedChat#f141b5e1 as nil")
	}
	b.PutID(InputEncryptedChatTypeID)
	b.PutInt(i.ChatID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedChat#f141b5e1 to nil")
	}
	if err := b.ConsumeID(InputEncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputEncryptedChat.
var (
	_ bin.Encoder = &InputEncryptedChat{}
	_ bin.Decoder = &InputEncryptedChat{}
)
