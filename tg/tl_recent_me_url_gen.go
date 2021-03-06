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

// RecentMeUrlUnknown represents TL type `recentMeUrlUnknown#46e1d13d`.
// Unknown t.me url
//
// See https://core.telegram.org/constructor/recentMeUrlUnknown for reference.
type RecentMeUrlUnknown struct {
	// URL
	URL string
}

// RecentMeUrlUnknownTypeID is TL type id of RecentMeUrlUnknown.
const RecentMeUrlUnknownTypeID = 0x46e1d13d

func (r *RecentMeUrlUnknown) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeUrlUnknown) String() string {
	if r == nil {
		return "RecentMeUrlUnknown(nil)"
	}
	type Alias RecentMeUrlUnknown
	return fmt.Sprintf("RecentMeUrlUnknown%+v", Alias(*r))
}

// FillFrom fills RecentMeUrlUnknown from given interface.
func (r *RecentMeUrlUnknown) FillFrom(from interface {
	GetURL() (value string)
}) {
	r.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeUrlUnknown) TypeID() uint32 {
	return RecentMeUrlUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeUrlUnknown) TypeName() string {
	return "recentMeUrlUnknown"
}

// TypeInfo returns info about TL type.
func (r *RecentMeUrlUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUnknown",
		ID:   RecentMeUrlUnknownTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutID(RecentMeUrlUnknownTypeID)
	b.PutString(r.URL)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeUrlUnknown) GetURL() (value string) {
	return r.URL
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	if err := b.ConsumeID(RecentMeUrlUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: field url: %w", err)
		}
		r.URL = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlUnknown) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlUnknown.
var (
	_ bin.Encoder = &RecentMeUrlUnknown{}
	_ bin.Decoder = &RecentMeUrlUnknown{}

	_ RecentMeUrlClass = &RecentMeUrlUnknown{}
)

// RecentMeUrlUser represents TL type `recentMeUrlUser#8dbc3336`.
// Recent t.me link to a user
//
// See https://core.telegram.org/constructor/recentMeUrlUser for reference.
type RecentMeUrlUser struct {
	// URL
	URL string
	// User ID
	UserID int
}

// RecentMeUrlUserTypeID is TL type id of RecentMeUrlUser.
const RecentMeUrlUserTypeID = 0x8dbc3336

func (r *RecentMeUrlUser) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeUrlUser) String() string {
	if r == nil {
		return "RecentMeUrlUser(nil)"
	}
	type Alias RecentMeUrlUser
	return fmt.Sprintf("RecentMeUrlUser%+v", Alias(*r))
}

// FillFrom fills RecentMeUrlUser from given interface.
func (r *RecentMeUrlUser) FillFrom(from interface {
	GetURL() (value string)
	GetUserID() (value int)
}) {
	r.URL = from.GetURL()
	r.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeUrlUser) TypeID() uint32 {
	return RecentMeUrlUserTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeUrlUser) TypeName() string {
	return "recentMeUrlUser"
}

// TypeInfo returns info about TL type.
func (r *RecentMeUrlUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUser",
		ID:   RecentMeUrlUserTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlUser) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#8dbc3336 as nil")
	}
	b.PutID(RecentMeUrlUserTypeID)
	b.PutString(r.URL)
	b.PutInt(r.UserID)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeUrlUser) GetURL() (value string) {
	return r.URL
}

// GetUserID returns value of UserID field.
func (r *RecentMeUrlUser) GetUserID() (value int) {
	return r.UserID
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlUser) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#8dbc3336 to nil")
	}
	if err := b.ConsumeID(RecentMeUrlUserTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field user_id: %w", err)
		}
		r.UserID = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlUser) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlUser.
var (
	_ bin.Encoder = &RecentMeUrlUser{}
	_ bin.Decoder = &RecentMeUrlUser{}

	_ RecentMeUrlClass = &RecentMeUrlUser{}
)

// RecentMeUrlChat represents TL type `recentMeUrlChat#a01b22f9`.
// Recent t.me link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChat for reference.
type RecentMeUrlChat struct {
	// t.me URL
	URL string
	// Chat ID
	ChatID int
}

// RecentMeUrlChatTypeID is TL type id of RecentMeUrlChat.
const RecentMeUrlChatTypeID = 0xa01b22f9

func (r *RecentMeUrlChat) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeUrlChat) String() string {
	if r == nil {
		return "RecentMeUrlChat(nil)"
	}
	type Alias RecentMeUrlChat
	return fmt.Sprintf("RecentMeUrlChat%+v", Alias(*r))
}

// FillFrom fills RecentMeUrlChat from given interface.
func (r *RecentMeUrlChat) FillFrom(from interface {
	GetURL() (value string)
	GetChatID() (value int)
}) {
	r.URL = from.GetURL()
	r.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeUrlChat) TypeID() uint32 {
	return RecentMeUrlChatTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeUrlChat) TypeName() string {
	return "recentMeUrlChat"
}

// TypeInfo returns info about TL type.
func (r *RecentMeUrlChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChat",
		ID:   RecentMeUrlChatTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlChat) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#a01b22f9 as nil")
	}
	b.PutID(RecentMeUrlChatTypeID)
	b.PutString(r.URL)
	b.PutInt(r.ChatID)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeUrlChat) GetURL() (value string) {
	return r.URL
}

// GetChatID returns value of ChatID field.
func (r *RecentMeUrlChat) GetChatID() (value int) {
	return r.ChatID
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlChat) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#a01b22f9 to nil")
	}
	if err := b.ConsumeID(RecentMeUrlChatTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlChat) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlChat.
var (
	_ bin.Encoder = &RecentMeUrlChat{}
	_ bin.Decoder = &RecentMeUrlChat{}

	_ RecentMeUrlClass = &RecentMeUrlChat{}
)

// RecentMeUrlChatInvite represents TL type `recentMeUrlChatInvite#eb49081d`.
// Recent t.me invite link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChatInvite for reference.
type RecentMeUrlChatInvite struct {
	// t.me URL
	URL string
	// Chat invitation
	ChatInvite ChatInviteClass
}

// RecentMeUrlChatInviteTypeID is TL type id of RecentMeUrlChatInvite.
const RecentMeUrlChatInviteTypeID = 0xeb49081d

func (r *RecentMeUrlChatInvite) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatInvite == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeUrlChatInvite) String() string {
	if r == nil {
		return "RecentMeUrlChatInvite(nil)"
	}
	type Alias RecentMeUrlChatInvite
	return fmt.Sprintf("RecentMeUrlChatInvite%+v", Alias(*r))
}

// FillFrom fills RecentMeUrlChatInvite from given interface.
func (r *RecentMeUrlChatInvite) FillFrom(from interface {
	GetURL() (value string)
	GetChatInvite() (value ChatInviteClass)
}) {
	r.URL = from.GetURL()
	r.ChatInvite = from.GetChatInvite()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeUrlChatInvite) TypeID() uint32 {
	return RecentMeUrlChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeUrlChatInvite) TypeName() string {
	return "recentMeUrlChatInvite"
}

// TypeInfo returns info about TL type.
func (r *RecentMeUrlChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChatInvite",
		ID:   RecentMeUrlChatInviteTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatInvite",
			SchemaName: "chat_invite",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlChatInvite) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutID(RecentMeUrlChatInviteTypeID)
	b.PutString(r.URL)
	if r.ChatInvite == nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite is nil")
	}
	if err := r.ChatInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeUrlChatInvite) GetURL() (value string) {
	return r.URL
}

// GetChatInvite returns value of ChatInvite field.
func (r *RecentMeUrlChatInvite) GetChatInvite() (value ChatInviteClass) {
	return r.ChatInvite
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlChatInvite) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	if err := b.ConsumeID(RecentMeUrlChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
		}
		r.ChatInvite = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlChatInvite) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlChatInvite.
var (
	_ bin.Encoder = &RecentMeUrlChatInvite{}
	_ bin.Decoder = &RecentMeUrlChatInvite{}

	_ RecentMeUrlClass = &RecentMeUrlChatInvite{}
)

// RecentMeUrlStickerSet represents TL type `recentMeUrlStickerSet#bc0a57dc`.
// Recent t.me stickerset installation URL
//
// See https://core.telegram.org/constructor/recentMeUrlStickerSet for reference.
type RecentMeUrlStickerSet struct {
	// t.me URL
	URL string
	// Stickerset
	Set StickerSetCoveredClass
}

// RecentMeUrlStickerSetTypeID is TL type id of RecentMeUrlStickerSet.
const RecentMeUrlStickerSetTypeID = 0xbc0a57dc

func (r *RecentMeUrlStickerSet) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.Set == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeUrlStickerSet) String() string {
	if r == nil {
		return "RecentMeUrlStickerSet(nil)"
	}
	type Alias RecentMeUrlStickerSet
	return fmt.Sprintf("RecentMeUrlStickerSet%+v", Alias(*r))
}

// FillFrom fills RecentMeUrlStickerSet from given interface.
func (r *RecentMeUrlStickerSet) FillFrom(from interface {
	GetURL() (value string)
	GetSet() (value StickerSetCoveredClass)
}) {
	r.URL = from.GetURL()
	r.Set = from.GetSet()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeUrlStickerSet) TypeID() uint32 {
	return RecentMeUrlStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeUrlStickerSet) TypeName() string {
	return "recentMeUrlStickerSet"
}

// TypeInfo returns info about TL type.
func (r *RecentMeUrlStickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlStickerSet",
		ID:   RecentMeUrlStickerSetTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Set",
			SchemaName: "set",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeUrlStickerSet) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutID(RecentMeUrlStickerSetTypeID)
	b.PutString(r.URL)
	if r.Set == nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set is nil")
	}
	if err := r.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeUrlStickerSet) GetURL() (value string) {
	return r.URL
}

// GetSet returns value of Set field.
func (r *RecentMeUrlStickerSet) GetSet() (value StickerSetCoveredClass) {
	return r.Set
}

// Decode implements bin.Decoder.
func (r *RecentMeUrlStickerSet) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	if err := b.ConsumeID(RecentMeUrlStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeStickerSetCovered(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
		}
		r.Set = value
	}
	return nil
}

// construct implements constructor of RecentMeUrlClass.
func (r RecentMeUrlStickerSet) construct() RecentMeUrlClass { return &r }

// Ensuring interfaces in compile-time for RecentMeUrlStickerSet.
var (
	_ bin.Encoder = &RecentMeUrlStickerSet{}
	_ bin.Decoder = &RecentMeUrlStickerSet{}

	_ RecentMeUrlClass = &RecentMeUrlStickerSet{}
)

// RecentMeUrlClass represents RecentMeUrl generic type.
//
// See https://core.telegram.org/type/RecentMeUrl for reference.
//
// Example:
//  g, err := tg.DecodeRecentMeUrl(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.RecentMeUrlUnknown: // recentMeUrlUnknown#46e1d13d
//  case *tg.RecentMeUrlUser: // recentMeUrlUser#8dbc3336
//  case *tg.RecentMeUrlChat: // recentMeUrlChat#a01b22f9
//  case *tg.RecentMeUrlChatInvite: // recentMeUrlChatInvite#eb49081d
//  case *tg.RecentMeUrlStickerSet: // recentMeUrlStickerSet#bc0a57dc
//  default: panic(v)
//  }
type RecentMeUrlClass interface {
	bin.Encoder
	bin.Decoder
	construct() RecentMeUrlClass

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

	// URL
	GetURL() (value string)
}

// DecodeRecentMeUrl implements binary de-serialization for RecentMeUrlClass.
func DecodeRecentMeUrl(buf *bin.Buffer) (RecentMeUrlClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RecentMeUrlUnknownTypeID:
		// Decoding recentMeUrlUnknown#46e1d13d.
		v := RecentMeUrlUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlUserTypeID:
		// Decoding recentMeUrlUser#8dbc3336.
		v := RecentMeUrlUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlChatTypeID:
		// Decoding recentMeUrlChat#a01b22f9.
		v := RecentMeUrlChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlChatInviteTypeID:
		// Decoding recentMeUrlChatInvite#eb49081d.
		v := RecentMeUrlChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	case RecentMeUrlStickerSetTypeID:
		// Decoding recentMeUrlStickerSet#bc0a57dc.
		v := RecentMeUrlStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RecentMeUrlClass: %w", bin.NewUnexpectedID(id))
	}
}

// RecentMeUrl boxes the RecentMeUrlClass providing a helper.
type RecentMeUrlBox struct {
	RecentMeUrl RecentMeUrlClass
}

// Decode implements bin.Decoder for RecentMeUrlBox.
func (b *RecentMeUrlBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RecentMeUrlBox to nil")
	}
	v, err := DecodeRecentMeUrl(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentMeUrl = v
	return nil
}

// Encode implements bin.Encode for RecentMeUrlBox.
func (b *RecentMeUrlBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentMeUrl == nil {
		return fmt.Errorf("unable to encode RecentMeUrlClass as nil")
	}
	return b.RecentMeUrl.Encode(buf)
}

// RecentMeUrlClassArray is adapter for slice of RecentMeUrlClass.
type RecentMeUrlClassArray []RecentMeUrlClass

// Sort sorts slice of RecentMeUrlClass.
func (s RecentMeUrlClassArray) Sort(less func(a, b RecentMeUrlClass) bool) RecentMeUrlClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlClass.
func (s RecentMeUrlClassArray) SortStable(less func(a, b RecentMeUrlClass) bool) RecentMeUrlClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlClass.
func (s RecentMeUrlClassArray) Retain(keep func(x RecentMeUrlClass) bool) RecentMeUrlClassArray {
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
func (s RecentMeUrlClassArray) First() (v RecentMeUrlClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlClassArray) Last() (v RecentMeUrlClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlClassArray) PopFirst() (v RecentMeUrlClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlClassArray) Pop() (v RecentMeUrlClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsRecentMeUrlUnknown returns copy with only RecentMeUrlUnknown constructors.
func (s RecentMeUrlClassArray) AsRecentMeUrlUnknown() (to RecentMeUrlUnknownArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeUrlUnknown)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeUrlUser returns copy with only RecentMeUrlUser constructors.
func (s RecentMeUrlClassArray) AsRecentMeUrlUser() (to RecentMeUrlUserArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeUrlUser)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeUrlChat returns copy with only RecentMeUrlChat constructors.
func (s RecentMeUrlClassArray) AsRecentMeUrlChat() (to RecentMeUrlChatArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeUrlChat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeUrlChatInvite returns copy with only RecentMeUrlChatInvite constructors.
func (s RecentMeUrlClassArray) AsRecentMeUrlChatInvite() (to RecentMeUrlChatInviteArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeUrlChatInvite)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeUrlStickerSet returns copy with only RecentMeUrlStickerSet constructors.
func (s RecentMeUrlClassArray) AsRecentMeUrlStickerSet() (to RecentMeUrlStickerSetArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeUrlStickerSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// RecentMeUrlUnknownArray is adapter for slice of RecentMeUrlUnknown.
type RecentMeUrlUnknownArray []RecentMeUrlUnknown

// Sort sorts slice of RecentMeUrlUnknown.
func (s RecentMeUrlUnknownArray) Sort(less func(a, b RecentMeUrlUnknown) bool) RecentMeUrlUnknownArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlUnknown.
func (s RecentMeUrlUnknownArray) SortStable(less func(a, b RecentMeUrlUnknown) bool) RecentMeUrlUnknownArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlUnknown.
func (s RecentMeUrlUnknownArray) Retain(keep func(x RecentMeUrlUnknown) bool) RecentMeUrlUnknownArray {
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
func (s RecentMeUrlUnknownArray) First() (v RecentMeUrlUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlUnknownArray) Last() (v RecentMeUrlUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlUnknownArray) PopFirst() (v RecentMeUrlUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlUnknown
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlUnknownArray) Pop() (v RecentMeUrlUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeUrlUserArray is adapter for slice of RecentMeUrlUser.
type RecentMeUrlUserArray []RecentMeUrlUser

// Sort sorts slice of RecentMeUrlUser.
func (s RecentMeUrlUserArray) Sort(less func(a, b RecentMeUrlUser) bool) RecentMeUrlUserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlUser.
func (s RecentMeUrlUserArray) SortStable(less func(a, b RecentMeUrlUser) bool) RecentMeUrlUserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlUser.
func (s RecentMeUrlUserArray) Retain(keep func(x RecentMeUrlUser) bool) RecentMeUrlUserArray {
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
func (s RecentMeUrlUserArray) First() (v RecentMeUrlUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlUserArray) Last() (v RecentMeUrlUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlUserArray) PopFirst() (v RecentMeUrlUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlUser
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlUserArray) Pop() (v RecentMeUrlUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeUrlChatArray is adapter for slice of RecentMeUrlChat.
type RecentMeUrlChatArray []RecentMeUrlChat

// Sort sorts slice of RecentMeUrlChat.
func (s RecentMeUrlChatArray) Sort(less func(a, b RecentMeUrlChat) bool) RecentMeUrlChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlChat.
func (s RecentMeUrlChatArray) SortStable(less func(a, b RecentMeUrlChat) bool) RecentMeUrlChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlChat.
func (s RecentMeUrlChatArray) Retain(keep func(x RecentMeUrlChat) bool) RecentMeUrlChatArray {
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
func (s RecentMeUrlChatArray) First() (v RecentMeUrlChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlChatArray) Last() (v RecentMeUrlChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlChatArray) PopFirst() (v RecentMeUrlChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlChat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlChatArray) Pop() (v RecentMeUrlChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeUrlChatInviteArray is adapter for slice of RecentMeUrlChatInvite.
type RecentMeUrlChatInviteArray []RecentMeUrlChatInvite

// Sort sorts slice of RecentMeUrlChatInvite.
func (s RecentMeUrlChatInviteArray) Sort(less func(a, b RecentMeUrlChatInvite) bool) RecentMeUrlChatInviteArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlChatInvite.
func (s RecentMeUrlChatInviteArray) SortStable(less func(a, b RecentMeUrlChatInvite) bool) RecentMeUrlChatInviteArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlChatInvite.
func (s RecentMeUrlChatInviteArray) Retain(keep func(x RecentMeUrlChatInvite) bool) RecentMeUrlChatInviteArray {
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
func (s RecentMeUrlChatInviteArray) First() (v RecentMeUrlChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlChatInviteArray) Last() (v RecentMeUrlChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlChatInviteArray) PopFirst() (v RecentMeUrlChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlChatInvite
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlChatInviteArray) Pop() (v RecentMeUrlChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeUrlStickerSetArray is adapter for slice of RecentMeUrlStickerSet.
type RecentMeUrlStickerSetArray []RecentMeUrlStickerSet

// Sort sorts slice of RecentMeUrlStickerSet.
func (s RecentMeUrlStickerSetArray) Sort(less func(a, b RecentMeUrlStickerSet) bool) RecentMeUrlStickerSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeUrlStickerSet.
func (s RecentMeUrlStickerSetArray) SortStable(less func(a, b RecentMeUrlStickerSet) bool) RecentMeUrlStickerSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeUrlStickerSet.
func (s RecentMeUrlStickerSetArray) Retain(keep func(x RecentMeUrlStickerSet) bool) RecentMeUrlStickerSetArray {
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
func (s RecentMeUrlStickerSetArray) First() (v RecentMeUrlStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeUrlStickerSetArray) Last() (v RecentMeUrlStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeUrlStickerSetArray) PopFirst() (v RecentMeUrlStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeUrlStickerSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeUrlStickerSetArray) Pop() (v RecentMeUrlStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
