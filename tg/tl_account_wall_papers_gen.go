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

// AccountWallPapersNotModified represents TL type `account.wallPapersNotModified#1c199183`.
type AccountWallPapersNotModified struct {
}

// AccountWallPapersNotModifiedTypeID is TL type id of AccountWallPapersNotModified.
const AccountWallPapersNotModifiedTypeID = 0x1c199183

// Encode implements bin.Encoder.
func (w *AccountWallPapersNotModified) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapersNotModified#1c199183 as nil")
	}
	b.PutID(AccountWallPapersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (w *AccountWallPapersNotModified) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapersNotModified#1c199183 to nil")
	}
	if err := b.ConsumeID(AccountWallPapersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.wallPapersNotModified#1c199183: %w", err)
	}
	return nil
}

// construct implements constructor of AccountWallPapersClass.
func (w AccountWallPapersNotModified) construct() AccountWallPapersClass { return &w }

// Ensuring interfaces in compile-time for AccountWallPapersNotModified.
var (
	_ bin.Encoder = &AccountWallPapersNotModified{}
	_ bin.Decoder = &AccountWallPapersNotModified{}

	_ AccountWallPapersClass = &AccountWallPapersNotModified{}
)

// AccountWallPapers represents TL type `account.wallPapers#702b65a9`.
type AccountWallPapers struct {
	// Hash field of AccountWallPapers.
	Hash int
	// Wallpapers field of AccountWallPapers.
	Wallpapers []WallPaperClass
}

// AccountWallPapersTypeID is TL type id of AccountWallPapers.
const AccountWallPapersTypeID = 0x702b65a9

// Encode implements bin.Encoder.
func (w *AccountWallPapers) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapers#702b65a9 as nil")
	}
	b.PutID(AccountWallPapersTypeID)
	b.PutInt(w.Hash)
	b.PutVectorHeader(len(w.Wallpapers))
	for idx, v := range w.Wallpapers {
		if v == nil {
			return fmt.Errorf("unable to encode account.wallPapers#702b65a9: field wallpapers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.wallPapers#702b65a9: field wallpapers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *AccountWallPapers) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapers#702b65a9 to nil")
	}
	if err := b.ConsumeID(AccountWallPapersTypeID); err != nil {
		return fmt.Errorf("unable to decode account.wallPapers#702b65a9: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field hash: %w", err)
		}
		w.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field wallpapers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field wallpapers: %w", err)
			}
			w.Wallpapers = append(w.Wallpapers, value)
		}
	}
	return nil
}

// construct implements constructor of AccountWallPapersClass.
func (w AccountWallPapers) construct() AccountWallPapersClass { return &w }

// Ensuring interfaces in compile-time for AccountWallPapers.
var (
	_ bin.Encoder = &AccountWallPapers{}
	_ bin.Decoder = &AccountWallPapers{}

	_ AccountWallPapersClass = &AccountWallPapers{}
)

// AccountWallPapersClass represents account.WallPapers generic type.
//
// Example:
//  g, err := DecodeAccountWallPapers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AccountWallPapersNotModified: // account.wallPapersNotModified#1c199183
//  case *AccountWallPapers: // account.wallPapers#702b65a9
//  default: panic(v)
//  }
type AccountWallPapersClass interface {
	bin.Encoder
	bin.Decoder
	construct() AccountWallPapersClass
}

// DecodeAccountWallPapers implements binary de-serialization for AccountWallPapersClass.
func DecodeAccountWallPapers(buf *bin.Buffer) (AccountWallPapersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountWallPapersNotModifiedTypeID:
		// Decoding account.wallPapersNotModified#1c199183.
		v := AccountWallPapersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", err)
		}
		return &v, nil
	case AccountWallPapersTypeID:
		// Decoding account.wallPapers#702b65a9.
		v := AccountWallPapers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountWallPapers boxes the AccountWallPapersClass providing a helper.
type AccountWallPapersBox struct {
	AccountWallPapers AccountWallPapersClass
}

// Decode implements bin.Decoder for AccountWallPapersBox.
func (b *AccountWallPapersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountWallPapersBox to nil")
	}
	v, err := DecodeAccountWallPapers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AccountWallPapers = v
	return nil
}

// Encode implements bin.Encode for AccountWallPapersBox.
func (b *AccountWallPapersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AccountWallPapers == nil {
		return fmt.Errorf("unable to encode AccountWallPapersClass as nil")
	}
	return b.AccountWallPapers.Encode(buf)
}
