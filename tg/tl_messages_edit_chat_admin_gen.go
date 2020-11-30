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

// MessagesEditChatAdminRequest represents TL type `messages.editChatAdmin#a9e69f2e`.
type MessagesEditChatAdminRequest struct {
	// ChatID field of MessagesEditChatAdminRequest.
	ChatID int
	// UserID field of MessagesEditChatAdminRequest.
	UserID InputUserClass
	// IsAdmin field of MessagesEditChatAdminRequest.
	IsAdmin bool
}

// MessagesEditChatAdminRequestTypeID is TL type id of MessagesEditChatAdminRequest.
const MessagesEditChatAdminRequestTypeID = 0xa9e69f2e

// Encode implements bin.Encoder.
func (e *MessagesEditChatAdminRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAdmin#a9e69f2e as nil")
	}
	b.PutID(MessagesEditChatAdminRequestTypeID)
	b.PutInt(e.ChatID)
	if e.UserID == nil {
		return fmt.Errorf("unable to encode messages.editChatAdmin#a9e69f2e: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatAdmin#a9e69f2e: field user_id: %w", err)
	}
	b.PutBool(e.IsAdmin)
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatAdminRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAdmin#a9e69f2e to nil")
	}
	if err := b.ConsumeID(MessagesEditChatAdminRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field is_admin: %w", err)
		}
		e.IsAdmin = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditChatAdminRequest.
var (
	_ bin.Encoder = &MessagesEditChatAdminRequest{}
	_ bin.Decoder = &MessagesEditChatAdminRequest{}
)
