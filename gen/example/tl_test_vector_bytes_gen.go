// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// TestVectorBytes represents TL type `testVectorBytes#a590fb25`.
type TestVectorBytes struct {
	// Value field of TestVectorBytes.
	Value [][]byte
}

// TestVectorBytesTypeID is TL type id of TestVectorBytes.
const TestVectorBytesTypeID = 0xa590fb25

// Encode implements bin.Encoder.
func (t *TestVectorBytes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorBytes#a590fb25 as nil")
	}
	b.PutID(TestVectorBytesTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutBytes(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorBytes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorBytes#a590fb25 to nil")
	}
	if err := b.ConsumeID(TestVectorBytesTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorBytes#a590fb25: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorBytes#a590fb25: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorBytes#a590fb25: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorBytes.
var (
	_ bin.Encoder = &TestVectorBytes{}
	_ bin.Decoder = &TestVectorBytes{}
)
