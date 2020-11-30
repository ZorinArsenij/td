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

// BotsSendCustomRequestRequest represents TL type `bots.sendCustomRequest#aa2769ed`.
type BotsSendCustomRequestRequest struct {
	// CustomMethod field of BotsSendCustomRequestRequest.
	CustomMethod string
	// Params field of BotsSendCustomRequestRequest.
	Params DataJSON
}

// BotsSendCustomRequestRequestTypeID is TL type id of BotsSendCustomRequestRequest.
const BotsSendCustomRequestRequestTypeID = 0xaa2769ed

// Encode implements bin.Encoder.
func (s *BotsSendCustomRequestRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.sendCustomRequest#aa2769ed as nil")
	}
	b.PutID(BotsSendCustomRequestRequestTypeID)
	b.PutString(s.CustomMethod)
	if err := s.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.sendCustomRequest#aa2769ed: field params: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *BotsSendCustomRequestRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.sendCustomRequest#aa2769ed to nil")
	}
	if err := b.ConsumeID(BotsSendCustomRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field custom_method: %w", err)
		}
		s.CustomMethod = value
	}
	{
		if err := s.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field params: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for BotsSendCustomRequestRequest.
var (
	_ bin.Encoder = &BotsSendCustomRequestRequest{}
	_ bin.Decoder = &BotsSendCustomRequestRequest{}
)
