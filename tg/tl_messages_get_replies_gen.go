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

// MessagesGetRepliesRequest represents TL type `messages.getReplies#24b581ba`.
type MessagesGetRepliesRequest struct {
	// Peer field of MessagesGetRepliesRequest.
	Peer InputPeerClass
	// MsgID field of MessagesGetRepliesRequest.
	MsgID int
	// OffsetID field of MessagesGetRepliesRequest.
	OffsetID int
	// OffsetDate field of MessagesGetRepliesRequest.
	OffsetDate int
	// AddOffset field of MessagesGetRepliesRequest.
	AddOffset int
	// Limit field of MessagesGetRepliesRequest.
	Limit int
	// MaxID field of MessagesGetRepliesRequest.
	MaxID int
	// MinID field of MessagesGetRepliesRequest.
	MinID int
	// Hash field of MessagesGetRepliesRequest.
	Hash int
}

// MessagesGetRepliesRequestTypeID is TL type id of MessagesGetRepliesRequest.
const MessagesGetRepliesRequestTypeID = 0x24b581ba

// Encode implements bin.Encoder.
func (g *MessagesGetRepliesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getReplies#24b581ba as nil")
	}
	b.PutID(MessagesGetRepliesRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getReplies#24b581ba: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getReplies#24b581ba: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetRepliesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getReplies#24b581ba to nil")
	}
	if err := b.ConsumeID(MessagesGetRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getReplies#24b581ba: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetRepliesRequest.
var (
	_ bin.Encoder = &MessagesGetRepliesRequest{}
	_ bin.Decoder = &MessagesGetRepliesRequest{}
)
