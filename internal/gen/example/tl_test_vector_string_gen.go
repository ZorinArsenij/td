// Code generated by gotdgen, DO NOT EDIT.

package td

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

// TestVectorString represents TL type `testVectorString#5d6f85bc`.
//
// See https://localhost:80/doc/constructor/testVectorString for reference.
type TestVectorString struct {
	// Vector of strings
	Value []string
}

// TestVectorStringTypeID is TL type id of TestVectorString.
const TestVectorStringTypeID = 0x5d6f85bc

func (t *TestVectorString) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorString) String() string {
	if t == nil {
		return "TestVectorString(nil)"
	}
	type Alias TestVectorString
	return fmt.Sprintf("TestVectorString%+v", Alias(*t))
}

// FillFrom fills TestVectorString from given interface.
func (t *TestVectorString) FillFrom(from interface {
	GetValue() (value []string)
}) {
	t.Value = from.GetValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestVectorString) TypeID() uint32 {
	return TestVectorStringTypeID
}

// TypeName returns name of type in TL schema.
func (*TestVectorString) TypeName() string {
	return "testVectorString"
}

// TypeInfo returns info about TL type.
func (t *TestVectorString) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testVectorString",
		ID:   TestVectorStringTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestVectorString) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorString#5d6f85bc as nil")
	}
	b.PutID(TestVectorStringTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutString(v)
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorString) GetValue() (value []string) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorString) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorString#5d6f85bc to nil")
	}
	if err := b.ConsumeID(TestVectorStringTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorString#5d6f85bc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorString#5d6f85bc: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorString#5d6f85bc: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorString.
var (
	_ bin.Encoder = &TestVectorString{}
	_ bin.Decoder = &TestVectorString{}
)
