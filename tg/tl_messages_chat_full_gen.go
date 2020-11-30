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

// MessagesChatFull represents TL type `messages.chatFull#e5d7d19c`.
type MessagesChatFull struct {
	// FullChat field of MessagesChatFull.
	FullChat ChatFullClass
	// Chats field of MessagesChatFull.
	Chats []ChatClass
	// Users field of MessagesChatFull.
	Users []UserClass
}

// MessagesChatFullTypeID is TL type id of MessagesChatFull.
const MessagesChatFullTypeID = 0xe5d7d19c

// Encode implements bin.Encoder.
func (c *MessagesChatFull) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatFull#e5d7d19c as nil")
	}
	b.PutID(MessagesChatFullTypeID)
	if c.FullChat == nil {
		return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field full_chat is nil")
	}
	if err := c.FullChat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field full_chat: %w", err)
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesChatFull) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatFull#e5d7d19c to nil")
	}
	if err := b.ConsumeID(MessagesChatFullTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: %w", err)
	}
	{
		value, err := DecodeChatFull(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field full_chat: %w", err)
		}
		c.FullChat = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesChatFull.
var (
	_ bin.Encoder = &MessagesChatFull{}
	_ bin.Decoder = &MessagesChatFull{}
)
