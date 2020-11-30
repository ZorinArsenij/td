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

// ContactsBlocked represents TL type `contacts.blocked#ade1591`.
type ContactsBlocked struct {
	// Blocked field of ContactsBlocked.
	Blocked []PeerBlocked
	// Chats field of ContactsBlocked.
	Chats []ChatClass
	// Users field of ContactsBlocked.
	Users []UserClass
}

// ContactsBlockedTypeID is TL type id of ContactsBlocked.
const ContactsBlockedTypeID = 0xade1591

// Encode implements bin.Encoder.
func (b *ContactsBlocked) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.blocked#ade1591 as nil")
	}
	buf.PutID(ContactsBlockedTypeID)
	buf.PutVectorHeader(len(b.Blocked))
	for idx, v := range b.Blocked {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blocked#ade1591: field blocked element with index %d: %w", idx, err)
		}
	}
	buf.PutVectorHeader(len(b.Chats))
	for idx, v := range b.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.blocked#ade1591: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blocked#ade1591: field chats element with index %d: %w", idx, err)
		}
	}
	buf.PutVectorHeader(len(b.Users))
	for idx, v := range b.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.blocked#ade1591: field users element with index %d is nil", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blocked#ade1591: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *ContactsBlocked) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.blocked#ade1591 to nil")
	}
	if err := buf.ConsumeID(ContactsBlockedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.blocked#ade1591: %w", err)
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blocked#ade1591: field blocked: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PeerBlocked
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode contacts.blocked#ade1591: field blocked: %w", err)
			}
			b.Blocked = append(b.Blocked, value)
		}
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blocked#ade1591: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(buf)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.blocked#ade1591: field chats: %w", err)
			}
			b.Chats = append(b.Chats, value)
		}
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blocked#ade1591: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(buf)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.blocked#ade1591: field users: %w", err)
			}
			b.Users = append(b.Users, value)
		}
	}
	return nil
}

// construct implements constructor of ContactsBlockedClass.
func (b ContactsBlocked) construct() ContactsBlockedClass { return &b }

// Ensuring interfaces in compile-time for ContactsBlocked.
var (
	_ bin.Encoder = &ContactsBlocked{}
	_ bin.Decoder = &ContactsBlocked{}

	_ ContactsBlockedClass = &ContactsBlocked{}
)

// ContactsBlockedSlice represents TL type `contacts.blockedSlice#e1664194`.
type ContactsBlockedSlice struct {
	// Count field of ContactsBlockedSlice.
	Count int
	// Blocked field of ContactsBlockedSlice.
	Blocked []PeerBlocked
	// Chats field of ContactsBlockedSlice.
	Chats []ChatClass
	// Users field of ContactsBlockedSlice.
	Users []UserClass
}

// ContactsBlockedSliceTypeID is TL type id of ContactsBlockedSlice.
const ContactsBlockedSliceTypeID = 0xe1664194

// Encode implements bin.Encoder.
func (b *ContactsBlockedSlice) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.blockedSlice#e1664194 as nil")
	}
	buf.PutID(ContactsBlockedSliceTypeID)
	buf.PutInt(b.Count)
	buf.PutVectorHeader(len(b.Blocked))
	for idx, v := range b.Blocked {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blockedSlice#e1664194: field blocked element with index %d: %w", idx, err)
		}
	}
	buf.PutVectorHeader(len(b.Chats))
	for idx, v := range b.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.blockedSlice#e1664194: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blockedSlice#e1664194: field chats element with index %d: %w", idx, err)
		}
	}
	buf.PutVectorHeader(len(b.Users))
	for idx, v := range b.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.blockedSlice#e1664194: field users element with index %d is nil", idx)
		}
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode contacts.blockedSlice#e1664194: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *ContactsBlockedSlice) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.blockedSlice#e1664194 to nil")
	}
	if err := buf.ConsumeID(ContactsBlockedSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: %w", err)
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field count: %w", err)
		}
		b.Count = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field blocked: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PeerBlocked
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field blocked: %w", err)
			}
			b.Blocked = append(b.Blocked, value)
		}
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(buf)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field chats: %w", err)
			}
			b.Chats = append(b.Chats, value)
		}
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(buf)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.blockedSlice#e1664194: field users: %w", err)
			}
			b.Users = append(b.Users, value)
		}
	}
	return nil
}

// construct implements constructor of ContactsBlockedClass.
func (b ContactsBlockedSlice) construct() ContactsBlockedClass { return &b }

// Ensuring interfaces in compile-time for ContactsBlockedSlice.
var (
	_ bin.Encoder = &ContactsBlockedSlice{}
	_ bin.Decoder = &ContactsBlockedSlice{}

	_ ContactsBlockedClass = &ContactsBlockedSlice{}
)

// ContactsBlockedClass represents contacts.Blocked generic type.
//
// Example:
//  g, err := DecodeContactsBlocked(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ContactsBlocked: // contacts.blocked#ade1591
//  case *ContactsBlockedSlice: // contacts.blockedSlice#e1664194
//  default: panic(v)
//  }
type ContactsBlockedClass interface {
	bin.Encoder
	bin.Decoder
	construct() ContactsBlockedClass
}

// DecodeContactsBlocked implements binary de-serialization for ContactsBlockedClass.
func DecodeContactsBlocked(buf *bin.Buffer) (ContactsBlockedClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsBlockedTypeID:
		// Decoding contacts.blocked#ade1591.
		v := ContactsBlocked{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsBlockedClass: %w", err)
		}
		return &v, nil
	case ContactsBlockedSliceTypeID:
		// Decoding contacts.blockedSlice#e1664194.
		v := ContactsBlockedSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsBlockedClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsBlockedClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsBlocked boxes the ContactsBlockedClass providing a helper.
type ContactsBlockedBox struct {
	ContactsBlocked ContactsBlockedClass
}

// Decode implements bin.Decoder for ContactsBlockedBox.
func (b *ContactsBlockedBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsBlockedBox to nil")
	}
	v, err := DecodeContactsBlocked(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ContactsBlocked = v
	return nil
}

// Encode implements bin.Encode for ContactsBlockedBox.
func (b *ContactsBlockedBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ContactsBlocked == nil {
		return fmt.Errorf("unable to encode ContactsBlockedClass as nil")
	}
	return b.ContactsBlocked.Encode(buf)
}
