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

// InputUserEmpty represents TL type `inputUserEmpty#b98886cf`.
type InputUserEmpty struct {
}

// InputUserEmptyTypeID is TL type id of InputUserEmpty.
const InputUserEmptyTypeID = 0xb98886cf

// Encode implements bin.Encoder.
func (i *InputUserEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserEmpty#b98886cf as nil")
	}
	b.PutID(InputUserEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserEmpty#b98886cf to nil")
	}
	if err := b.ConsumeID(InputUserEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserEmpty#b98886cf: %w", err)
	}
	return nil
}

// construct implements constructor of InputUserClass.
func (i InputUserEmpty) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserEmpty.
var (
	_ bin.Encoder = &InputUserEmpty{}
	_ bin.Decoder = &InputUserEmpty{}

	_ InputUserClass = &InputUserEmpty{}
)

// InputUserSelf represents TL type `inputUserSelf#f7c1b13f`.
type InputUserSelf struct {
}

// InputUserSelfTypeID is TL type id of InputUserSelf.
const InputUserSelfTypeID = 0xf7c1b13f

// Encode implements bin.Encoder.
func (i *InputUserSelf) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserSelf#f7c1b13f as nil")
	}
	b.PutID(InputUserSelfTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserSelf) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserSelf#f7c1b13f to nil")
	}
	if err := b.ConsumeID(InputUserSelfTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserSelf#f7c1b13f: %w", err)
	}
	return nil
}

// construct implements constructor of InputUserClass.
func (i InputUserSelf) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserSelf.
var (
	_ bin.Encoder = &InputUserSelf{}
	_ bin.Decoder = &InputUserSelf{}

	_ InputUserClass = &InputUserSelf{}
)

// InputUser represents TL type `inputUser#d8292816`.
type InputUser struct {
	// UserID field of InputUser.
	UserID int
	// AccessHash field of InputUser.
	AccessHash int64
}

// InputUserTypeID is TL type id of InputUser.
const InputUserTypeID = 0xd8292816

// Encode implements bin.Encoder.
func (i *InputUser) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUser#d8292816 as nil")
	}
	b.PutID(InputUserTypeID)
	b.PutInt(i.UserID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUser) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUser#d8292816 to nil")
	}
	if err := b.ConsumeID(InputUserTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUser#d8292816: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputUser#d8292816: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputUser#d8292816: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputUserClass.
func (i InputUser) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUser.
var (
	_ bin.Encoder = &InputUser{}
	_ bin.Decoder = &InputUser{}

	_ InputUserClass = &InputUser{}
)

// InputUserFromMessage represents TL type `inputUserFromMessage#2d117597`.
type InputUserFromMessage struct {
	// Peer field of InputUserFromMessage.
	Peer InputPeerClass
	// MsgID field of InputUserFromMessage.
	MsgID int
	// UserID field of InputUserFromMessage.
	UserID int
}

// InputUserFromMessageTypeID is TL type id of InputUserFromMessage.
const InputUserFromMessageTypeID = 0x2d117597

// Encode implements bin.Encoder.
func (i *InputUserFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserFromMessage#2d117597 as nil")
	}
	b.PutID(InputUserFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputUserFromMessage#2d117597: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputUserFromMessage#2d117597: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserFromMessage#2d117597 to nil")
	}
	if err := b.ConsumeID(InputUserFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserFromMessage#2d117597: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#2d117597: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#2d117597: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#2d117597: field user_id: %w", err)
		}
		i.UserID = value
	}
	return nil
}

// construct implements constructor of InputUserClass.
func (i InputUserFromMessage) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserFromMessage.
var (
	_ bin.Encoder = &InputUserFromMessage{}
	_ bin.Decoder = &InputUserFromMessage{}

	_ InputUserClass = &InputUserFromMessage{}
)

// InputUserClass represents InputUser generic type.
//
// Example:
//  g, err := DecodeInputUser(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputUserEmpty: // inputUserEmpty#b98886cf
//  case *InputUserSelf: // inputUserSelf#f7c1b13f
//  case *InputUser: // inputUser#d8292816
//  case *InputUserFromMessage: // inputUserFromMessage#2d117597
//  default: panic(v)
//  }
type InputUserClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputUserClass
}

// DecodeInputUser implements binary de-serialization for InputUserClass.
func DecodeInputUser(buf *bin.Buffer) (InputUserClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputUserEmptyTypeID:
		// Decoding inputUserEmpty#b98886cf.
		v := InputUserEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserSelfTypeID:
		// Decoding inputUserSelf#f7c1b13f.
		v := InputUserSelf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserTypeID:
		// Decoding inputUser#d8292816.
		v := InputUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserFromMessageTypeID:
		// Decoding inputUserFromMessage#2d117597.
		v := InputUserFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputUserClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputUser boxes the InputUserClass providing a helper.
type InputUserBox struct {
	InputUser InputUserClass
}

// Decode implements bin.Decoder for InputUserBox.
func (b *InputUserBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputUserBox to nil")
	}
	v, err := DecodeInputUser(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputUser = v
	return nil
}

// Encode implements bin.Encode for InputUserBox.
func (b *InputUserBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputUser == nil {
		return fmt.Errorf("unable to encode InputUserClass as nil")
	}
	return b.InputUser.Encode(buf)
}
