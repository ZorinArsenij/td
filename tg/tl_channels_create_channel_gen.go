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

// ChannelsCreateChannelRequest represents TL type `channels.createChannel#3d5fb10f`.
type ChannelsCreateChannelRequest struct {
	// Flags field of ChannelsCreateChannelRequest.
	Flags bin.Fields
	// Broadcast field of ChannelsCreateChannelRequest.
	Broadcast bool
	// Megagroup field of ChannelsCreateChannelRequest.
	Megagroup bool
	// Title field of ChannelsCreateChannelRequest.
	Title string
	// About field of ChannelsCreateChannelRequest.
	About string
	// GeoPoint field of ChannelsCreateChannelRequest.
	//
	// Use SetGeoPoint and GetGeoPoint helpers.
	GeoPoint InputGeoPointClass
	// Address field of ChannelsCreateChannelRequest.
	//
	// Use SetAddress and GetAddress helpers.
	Address string
}

// ChannelsCreateChannelRequestTypeID is TL type id of ChannelsCreateChannelRequest.
const ChannelsCreateChannelRequestTypeID = 0x3d5fb10f

// Encode implements bin.Encoder.
func (c *ChannelsCreateChannelRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.createChannel#3d5fb10f as nil")
	}
	b.PutID(ChannelsCreateChannelRequestTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field flags: %w", err)
	}
	b.PutString(c.Title)
	b.PutString(c.About)
	if c.Flags.Has(2) {
		if c.GeoPoint == nil {
			return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field geo_point is nil")
		}
		if err := c.GeoPoint.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.createChannel#3d5fb10f: field geo_point: %w", err)
		}
	}
	if c.Flags.Has(2) {
		b.PutString(c.Address)
	}
	return nil
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChannelsCreateChannelRequest) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChannelsCreateChannelRequest) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// SetGeoPoint sets value of GeoPoint conditional field.
func (c *ChannelsCreateChannelRequest) SetGeoPoint(value InputGeoPointClass) {
	c.Flags.Set(2)
	c.GeoPoint = value
}

// GetGeoPoint returns value of GeoPoint conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetGeoPoint() (value InputGeoPointClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.GeoPoint, true
}

// SetAddress sets value of Address conditional field.
func (c *ChannelsCreateChannelRequest) SetAddress(value string) {
	c.Flags.Set(2)
	c.Address = value
}

// GetAddress returns value of Address conditional field and
// boolean which is true if field was set.
func (c *ChannelsCreateChannelRequest) GetAddress() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Address, true
}

// Decode implements bin.Decoder.
func (c *ChannelsCreateChannelRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.createChannel#3d5fb10f to nil")
	}
	if err := b.ConsumeID(ChannelsCreateChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field flags: %w", err)
		}
	}
	c.Broadcast = c.Flags.Has(0)
	c.Megagroup = c.Flags.Has(1)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field about: %w", err)
		}
		c.About = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.createChannel#3d5fb10f: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsCreateChannelRequest.
var (
	_ bin.Encoder = &ChannelsCreateChannelRequest{}
	_ bin.Decoder = &ChannelsCreateChannelRequest{}
)
