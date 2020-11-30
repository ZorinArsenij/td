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

// InputEncryptedFileEmpty represents TL type `inputEncryptedFileEmpty#1837c364`.
type InputEncryptedFileEmpty struct {
}

// InputEncryptedFileEmptyTypeID is TL type id of InputEncryptedFileEmpty.
const InputEncryptedFileEmptyTypeID = 0x1837c364

// Encode implements bin.Encoder.
func (i *InputEncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileEmpty#1837c364 as nil")
	}
	b.PutID(InputEncryptedFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileEmpty#1837c364 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileEmpty#1837c364: %w", err)
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileEmpty) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileEmpty.
var (
	_ bin.Encoder = &InputEncryptedFileEmpty{}
	_ bin.Decoder = &InputEncryptedFileEmpty{}

	_ InputEncryptedFileClass = &InputEncryptedFileEmpty{}
)

// InputEncryptedFileUploaded represents TL type `inputEncryptedFileUploaded#64bd0306`.
type InputEncryptedFileUploaded struct {
	// ID field of InputEncryptedFileUploaded.
	ID int64
	// Parts field of InputEncryptedFileUploaded.
	Parts int
	// Md5Checksum field of InputEncryptedFileUploaded.
	Md5Checksum string
	// KeyFingerprint field of InputEncryptedFileUploaded.
	KeyFingerprint int
}

// InputEncryptedFileUploadedTypeID is TL type id of InputEncryptedFileUploaded.
const InputEncryptedFileUploadedTypeID = 0x64bd0306

// Encode implements bin.Encoder.
func (i *InputEncryptedFileUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileUploaded#64bd0306 as nil")
	}
	b.PutID(InputEncryptedFileUploadedTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Md5Checksum)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileUploaded#64bd0306 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field md5_checksum: %w", err)
		}
		i.Md5Checksum = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileUploaded.
var (
	_ bin.Encoder = &InputEncryptedFileUploaded{}
	_ bin.Decoder = &InputEncryptedFileUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileUploaded{}
)

// InputEncryptedFile represents TL type `inputEncryptedFile#5a17b5e5`.
type InputEncryptedFile struct {
	// ID field of InputEncryptedFile.
	ID int64
	// AccessHash field of InputEncryptedFile.
	AccessHash int64
}

// InputEncryptedFileTypeID is TL type id of InputEncryptedFile.
const InputEncryptedFileTypeID = 0x5a17b5e5

// Encode implements bin.Encoder.
func (i *InputEncryptedFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFile#5a17b5e5 as nil")
	}
	b.PutID(InputEncryptedFileTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFile#5a17b5e5 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFile) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFile.
var (
	_ bin.Encoder = &InputEncryptedFile{}
	_ bin.Decoder = &InputEncryptedFile{}

	_ InputEncryptedFileClass = &InputEncryptedFile{}
)

// InputEncryptedFileBigUploaded represents TL type `inputEncryptedFileBigUploaded#2dc173c8`.
type InputEncryptedFileBigUploaded struct {
	// ID field of InputEncryptedFileBigUploaded.
	ID int64
	// Parts field of InputEncryptedFileBigUploaded.
	Parts int
	// KeyFingerprint field of InputEncryptedFileBigUploaded.
	KeyFingerprint int
}

// InputEncryptedFileBigUploadedTypeID is TL type id of InputEncryptedFileBigUploaded.
const InputEncryptedFileBigUploadedTypeID = 0x2dc173c8

// Encode implements bin.Encoder.
func (i *InputEncryptedFileBigUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileBigUploaded#2dc173c8 as nil")
	}
	b.PutID(InputEncryptedFileBigUploadedTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileBigUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileBigUploaded#2dc173c8 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileBigUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileBigUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileBigUploaded.
var (
	_ bin.Encoder = &InputEncryptedFileBigUploaded{}
	_ bin.Decoder = &InputEncryptedFileBigUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileBigUploaded{}
)

// InputEncryptedFileClass represents InputEncryptedFile generic type.
//
// Example:
//  g, err := DecodeInputEncryptedFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputEncryptedFileEmpty: // inputEncryptedFileEmpty#1837c364
//  case *InputEncryptedFileUploaded: // inputEncryptedFileUploaded#64bd0306
//  case *InputEncryptedFile: // inputEncryptedFile#5a17b5e5
//  case *InputEncryptedFileBigUploaded: // inputEncryptedFileBigUploaded#2dc173c8
//  default: panic(v)
//  }
type InputEncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputEncryptedFileClass
}

// DecodeInputEncryptedFile implements binary de-serialization for InputEncryptedFileClass.
func DecodeInputEncryptedFile(buf *bin.Buffer) (InputEncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputEncryptedFileEmptyTypeID:
		// Decoding inputEncryptedFileEmpty#1837c364.
		v := InputEncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileUploadedTypeID:
		// Decoding inputEncryptedFileUploaded#64bd0306.
		v := InputEncryptedFileUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileTypeID:
		// Decoding inputEncryptedFile#5a17b5e5.
		v := InputEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileBigUploadedTypeID:
		// Decoding inputEncryptedFileBigUploaded#2dc173c8.
		v := InputEncryptedFileBigUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputEncryptedFile boxes the InputEncryptedFileClass providing a helper.
type InputEncryptedFileBox struct {
	InputEncryptedFile InputEncryptedFileClass
}

// Decode implements bin.Decoder for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputEncryptedFileBox to nil")
	}
	v, err := DecodeInputEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputEncryptedFile = v
	return nil
}

// Encode implements bin.Encode for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputEncryptedFile == nil {
		return fmt.Errorf("unable to encode InputEncryptedFileClass as nil")
	}
	return b.InputEncryptedFile.Encode(buf)
}
