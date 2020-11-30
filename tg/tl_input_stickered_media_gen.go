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

// InputStickeredMediaPhoto represents TL type `inputStickeredMediaPhoto#4a992157`.
type InputStickeredMediaPhoto struct {
	// ID field of InputStickeredMediaPhoto.
	ID InputPhotoClass
}

// InputStickeredMediaPhotoTypeID is TL type id of InputStickeredMediaPhoto.
const InputStickeredMediaPhotoTypeID = 0x4a992157

// Encode implements bin.Encoder.
func (i *InputStickeredMediaPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaPhoto#4a992157 as nil")
	}
	b.PutID(InputStickeredMediaPhotoTypeID)
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputStickeredMediaPhoto#4a992157: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickeredMediaPhoto#4a992157: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickeredMediaPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaPhoto#4a992157 to nil")
	}
	if err := b.ConsumeID(InputStickeredMediaPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickeredMediaPhoto#4a992157: %w", err)
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickeredMediaPhoto#4a992157: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputStickeredMediaClass.
func (i InputStickeredMediaPhoto) construct() InputStickeredMediaClass { return &i }

// Ensuring interfaces in compile-time for InputStickeredMediaPhoto.
var (
	_ bin.Encoder = &InputStickeredMediaPhoto{}
	_ bin.Decoder = &InputStickeredMediaPhoto{}

	_ InputStickeredMediaClass = &InputStickeredMediaPhoto{}
)

// InputStickeredMediaDocument represents TL type `inputStickeredMediaDocument#438865b`.
type InputStickeredMediaDocument struct {
	// ID field of InputStickeredMediaDocument.
	ID InputDocumentClass
}

// InputStickeredMediaDocumentTypeID is TL type id of InputStickeredMediaDocument.
const InputStickeredMediaDocumentTypeID = 0x438865b

// Encode implements bin.Encoder.
func (i *InputStickeredMediaDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickeredMediaDocument#438865b as nil")
	}
	b.PutID(InputStickeredMediaDocumentTypeID)
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputStickeredMediaDocument#438865b: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickeredMediaDocument#438865b: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickeredMediaDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickeredMediaDocument#438865b to nil")
	}
	if err := b.ConsumeID(InputStickeredMediaDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickeredMediaDocument#438865b: %w", err)
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickeredMediaDocument#438865b: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputStickeredMediaClass.
func (i InputStickeredMediaDocument) construct() InputStickeredMediaClass { return &i }

// Ensuring interfaces in compile-time for InputStickeredMediaDocument.
var (
	_ bin.Encoder = &InputStickeredMediaDocument{}
	_ bin.Decoder = &InputStickeredMediaDocument{}

	_ InputStickeredMediaClass = &InputStickeredMediaDocument{}
)

// InputStickeredMediaClass represents InputStickeredMedia generic type.
//
// Example:
//  g, err := DecodeInputStickeredMedia(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputStickeredMediaPhoto: // inputStickeredMediaPhoto#4a992157
//  case *InputStickeredMediaDocument: // inputStickeredMediaDocument#438865b
//  default: panic(v)
//  }
type InputStickeredMediaClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputStickeredMediaClass
}

// DecodeInputStickeredMedia implements binary de-serialization for InputStickeredMediaClass.
func DecodeInputStickeredMedia(buf *bin.Buffer) (InputStickeredMediaClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickeredMediaPhotoTypeID:
		// Decoding inputStickeredMediaPhoto#4a992157.
		v := InputStickeredMediaPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", err)
		}
		return &v, nil
	case InputStickeredMediaDocumentTypeID:
		// Decoding inputStickeredMediaDocument#438865b.
		v := InputStickeredMediaDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickeredMediaClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStickeredMedia boxes the InputStickeredMediaClass providing a helper.
type InputStickeredMediaBox struct {
	InputStickeredMedia InputStickeredMediaClass
}

// Decode implements bin.Decoder for InputStickeredMediaBox.
func (b *InputStickeredMediaBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickeredMediaBox to nil")
	}
	v, err := DecodeInputStickeredMedia(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStickeredMedia = v
	return nil
}

// Encode implements bin.Encode for InputStickeredMediaBox.
func (b *InputStickeredMediaBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStickeredMedia == nil {
		return fmt.Errorf("unable to encode InputStickeredMediaClass as nil")
	}
	return b.InputStickeredMedia.Encode(buf)
}
