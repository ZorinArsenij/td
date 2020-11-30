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

// BotInlineMessageMediaAuto represents TL type `botInlineMessageMediaAuto#764cf810`.
type BotInlineMessageMediaAuto struct {
	// Flags field of BotInlineMessageMediaAuto.
	Flags bin.Fields
	// Message field of BotInlineMessageMediaAuto.
	Message string
	// Entities field of BotInlineMessageMediaAuto.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// ReplyMarkup field of BotInlineMessageMediaAuto.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaAutoTypeID is TL type id of BotInlineMessageMediaAuto.
const BotInlineMessageMediaAutoTypeID = 0x764cf810

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaAuto) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaAuto#764cf810 as nil")
	}
	buf.PutID(BotInlineMessageMediaAutoTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field flags: %w", err)
	}
	buf.PutString(b.Message)
	if b.Flags.Has(1) {
		buf.PutVectorHeader(len(b.Entities))
		for idx, v := range b.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(buf); err != nil {
				return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetEntities sets value of Entities conditional field.
func (b *BotInlineMessageMediaAuto) SetEntities(value []MessageEntityClass) {
	b.Flags.Set(1)
	b.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaAuto) GetEntities() (value []MessageEntityClass, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaAuto) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaAuto) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaAuto) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaAuto#764cf810 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaAutoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(1) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field entities: %w", err)
			}
			b.Entities = append(b.Entities, value)
		}
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaAuto) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaAuto.
var (
	_ bin.Encoder = &BotInlineMessageMediaAuto{}
	_ bin.Decoder = &BotInlineMessageMediaAuto{}

	_ BotInlineMessageClass = &BotInlineMessageMediaAuto{}
)

// BotInlineMessageText represents TL type `botInlineMessageText#8c7f65e2`.
type BotInlineMessageText struct {
	// Flags field of BotInlineMessageText.
	Flags bin.Fields
	// NoWebpage field of BotInlineMessageText.
	NoWebpage bool
	// Message field of BotInlineMessageText.
	Message string
	// Entities field of BotInlineMessageText.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// ReplyMarkup field of BotInlineMessageText.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageTextTypeID is TL type id of BotInlineMessageText.
const BotInlineMessageTextTypeID = 0x8c7f65e2

// Encode implements bin.Encoder.
func (b *BotInlineMessageText) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageText#8c7f65e2 as nil")
	}
	buf.PutID(BotInlineMessageTextTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field flags: %w", err)
	}
	buf.PutString(b.Message)
	if b.Flags.Has(1) {
		buf.PutVectorHeader(len(b.Entities))
		for idx, v := range b.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(buf); err != nil {
				return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (b *BotInlineMessageText) SetNoWebpage(value bool) {
	if value {
		b.Flags.Set(0)
	} else {
		b.Flags.Unset(0)
	}
}

// SetEntities sets value of Entities conditional field.
func (b *BotInlineMessageText) SetEntities(value []MessageEntityClass) {
	b.Flags.Set(1)
	b.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageText) GetEntities() (value []MessageEntityClass, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageText) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageText) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageText) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageText#8c7f65e2 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageTextTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field flags: %w", err)
		}
	}
	b.NoWebpage = b.Flags.Has(0)
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(1) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field entities: %w", err)
			}
			b.Entities = append(b.Entities, value)
		}
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageText) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageText.
var (
	_ bin.Encoder = &BotInlineMessageText{}
	_ bin.Decoder = &BotInlineMessageText{}

	_ BotInlineMessageClass = &BotInlineMessageText{}
)

// BotInlineMessageMediaGeo represents TL type `botInlineMessageMediaGeo#51846fd`.
type BotInlineMessageMediaGeo struct {
	// Flags field of BotInlineMessageMediaGeo.
	Flags bin.Fields
	// Geo field of BotInlineMessageMediaGeo.
	Geo GeoPointClass
	// Heading field of BotInlineMessageMediaGeo.
	//
	// Use SetHeading and GetHeading helpers.
	Heading int
	// Period field of BotInlineMessageMediaGeo.
	//
	// Use SetPeriod and GetPeriod helpers.
	Period int
	// ProximityNotificationRadius field of BotInlineMessageMediaGeo.
	//
	// Use SetProximityNotificationRadius and GetProximityNotificationRadius helpers.
	ProximityNotificationRadius int
	// ReplyMarkup field of BotInlineMessageMediaGeo.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaGeoTypeID is TL type id of BotInlineMessageMediaGeo.
const BotInlineMessageMediaGeoTypeID = 0x51846fd

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaGeo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaGeo#51846fd as nil")
	}
	buf.PutID(BotInlineMessageMediaGeoTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field flags: %w", err)
	}
	if b.Geo == nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field geo is nil")
	}
	if err := b.Geo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field geo: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutInt(b.Heading)
	}
	if b.Flags.Has(1) {
		buf.PutInt(b.Period)
	}
	if b.Flags.Has(3) {
		buf.PutInt(b.ProximityNotificationRadius)
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetHeading sets value of Heading conditional field.
func (b *BotInlineMessageMediaGeo) SetHeading(value int) {
	b.Flags.Set(0)
	b.Heading = value
}

// GetHeading returns value of Heading conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetHeading() (value int, ok bool) {
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Heading, true
}

// SetPeriod sets value of Period conditional field.
func (b *BotInlineMessageMediaGeo) SetPeriod(value int) {
	b.Flags.Set(1)
	b.Period = value
}

// GetPeriod returns value of Period conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetPeriod() (value int, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Period, true
}

// SetProximityNotificationRadius sets value of ProximityNotificationRadius conditional field.
func (b *BotInlineMessageMediaGeo) SetProximityNotificationRadius(value int) {
	b.Flags.Set(3)
	b.ProximityNotificationRadius = value
}

// GetProximityNotificationRadius returns value of ProximityNotificationRadius conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetProximityNotificationRadius() (value int, ok bool) {
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.ProximityNotificationRadius, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaGeo) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaGeo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaGeo#51846fd to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaGeoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field flags: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field geo: %w", err)
		}
		b.Geo = value
	}
	if b.Flags.Has(0) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field heading: %w", err)
		}
		b.Heading = value
	}
	if b.Flags.Has(1) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field period: %w", err)
		}
		b.Period = value
	}
	if b.Flags.Has(3) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field proximity_notification_radius: %w", err)
		}
		b.ProximityNotificationRadius = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaGeo) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaGeo.
var (
	_ bin.Encoder = &BotInlineMessageMediaGeo{}
	_ bin.Decoder = &BotInlineMessageMediaGeo{}

	_ BotInlineMessageClass = &BotInlineMessageMediaGeo{}
)

// BotInlineMessageMediaVenue represents TL type `botInlineMessageMediaVenue#8a86659c`.
type BotInlineMessageMediaVenue struct {
	// Flags field of BotInlineMessageMediaVenue.
	Flags bin.Fields
	// Geo field of BotInlineMessageMediaVenue.
	Geo GeoPointClass
	// Title field of BotInlineMessageMediaVenue.
	Title string
	// Address field of BotInlineMessageMediaVenue.
	Address string
	// Provider field of BotInlineMessageMediaVenue.
	Provider string
	// VenueID field of BotInlineMessageMediaVenue.
	VenueID string
	// VenueType field of BotInlineMessageMediaVenue.
	VenueType string
	// ReplyMarkup field of BotInlineMessageMediaVenue.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaVenueTypeID is TL type id of BotInlineMessageMediaVenue.
const BotInlineMessageMediaVenueTypeID = 0x8a86659c

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaVenue) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaVenue#8a86659c as nil")
	}
	buf.PutID(BotInlineMessageMediaVenueTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field flags: %w", err)
	}
	if b.Geo == nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field geo is nil")
	}
	if err := b.Geo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field geo: %w", err)
	}
	buf.PutString(b.Title)
	buf.PutString(b.Address)
	buf.PutString(b.Provider)
	buf.PutString(b.VenueID)
	buf.PutString(b.VenueType)
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaVenue) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaVenue) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaVenue) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaVenue#8a86659c to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field flags: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field geo: %w", err)
		}
		b.Geo = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field title: %w", err)
		}
		b.Title = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field address: %w", err)
		}
		b.Address = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field provider: %w", err)
		}
		b.Provider = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field venue_id: %w", err)
		}
		b.VenueID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field venue_type: %w", err)
		}
		b.VenueType = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaVenue) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaVenue.
var (
	_ bin.Encoder = &BotInlineMessageMediaVenue{}
	_ bin.Decoder = &BotInlineMessageMediaVenue{}

	_ BotInlineMessageClass = &BotInlineMessageMediaVenue{}
)

// BotInlineMessageMediaContact represents TL type `botInlineMessageMediaContact#18d1cdc2`.
type BotInlineMessageMediaContact struct {
	// Flags field of BotInlineMessageMediaContact.
	Flags bin.Fields
	// PhoneNumber field of BotInlineMessageMediaContact.
	PhoneNumber string
	// FirstName field of BotInlineMessageMediaContact.
	FirstName string
	// LastName field of BotInlineMessageMediaContact.
	LastName string
	// Vcard field of BotInlineMessageMediaContact.
	Vcard string
	// ReplyMarkup field of BotInlineMessageMediaContact.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaContactTypeID is TL type id of BotInlineMessageMediaContact.
const BotInlineMessageMediaContactTypeID = 0x18d1cdc2

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaContact) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaContact#18d1cdc2 as nil")
	}
	buf.PutID(BotInlineMessageMediaContactTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field flags: %w", err)
	}
	buf.PutString(b.PhoneNumber)
	buf.PutString(b.FirstName)
	buf.PutString(b.LastName)
	buf.PutString(b.Vcard)
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaContact) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaContact) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaContact) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaContact#18d1cdc2 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaContactTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field phone_number: %w", err)
		}
		b.PhoneNumber = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field first_name: %w", err)
		}
		b.FirstName = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field last_name: %w", err)
		}
		b.LastName = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field vcard: %w", err)
		}
		b.Vcard = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaContact) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaContact.
var (
	_ bin.Encoder = &BotInlineMessageMediaContact{}
	_ bin.Decoder = &BotInlineMessageMediaContact{}

	_ BotInlineMessageClass = &BotInlineMessageMediaContact{}
)

// BotInlineMessageClass represents BotInlineMessage generic type.
//
// Example:
//  g, err := DecodeBotInlineMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BotInlineMessageMediaAuto: // botInlineMessageMediaAuto#764cf810
//  case *BotInlineMessageText: // botInlineMessageText#8c7f65e2
//  case *BotInlineMessageMediaGeo: // botInlineMessageMediaGeo#51846fd
//  case *BotInlineMessageMediaVenue: // botInlineMessageMediaVenue#8a86659c
//  case *BotInlineMessageMediaContact: // botInlineMessageMediaContact#18d1cdc2
//  default: panic(v)
//  }
type BotInlineMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() BotInlineMessageClass
}

// DecodeBotInlineMessage implements binary de-serialization for BotInlineMessageClass.
func DecodeBotInlineMessage(buf *bin.Buffer) (BotInlineMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotInlineMessageMediaAutoTypeID:
		// Decoding botInlineMessageMediaAuto#764cf810.
		v := BotInlineMessageMediaAuto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageTextTypeID:
		// Decoding botInlineMessageText#8c7f65e2.
		v := BotInlineMessageText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaGeoTypeID:
		// Decoding botInlineMessageMediaGeo#51846fd.
		v := BotInlineMessageMediaGeo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaVenueTypeID:
		// Decoding botInlineMessageMediaVenue#8a86659c.
		v := BotInlineMessageMediaVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaContactTypeID:
		// Decoding botInlineMessageMediaContact#18d1cdc2.
		v := BotInlineMessageMediaContact{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// BotInlineMessage boxes the BotInlineMessageClass providing a helper.
type BotInlineMessageBox struct {
	BotInlineMessage BotInlineMessageClass
}

// Decode implements bin.Decoder for BotInlineMessageBox.
func (b *BotInlineMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotInlineMessageBox to nil")
	}
	v, err := DecodeBotInlineMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotInlineMessage = v
	return nil
}

// Encode implements bin.Encode for BotInlineMessageBox.
func (b *BotInlineMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotInlineMessage == nil {
		return fmt.Errorf("unable to encode BotInlineMessageClass as nil")
	}
	return b.BotInlineMessage.Encode(buf)
}
