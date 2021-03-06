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

// ChatParticipantsForbidden represents TL type `chatParticipantsForbidden#fc900c2b`.
// Info on members is unavailable
//
// See https://core.telegram.org/constructor/chatParticipantsForbidden for reference.
type ChatParticipantsForbidden struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Group ID
	ChatID int
	// Info about the group membership of the current user
	//
	// Use SetSelfParticipant and GetSelfParticipant helpers.
	SelfParticipant ChatParticipantClass
}

// ChatParticipantsForbiddenTypeID is TL type id of ChatParticipantsForbidden.
const ChatParticipantsForbiddenTypeID = 0xfc900c2b

func (c *ChatParticipantsForbidden) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.SelfParticipant == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantsForbidden) String() string {
	if c == nil {
		return "ChatParticipantsForbidden(nil)"
	}
	type Alias ChatParticipantsForbidden
	return fmt.Sprintf("ChatParticipantsForbidden%+v", Alias(*c))
}

// FillFrom fills ChatParticipantsForbidden from given interface.
func (c *ChatParticipantsForbidden) FillFrom(from interface {
	GetChatID() (value int)
	GetSelfParticipant() (value ChatParticipantClass, ok bool)
}) {
	c.ChatID = from.GetChatID()
	if val, ok := from.GetSelfParticipant(); ok {
		c.SelfParticipant = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantsForbidden) TypeID() uint32 {
	return ChatParticipantsForbiddenTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantsForbidden) TypeName() string {
	return "chatParticipantsForbidden"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantsForbidden) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantsForbidden",
		ID:   ChatParticipantsForbiddenTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "SelfParticipant",
			SchemaName: "self_participant",
			Null:       !c.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipantsForbidden) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantsForbidden#fc900c2b as nil")
	}
	b.PutID(ChatParticipantsForbiddenTypeID)
	if !(c.SelfParticipant == nil) {
		c.Flags.Set(0)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatParticipantsForbidden#fc900c2b: field flags: %w", err)
	}
	b.PutInt(c.ChatID)
	if c.Flags.Has(0) {
		if c.SelfParticipant == nil {
			return fmt.Errorf("unable to encode chatParticipantsForbidden#fc900c2b: field self_participant is nil")
		}
		if err := c.SelfParticipant.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatParticipantsForbidden#fc900c2b: field self_participant: %w", err)
		}
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (c *ChatParticipantsForbidden) GetChatID() (value int) {
	return c.ChatID
}

// SetSelfParticipant sets value of SelfParticipant conditional field.
func (c *ChatParticipantsForbidden) SetSelfParticipant(value ChatParticipantClass) {
	c.Flags.Set(0)
	c.SelfParticipant = value
}

// GetSelfParticipant returns value of SelfParticipant conditional field and
// boolean which is true if field was set.
func (c *ChatParticipantsForbidden) GetSelfParticipant() (value ChatParticipantClass, ok bool) {
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.SelfParticipant, true
}

// Decode implements bin.Decoder.
func (c *ChatParticipantsForbidden) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantsForbidden#fc900c2b to nil")
	}
	if err := b.ConsumeID(ChatParticipantsForbiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantsForbidden#fc900c2b: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#fc900c2b: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#fc900c2b: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	if c.Flags.Has(0) {
		value, err := DecodeChatParticipant(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#fc900c2b: field self_participant: %w", err)
		}
		c.SelfParticipant = value
	}
	return nil
}

// construct implements constructor of ChatParticipantsClass.
func (c ChatParticipantsForbidden) construct() ChatParticipantsClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantsForbidden.
var (
	_ bin.Encoder = &ChatParticipantsForbidden{}
	_ bin.Decoder = &ChatParticipantsForbidden{}

	_ ChatParticipantsClass = &ChatParticipantsForbidden{}
)

// ChatParticipants represents TL type `chatParticipants#3f460fed`.
// Group members.
//
// See https://core.telegram.org/constructor/chatParticipants for reference.
type ChatParticipants struct {
	// Group identifier
	ChatID int
	// List of group members
	Participants []ChatParticipantClass
	// Group version number
	Version int
}

// ChatParticipantsTypeID is TL type id of ChatParticipants.
const ChatParticipantsTypeID = 0x3f460fed

func (c *ChatParticipants) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.Participants == nil) {
		return false
	}
	if !(c.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipants) String() string {
	if c == nil {
		return "ChatParticipants(nil)"
	}
	type Alias ChatParticipants
	return fmt.Sprintf("ChatParticipants%+v", Alias(*c))
}

// FillFrom fills ChatParticipants from given interface.
func (c *ChatParticipants) FillFrom(from interface {
	GetChatID() (value int)
	GetParticipants() (value []ChatParticipantClass)
	GetVersion() (value int)
}) {
	c.ChatID = from.GetChatID()
	c.Participants = from.GetParticipants()
	c.Version = from.GetVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipants) TypeID() uint32 {
	return ChatParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipants) TypeName() string {
	return "chatParticipants"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipants",
		ID:   ChatParticipantsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipants) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipants#3f460fed as nil")
	}
	b.PutID(ChatParticipantsTypeID)
	b.PutInt(c.ChatID)
	b.PutVectorHeader(len(c.Participants))
	for idx, v := range c.Participants {
		if v == nil {
			return fmt.Errorf("unable to encode chatParticipants#3f460fed: field participants element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatParticipants#3f460fed: field participants element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.Version)
	return nil
}

// GetChatID returns value of ChatID field.
func (c *ChatParticipants) GetChatID() (value int) {
	return c.ChatID
}

// GetParticipants returns value of Participants field.
func (c *ChatParticipants) GetParticipants() (value []ChatParticipantClass) {
	return c.Participants
}

// MapParticipants returns field Participants wrapped in ChatParticipantClassArray helper.
func (c *ChatParticipants) MapParticipants() (value ChatParticipantClassArray) {
	return ChatParticipantClassArray(c.Participants)
}

// GetVersion returns value of Version field.
func (c *ChatParticipants) GetVersion() (value int) {
	return c.Version
}

// Decode implements bin.Decoder.
func (c *ChatParticipants) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipants#3f460fed to nil")
	}
	if err := b.ConsumeID(ChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipants#3f460fed: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3f460fed: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3f460fed: field participants: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChatParticipant(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatParticipants#3f460fed: field participants: %w", err)
			}
			c.Participants = append(c.Participants, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3f460fed: field version: %w", err)
		}
		c.Version = value
	}
	return nil
}

// construct implements constructor of ChatParticipantsClass.
func (c ChatParticipants) construct() ChatParticipantsClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipants.
var (
	_ bin.Encoder = &ChatParticipants{}
	_ bin.Decoder = &ChatParticipants{}

	_ ChatParticipantsClass = &ChatParticipants{}
)

// ChatParticipantsClass represents ChatParticipants generic type.
//
// See https://core.telegram.org/type/ChatParticipants for reference.
//
// Example:
//  g, err := tg.DecodeChatParticipants(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChatParticipantsForbidden: // chatParticipantsForbidden#fc900c2b
//  case *tg.ChatParticipants: // chatParticipants#3f460fed
//  default: panic(v)
//  }
type ChatParticipantsClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatParticipantsClass

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

	// Group ID
	GetChatID() (value int)
	// AsNotForbidden tries to map ChatParticipantsClass to ChatParticipants.
	AsNotForbidden() (*ChatParticipants, bool)
}

// AsNotForbidden tries to map ChatParticipantsForbidden to ChatParticipants.
func (c *ChatParticipantsForbidden) AsNotForbidden() (*ChatParticipants, bool) {
	return nil, false
}

// AsNotForbidden tries to map ChatParticipants to ChatParticipants.
func (c *ChatParticipants) AsNotForbidden() (*ChatParticipants, bool) {
	return c, true
}

// DecodeChatParticipants implements binary de-serialization for ChatParticipantsClass.
func DecodeChatParticipants(buf *bin.Buffer) (ChatParticipantsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatParticipantsForbiddenTypeID:
		// Decoding chatParticipantsForbidden#fc900c2b.
		v := ChatParticipantsForbidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", err)
		}
		return &v, nil
	case ChatParticipantsTypeID:
		// Decoding chatParticipants#3f460fed.
		v := ChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatParticipants boxes the ChatParticipantsClass providing a helper.
type ChatParticipantsBox struct {
	ChatParticipants ChatParticipantsClass
}

// Decode implements bin.Decoder for ChatParticipantsBox.
func (b *ChatParticipantsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatParticipantsBox to nil")
	}
	v, err := DecodeChatParticipants(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatParticipants = v
	return nil
}

// Encode implements bin.Encode for ChatParticipantsBox.
func (b *ChatParticipantsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatParticipants == nil {
		return fmt.Errorf("unable to encode ChatParticipantsClass as nil")
	}
	return b.ChatParticipants.Encode(buf)
}

// ChatParticipantsClassArray is adapter for slice of ChatParticipantsClass.
type ChatParticipantsClassArray []ChatParticipantsClass

// Sort sorts slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) Sort(less func(a, b ChatParticipantsClass) bool) ChatParticipantsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) SortStable(less func(a, b ChatParticipantsClass) bool) ChatParticipantsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantsClass.
func (s ChatParticipantsClassArray) Retain(keep func(x ChatParticipantsClass) bool) ChatParticipantsClassArray {
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
func (s ChatParticipantsClassArray) First() (v ChatParticipantsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsClassArray) Last() (v ChatParticipantsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsClassArray) PopFirst() (v ChatParticipantsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsClassArray) Pop() (v ChatParticipantsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChatParticipantsForbidden returns copy with only ChatParticipantsForbidden constructors.
func (s ChatParticipantsClassArray) AsChatParticipantsForbidden() (to ChatParticipantsForbiddenArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantsForbidden)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipants returns copy with only ChatParticipants constructors.
func (s ChatParticipantsClassArray) AsChatParticipants() (to ChatParticipantsArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipants)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotForbidden appends only NotForbidden constructors to
// given slice.
func (s ChatParticipantsClassArray) AppendOnlyNotForbidden(to []*ChatParticipants) []*ChatParticipants {
	for _, elem := range s {
		value, ok := elem.AsNotForbidden()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotForbidden returns copy with only NotForbidden constructors.
func (s ChatParticipantsClassArray) AsNotForbidden() (to []*ChatParticipants) {
	return s.AppendOnlyNotForbidden(to)
}

// FirstAsNotForbidden returns first element of slice (if exists).
func (s ChatParticipantsClassArray) FirstAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// LastAsNotForbidden returns last element of slice (if exists).
func (s ChatParticipantsClassArray) LastAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopFirstAsNotForbidden returns element of slice (if exists).
func (s *ChatParticipantsClassArray) PopFirstAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopAsNotForbidden returns element of slice (if exists).
func (s *ChatParticipantsClassArray) PopAsNotForbidden() (v *ChatParticipants, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// ChatParticipantsForbiddenArray is adapter for slice of ChatParticipantsForbidden.
type ChatParticipantsForbiddenArray []ChatParticipantsForbidden

// Sort sorts slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) Sort(less func(a, b ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) SortStable(less func(a, b ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantsForbidden.
func (s ChatParticipantsForbiddenArray) Retain(keep func(x ChatParticipantsForbidden) bool) ChatParticipantsForbiddenArray {
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
func (s ChatParticipantsForbiddenArray) First() (v ChatParticipantsForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsForbiddenArray) Last() (v ChatParticipantsForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsForbiddenArray) PopFirst() (v ChatParticipantsForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantsForbidden
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsForbiddenArray) Pop() (v ChatParticipantsForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatParticipantsArray is adapter for slice of ChatParticipants.
type ChatParticipantsArray []ChatParticipants

// Sort sorts slice of ChatParticipants.
func (s ChatParticipantsArray) Sort(less func(a, b ChatParticipants) bool) ChatParticipantsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipants.
func (s ChatParticipantsArray) SortStable(less func(a, b ChatParticipants) bool) ChatParticipantsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipants.
func (s ChatParticipantsArray) Retain(keep func(x ChatParticipants) bool) ChatParticipantsArray {
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
func (s ChatParticipantsArray) First() (v ChatParticipants, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantsArray) Last() (v ChatParticipants, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantsArray) PopFirst() (v ChatParticipants, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipants
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantsArray) Pop() (v ChatParticipants, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
