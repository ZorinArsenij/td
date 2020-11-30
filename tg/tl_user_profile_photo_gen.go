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

// UserProfilePhotoEmpty represents TL type `userProfilePhotoEmpty#4f11bae1`.
type UserProfilePhotoEmpty struct {
}

// UserProfilePhotoEmptyTypeID is TL type id of UserProfilePhotoEmpty.
const UserProfilePhotoEmptyTypeID = 0x4f11bae1

// Encode implements bin.Encoder.
func (u *UserProfilePhotoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhotoEmpty#4f11bae1 as nil")
	}
	b.PutID(UserProfilePhotoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserProfilePhotoEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhotoEmpty#4f11bae1 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhotoEmpty#4f11bae1: %w", err)
	}
	return nil
}

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhotoEmpty) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhotoEmpty.
var (
	_ bin.Encoder = &UserProfilePhotoEmpty{}
	_ bin.Decoder = &UserProfilePhotoEmpty{}

	_ UserProfilePhotoClass = &UserProfilePhotoEmpty{}
)

// UserProfilePhoto represents TL type `userProfilePhoto#69d3ab26`.
type UserProfilePhoto struct {
	// Flags field of UserProfilePhoto.
	Flags bin.Fields
	// HasVideo field of UserProfilePhoto.
	HasVideo bool
	// PhotoID field of UserProfilePhoto.
	PhotoID int64
	// PhotoSmall field of UserProfilePhoto.
	PhotoSmall FileLocationToBeDeprecated
	// PhotoBig field of UserProfilePhoto.
	PhotoBig FileLocationToBeDeprecated
	// DCID field of UserProfilePhoto.
	DCID int
}

// UserProfilePhotoTypeID is TL type id of UserProfilePhoto.
const UserProfilePhotoTypeID = 0x69d3ab26

// Encode implements bin.Encoder.
func (u *UserProfilePhoto) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhoto#69d3ab26 as nil")
	}
	b.PutID(UserProfilePhotoTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field flags: %w", err)
	}
	b.PutLong(u.PhotoID)
	if err := u.PhotoSmall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field photo_small: %w", err)
	}
	if err := u.PhotoBig.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field photo_big: %w", err)
	}
	b.PutInt(u.DCID)
	return nil
}

// SetHasVideo sets value of HasVideo conditional field.
func (u *UserProfilePhoto) SetHasVideo(value bool) {
	if value {
		u.Flags.Set(0)
	} else {
		u.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (u *UserProfilePhoto) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhoto#69d3ab26 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field flags: %w", err)
		}
	}
	u.HasVideo = u.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_id: %w", err)
		}
		u.PhotoID = value
	}
	{
		if err := u.PhotoSmall.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_small: %w", err)
		}
	}
	{
		if err := u.PhotoBig.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_big: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field dc_id: %w", err)
		}
		u.DCID = value
	}
	return nil
}

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhoto) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhoto.
var (
	_ bin.Encoder = &UserProfilePhoto{}
	_ bin.Decoder = &UserProfilePhoto{}

	_ UserProfilePhotoClass = &UserProfilePhoto{}
)

// UserProfilePhotoClass represents UserProfilePhoto generic type.
//
// Example:
//  g, err := DecodeUserProfilePhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserProfilePhotoEmpty: // userProfilePhotoEmpty#4f11bae1
//  case *UserProfilePhoto: // userProfilePhoto#69d3ab26
//  default: panic(v)
//  }
type UserProfilePhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserProfilePhotoClass
}

// DecodeUserProfilePhoto implements binary de-serialization for UserProfilePhotoClass.
func DecodeUserProfilePhoto(buf *bin.Buffer) (UserProfilePhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserProfilePhotoEmptyTypeID:
		// Decoding userProfilePhotoEmpty#4f11bae1.
		v := UserProfilePhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	case UserProfilePhotoTypeID:
		// Decoding userProfilePhoto#69d3ab26.
		v := UserProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserProfilePhoto boxes the UserProfilePhotoClass providing a helper.
type UserProfilePhotoBox struct {
	UserProfilePhoto UserProfilePhotoClass
}

// Decode implements bin.Decoder for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserProfilePhotoBox to nil")
	}
	v, err := DecodeUserProfilePhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserProfilePhoto = v
	return nil
}

// Encode implements bin.Encode for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserProfilePhoto == nil {
		return fmt.Errorf("unable to encode UserProfilePhotoClass as nil")
	}
	return b.UserProfilePhoto.Encode(buf)
}
