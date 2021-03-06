// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// InputGroupCall represents TL type `inputGroupCall#d8aa840f`.
//
// See https://core.telegram.org/constructor/inputGroupCall for reference.
type InputGroupCall struct {
	// ID field of InputGroupCall.
	ID int64
	// AccessHash field of InputGroupCall.
	AccessHash int64
}

// InputGroupCallTypeID is TL type id of InputGroupCall.
const InputGroupCallTypeID = 0xd8aa840f

func (i *InputGroupCall) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputGroupCall) String() string {
	if i == nil {
		return "InputGroupCall(nil)"
	}
	type Alias InputGroupCall
	return fmt.Sprintf("InputGroupCall%+v", Alias(*i))
}

// FillFrom fills InputGroupCall from given interface.
func (i *InputGroupCall) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputGroupCall) TypeID() uint32 {
	return InputGroupCallTypeID
}

// TypeName returns name of type in TL schema.
func (*InputGroupCall) TypeName() string {
	return "inputGroupCall"
}

// TypeInfo returns info about TL type.
func (i *InputGroupCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputGroupCall",
		ID:   InputGroupCallTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputGroupCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGroupCall#d8aa840f as nil")
	}
	b.PutID(InputGroupCallTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetID returns value of ID field.
func (i *InputGroupCall) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputGroupCall) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputGroupCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGroupCall#d8aa840f to nil")
	}
	if err := b.ConsumeID(InputGroupCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGroupCall#d8aa840f: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGroupCall#d8aa840f: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGroupCall#d8aa840f: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputGroupCall.
var (
	_ bin.Encoder = &InputGroupCall{}
	_ bin.Decoder = &InputGroupCall{}
)
