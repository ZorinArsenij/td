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

// MessagesEditChatAboutRequest represents TL type `messages.editChatAbout#def60797`.
type MessagesEditChatAboutRequest struct {
	// Peer field of MessagesEditChatAboutRequest.
	Peer InputPeerClass
	// About field of MessagesEditChatAboutRequest.
	About string
}

// MessagesEditChatAboutRequestTypeID is TL type id of MessagesEditChatAboutRequest.
const MessagesEditChatAboutRequestTypeID = 0xdef60797

// Encode implements bin.Encoder.
func (e *MessagesEditChatAboutRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAbout#def60797 as nil")
	}
	b.PutID(MessagesEditChatAboutRequestTypeID)
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer: %w", err)
	}
	b.PutString(e.About)
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatAboutRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAbout#def60797 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatAboutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatAbout#def60797: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field about: %w", err)
		}
		e.About = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditChatAboutRequest.
var (
	_ bin.Encoder = &MessagesEditChatAboutRequest{}
	_ bin.Decoder = &MessagesEditChatAboutRequest{}
)
