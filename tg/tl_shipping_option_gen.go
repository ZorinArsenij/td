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

// ShippingOption represents TL type `shippingOption#b6213cdf`.
type ShippingOption struct {
	// ID field of ShippingOption.
	ID string
	// Title field of ShippingOption.
	Title string
	// Prices field of ShippingOption.
	Prices []LabeledPrice
}

// ShippingOptionTypeID is TL type id of ShippingOption.
const ShippingOptionTypeID = 0xb6213cdf

// Encode implements bin.Encoder.
func (s *ShippingOption) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shippingOption#b6213cdf as nil")
	}
	b.PutID(ShippingOptionTypeID)
	b.PutString(s.ID)
	b.PutString(s.Title)
	b.PutVectorHeader(len(s.Prices))
	for idx, v := range s.Prices {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode shippingOption#b6213cdf: field prices element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ShippingOption) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shippingOption#b6213cdf to nil")
	}
	if err := b.ConsumeID(ShippingOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode shippingOption#b6213cdf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field title: %w", err)
		}
		s.Title = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field prices: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPrice
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode shippingOption#b6213cdf: field prices: %w", err)
			}
			s.Prices = append(s.Prices, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ShippingOption.
var (
	_ bin.Encoder = &ShippingOption{}
	_ bin.Decoder = &ShippingOption{}
)
