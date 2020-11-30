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

// PhotoEmpty represents TL type `photoEmpty#2331b22d`.
type PhotoEmpty struct {
	// ID field of PhotoEmpty.
	ID int64
}

// PhotoEmptyTypeID is TL type id of PhotoEmpty.
const PhotoEmptyTypeID = 0x2331b22d

// Encode implements bin.Encoder.
func (p *PhotoEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoEmpty#2331b22d as nil")
	}
	b.PutID(PhotoEmptyTypeID)
	b.PutLong(p.ID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoEmpty#2331b22d to nil")
	}
	if err := b.ConsumeID(PhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode photoEmpty#2331b22d: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photoEmpty#2331b22d: field id: %w", err)
		}
		p.ID = value
	}
	return nil
}

// construct implements constructor of PhotoClass.
func (p PhotoEmpty) construct() PhotoClass { return &p }

// Ensuring interfaces in compile-time for PhotoEmpty.
var (
	_ bin.Encoder = &PhotoEmpty{}
	_ bin.Decoder = &PhotoEmpty{}

	_ PhotoClass = &PhotoEmpty{}
)

// Photo represents TL type `photo#fb197a65`.
type Photo struct {
	// Flags field of Photo.
	Flags bin.Fields
	// HasStickers field of Photo.
	HasStickers bool
	// ID field of Photo.
	ID int64
	// AccessHash field of Photo.
	AccessHash int64
	// FileReference field of Photo.
	FileReference []byte
	// Date field of Photo.
	Date int
	// Sizes field of Photo.
	Sizes []PhotoSizeClass
	// VideoSizes field of Photo.
	//
	// Use SetVideoSizes and GetVideoSizes helpers.
	VideoSizes []VideoSize
	// DCID field of Photo.
	DCID int
}

// PhotoTypeID is TL type id of Photo.
const PhotoTypeID = 0xfb197a65

// Encode implements bin.Encoder.
func (p *Photo) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photo#fb197a65 as nil")
	}
	b.PutID(PhotoTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photo#fb197a65: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutBytes(p.FileReference)
	b.PutInt(p.Date)
	b.PutVectorHeader(len(p.Sizes))
	for idx, v := range p.Sizes {
		if v == nil {
			return fmt.Errorf("unable to encode photo#fb197a65: field sizes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photo#fb197a65: field sizes element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(1) {
		b.PutVectorHeader(len(p.VideoSizes))
		for idx, v := range p.VideoSizes {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode photo#fb197a65: field video_sizes element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(p.DCID)
	return nil
}

// SetHasStickers sets value of HasStickers conditional field.
func (p *Photo) SetHasStickers(value bool) {
	if value {
		p.Flags.Set(0)
	} else {
		p.Flags.Unset(0)
	}
}

// SetVideoSizes sets value of VideoSizes conditional field.
func (p *Photo) SetVideoSizes(value []VideoSize) {
	p.Flags.Set(1)
	p.VideoSizes = value
}

// GetVideoSizes returns value of VideoSizes conditional field and
// boolean which is true if field was set.
func (p *Photo) GetVideoSizes() (value []VideoSize, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.VideoSizes, true
}

// Decode implements bin.Decoder.
func (p *Photo) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photo#fb197a65 to nil")
	}
	if err := b.ConsumeID(PhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode photo#fb197a65: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field flags: %w", err)
		}
	}
	p.HasStickers = p.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field file_reference: %w", err)
		}
		p.FileReference = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field date: %w", err)
		}
		p.Date = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field sizes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode photo#fb197a65: field sizes: %w", err)
			}
			p.Sizes = append(p.Sizes, value)
		}
	}
	if p.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field video_sizes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value VideoSize
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode photo#fb197a65: field video_sizes: %w", err)
			}
			p.VideoSizes = append(p.VideoSizes, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field dc_id: %w", err)
		}
		p.DCID = value
	}
	return nil
}

// construct implements constructor of PhotoClass.
func (p Photo) construct() PhotoClass { return &p }

// Ensuring interfaces in compile-time for Photo.
var (
	_ bin.Encoder = &Photo{}
	_ bin.Decoder = &Photo{}

	_ PhotoClass = &Photo{}
)

// PhotoClass represents Photo generic type.
//
// Example:
//  g, err := DecodePhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PhotoEmpty: // photoEmpty#2331b22d
//  case *Photo: // photo#fb197a65
//  default: panic(v)
//  }
type PhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhotoClass
}

// DecodePhoto implements binary de-serialization for PhotoClass.
func DecodePhoto(buf *bin.Buffer) (PhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotoEmptyTypeID:
		// Decoding photoEmpty#2331b22d.
		v := PhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoClass: %w", err)
		}
		return &v, nil
	case PhotoTypeID:
		// Decoding photo#fb197a65.
		v := Photo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// Photo boxes the PhotoClass providing a helper.
type PhotoBox struct {
	Photo PhotoClass
}

// Decode implements bin.Decoder for PhotoBox.
func (b *PhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotoBox to nil")
	}
	v, err := DecodePhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Photo = v
	return nil
}

// Encode implements bin.Encode for PhotoBox.
func (b *PhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Photo == nil {
		return fmt.Errorf("unable to encode PhotoClass as nil")
	}
	return b.Photo.Encode(buf)
}
