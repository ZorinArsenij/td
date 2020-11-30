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

// ContactsContactsNotModified represents TL type `contacts.contactsNotModified#b74ba9d2`.
type ContactsContactsNotModified struct {
}

// ContactsContactsNotModifiedTypeID is TL type id of ContactsContactsNotModified.
const ContactsContactsNotModifiedTypeID = 0xb74ba9d2

// Encode implements bin.Encoder.
func (c *ContactsContactsNotModified) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contactsNotModified#b74ba9d2 as nil")
	}
	b.PutID(ContactsContactsNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactsContactsNotModified) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contactsNotModified#b74ba9d2 to nil")
	}
	if err := b.ConsumeID(ContactsContactsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.contactsNotModified#b74ba9d2: %w", err)
	}
	return nil
}

// construct implements constructor of ContactsContactsClass.
func (c ContactsContactsNotModified) construct() ContactsContactsClass { return &c }

// Ensuring interfaces in compile-time for ContactsContactsNotModified.
var (
	_ bin.Encoder = &ContactsContactsNotModified{}
	_ bin.Decoder = &ContactsContactsNotModified{}

	_ ContactsContactsClass = &ContactsContactsNotModified{}
)

// ContactsContacts represents TL type `contacts.contacts#eae87e42`.
type ContactsContacts struct {
	// Contacts field of ContactsContacts.
	Contacts []Contact
	// SavedCount field of ContactsContacts.
	SavedCount int
	// Users field of ContactsContacts.
	Users []UserClass
}

// ContactsContactsTypeID is TL type id of ContactsContacts.
const ContactsContactsTypeID = 0xeae87e42

// Encode implements bin.Encoder.
func (c *ContactsContacts) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contacts#eae87e42 as nil")
	}
	b.PutID(ContactsContactsTypeID)
	b.PutVectorHeader(len(c.Contacts))
	for idx, v := range c.Contacts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field contacts element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.SavedCount)
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactsContacts) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contacts#eae87e42 to nil")
	}
	if err := b.ConsumeID(ContactsContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.contacts#eae87e42: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field contacts: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Contact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field contacts: %w", err)
			}
			c.Contacts = append(c.Contacts, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field saved_count: %w", err)
		}
		c.SavedCount = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// construct implements constructor of ContactsContactsClass.
func (c ContactsContacts) construct() ContactsContactsClass { return &c }

// Ensuring interfaces in compile-time for ContactsContacts.
var (
	_ bin.Encoder = &ContactsContacts{}
	_ bin.Decoder = &ContactsContacts{}

	_ ContactsContactsClass = &ContactsContacts{}
)

// ContactsContactsClass represents contacts.Contacts generic type.
//
// Example:
//  g, err := DecodeContactsContacts(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ContactsContactsNotModified: // contacts.contactsNotModified#b74ba9d2
//  case *ContactsContacts: // contacts.contacts#eae87e42
//  default: panic(v)
//  }
type ContactsContactsClass interface {
	bin.Encoder
	bin.Decoder
	construct() ContactsContactsClass
}

// DecodeContactsContacts implements binary de-serialization for ContactsContactsClass.
func DecodeContactsContacts(buf *bin.Buffer) (ContactsContactsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsContactsNotModifiedTypeID:
		// Decoding contacts.contactsNotModified#b74ba9d2.
		v := ContactsContactsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", err)
		}
		return &v, nil
	case ContactsContactsTypeID:
		// Decoding contacts.contacts#eae87e42.
		v := ContactsContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsContacts boxes the ContactsContactsClass providing a helper.
type ContactsContactsBox struct {
	ContactsContacts ContactsContactsClass
}

// Decode implements bin.Decoder for ContactsContactsBox.
func (b *ContactsContactsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsContactsBox to nil")
	}
	v, err := DecodeContactsContacts(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ContactsContacts = v
	return nil
}

// Encode implements bin.Encode for ContactsContactsBox.
func (b *ContactsContactsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ContactsContacts == nil {
		return fmt.Errorf("unable to encode ContactsContactsClass as nil")
	}
	return b.ContactsContacts.Encode(buf)
}
