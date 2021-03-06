// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// ContactsAcceptContactRequest represents TL type `contacts.acceptContact#f831a20f`.
// If the peer settings¹ of a new user allow us to add him as contact, add that user as contact
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
type ContactsAcceptContactRequest struct {
	// The user to add as contact
	ID InputUserClass
}

// ContactsAcceptContactRequestTypeID is TL type id of ContactsAcceptContactRequest.
const ContactsAcceptContactRequestTypeID = 0xf831a20f

func (a *ContactsAcceptContactRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *ContactsAcceptContactRequest) String() string {
	if a == nil {
		return "ContactsAcceptContactRequest(nil)"
	}
	type Alias ContactsAcceptContactRequest
	return fmt.Sprintf("ContactsAcceptContactRequest%+v", Alias(*a))
}

// FillFrom fills ContactsAcceptContactRequest from given interface.
func (a *ContactsAcceptContactRequest) FillFrom(from interface {
	GetID() (value InputUserClass)
}) {
	a.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsAcceptContactRequest) TypeID() uint32 {
	return ContactsAcceptContactRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsAcceptContactRequest) TypeName() string {
	return "contacts.acceptContact"
}

// TypeInfo returns info about TL type.
func (a *ContactsAcceptContactRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.acceptContact",
		ID:   ContactsAcceptContactRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *ContactsAcceptContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.acceptContact#f831a20f as nil")
	}
	b.PutID(ContactsAcceptContactRequestTypeID)
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (a *ContactsAcceptContactRequest) GetID() (value InputUserClass) {
	return a.ID
}

// Decode implements bin.Decoder.
func (a *ContactsAcceptContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.acceptContact#f831a20f to nil")
	}
	if err := b.ConsumeID(ContactsAcceptContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: field id: %w", err)
		}
		a.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsAcceptContactRequest.
var (
	_ bin.Encoder = &ContactsAcceptContactRequest{}
	_ bin.Decoder = &ContactsAcceptContactRequest{}
)

// ContactsAcceptContact invokes method contacts.acceptContact#f831a20f returning error if any.
// If the peer settings¹ of a new user allow us to add him as contact, add that user as contact
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// Possible errors:
//  400 CONTACT_ADD_MISSING: Contact to add is missing
//  400 CONTACT_ID_INVALID: The provided contact ID is invalid
//  400 CONTACT_REQ_MISSING: Missing contact request
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
func (c *Client) ContactsAcceptContact(ctx context.Context, id InputUserClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ContactsAcceptContactRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
