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

// SecurePlainPhone represents TL type `securePlainPhone#7d6099dd`.
type SecurePlainPhone struct {
	// Phone field of SecurePlainPhone.
	Phone string
}

// SecurePlainPhoneTypeID is TL type id of SecurePlainPhone.
const SecurePlainPhoneTypeID = 0x7d6099dd

// Encode implements bin.Encoder.
func (s *SecurePlainPhone) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePlainPhone#7d6099dd as nil")
	}
	b.PutID(SecurePlainPhoneTypeID)
	b.PutString(s.Phone)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecurePlainPhone) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePlainPhone#7d6099dd to nil")
	}
	if err := b.ConsumeID(SecurePlainPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode securePlainPhone#7d6099dd: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode securePlainPhone#7d6099dd: field phone: %w", err)
		}
		s.Phone = value
	}
	return nil
}

// construct implements constructor of SecurePlainDataClass.
func (s SecurePlainPhone) construct() SecurePlainDataClass { return &s }

// Ensuring interfaces in compile-time for SecurePlainPhone.
var (
	_ bin.Encoder = &SecurePlainPhone{}
	_ bin.Decoder = &SecurePlainPhone{}

	_ SecurePlainDataClass = &SecurePlainPhone{}
)

// SecurePlainEmail represents TL type `securePlainEmail#21ec5a5f`.
type SecurePlainEmail struct {
	// Email field of SecurePlainEmail.
	Email string
}

// SecurePlainEmailTypeID is TL type id of SecurePlainEmail.
const SecurePlainEmailTypeID = 0x21ec5a5f

// Encode implements bin.Encoder.
func (s *SecurePlainEmail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePlainEmail#21ec5a5f as nil")
	}
	b.PutID(SecurePlainEmailTypeID)
	b.PutString(s.Email)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecurePlainEmail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePlainEmail#21ec5a5f to nil")
	}
	if err := b.ConsumeID(SecurePlainEmailTypeID); err != nil {
		return fmt.Errorf("unable to decode securePlainEmail#21ec5a5f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode securePlainEmail#21ec5a5f: field email: %w", err)
		}
		s.Email = value
	}
	return nil
}

// construct implements constructor of SecurePlainDataClass.
func (s SecurePlainEmail) construct() SecurePlainDataClass { return &s }

// Ensuring interfaces in compile-time for SecurePlainEmail.
var (
	_ bin.Encoder = &SecurePlainEmail{}
	_ bin.Decoder = &SecurePlainEmail{}

	_ SecurePlainDataClass = &SecurePlainEmail{}
)

// SecurePlainDataClass represents SecurePlainData generic type.
//
// Example:
//  g, err := DecodeSecurePlainData(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SecurePlainPhone: // securePlainPhone#7d6099dd
//  case *SecurePlainEmail: // securePlainEmail#21ec5a5f
//  default: panic(v)
//  }
type SecurePlainDataClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecurePlainDataClass
}

// DecodeSecurePlainData implements binary de-serialization for SecurePlainDataClass.
func DecodeSecurePlainData(buf *bin.Buffer) (SecurePlainDataClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecurePlainPhoneTypeID:
		// Decoding securePlainPhone#7d6099dd.
		v := SecurePlainPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", err)
		}
		return &v, nil
	case SecurePlainEmailTypeID:
		// Decoding securePlainEmail#21ec5a5f.
		v := SecurePlainEmail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecurePlainDataClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecurePlainData boxes the SecurePlainDataClass providing a helper.
type SecurePlainDataBox struct {
	SecurePlainData SecurePlainDataClass
}

// Decode implements bin.Decoder for SecurePlainDataBox.
func (b *SecurePlainDataBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecurePlainDataBox to nil")
	}
	v, err := DecodeSecurePlainData(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecurePlainData = v
	return nil
}

// Encode implements bin.Encode for SecurePlainDataBox.
func (b *SecurePlainDataBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecurePlainData == nil {
		return fmt.Errorf("unable to encode SecurePlainDataClass as nil")
	}
	return b.SecurePlainData.Encode(buf)
}
