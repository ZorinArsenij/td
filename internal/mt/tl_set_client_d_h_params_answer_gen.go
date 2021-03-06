// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// DhGenOk represents TL type `dh_gen_ok#3bcbf734`.
type DhGenOk struct {
	// Nonce field of DhGenOk.
	Nonce bin.Int128
	// ServerNonce field of DhGenOk.
	ServerNonce bin.Int128
	// NewNonceHash1 field of DhGenOk.
	NewNonceHash1 bin.Int128
}

// DhGenOkTypeID is TL type id of DhGenOk.
const DhGenOkTypeID = 0x3bcbf734

func (d *DhGenOk) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Nonce == bin.Int128{}) {
		return false
	}
	if !(d.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(d.NewNonceHash1 == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DhGenOk) String() string {
	if d == nil {
		return "DhGenOk(nil)"
	}
	type Alias DhGenOk
	return fmt.Sprintf("DhGenOk%+v", Alias(*d))
}

// FillFrom fills DhGenOk from given interface.
func (d *DhGenOk) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonceHash1() (value bin.Int128)
}) {
	d.Nonce = from.GetNonce()
	d.ServerNonce = from.GetServerNonce()
	d.NewNonceHash1 = from.GetNewNonceHash1()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DhGenOk) TypeID() uint32 {
	return DhGenOkTypeID
}

// TypeName returns name of type in TL schema.
func (*DhGenOk) TypeName() string {
	return "dh_gen_ok"
}

// TypeInfo returns info about TL type.
func (d *DhGenOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dh_gen_ok",
		ID:   DhGenOkTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "NewNonceHash1",
			SchemaName: "new_nonce_hash1",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DhGenOk) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_ok#3bcbf734 as nil")
	}
	b.PutID(DhGenOkTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash1)
	return nil
}

// GetNonce returns value of Nonce field.
func (d *DhGenOk) GetNonce() (value bin.Int128) {
	return d.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (d *DhGenOk) GetServerNonce() (value bin.Int128) {
	return d.ServerNonce
}

// GetNewNonceHash1 returns value of NewNonceHash1 field.
func (d *DhGenOk) GetNewNonceHash1() (value bin.Int128) {
	return d.NewNonceHash1
}

// Decode implements bin.Decoder.
func (d *DhGenOk) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_ok#3bcbf734 to nil")
	}
	if err := b.ConsumeID(DhGenOkTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field new_nonce_hash1: %w", err)
		}
		d.NewNonceHash1 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswerClass.
func (d DhGenOk) construct() SetClientDHParamsAnswerClass { return &d }

// Ensuring interfaces in compile-time for DhGenOk.
var (
	_ bin.Encoder = &DhGenOk{}
	_ bin.Decoder = &DhGenOk{}

	_ SetClientDHParamsAnswerClass = &DhGenOk{}
)

// DhGenRetry represents TL type `dh_gen_retry#46dc1fb9`.
type DhGenRetry struct {
	// Nonce field of DhGenRetry.
	Nonce bin.Int128
	// ServerNonce field of DhGenRetry.
	ServerNonce bin.Int128
	// NewNonceHash2 field of DhGenRetry.
	NewNonceHash2 bin.Int128
}

// DhGenRetryTypeID is TL type id of DhGenRetry.
const DhGenRetryTypeID = 0x46dc1fb9

func (d *DhGenRetry) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Nonce == bin.Int128{}) {
		return false
	}
	if !(d.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(d.NewNonceHash2 == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DhGenRetry) String() string {
	if d == nil {
		return "DhGenRetry(nil)"
	}
	type Alias DhGenRetry
	return fmt.Sprintf("DhGenRetry%+v", Alias(*d))
}

// FillFrom fills DhGenRetry from given interface.
func (d *DhGenRetry) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonceHash2() (value bin.Int128)
}) {
	d.Nonce = from.GetNonce()
	d.ServerNonce = from.GetServerNonce()
	d.NewNonceHash2 = from.GetNewNonceHash2()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DhGenRetry) TypeID() uint32 {
	return DhGenRetryTypeID
}

// TypeName returns name of type in TL schema.
func (*DhGenRetry) TypeName() string {
	return "dh_gen_retry"
}

// TypeInfo returns info about TL type.
func (d *DhGenRetry) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dh_gen_retry",
		ID:   DhGenRetryTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "NewNonceHash2",
			SchemaName: "new_nonce_hash2",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DhGenRetry) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_retry#46dc1fb9 as nil")
	}
	b.PutID(DhGenRetryTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash2)
	return nil
}

// GetNonce returns value of Nonce field.
func (d *DhGenRetry) GetNonce() (value bin.Int128) {
	return d.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (d *DhGenRetry) GetServerNonce() (value bin.Int128) {
	return d.ServerNonce
}

// GetNewNonceHash2 returns value of NewNonceHash2 field.
func (d *DhGenRetry) GetNewNonceHash2() (value bin.Int128) {
	return d.NewNonceHash2
}

// Decode implements bin.Decoder.
func (d *DhGenRetry) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_retry#46dc1fb9 to nil")
	}
	if err := b.ConsumeID(DhGenRetryTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field new_nonce_hash2: %w", err)
		}
		d.NewNonceHash2 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswerClass.
func (d DhGenRetry) construct() SetClientDHParamsAnswerClass { return &d }

// Ensuring interfaces in compile-time for DhGenRetry.
var (
	_ bin.Encoder = &DhGenRetry{}
	_ bin.Decoder = &DhGenRetry{}

	_ SetClientDHParamsAnswerClass = &DhGenRetry{}
)

// DhGenFail represents TL type `dh_gen_fail#a69dae02`.
type DhGenFail struct {
	// Nonce field of DhGenFail.
	Nonce bin.Int128
	// ServerNonce field of DhGenFail.
	ServerNonce bin.Int128
	// NewNonceHash3 field of DhGenFail.
	NewNonceHash3 bin.Int128
}

// DhGenFailTypeID is TL type id of DhGenFail.
const DhGenFailTypeID = 0xa69dae02

func (d *DhGenFail) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Nonce == bin.Int128{}) {
		return false
	}
	if !(d.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(d.NewNonceHash3 == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DhGenFail) String() string {
	if d == nil {
		return "DhGenFail(nil)"
	}
	type Alias DhGenFail
	return fmt.Sprintf("DhGenFail%+v", Alias(*d))
}

// FillFrom fills DhGenFail from given interface.
func (d *DhGenFail) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonceHash3() (value bin.Int128)
}) {
	d.Nonce = from.GetNonce()
	d.ServerNonce = from.GetServerNonce()
	d.NewNonceHash3 = from.GetNewNonceHash3()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DhGenFail) TypeID() uint32 {
	return DhGenFailTypeID
}

// TypeName returns name of type in TL schema.
func (*DhGenFail) TypeName() string {
	return "dh_gen_fail"
}

// TypeInfo returns info about TL type.
func (d *DhGenFail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dh_gen_fail",
		ID:   DhGenFailTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "NewNonceHash3",
			SchemaName: "new_nonce_hash3",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DhGenFail) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_fail#a69dae02 as nil")
	}
	b.PutID(DhGenFailTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash3)
	return nil
}

// GetNonce returns value of Nonce field.
func (d *DhGenFail) GetNonce() (value bin.Int128) {
	return d.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (d *DhGenFail) GetServerNonce() (value bin.Int128) {
	return d.ServerNonce
}

// GetNewNonceHash3 returns value of NewNonceHash3 field.
func (d *DhGenFail) GetNewNonceHash3() (value bin.Int128) {
	return d.NewNonceHash3
}

// Decode implements bin.Decoder.
func (d *DhGenFail) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_fail#a69dae02 to nil")
	}
	if err := b.ConsumeID(DhGenFailTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field new_nonce_hash3: %w", err)
		}
		d.NewNonceHash3 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswerClass.
func (d DhGenFail) construct() SetClientDHParamsAnswerClass { return &d }

// Ensuring interfaces in compile-time for DhGenFail.
var (
	_ bin.Encoder = &DhGenFail{}
	_ bin.Decoder = &DhGenFail{}

	_ SetClientDHParamsAnswerClass = &DhGenFail{}
)

// SetClientDHParamsAnswerClass represents Set_client_DH_params_answer generic type.
//
// Example:
//  g, err := mt.DecodeSetClientDHParamsAnswer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.DhGenOk: // dh_gen_ok#3bcbf734
//  case *mt.DhGenRetry: // dh_gen_retry#46dc1fb9
//  case *mt.DhGenFail: // dh_gen_fail#a69dae02
//  default: panic(v)
//  }
type SetClientDHParamsAnswerClass interface {
	bin.Encoder
	bin.Decoder
	construct() SetClientDHParamsAnswerClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Nonce field of DhGenOk.
	GetNonce() (value bin.Int128)
	// ServerNonce field of DhGenOk.
	GetServerNonce() (value bin.Int128)
}

// DecodeSetClientDHParamsAnswer implements binary de-serialization for SetClientDHParamsAnswerClass.
func DecodeSetClientDHParamsAnswer(buf *bin.Buffer) (SetClientDHParamsAnswerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DhGenOkTypeID:
		// Decoding dh_gen_ok#3bcbf734.
		v := DhGenOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswerClass: %w", err)
		}
		return &v, nil
	case DhGenRetryTypeID:
		// Decoding dh_gen_retry#46dc1fb9.
		v := DhGenRetry{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswerClass: %w", err)
		}
		return &v, nil
	case DhGenFailTypeID:
		// Decoding dh_gen_fail#a69dae02.
		v := DhGenFail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswerClass: %w", bin.NewUnexpectedID(id))
	}
}

// SetClientDHParamsAnswer boxes the SetClientDHParamsAnswerClass providing a helper.
type SetClientDHParamsAnswerBox struct {
	Set_client_DH_params_answer SetClientDHParamsAnswerClass
}

// Decode implements bin.Decoder for SetClientDHParamsAnswerBox.
func (b *SetClientDHParamsAnswerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SetClientDHParamsAnswerBox to nil")
	}
	v, err := DecodeSetClientDHParamsAnswer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Set_client_DH_params_answer = v
	return nil
}

// Encode implements bin.Encode for SetClientDHParamsAnswerBox.
func (b *SetClientDHParamsAnswerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Set_client_DH_params_answer == nil {
		return fmt.Errorf("unable to encode SetClientDHParamsAnswerClass as nil")
	}
	return b.Set_client_DH_params_answer.Encode(buf)
}

// SetClientDHParamsAnswerClassArray is adapter for slice of SetClientDHParamsAnswerClass.
type SetClientDHParamsAnswerClassArray []SetClientDHParamsAnswerClass

// Sort sorts slice of SetClientDHParamsAnswerClass.
func (s SetClientDHParamsAnswerClassArray) Sort(less func(a, b SetClientDHParamsAnswerClass) bool) SetClientDHParamsAnswerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SetClientDHParamsAnswerClass.
func (s SetClientDHParamsAnswerClassArray) SortStable(less func(a, b SetClientDHParamsAnswerClass) bool) SetClientDHParamsAnswerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SetClientDHParamsAnswerClass.
func (s SetClientDHParamsAnswerClassArray) Retain(keep func(x SetClientDHParamsAnswerClass) bool) SetClientDHParamsAnswerClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SetClientDHParamsAnswerClassArray) First() (v SetClientDHParamsAnswerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SetClientDHParamsAnswerClassArray) Last() (v SetClientDHParamsAnswerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SetClientDHParamsAnswerClassArray) PopFirst() (v SetClientDHParamsAnswerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SetClientDHParamsAnswerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SetClientDHParamsAnswerClassArray) Pop() (v SetClientDHParamsAnswerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsDhGenOk returns copy with only DhGenOk constructors.
func (s SetClientDHParamsAnswerClassArray) AsDhGenOk() (to DhGenOkArray) {
	for _, elem := range s {
		value, ok := elem.(*DhGenOk)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDhGenRetry returns copy with only DhGenRetry constructors.
func (s SetClientDHParamsAnswerClassArray) AsDhGenRetry() (to DhGenRetryArray) {
	for _, elem := range s {
		value, ok := elem.(*DhGenRetry)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDhGenFail returns copy with only DhGenFail constructors.
func (s SetClientDHParamsAnswerClassArray) AsDhGenFail() (to DhGenFailArray) {
	for _, elem := range s {
		value, ok := elem.(*DhGenFail)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// DhGenOkArray is adapter for slice of DhGenOk.
type DhGenOkArray []DhGenOk

// Sort sorts slice of DhGenOk.
func (s DhGenOkArray) Sort(less func(a, b DhGenOk) bool) DhGenOkArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DhGenOk.
func (s DhGenOkArray) SortStable(less func(a, b DhGenOk) bool) DhGenOkArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DhGenOk.
func (s DhGenOkArray) Retain(keep func(x DhGenOk) bool) DhGenOkArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DhGenOkArray) First() (v DhGenOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DhGenOkArray) Last() (v DhGenOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DhGenOkArray) PopFirst() (v DhGenOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DhGenOk
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DhGenOkArray) Pop() (v DhGenOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DhGenRetryArray is adapter for slice of DhGenRetry.
type DhGenRetryArray []DhGenRetry

// Sort sorts slice of DhGenRetry.
func (s DhGenRetryArray) Sort(less func(a, b DhGenRetry) bool) DhGenRetryArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DhGenRetry.
func (s DhGenRetryArray) SortStable(less func(a, b DhGenRetry) bool) DhGenRetryArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DhGenRetry.
func (s DhGenRetryArray) Retain(keep func(x DhGenRetry) bool) DhGenRetryArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DhGenRetryArray) First() (v DhGenRetry, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DhGenRetryArray) Last() (v DhGenRetry, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DhGenRetryArray) PopFirst() (v DhGenRetry, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DhGenRetry
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DhGenRetryArray) Pop() (v DhGenRetry, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DhGenFailArray is adapter for slice of DhGenFail.
type DhGenFailArray []DhGenFail

// Sort sorts slice of DhGenFail.
func (s DhGenFailArray) Sort(less func(a, b DhGenFail) bool) DhGenFailArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DhGenFail.
func (s DhGenFailArray) SortStable(less func(a, b DhGenFail) bool) DhGenFailArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DhGenFail.
func (s DhGenFailArray) Retain(keep func(x DhGenFail) bool) DhGenFailArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DhGenFailArray) First() (v DhGenFail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DhGenFailArray) Last() (v DhGenFail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DhGenFailArray) PopFirst() (v DhGenFail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DhGenFail
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DhGenFailArray) Pop() (v DhGenFail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
