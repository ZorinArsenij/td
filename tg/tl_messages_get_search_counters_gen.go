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

// MessagesGetSearchCountersRequest represents TL type `messages.getSearchCounters#732eef00`.
type MessagesGetSearchCountersRequest struct {
	// Peer field of MessagesGetSearchCountersRequest.
	Peer InputPeerClass
	// Filters field of MessagesGetSearchCountersRequest.
	Filters []MessagesFilterClass
}

// MessagesGetSearchCountersRequestTypeID is TL type id of MessagesGetSearchCountersRequest.
const MessagesGetSearchCountersRequestTypeID = 0x732eef00

// Encode implements bin.Encoder.
func (g *MessagesGetSearchCountersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchCounters#732eef00 as nil")
	}
	b.PutID(MessagesGetSearchCountersRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.Filters))
	for idx, v := range g.Filters {
		if v == nil {
			return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field filters element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field filters element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSearchCountersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchCounters#732eef00 to nil")
	}
	if err := b.ConsumeID(MessagesGetSearchCountersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field filters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field filters: %w", err)
			}
			g.Filters = append(g.Filters, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetSearchCountersRequest.
var (
	_ bin.Encoder = &MessagesGetSearchCountersRequest{}
	_ bin.Decoder = &MessagesGetSearchCountersRequest{}
)
