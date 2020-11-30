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

// SendMessageTypingAction represents TL type `sendMessageTypingAction#16bf744e`.
type SendMessageTypingAction struct {
}

// SendMessageTypingActionTypeID is TL type id of SendMessageTypingAction.
const SendMessageTypingActionTypeID = 0x16bf744e

// Encode implements bin.Encoder.
func (s *SendMessageTypingAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageTypingAction#16bf744e as nil")
	}
	b.PutID(SendMessageTypingActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageTypingAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageTypingAction#16bf744e to nil")
	}
	if err := b.ConsumeID(SendMessageTypingActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageTypingAction#16bf744e: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageTypingAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageTypingAction.
var (
	_ bin.Encoder = &SendMessageTypingAction{}
	_ bin.Decoder = &SendMessageTypingAction{}

	_ SendMessageActionClass = &SendMessageTypingAction{}
)

// SendMessageCancelAction represents TL type `sendMessageCancelAction#fd5ec8f5`.
type SendMessageCancelAction struct {
}

// SendMessageCancelActionTypeID is TL type id of SendMessageCancelAction.
const SendMessageCancelActionTypeID = 0xfd5ec8f5

// Encode implements bin.Encoder.
func (s *SendMessageCancelAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageCancelAction#fd5ec8f5 as nil")
	}
	b.PutID(SendMessageCancelActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageCancelAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageCancelAction#fd5ec8f5 to nil")
	}
	if err := b.ConsumeID(SendMessageCancelActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageCancelAction#fd5ec8f5: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageCancelAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageCancelAction.
var (
	_ bin.Encoder = &SendMessageCancelAction{}
	_ bin.Decoder = &SendMessageCancelAction{}

	_ SendMessageActionClass = &SendMessageCancelAction{}
)

// SendMessageRecordVideoAction represents TL type `sendMessageRecordVideoAction#a187d66f`.
type SendMessageRecordVideoAction struct {
}

// SendMessageRecordVideoActionTypeID is TL type id of SendMessageRecordVideoAction.
const SendMessageRecordVideoActionTypeID = 0xa187d66f

// Encode implements bin.Encoder.
func (s *SendMessageRecordVideoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordVideoAction#a187d66f as nil")
	}
	b.PutID(SendMessageRecordVideoActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordVideoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordVideoAction#a187d66f to nil")
	}
	if err := b.ConsumeID(SendMessageRecordVideoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordVideoAction#a187d66f: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordVideoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordVideoAction.
var (
	_ bin.Encoder = &SendMessageRecordVideoAction{}
	_ bin.Decoder = &SendMessageRecordVideoAction{}

	_ SendMessageActionClass = &SendMessageRecordVideoAction{}
)

// SendMessageUploadVideoAction represents TL type `sendMessageUploadVideoAction#e9763aec`.
type SendMessageUploadVideoAction struct {
	// Progress field of SendMessageUploadVideoAction.
	Progress int
}

// SendMessageUploadVideoActionTypeID is TL type id of SendMessageUploadVideoAction.
const SendMessageUploadVideoActionTypeID = 0xe9763aec

// Encode implements bin.Encoder.
func (s *SendMessageUploadVideoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadVideoAction#e9763aec as nil")
	}
	b.PutID(SendMessageUploadVideoActionTypeID)
	b.PutInt(s.Progress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadVideoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadVideoAction#e9763aec to nil")
	}
	if err := b.ConsumeID(SendMessageUploadVideoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadVideoAction#e9763aec: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageUploadVideoAction#e9763aec: field progress: %w", err)
		}
		s.Progress = value
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadVideoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadVideoAction.
var (
	_ bin.Encoder = &SendMessageUploadVideoAction{}
	_ bin.Decoder = &SendMessageUploadVideoAction{}

	_ SendMessageActionClass = &SendMessageUploadVideoAction{}
)

// SendMessageRecordAudioAction represents TL type `sendMessageRecordAudioAction#d52f73f7`.
type SendMessageRecordAudioAction struct {
}

// SendMessageRecordAudioActionTypeID is TL type id of SendMessageRecordAudioAction.
const SendMessageRecordAudioActionTypeID = 0xd52f73f7

// Encode implements bin.Encoder.
func (s *SendMessageRecordAudioAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordAudioAction#d52f73f7 as nil")
	}
	b.PutID(SendMessageRecordAudioActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordAudioAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordAudioAction#d52f73f7 to nil")
	}
	if err := b.ConsumeID(SendMessageRecordAudioActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordAudioAction#d52f73f7: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordAudioAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordAudioAction.
var (
	_ bin.Encoder = &SendMessageRecordAudioAction{}
	_ bin.Decoder = &SendMessageRecordAudioAction{}

	_ SendMessageActionClass = &SendMessageRecordAudioAction{}
)

// SendMessageUploadAudioAction represents TL type `sendMessageUploadAudioAction#f351d7ab`.
type SendMessageUploadAudioAction struct {
	// Progress field of SendMessageUploadAudioAction.
	Progress int
}

// SendMessageUploadAudioActionTypeID is TL type id of SendMessageUploadAudioAction.
const SendMessageUploadAudioActionTypeID = 0xf351d7ab

// Encode implements bin.Encoder.
func (s *SendMessageUploadAudioAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadAudioAction#f351d7ab as nil")
	}
	b.PutID(SendMessageUploadAudioActionTypeID)
	b.PutInt(s.Progress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadAudioAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadAudioAction#f351d7ab to nil")
	}
	if err := b.ConsumeID(SendMessageUploadAudioActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadAudioAction#f351d7ab: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageUploadAudioAction#f351d7ab: field progress: %w", err)
		}
		s.Progress = value
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadAudioAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadAudioAction.
var (
	_ bin.Encoder = &SendMessageUploadAudioAction{}
	_ bin.Decoder = &SendMessageUploadAudioAction{}

	_ SendMessageActionClass = &SendMessageUploadAudioAction{}
)

// SendMessageUploadPhotoAction represents TL type `sendMessageUploadPhotoAction#d1d34a26`.
type SendMessageUploadPhotoAction struct {
	// Progress field of SendMessageUploadPhotoAction.
	Progress int
}

// SendMessageUploadPhotoActionTypeID is TL type id of SendMessageUploadPhotoAction.
const SendMessageUploadPhotoActionTypeID = 0xd1d34a26

// Encode implements bin.Encoder.
func (s *SendMessageUploadPhotoAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadPhotoAction#d1d34a26 as nil")
	}
	b.PutID(SendMessageUploadPhotoActionTypeID)
	b.PutInt(s.Progress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadPhotoAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadPhotoAction#d1d34a26 to nil")
	}
	if err := b.ConsumeID(SendMessageUploadPhotoActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadPhotoAction#d1d34a26: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageUploadPhotoAction#d1d34a26: field progress: %w", err)
		}
		s.Progress = value
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadPhotoAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadPhotoAction.
var (
	_ bin.Encoder = &SendMessageUploadPhotoAction{}
	_ bin.Decoder = &SendMessageUploadPhotoAction{}

	_ SendMessageActionClass = &SendMessageUploadPhotoAction{}
)

// SendMessageUploadDocumentAction represents TL type `sendMessageUploadDocumentAction#aa0cd9e4`.
type SendMessageUploadDocumentAction struct {
	// Progress field of SendMessageUploadDocumentAction.
	Progress int
}

// SendMessageUploadDocumentActionTypeID is TL type id of SendMessageUploadDocumentAction.
const SendMessageUploadDocumentActionTypeID = 0xaa0cd9e4

// Encode implements bin.Encoder.
func (s *SendMessageUploadDocumentAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadDocumentAction#aa0cd9e4 as nil")
	}
	b.PutID(SendMessageUploadDocumentActionTypeID)
	b.PutInt(s.Progress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadDocumentAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadDocumentAction#aa0cd9e4 to nil")
	}
	if err := b.ConsumeID(SendMessageUploadDocumentActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadDocumentAction#aa0cd9e4: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageUploadDocumentAction#aa0cd9e4: field progress: %w", err)
		}
		s.Progress = value
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadDocumentAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadDocumentAction.
var (
	_ bin.Encoder = &SendMessageUploadDocumentAction{}
	_ bin.Decoder = &SendMessageUploadDocumentAction{}

	_ SendMessageActionClass = &SendMessageUploadDocumentAction{}
)

// SendMessageGeoLocationAction represents TL type `sendMessageGeoLocationAction#176f8ba1`.
type SendMessageGeoLocationAction struct {
}

// SendMessageGeoLocationActionTypeID is TL type id of SendMessageGeoLocationAction.
const SendMessageGeoLocationActionTypeID = 0x176f8ba1

// Encode implements bin.Encoder.
func (s *SendMessageGeoLocationAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageGeoLocationAction#176f8ba1 as nil")
	}
	b.PutID(SendMessageGeoLocationActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageGeoLocationAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageGeoLocationAction#176f8ba1 to nil")
	}
	if err := b.ConsumeID(SendMessageGeoLocationActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageGeoLocationAction#176f8ba1: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageGeoLocationAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageGeoLocationAction.
var (
	_ bin.Encoder = &SendMessageGeoLocationAction{}
	_ bin.Decoder = &SendMessageGeoLocationAction{}

	_ SendMessageActionClass = &SendMessageGeoLocationAction{}
)

// SendMessageChooseContactAction represents TL type `sendMessageChooseContactAction#628cbc6f`.
type SendMessageChooseContactAction struct {
}

// SendMessageChooseContactActionTypeID is TL type id of SendMessageChooseContactAction.
const SendMessageChooseContactActionTypeID = 0x628cbc6f

// Encode implements bin.Encoder.
func (s *SendMessageChooseContactAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageChooseContactAction#628cbc6f as nil")
	}
	b.PutID(SendMessageChooseContactActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageChooseContactAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageChooseContactAction#628cbc6f to nil")
	}
	if err := b.ConsumeID(SendMessageChooseContactActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageChooseContactAction#628cbc6f: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageChooseContactAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageChooseContactAction.
var (
	_ bin.Encoder = &SendMessageChooseContactAction{}
	_ bin.Decoder = &SendMessageChooseContactAction{}

	_ SendMessageActionClass = &SendMessageChooseContactAction{}
)

// SendMessageGamePlayAction represents TL type `sendMessageGamePlayAction#dd6a8f48`.
type SendMessageGamePlayAction struct {
}

// SendMessageGamePlayActionTypeID is TL type id of SendMessageGamePlayAction.
const SendMessageGamePlayActionTypeID = 0xdd6a8f48

// Encode implements bin.Encoder.
func (s *SendMessageGamePlayAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageGamePlayAction#dd6a8f48 as nil")
	}
	b.PutID(SendMessageGamePlayActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageGamePlayAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageGamePlayAction#dd6a8f48 to nil")
	}
	if err := b.ConsumeID(SendMessageGamePlayActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageGamePlayAction#dd6a8f48: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageGamePlayAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageGamePlayAction.
var (
	_ bin.Encoder = &SendMessageGamePlayAction{}
	_ bin.Decoder = &SendMessageGamePlayAction{}

	_ SendMessageActionClass = &SendMessageGamePlayAction{}
)

// SendMessageRecordRoundAction represents TL type `sendMessageRecordRoundAction#88f27fbc`.
type SendMessageRecordRoundAction struct {
}

// SendMessageRecordRoundActionTypeID is TL type id of SendMessageRecordRoundAction.
const SendMessageRecordRoundActionTypeID = 0x88f27fbc

// Encode implements bin.Encoder.
func (s *SendMessageRecordRoundAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageRecordRoundAction#88f27fbc as nil")
	}
	b.PutID(SendMessageRecordRoundActionTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageRecordRoundAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageRecordRoundAction#88f27fbc to nil")
	}
	if err := b.ConsumeID(SendMessageRecordRoundActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageRecordRoundAction#88f27fbc: %w", err)
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageRecordRoundAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageRecordRoundAction.
var (
	_ bin.Encoder = &SendMessageRecordRoundAction{}
	_ bin.Decoder = &SendMessageRecordRoundAction{}

	_ SendMessageActionClass = &SendMessageRecordRoundAction{}
)

// SendMessageUploadRoundAction represents TL type `sendMessageUploadRoundAction#243e1c66`.
type SendMessageUploadRoundAction struct {
	// Progress field of SendMessageUploadRoundAction.
	Progress int
}

// SendMessageUploadRoundActionTypeID is TL type id of SendMessageUploadRoundAction.
const SendMessageUploadRoundActionTypeID = 0x243e1c66

// Encode implements bin.Encoder.
func (s *SendMessageUploadRoundAction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageUploadRoundAction#243e1c66 as nil")
	}
	b.PutID(SendMessageUploadRoundActionTypeID)
	b.PutInt(s.Progress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageUploadRoundAction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageUploadRoundAction#243e1c66 to nil")
	}
	if err := b.ConsumeID(SendMessageUploadRoundActionTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageUploadRoundAction#243e1c66: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageUploadRoundAction#243e1c66: field progress: %w", err)
		}
		s.Progress = value
	}
	return nil
}

// construct implements constructor of SendMessageActionClass.
func (s SendMessageUploadRoundAction) construct() SendMessageActionClass { return &s }

// Ensuring interfaces in compile-time for SendMessageUploadRoundAction.
var (
	_ bin.Encoder = &SendMessageUploadRoundAction{}
	_ bin.Decoder = &SendMessageUploadRoundAction{}

	_ SendMessageActionClass = &SendMessageUploadRoundAction{}
)

// SendMessageActionClass represents SendMessageAction generic type.
//
// Example:
//  g, err := DecodeSendMessageAction(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SendMessageTypingAction: // sendMessageTypingAction#16bf744e
//  case *SendMessageCancelAction: // sendMessageCancelAction#fd5ec8f5
//  case *SendMessageRecordVideoAction: // sendMessageRecordVideoAction#a187d66f
//  case *SendMessageUploadVideoAction: // sendMessageUploadVideoAction#e9763aec
//  case *SendMessageRecordAudioAction: // sendMessageRecordAudioAction#d52f73f7
//  case *SendMessageUploadAudioAction: // sendMessageUploadAudioAction#f351d7ab
//  case *SendMessageUploadPhotoAction: // sendMessageUploadPhotoAction#d1d34a26
//  case *SendMessageUploadDocumentAction: // sendMessageUploadDocumentAction#aa0cd9e4
//  case *SendMessageGeoLocationAction: // sendMessageGeoLocationAction#176f8ba1
//  case *SendMessageChooseContactAction: // sendMessageChooseContactAction#628cbc6f
//  case *SendMessageGamePlayAction: // sendMessageGamePlayAction#dd6a8f48
//  case *SendMessageRecordRoundAction: // sendMessageRecordRoundAction#88f27fbc
//  case *SendMessageUploadRoundAction: // sendMessageUploadRoundAction#243e1c66
//  default: panic(v)
//  }
type SendMessageActionClass interface {
	bin.Encoder
	bin.Decoder
	construct() SendMessageActionClass
}

// DecodeSendMessageAction implements binary de-serialization for SendMessageActionClass.
func DecodeSendMessageAction(buf *bin.Buffer) (SendMessageActionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SendMessageTypingActionTypeID:
		// Decoding sendMessageTypingAction#16bf744e.
		v := SendMessageTypingAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageCancelActionTypeID:
		// Decoding sendMessageCancelAction#fd5ec8f5.
		v := SendMessageCancelAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordVideoActionTypeID:
		// Decoding sendMessageRecordVideoAction#a187d66f.
		v := SendMessageRecordVideoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadVideoActionTypeID:
		// Decoding sendMessageUploadVideoAction#e9763aec.
		v := SendMessageUploadVideoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordAudioActionTypeID:
		// Decoding sendMessageRecordAudioAction#d52f73f7.
		v := SendMessageRecordAudioAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadAudioActionTypeID:
		// Decoding sendMessageUploadAudioAction#f351d7ab.
		v := SendMessageUploadAudioAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadPhotoActionTypeID:
		// Decoding sendMessageUploadPhotoAction#d1d34a26.
		v := SendMessageUploadPhotoAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadDocumentActionTypeID:
		// Decoding sendMessageUploadDocumentAction#aa0cd9e4.
		v := SendMessageUploadDocumentAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageGeoLocationActionTypeID:
		// Decoding sendMessageGeoLocationAction#176f8ba1.
		v := SendMessageGeoLocationAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageChooseContactActionTypeID:
		// Decoding sendMessageChooseContactAction#628cbc6f.
		v := SendMessageChooseContactAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageGamePlayActionTypeID:
		// Decoding sendMessageGamePlayAction#dd6a8f48.
		v := SendMessageGamePlayAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageRecordRoundActionTypeID:
		// Decoding sendMessageRecordRoundAction#88f27fbc.
		v := SendMessageRecordRoundAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	case SendMessageUploadRoundActionTypeID:
		// Decoding sendMessageUploadRoundAction#243e1c66.
		v := SendMessageUploadRoundAction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SendMessageActionClass: %w", bin.NewUnexpectedID(id))
	}
}

// SendMessageAction boxes the SendMessageActionClass providing a helper.
type SendMessageActionBox struct {
	SendMessageAction SendMessageActionClass
}

// Decode implements bin.Decoder for SendMessageActionBox.
func (b *SendMessageActionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SendMessageActionBox to nil")
	}
	v, err := DecodeSendMessageAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SendMessageAction = v
	return nil
}

// Encode implements bin.Encode for SendMessageActionBox.
func (b *SendMessageActionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SendMessageAction == nil {
		return fmt.Errorf("unable to encode SendMessageActionClass as nil")
	}
	return b.SendMessageAction.Encode(buf)
}
