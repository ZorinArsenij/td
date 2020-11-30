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

// ContactsResolveUsernameRequest represents TL type `contacts.resolveUsername#f93ccba3`.
type ContactsResolveUsernameRequest struct {
	// Username field of ContactsResolveUsernameRequest.
	Username string
}

// ContactsResolveUsernameRequestTypeID is TL type id of ContactsResolveUsernameRequest.
const ContactsResolveUsernameRequestTypeID = 0xf93ccba3

// Encode implements bin.Encoder.
func (r *ContactsResolveUsernameRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolveUsername#f93ccba3 as nil")
	}
	b.PutID(ContactsResolveUsernameRequestTypeID)
	b.PutString(r.Username)
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResolveUsernameRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolveUsername#f93ccba3 to nil")
	}
	if err := b.ConsumeID(ContactsResolveUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resolveUsername#f93ccba3: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolveUsername#f93ccba3: field username: %w", err)
		}
		r.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsResolveUsernameRequest.
var (
	_ bin.Encoder = &ContactsResolveUsernameRequest{}
	_ bin.Decoder = &ContactsResolveUsernameRequest{}
)
