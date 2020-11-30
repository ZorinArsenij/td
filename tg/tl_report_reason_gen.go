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

// InputReportReasonSpam represents TL type `inputReportReasonSpam#58dbcab8`.
type InputReportReasonSpam struct {
}

// InputReportReasonSpamTypeID is TL type id of InputReportReasonSpam.
const InputReportReasonSpamTypeID = 0x58dbcab8

// Encode implements bin.Encoder.
func (i *InputReportReasonSpam) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonSpam#58dbcab8 as nil")
	}
	b.PutID(InputReportReasonSpamTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonSpam) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonSpam#58dbcab8 to nil")
	}
	if err := b.ConsumeID(InputReportReasonSpamTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonSpam#58dbcab8: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonSpam) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonSpam.
var (
	_ bin.Encoder = &InputReportReasonSpam{}
	_ bin.Decoder = &InputReportReasonSpam{}

	_ ReportReasonClass = &InputReportReasonSpam{}
)

// InputReportReasonViolence represents TL type `inputReportReasonViolence#1e22c78d`.
type InputReportReasonViolence struct {
}

// InputReportReasonViolenceTypeID is TL type id of InputReportReasonViolence.
const InputReportReasonViolenceTypeID = 0x1e22c78d

// Encode implements bin.Encoder.
func (i *InputReportReasonViolence) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonViolence#1e22c78d as nil")
	}
	b.PutID(InputReportReasonViolenceTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonViolence) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonViolence#1e22c78d to nil")
	}
	if err := b.ConsumeID(InputReportReasonViolenceTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonViolence#1e22c78d: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonViolence) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonViolence.
var (
	_ bin.Encoder = &InputReportReasonViolence{}
	_ bin.Decoder = &InputReportReasonViolence{}

	_ ReportReasonClass = &InputReportReasonViolence{}
)

// InputReportReasonPornography represents TL type `inputReportReasonPornography#2e59d922`.
type InputReportReasonPornography struct {
}

// InputReportReasonPornographyTypeID is TL type id of InputReportReasonPornography.
const InputReportReasonPornographyTypeID = 0x2e59d922

// Encode implements bin.Encoder.
func (i *InputReportReasonPornography) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonPornography#2e59d922 as nil")
	}
	b.PutID(InputReportReasonPornographyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonPornography) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonPornography#2e59d922 to nil")
	}
	if err := b.ConsumeID(InputReportReasonPornographyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonPornography#2e59d922: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonPornography) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonPornography.
var (
	_ bin.Encoder = &InputReportReasonPornography{}
	_ bin.Decoder = &InputReportReasonPornography{}

	_ ReportReasonClass = &InputReportReasonPornography{}
)

// InputReportReasonChildAbuse represents TL type `inputReportReasonChildAbuse#adf44ee3`.
type InputReportReasonChildAbuse struct {
}

// InputReportReasonChildAbuseTypeID is TL type id of InputReportReasonChildAbuse.
const InputReportReasonChildAbuseTypeID = 0xadf44ee3

// Encode implements bin.Encoder.
func (i *InputReportReasonChildAbuse) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonChildAbuse#adf44ee3 as nil")
	}
	b.PutID(InputReportReasonChildAbuseTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonChildAbuse) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonChildAbuse#adf44ee3 to nil")
	}
	if err := b.ConsumeID(InputReportReasonChildAbuseTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonChildAbuse#adf44ee3: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonChildAbuse) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonChildAbuse.
var (
	_ bin.Encoder = &InputReportReasonChildAbuse{}
	_ bin.Decoder = &InputReportReasonChildAbuse{}

	_ ReportReasonClass = &InputReportReasonChildAbuse{}
)

// InputReportReasonOther represents TL type `inputReportReasonOther#e1746d0a`.
type InputReportReasonOther struct {
	// Text field of InputReportReasonOther.
	Text string
}

// InputReportReasonOtherTypeID is TL type id of InputReportReasonOther.
const InputReportReasonOtherTypeID = 0xe1746d0a

// Encode implements bin.Encoder.
func (i *InputReportReasonOther) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonOther#e1746d0a as nil")
	}
	b.PutID(InputReportReasonOtherTypeID)
	b.PutString(i.Text)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonOther) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonOther#e1746d0a to nil")
	}
	if err := b.ConsumeID(InputReportReasonOtherTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonOther#e1746d0a: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputReportReasonOther#e1746d0a: field text: %w", err)
		}
		i.Text = value
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonOther) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonOther.
var (
	_ bin.Encoder = &InputReportReasonOther{}
	_ bin.Decoder = &InputReportReasonOther{}

	_ ReportReasonClass = &InputReportReasonOther{}
)

// InputReportReasonCopyright represents TL type `inputReportReasonCopyright#9b89f93a`.
type InputReportReasonCopyright struct {
}

// InputReportReasonCopyrightTypeID is TL type id of InputReportReasonCopyright.
const InputReportReasonCopyrightTypeID = 0x9b89f93a

// Encode implements bin.Encoder.
func (i *InputReportReasonCopyright) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonCopyright#9b89f93a as nil")
	}
	b.PutID(InputReportReasonCopyrightTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonCopyright) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonCopyright#9b89f93a to nil")
	}
	if err := b.ConsumeID(InputReportReasonCopyrightTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonCopyright#9b89f93a: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonCopyright) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonCopyright.
var (
	_ bin.Encoder = &InputReportReasonCopyright{}
	_ bin.Decoder = &InputReportReasonCopyright{}

	_ ReportReasonClass = &InputReportReasonCopyright{}
)

// InputReportReasonGeoIrrelevant represents TL type `inputReportReasonGeoIrrelevant#dbd4feed`.
type InputReportReasonGeoIrrelevant struct {
}

// InputReportReasonGeoIrrelevantTypeID is TL type id of InputReportReasonGeoIrrelevant.
const InputReportReasonGeoIrrelevantTypeID = 0xdbd4feed

// Encode implements bin.Encoder.
func (i *InputReportReasonGeoIrrelevant) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonGeoIrrelevant#dbd4feed as nil")
	}
	b.PutID(InputReportReasonGeoIrrelevantTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonGeoIrrelevant) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonGeoIrrelevant#dbd4feed to nil")
	}
	if err := b.ConsumeID(InputReportReasonGeoIrrelevantTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonGeoIrrelevant#dbd4feed: %w", err)
	}
	return nil
}

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonGeoIrrelevant) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonGeoIrrelevant.
var (
	_ bin.Encoder = &InputReportReasonGeoIrrelevant{}
	_ bin.Decoder = &InputReportReasonGeoIrrelevant{}

	_ ReportReasonClass = &InputReportReasonGeoIrrelevant{}
)

// ReportReasonClass represents ReportReason generic type.
//
// Example:
//  g, err := DecodeReportReason(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputReportReasonSpam: // inputReportReasonSpam#58dbcab8
//  case *InputReportReasonViolence: // inputReportReasonViolence#1e22c78d
//  case *InputReportReasonPornography: // inputReportReasonPornography#2e59d922
//  case *InputReportReasonChildAbuse: // inputReportReasonChildAbuse#adf44ee3
//  case *InputReportReasonOther: // inputReportReasonOther#e1746d0a
//  case *InputReportReasonCopyright: // inputReportReasonCopyright#9b89f93a
//  case *InputReportReasonGeoIrrelevant: // inputReportReasonGeoIrrelevant#dbd4feed
//  default: panic(v)
//  }
type ReportReasonClass interface {
	bin.Encoder
	bin.Decoder
	construct() ReportReasonClass
}

// DecodeReportReason implements binary de-serialization for ReportReasonClass.
func DecodeReportReason(buf *bin.Buffer) (ReportReasonClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputReportReasonSpamTypeID:
		// Decoding inputReportReasonSpam#58dbcab8.
		v := InputReportReasonSpam{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonViolenceTypeID:
		// Decoding inputReportReasonViolence#1e22c78d.
		v := InputReportReasonViolence{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonPornographyTypeID:
		// Decoding inputReportReasonPornography#2e59d922.
		v := InputReportReasonPornography{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonChildAbuseTypeID:
		// Decoding inputReportReasonChildAbuse#adf44ee3.
		v := InputReportReasonChildAbuse{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonOtherTypeID:
		// Decoding inputReportReasonOther#e1746d0a.
		v := InputReportReasonOther{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonCopyrightTypeID:
		// Decoding inputReportReasonCopyright#9b89f93a.
		v := InputReportReasonCopyright{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonGeoIrrelevantTypeID:
		// Decoding inputReportReasonGeoIrrelevant#dbd4feed.
		v := InputReportReasonGeoIrrelevant{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReportReason boxes the ReportReasonClass providing a helper.
type ReportReasonBox struct {
	ReportReason ReportReasonClass
}

// Decode implements bin.Decoder for ReportReasonBox.
func (b *ReportReasonBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReportReasonBox to nil")
	}
	v, err := DecodeReportReason(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReportReason = v
	return nil
}

// Encode implements bin.Encode for ReportReasonBox.
func (b *ReportReasonBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReportReason == nil {
		return fmt.Errorf("unable to encode ReportReasonClass as nil")
	}
	return b.ReportReason.Encode(buf)
}
