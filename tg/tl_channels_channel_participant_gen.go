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

// ChannelsChannelParticipant represents TL type `channels.channelParticipant#d0d9b163`.
type ChannelsChannelParticipant struct {
	// Participant field of ChannelsChannelParticipant.
	Participant ChannelParticipantClass
	// Users field of ChannelsChannelParticipant.
	Users []UserClass
}

// ChannelsChannelParticipantTypeID is TL type id of ChannelsChannelParticipant.
const ChannelsChannelParticipantTypeID = 0xd0d9b163

// Encode implements bin.Encoder.
func (c *ChannelsChannelParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.channelParticipant#d0d9b163 as nil")
	}
	b.PutID(ChannelsChannelParticipantTypeID)
	if c.Participant == nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field participant is nil")
	}
	if err := c.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field participant: %w", err)
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#d0d9b163: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsChannelParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.channelParticipant#d0d9b163 to nil")
	}
	if err := b.ConsumeID(ChannelsChannelParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: %w", err)
	}
	{
		value, err := DecodeChannelParticipant(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field participant: %w", err)
		}
		c.Participant = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.channelParticipant#d0d9b163: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsChannelParticipant.
var (
	_ bin.Encoder = &ChannelsChannelParticipant{}
	_ bin.Decoder = &ChannelsChannelParticipant{}
)
