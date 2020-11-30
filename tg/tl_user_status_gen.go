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

// UserStatusEmpty represents TL type `userStatusEmpty#9d05049`.
type UserStatusEmpty struct {
}

// UserStatusEmptyTypeID is TL type id of UserStatusEmpty.
const UserStatusEmptyTypeID = 0x9d05049

// Encode implements bin.Encoder.
func (u *UserStatusEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusEmpty#9d05049 as nil")
	}
	b.PutID(UserStatusEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusEmpty#9d05049 to nil")
	}
	if err := b.ConsumeID(UserStatusEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusEmpty#9d05049: %w", err)
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusEmpty) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusEmpty.
var (
	_ bin.Encoder = &UserStatusEmpty{}
	_ bin.Decoder = &UserStatusEmpty{}

	_ UserStatusClass = &UserStatusEmpty{}
)

// UserStatusOnline represents TL type `userStatusOnline#edb93949`.
type UserStatusOnline struct {
	// Expires field of UserStatusOnline.
	Expires int
}

// UserStatusOnlineTypeID is TL type id of UserStatusOnline.
const UserStatusOnlineTypeID = 0xedb93949

// Encode implements bin.Encoder.
func (u *UserStatusOnline) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOnline#edb93949 as nil")
	}
	b.PutID(UserStatusOnlineTypeID)
	b.PutInt(u.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusOnline) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOnline#edb93949 to nil")
	}
	if err := b.ConsumeID(UserStatusOnlineTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusOnline#edb93949: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userStatusOnline#edb93949: field expires: %w", err)
		}
		u.Expires = value
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusOnline) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusOnline.
var (
	_ bin.Encoder = &UserStatusOnline{}
	_ bin.Decoder = &UserStatusOnline{}

	_ UserStatusClass = &UserStatusOnline{}
)

// UserStatusOffline represents TL type `userStatusOffline#8c703f`.
type UserStatusOffline struct {
	// WasOnline field of UserStatusOffline.
	WasOnline int
}

// UserStatusOfflineTypeID is TL type id of UserStatusOffline.
const UserStatusOfflineTypeID = 0x8c703f

// Encode implements bin.Encoder.
func (u *UserStatusOffline) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOffline#8c703f as nil")
	}
	b.PutID(UserStatusOfflineTypeID)
	b.PutInt(u.WasOnline)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusOffline) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOffline#8c703f to nil")
	}
	if err := b.ConsumeID(UserStatusOfflineTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusOffline#8c703f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userStatusOffline#8c703f: field was_online: %w", err)
		}
		u.WasOnline = value
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusOffline) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusOffline.
var (
	_ bin.Encoder = &UserStatusOffline{}
	_ bin.Decoder = &UserStatusOffline{}

	_ UserStatusClass = &UserStatusOffline{}
)

// UserStatusRecently represents TL type `userStatusRecently#e26f42f1`.
type UserStatusRecently struct {
}

// UserStatusRecentlyTypeID is TL type id of UserStatusRecently.
const UserStatusRecentlyTypeID = 0xe26f42f1

// Encode implements bin.Encoder.
func (u *UserStatusRecently) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusRecently#e26f42f1 as nil")
	}
	b.PutID(UserStatusRecentlyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusRecently) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusRecently#e26f42f1 to nil")
	}
	if err := b.ConsumeID(UserStatusRecentlyTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusRecently#e26f42f1: %w", err)
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusRecently) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusRecently.
var (
	_ bin.Encoder = &UserStatusRecently{}
	_ bin.Decoder = &UserStatusRecently{}

	_ UserStatusClass = &UserStatusRecently{}
)

// UserStatusLastWeek represents TL type `userStatusLastWeek#7bf09fc`.
type UserStatusLastWeek struct {
}

// UserStatusLastWeekTypeID is TL type id of UserStatusLastWeek.
const UserStatusLastWeekTypeID = 0x7bf09fc

// Encode implements bin.Encoder.
func (u *UserStatusLastWeek) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastWeek#7bf09fc as nil")
	}
	b.PutID(UserStatusLastWeekTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusLastWeek) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastWeek#7bf09fc to nil")
	}
	if err := b.ConsumeID(UserStatusLastWeekTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusLastWeek#7bf09fc: %w", err)
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusLastWeek) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusLastWeek.
var (
	_ bin.Encoder = &UserStatusLastWeek{}
	_ bin.Decoder = &UserStatusLastWeek{}

	_ UserStatusClass = &UserStatusLastWeek{}
)

// UserStatusLastMonth represents TL type `userStatusLastMonth#77ebc742`.
type UserStatusLastMonth struct {
}

// UserStatusLastMonthTypeID is TL type id of UserStatusLastMonth.
const UserStatusLastMonthTypeID = 0x77ebc742

// Encode implements bin.Encoder.
func (u *UserStatusLastMonth) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastMonth#77ebc742 as nil")
	}
	b.PutID(UserStatusLastMonthTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusLastMonth) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastMonth#77ebc742 to nil")
	}
	if err := b.ConsumeID(UserStatusLastMonthTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusLastMonth#77ebc742: %w", err)
	}
	return nil
}

// construct implements constructor of UserStatusClass.
func (u UserStatusLastMonth) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusLastMonth.
var (
	_ bin.Encoder = &UserStatusLastMonth{}
	_ bin.Decoder = &UserStatusLastMonth{}

	_ UserStatusClass = &UserStatusLastMonth{}
)

// UserStatusClass represents UserStatus generic type.
//
// Example:
//  g, err := DecodeUserStatus(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserStatusEmpty: // userStatusEmpty#9d05049
//  case *UserStatusOnline: // userStatusOnline#edb93949
//  case *UserStatusOffline: // userStatusOffline#8c703f
//  case *UserStatusRecently: // userStatusRecently#e26f42f1
//  case *UserStatusLastWeek: // userStatusLastWeek#7bf09fc
//  case *UserStatusLastMonth: // userStatusLastMonth#77ebc742
//  default: panic(v)
//  }
type UserStatusClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserStatusClass
}

// DecodeUserStatus implements binary de-serialization for UserStatusClass.
func DecodeUserStatus(buf *bin.Buffer) (UserStatusClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserStatusEmptyTypeID:
		// Decoding userStatusEmpty#9d05049.
		v := UserStatusEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusOnlineTypeID:
		// Decoding userStatusOnline#edb93949.
		v := UserStatusOnline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusOfflineTypeID:
		// Decoding userStatusOffline#8c703f.
		v := UserStatusOffline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusRecentlyTypeID:
		// Decoding userStatusRecently#e26f42f1.
		v := UserStatusRecently{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusLastWeekTypeID:
		// Decoding userStatusLastWeek#7bf09fc.
		v := UserStatusLastWeek{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusLastMonthTypeID:
		// Decoding userStatusLastMonth#77ebc742.
		v := UserStatusLastMonth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserStatusClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserStatus boxes the UserStatusClass providing a helper.
type UserStatusBox struct {
	UserStatus UserStatusClass
}

// Decode implements bin.Decoder for UserStatusBox.
func (b *UserStatusBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserStatusBox to nil")
	}
	v, err := DecodeUserStatus(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserStatus = v
	return nil
}

// Encode implements bin.Encode for UserStatusBox.
func (b *UserStatusBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserStatus == nil {
		return fmt.Errorf("unable to encode UserStatusClass as nil")
	}
	return b.UserStatus.Encode(buf)
}
