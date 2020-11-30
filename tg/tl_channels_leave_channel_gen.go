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

// ChannelsLeaveChannelRequest represents TL type `channels.leaveChannel#f836aa95`.
type ChannelsLeaveChannelRequest struct {
	// Channel field of ChannelsLeaveChannelRequest.
	Channel InputChannelClass
}

// ChannelsLeaveChannelRequestTypeID is TL type id of ChannelsLeaveChannelRequest.
const ChannelsLeaveChannelRequestTypeID = 0xf836aa95

// Encode implements bin.Encoder.
func (l *ChannelsLeaveChannelRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode channels.leaveChannel#f836aa95 as nil")
	}
	b.PutID(ChannelsLeaveChannelRequestTypeID)
	if l.Channel == nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel is nil")
	}
	if err := l.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *ChannelsLeaveChannelRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode channels.leaveChannel#f836aa95 to nil")
	}
	if err := b.ConsumeID(ChannelsLeaveChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: field channel: %w", err)
		}
		l.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsLeaveChannelRequest.
var (
	_ bin.Encoder = &ChannelsLeaveChannelRequest{}
	_ bin.Decoder = &ChannelsLeaveChannelRequest{}
)
