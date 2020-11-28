// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesGetDocumentByHashRequest represents TL type `messages.getDocumentByHash#338e2464`.
type MessagesGetDocumentByHashRequest struct {
	// Sha256 field of MessagesGetDocumentByHashRequest.
	Sha256 []byte
	// Size field of MessagesGetDocumentByHashRequest.
	Size int
	// MimeType field of MessagesGetDocumentByHashRequest.
	MimeType string
}

// MessagesGetDocumentByHashRequestTypeID is TL type id of MessagesGetDocumentByHashRequest.
const MessagesGetDocumentByHashRequestTypeID = 0x338e2464

// Encode implements bin.Encoder.
func (g *MessagesGetDocumentByHashRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDocumentByHash#338e2464 as nil")
	}
	b.PutID(MessagesGetDocumentByHashRequestTypeID)
	b.PutBytes(g.Sha256)
	b.PutInt(g.Size)
	b.PutString(g.MimeType)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDocumentByHashRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDocumentByHash#338e2464 to nil")
	}
	if err := b.ConsumeID(MessagesGetDocumentByHashRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field sha256: %w", err)
		}
		g.Sha256 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field size: %w", err)
		}
		g.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field mime_type: %w", err)
		}
		g.MimeType = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDocumentByHashRequest.
var (
	_ bin.Encoder = &MessagesGetDocumentByHashRequest{}
	_ bin.Decoder = &MessagesGetDocumentByHashRequest{}
)