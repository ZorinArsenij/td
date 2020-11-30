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

// AuthCodeTypeSms represents TL type `auth.codeTypeSms#72a3158c`.
type AuthCodeTypeSms struct {
}

// AuthCodeTypeSmsTypeID is TL type id of AuthCodeTypeSms.
const AuthCodeTypeSmsTypeID = 0x72a3158c

// Encode implements bin.Encoder.
func (c *AuthCodeTypeSms) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeSms#72a3158c as nil")
	}
	b.PutID(AuthCodeTypeSmsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeSms) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeSms#72a3158c to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeSmsTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeSms#72a3158c: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeSms) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeSms.
var (
	_ bin.Encoder = &AuthCodeTypeSms{}
	_ bin.Decoder = &AuthCodeTypeSms{}

	_ AuthCodeTypeClass = &AuthCodeTypeSms{}
)

// AuthCodeTypeCall represents TL type `auth.codeTypeCall#741cd3e3`.
type AuthCodeTypeCall struct {
}

// AuthCodeTypeCallTypeID is TL type id of AuthCodeTypeCall.
const AuthCodeTypeCallTypeID = 0x741cd3e3

// Encode implements bin.Encoder.
func (c *AuthCodeTypeCall) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeCall#741cd3e3 as nil")
	}
	b.PutID(AuthCodeTypeCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeCall) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeCall#741cd3e3 to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeCall#741cd3e3: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeCall) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeCall.
var (
	_ bin.Encoder = &AuthCodeTypeCall{}
	_ bin.Decoder = &AuthCodeTypeCall{}

	_ AuthCodeTypeClass = &AuthCodeTypeCall{}
)

// AuthCodeTypeFlashCall represents TL type `auth.codeTypeFlashCall#226ccefb`.
type AuthCodeTypeFlashCall struct {
}

// AuthCodeTypeFlashCallTypeID is TL type id of AuthCodeTypeFlashCall.
const AuthCodeTypeFlashCallTypeID = 0x226ccefb

// Encode implements bin.Encoder.
func (c *AuthCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeFlashCall#226ccefb as nil")
	}
	b.PutID(AuthCodeTypeFlashCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeFlashCall#226ccefb to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeFlashCall#226ccefb: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeFlashCall) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeFlashCall.
var (
	_ bin.Encoder = &AuthCodeTypeFlashCall{}
	_ bin.Decoder = &AuthCodeTypeFlashCall{}

	_ AuthCodeTypeClass = &AuthCodeTypeFlashCall{}
)

// AuthCodeTypeClass represents auth.CodeType generic type.
//
// Example:
//  g, err := DecodeAuthCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthCodeTypeSms: // auth.codeTypeSms#72a3158c
//  case *AuthCodeTypeCall: // auth.codeTypeCall#741cd3e3
//  case *AuthCodeTypeFlashCall: // auth.codeTypeFlashCall#226ccefb
//  default: panic(v)
//  }
type AuthCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthCodeTypeClass
}

// DecodeAuthCodeType implements binary de-serialization for AuthCodeTypeClass.
func DecodeAuthCodeType(buf *bin.Buffer) (AuthCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthCodeTypeSmsTypeID:
		// Decoding auth.codeTypeSms#72a3158c.
		v := AuthCodeTypeSms{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthCodeTypeCallTypeID:
		// Decoding auth.codeTypeCall#741cd3e3.
		v := AuthCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthCodeTypeFlashCallTypeID:
		// Decoding auth.codeTypeFlashCall#226ccefb.
		v := AuthCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthCodeType boxes the AuthCodeTypeClass providing a helper.
type AuthCodeTypeBox struct {
	AuthCodeType AuthCodeTypeClass
}

// Decode implements bin.Decoder for AuthCodeTypeBox.
func (b *AuthCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthCodeTypeBox to nil")
	}
	v, err := DecodeAuthCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AuthCodeType = v
	return nil
}

// Encode implements bin.Encode for AuthCodeTypeBox.
func (b *AuthCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AuthCodeType == nil {
		return fmt.Errorf("unable to encode AuthCodeTypeClass as nil")
	}
	return b.AuthCodeType.Encode(buf)
}
