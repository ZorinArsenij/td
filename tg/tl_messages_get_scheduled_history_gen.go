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

// MessagesGetScheduledHistoryRequest represents TL type `messages.getScheduledHistory#e2c2685b`.
type MessagesGetScheduledHistoryRequest struct {
	// Peer field of MessagesGetScheduledHistoryRequest.
	Peer InputPeerClass
	// Hash field of MessagesGetScheduledHistoryRequest.
	Hash int
}

// MessagesGetScheduledHistoryRequestTypeID is TL type id of MessagesGetScheduledHistoryRequest.
const MessagesGetScheduledHistoryRequestTypeID = 0xe2c2685b

// Encode implements bin.Encoder.
func (g *MessagesGetScheduledHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledHistory#e2c2685b as nil")
	}
	b.PutID(MessagesGetScheduledHistoryRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#e2c2685b: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#e2c2685b: field peer: %w", err)
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetScheduledHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledHistory#e2c2685b to nil")
	}
	if err := b.ConsumeID(MessagesGetScheduledHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetScheduledHistoryRequest.
var (
	_ bin.Encoder = &MessagesGetScheduledHistoryRequest{}
	_ bin.Decoder = &MessagesGetScheduledHistoryRequest{}
)
