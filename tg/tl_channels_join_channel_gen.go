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

// ChannelsJoinChannelRequest represents TL type `channels.joinChannel#24b524c5`.
type ChannelsJoinChannelRequest struct {
	// Channel field of ChannelsJoinChannelRequest.
	Channel InputChannelClass
}

// ChannelsJoinChannelRequestTypeID is TL type id of ChannelsJoinChannelRequest.
const ChannelsJoinChannelRequestTypeID = 0x24b524c5

// Encode implements bin.Encoder.
func (j *ChannelsJoinChannelRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode channels.joinChannel#24b524c5 as nil")
	}
	b.PutID(ChannelsJoinChannelRequestTypeID)
	if j.Channel == nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel is nil")
	}
	if err := j.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *ChannelsJoinChannelRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode channels.joinChannel#24b524c5 to nil")
	}
	if err := b.ConsumeID(ChannelsJoinChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: field channel: %w", err)
		}
		j.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsJoinChannelRequest.
var (
	_ bin.Encoder = &ChannelsJoinChannelRequest{}
	_ bin.Decoder = &ChannelsJoinChannelRequest{}
)
