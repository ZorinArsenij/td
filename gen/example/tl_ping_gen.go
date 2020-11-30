// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// PingRequest represents TL type `ping#ce73048f`.
type PingRequest struct {
	// ID field of PingRequest.
	ID int32
}

// PingRequestTypeID is TL type id of PingRequest.
const PingRequestTypeID = 0xce73048f

// Encode implements bin.Encoder.
func (p *PingRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping#ce73048f as nil")
	}
	b.PutID(PingRequestTypeID)
	b.PutInt32(p.ID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PingRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping#ce73048f to nil")
	}
	if err := b.ConsumeID(PingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode ping#ce73048f: %w", err)
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode ping#ce73048f: field id: %w", err)
		}
		p.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PingRequest.
var (
	_ bin.Encoder = &PingRequest{}
	_ bin.Decoder = &PingRequest{}
)
