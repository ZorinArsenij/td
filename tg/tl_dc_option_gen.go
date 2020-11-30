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

// DcOption represents TL type `dcOption#18b7a10d`.
type DcOption struct {
	// Flags field of DcOption.
	Flags bin.Fields
	// Ipv6 field of DcOption.
	Ipv6 bool
	// MediaOnly field of DcOption.
	MediaOnly bool
	// TcpoOnly field of DcOption.
	TcpoOnly bool
	// CDN field of DcOption.
	CDN bool
	// Static field of DcOption.
	Static bool
	// ID field of DcOption.
	ID int
	// IPAddress field of DcOption.
	IPAddress string
	// Port field of DcOption.
	Port int
	// Secret field of DcOption.
	//
	// Use SetSecret and GetSecret helpers.
	Secret []byte
}

// DcOptionTypeID is TL type id of DcOption.
const DcOptionTypeID = 0x18b7a10d

// Encode implements bin.Encoder.
func (d *DcOption) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dcOption#18b7a10d as nil")
	}
	b.PutID(DcOptionTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dcOption#18b7a10d: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.IPAddress)
	b.PutInt(d.Port)
	if d.Flags.Has(10) {
		b.PutBytes(d.Secret)
	}
	return nil
}

// SetIpv6 sets value of Ipv6 conditional field.
func (d *DcOption) SetIpv6(value bool) {
	if value {
		d.Flags.Set(0)
	} else {
		d.Flags.Unset(0)
	}
}

// SetMediaOnly sets value of MediaOnly conditional field.
func (d *DcOption) SetMediaOnly(value bool) {
	if value {
		d.Flags.Set(1)
	} else {
		d.Flags.Unset(1)
	}
}

// SetTcpoOnly sets value of TcpoOnly conditional field.
func (d *DcOption) SetTcpoOnly(value bool) {
	if value {
		d.Flags.Set(2)
	} else {
		d.Flags.Unset(2)
	}
}

// SetCDN sets value of CDN conditional field.
func (d *DcOption) SetCDN(value bool) {
	if value {
		d.Flags.Set(3)
	} else {
		d.Flags.Unset(3)
	}
}

// SetStatic sets value of Static conditional field.
func (d *DcOption) SetStatic(value bool) {
	if value {
		d.Flags.Set(4)
	} else {
		d.Flags.Unset(4)
	}
}

// SetSecret sets value of Secret conditional field.
func (d *DcOption) SetSecret(value []byte) {
	d.Flags.Set(10)
	d.Secret = value
}

// GetSecret returns value of Secret conditional field and
// boolean which is true if field was set.
func (d *DcOption) GetSecret() (value []byte, ok bool) {
	if !d.Flags.Has(10) {
		return value, false
	}
	return d.Secret, true
}

// Decode implements bin.Decoder.
func (d *DcOption) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dcOption#18b7a10d to nil")
	}
	if err := b.ConsumeID(DcOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode dcOption#18b7a10d: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field flags: %w", err)
		}
	}
	d.Ipv6 = d.Flags.Has(0)
	d.MediaOnly = d.Flags.Has(1)
	d.TcpoOnly = d.Flags.Has(2)
	d.CDN = d.Flags.Has(3)
	d.Static = d.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field ip_address: %w", err)
		}
		d.IPAddress = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field port: %w", err)
		}
		d.Port = value
	}
	if d.Flags.Has(10) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field secret: %w", err)
		}
		d.Secret = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DcOption.
var (
	_ bin.Encoder = &DcOption{}
	_ bin.Decoder = &DcOption{}
)
