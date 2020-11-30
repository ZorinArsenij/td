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

// HelpPassportConfigNotModified represents TL type `help.passportConfigNotModified#bfb9f457`.
type HelpPassportConfigNotModified struct {
}

// HelpPassportConfigNotModifiedTypeID is TL type id of HelpPassportConfigNotModified.
const HelpPassportConfigNotModifiedTypeID = 0xbfb9f457

// Encode implements bin.Encoder.
func (p *HelpPassportConfigNotModified) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfigNotModified#bfb9f457 as nil")
	}
	b.PutID(HelpPassportConfigNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfigNotModified) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfigNotModified#bfb9f457 to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfigNotModified#bfb9f457: %w", err)
	}
	return nil
}

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfigNotModified) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfigNotModified.
var (
	_ bin.Encoder = &HelpPassportConfigNotModified{}
	_ bin.Decoder = &HelpPassportConfigNotModified{}

	_ HelpPassportConfigClass = &HelpPassportConfigNotModified{}
)

// HelpPassportConfig represents TL type `help.passportConfig#a098d6af`.
type HelpPassportConfig struct {
	// Hash field of HelpPassportConfig.
	Hash int
	// CountriesLangs field of HelpPassportConfig.
	CountriesLangs DataJSON
}

// HelpPassportConfigTypeID is TL type id of HelpPassportConfig.
const HelpPassportConfigTypeID = 0xa098d6af

// Encode implements bin.Encoder.
func (p *HelpPassportConfig) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfig#a098d6af as nil")
	}
	b.PutID(HelpPassportConfigTypeID)
	b.PutInt(p.Hash)
	if err := p.CountriesLangs.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.passportConfig#a098d6af: field countries_langs: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfig) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfig#a098d6af to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfig#a098d6af: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field hash: %w", err)
		}
		p.Hash = value
	}
	{
		if err := p.CountriesLangs.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field countries_langs: %w", err)
		}
	}
	return nil
}

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfig) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfig.
var (
	_ bin.Encoder = &HelpPassportConfig{}
	_ bin.Decoder = &HelpPassportConfig{}

	_ HelpPassportConfigClass = &HelpPassportConfig{}
)

// HelpPassportConfigClass represents help.PassportConfig generic type.
//
// Example:
//  g, err := DecodeHelpPassportConfig(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpPassportConfigNotModified: // help.passportConfigNotModified#bfb9f457
//  case *HelpPassportConfig: // help.passportConfig#a098d6af
//  default: panic(v)
//  }
type HelpPassportConfigClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpPassportConfigClass
}

// DecodeHelpPassportConfig implements binary de-serialization for HelpPassportConfigClass.
func DecodeHelpPassportConfig(buf *bin.Buffer) (HelpPassportConfigClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPassportConfigNotModifiedTypeID:
		// Decoding help.passportConfigNotModified#bfb9f457.
		v := HelpPassportConfigNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	case HelpPassportConfigTypeID:
		// Decoding help.passportConfig#a098d6af.
		v := HelpPassportConfig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPassportConfig boxes the HelpPassportConfigClass providing a helper.
type HelpPassportConfigBox struct {
	HelpPassportConfig HelpPassportConfigClass
}

// Decode implements bin.Decoder for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPassportConfigBox to nil")
	}
	v, err := DecodeHelpPassportConfig(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.HelpPassportConfig = v
	return nil
}

// Encode implements bin.Encode for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.HelpPassportConfig == nil {
		return fmt.Errorf("unable to encode HelpPassportConfigClass as nil")
	}
	return b.HelpPassportConfig.Encode(buf)
}
