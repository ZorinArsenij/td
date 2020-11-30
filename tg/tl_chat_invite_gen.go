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

// ChatInviteAlready represents TL type `chatInviteAlready#5a686d7c`.
type ChatInviteAlready struct {
	// Chat field of ChatInviteAlready.
	Chat ChatClass
}

// ChatInviteAlreadyTypeID is TL type id of ChatInviteAlready.
const ChatInviteAlreadyTypeID = 0x5a686d7c

// Encode implements bin.Encoder.
func (c *ChatInviteAlready) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteAlready#5a686d7c as nil")
	}
	b.PutID(ChatInviteAlreadyTypeID)
	if c.Chat == nil {
		return fmt.Errorf("unable to encode chatInviteAlready#5a686d7c: field chat is nil")
	}
	if err := c.Chat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteAlready#5a686d7c: field chat: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteAlready) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteAlready#5a686d7c to nil")
	}
	if err := b.ConsumeID(ChatInviteAlreadyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteAlready#5a686d7c: %w", err)
	}
	{
		value, err := DecodeChat(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteAlready#5a686d7c: field chat: %w", err)
		}
		c.Chat = value
	}
	return nil
}

// construct implements constructor of ChatInviteClass.
func (c ChatInviteAlready) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInviteAlready.
var (
	_ bin.Encoder = &ChatInviteAlready{}
	_ bin.Decoder = &ChatInviteAlready{}

	_ ChatInviteClass = &ChatInviteAlready{}
)

// ChatInvite represents TL type `chatInvite#dfc2f58e`.
type ChatInvite struct {
	// Flags field of ChatInvite.
	Flags bin.Fields
	// Channel field of ChatInvite.
	Channel bool
	// Broadcast field of ChatInvite.
	Broadcast bool
	// Public field of ChatInvite.
	Public bool
	// Megagroup field of ChatInvite.
	Megagroup bool
	// Title field of ChatInvite.
	Title string
	// Photo field of ChatInvite.
	Photo PhotoClass
	// ParticipantsCount field of ChatInvite.
	ParticipantsCount int
	// Participants field of ChatInvite.
	//
	// Use SetParticipants and GetParticipants helpers.
	Participants []UserClass
}

// ChatInviteTypeID is TL type id of ChatInvite.
const ChatInviteTypeID = 0xdfc2f58e

// Encode implements bin.Encoder.
func (c *ChatInvite) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvite#dfc2f58e as nil")
	}
	b.PutID(ChatInviteTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvite#dfc2f58e: field flags: %w", err)
	}
	b.PutString(c.Title)
	if c.Photo == nil {
		return fmt.Errorf("unable to encode chatInvite#dfc2f58e: field photo is nil")
	}
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvite#dfc2f58e: field photo: %w", err)
	}
	b.PutInt(c.ParticipantsCount)
	if c.Flags.Has(4) {
		b.PutVectorHeader(len(c.Participants))
		for idx, v := range c.Participants {
			if v == nil {
				return fmt.Errorf("unable to encode chatInvite#dfc2f58e: field participants element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode chatInvite#dfc2f58e: field participants element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetChannel sets value of Channel conditional field.
func (c *ChatInvite) SetChannel(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChatInvite) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// SetPublic sets value of Public conditional field.
func (c *ChatInvite) SetPublic(value bool) {
	if value {
		c.Flags.Set(2)
	} else {
		c.Flags.Unset(2)
	}
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChatInvite) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(3)
	} else {
		c.Flags.Unset(3)
	}
}

// SetParticipants sets value of Participants conditional field.
func (c *ChatInvite) SetParticipants(value []UserClass) {
	c.Flags.Set(4)
	c.Participants = value
}

// GetParticipants returns value of Participants conditional field and
// boolean which is true if field was set.
func (c *ChatInvite) GetParticipants() (value []UserClass, ok bool) {
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.Participants, true
}

// Decode implements bin.Decoder.
func (c *ChatInvite) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvite#dfc2f58e to nil")
	}
	if err := b.ConsumeID(ChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInvite#dfc2f58e: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field flags: %w", err)
		}
	}
	c.Channel = c.Flags.Has(0)
	c.Broadcast = c.Flags.Has(1)
	c.Public = c.Flags.Has(2)
	c.Megagroup = c.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field photo: %w", err)
		}
		c.Photo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field participants_count: %w", err)
		}
		c.ParticipantsCount = value
	}
	if c.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field participants: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatInvite#dfc2f58e: field participants: %w", err)
			}
			c.Participants = append(c.Participants, value)
		}
	}
	return nil
}

// construct implements constructor of ChatInviteClass.
func (c ChatInvite) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInvite.
var (
	_ bin.Encoder = &ChatInvite{}
	_ bin.Decoder = &ChatInvite{}

	_ ChatInviteClass = &ChatInvite{}
)

// ChatInvitePeek represents TL type `chatInvitePeek#61695cb0`.
type ChatInvitePeek struct {
	// Chat field of ChatInvitePeek.
	Chat ChatClass
	// Expires field of ChatInvitePeek.
	Expires int
}

// ChatInvitePeekTypeID is TL type id of ChatInvitePeek.
const ChatInvitePeekTypeID = 0x61695cb0

// Encode implements bin.Encoder.
func (c *ChatInvitePeek) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvitePeek#61695cb0 as nil")
	}
	b.PutID(ChatInvitePeekTypeID)
	if c.Chat == nil {
		return fmt.Errorf("unable to encode chatInvitePeek#61695cb0: field chat is nil")
	}
	if err := c.Chat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvitePeek#61695cb0: field chat: %w", err)
	}
	b.PutInt(c.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInvitePeek) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvitePeek#61695cb0 to nil")
	}
	if err := b.ConsumeID(ChatInvitePeekTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: %w", err)
	}
	{
		value, err := DecodeChat(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: field chat: %w", err)
		}
		c.Chat = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: field expires: %w", err)
		}
		c.Expires = value
	}
	return nil
}

// construct implements constructor of ChatInviteClass.
func (c ChatInvitePeek) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInvitePeek.
var (
	_ bin.Encoder = &ChatInvitePeek{}
	_ bin.Decoder = &ChatInvitePeek{}

	_ ChatInviteClass = &ChatInvitePeek{}
)

// ChatInviteClass represents ChatInvite generic type.
//
// Example:
//  g, err := DecodeChatInvite(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatInviteAlready: // chatInviteAlready#5a686d7c
//  case *ChatInvite: // chatInvite#dfc2f58e
//  case *ChatInvitePeek: // chatInvitePeek#61695cb0
//  default: panic(v)
//  }
type ChatInviteClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatInviteClass
}

// DecodeChatInvite implements binary de-serialization for ChatInviteClass.
func DecodeChatInvite(buf *bin.Buffer) (ChatInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatInviteAlreadyTypeID:
		// Decoding chatInviteAlready#5a686d7c.
		v := ChatInviteAlready{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInviteTypeID:
		// Decoding chatInvite#dfc2f58e.
		v := ChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInvitePeekTypeID:
		// Decoding chatInvitePeek#61695cb0.
		v := ChatInvitePeek{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatInvite boxes the ChatInviteClass providing a helper.
type ChatInviteBox struct {
	ChatInvite ChatInviteClass
}

// Decode implements bin.Decoder for ChatInviteBox.
func (b *ChatInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatInviteBox to nil")
	}
	v, err := DecodeChatInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatInvite = v
	return nil
}

// Encode implements bin.Encode for ChatInviteBox.
func (b *ChatInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatInvite == nil {
		return fmt.Errorf("unable to encode ChatInviteClass as nil")
	}
	return b.ChatInvite.Encode(buf)
}
