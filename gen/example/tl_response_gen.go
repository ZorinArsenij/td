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

// ResponseID represents TL type `responseID#85d7fd8b`.
type ResponseID struct {
	// ID field of ResponseID.
	ID int32
}

// ResponseIDTypeID is TL type id of ResponseID.
const ResponseIDTypeID = 0x85d7fd8b

// Encode implements bin.Encoder.
func (r *ResponseID) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode responseID#85d7fd8b as nil")
	}
	b.PutID(ResponseIDTypeID)
	b.PutInt32(r.ID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ResponseID) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode responseID#85d7fd8b to nil")
	}
	if err := b.ConsumeID(ResponseIDTypeID); err != nil {
		return fmt.Errorf("unable to decode responseID#85d7fd8b: %w", err)
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode responseID#85d7fd8b: field id: %w", err)
		}
		r.ID = value
	}
	return nil
}

// construct implements constructor of ResponseClass.
func (r ResponseID) construct() ResponseClass { return &r }

// Ensuring interfaces in compile-time for ResponseID.
var (
	_ bin.Encoder = &ResponseID{}
	_ bin.Decoder = &ResponseID{}

	_ ResponseClass = &ResponseID{}
)

// ResponseText represents TL type `responseText#cb0244f2`.
type ResponseText struct {
	// Text field of ResponseText.
	Text string
}

// ResponseTextTypeID is TL type id of ResponseText.
const ResponseTextTypeID = 0xcb0244f2

// Encode implements bin.Encoder.
func (r *ResponseText) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode responseText#cb0244f2 as nil")
	}
	b.PutID(ResponseTextTypeID)
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ResponseText) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode responseText#cb0244f2 to nil")
	}
	if err := b.ConsumeID(ResponseTextTypeID); err != nil {
		return fmt.Errorf("unable to decode responseText#cb0244f2: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode responseText#cb0244f2: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// construct implements constructor of ResponseClass.
func (r ResponseText) construct() ResponseClass { return &r }

// Ensuring interfaces in compile-time for ResponseText.
var (
	_ bin.Encoder = &ResponseText{}
	_ bin.Decoder = &ResponseText{}

	_ ResponseClass = &ResponseText{}
)

// ResponseClass represents Response generic type.
//
// Example:
//  g, err := DecodeResponse(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ResponseID: // responseID#85d7fd8b
//  case *ResponseText: // responseText#cb0244f2
//  default: panic(v)
//  }
type ResponseClass interface {
	bin.Encoder
	bin.Decoder
	construct() ResponseClass
}

// DecodeResponse implements binary de-serialization for ResponseClass.
func DecodeResponse(buf *bin.Buffer) (ResponseClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ResponseIDTypeID:
		// Decoding responseID#85d7fd8b.
		v := ResponseID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResponseClass: %w", err)
		}
		return &v, nil
	case ResponseTextTypeID:
		// Decoding responseText#cb0244f2.
		v := ResponseText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResponseClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ResponseClass: %w", bin.NewUnexpectedID(id))
	}
}

// Response boxes the ResponseClass providing a helper.
type ResponseBox struct {
	Response ResponseClass
}

// Decode implements bin.Decoder for ResponseBox.
func (b *ResponseBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ResponseBox to nil")
	}
	v, err := DecodeResponse(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Response = v
	return nil
}

// Encode implements bin.Encode for ResponseBox.
func (b *ResponseBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Response == nil {
		return fmt.Errorf("unable to encode ResponseClass as nil")
	}
	return b.Response.Encode(buf)
}
