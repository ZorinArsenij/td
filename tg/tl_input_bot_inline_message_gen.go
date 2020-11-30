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

// InputBotInlineMessageMediaAuto represents TL type `inputBotInlineMessageMediaAuto#3380c786`.
type InputBotInlineMessageMediaAuto struct {
	// Flags field of InputBotInlineMessageMediaAuto.
	Flags bin.Fields
	// Message field of InputBotInlineMessageMediaAuto.
	Message string
	// Entities field of InputBotInlineMessageMediaAuto.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// ReplyMarkup field of InputBotInlineMessageMediaAuto.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageMediaAutoTypeID is TL type id of InputBotInlineMessageMediaAuto.
const InputBotInlineMessageMediaAutoTypeID = 0x3380c786

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageMediaAuto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageMediaAuto#3380c786 as nil")
	}
	b.PutID(InputBotInlineMessageMediaAutoTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaAuto#3380c786: field flags: %w", err)
	}
	b.PutString(i.Message)
	if i.Flags.Has(1) {
		b.PutVectorHeader(len(i.Entities))
		for idx, v := range i.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode inputBotInlineMessageMediaAuto#3380c786: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputBotInlineMessageMediaAuto#3380c786: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaAuto#3380c786: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaAuto#3380c786: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetEntities sets value of Entities conditional field.
func (i *InputBotInlineMessageMediaAuto) SetEntities(value []MessageEntityClass) {
	i.Flags.Set(1)
	i.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaAuto) GetEntities() (value []MessageEntityClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageMediaAuto) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaAuto) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageMediaAuto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageMediaAuto#3380c786 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageMediaAutoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: field message: %w", err)
		}
		i.Message = value
	}
	if i.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: field entities: %w", err)
			}
			i.Entities = append(i.Entities, value)
		}
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaAuto#3380c786: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageMediaAuto) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageMediaAuto.
var (
	_ bin.Encoder = &InputBotInlineMessageMediaAuto{}
	_ bin.Decoder = &InputBotInlineMessageMediaAuto{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageMediaAuto{}
)

// InputBotInlineMessageText represents TL type `inputBotInlineMessageText#3dcd7a87`.
type InputBotInlineMessageText struct {
	// Flags field of InputBotInlineMessageText.
	Flags bin.Fields
	// NoWebpage field of InputBotInlineMessageText.
	NoWebpage bool
	// Message field of InputBotInlineMessageText.
	Message string
	// Entities field of InputBotInlineMessageText.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// ReplyMarkup field of InputBotInlineMessageText.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageTextTypeID is TL type id of InputBotInlineMessageText.
const InputBotInlineMessageTextTypeID = 0x3dcd7a87

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageText) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageText#3dcd7a87 as nil")
	}
	b.PutID(InputBotInlineMessageTextTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageText#3dcd7a87: field flags: %w", err)
	}
	b.PutString(i.Message)
	if i.Flags.Has(1) {
		b.PutVectorHeader(len(i.Entities))
		for idx, v := range i.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode inputBotInlineMessageText#3dcd7a87: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputBotInlineMessageText#3dcd7a87: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageText#3dcd7a87: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageText#3dcd7a87: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (i *InputBotInlineMessageText) SetNoWebpage(value bool) {
	if value {
		i.Flags.Set(0)
	} else {
		i.Flags.Unset(0)
	}
}

// SetEntities sets value of Entities conditional field.
func (i *InputBotInlineMessageText) SetEntities(value []MessageEntityClass) {
	i.Flags.Set(1)
	i.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageText) GetEntities() (value []MessageEntityClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageText) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageText) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageText) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageText#3dcd7a87 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageTextTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: field flags: %w", err)
		}
	}
	i.NoWebpage = i.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: field message: %w", err)
		}
		i.Message = value
	}
	if i.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: field entities: %w", err)
			}
			i.Entities = append(i.Entities, value)
		}
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageText#3dcd7a87: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageText) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageText.
var (
	_ bin.Encoder = &InputBotInlineMessageText{}
	_ bin.Decoder = &InputBotInlineMessageText{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageText{}
)

// InputBotInlineMessageMediaGeo represents TL type `inputBotInlineMessageMediaGeo#96929a85`.
type InputBotInlineMessageMediaGeo struct {
	// Flags field of InputBotInlineMessageMediaGeo.
	Flags bin.Fields
	// GeoPoint field of InputBotInlineMessageMediaGeo.
	GeoPoint InputGeoPointClass
	// Heading field of InputBotInlineMessageMediaGeo.
	//
	// Use SetHeading and GetHeading helpers.
	Heading int
	// Period field of InputBotInlineMessageMediaGeo.
	//
	// Use SetPeriod and GetPeriod helpers.
	Period int
	// ProximityNotificationRadius field of InputBotInlineMessageMediaGeo.
	//
	// Use SetProximityNotificationRadius and GetProximityNotificationRadius helpers.
	ProximityNotificationRadius int
	// ReplyMarkup field of InputBotInlineMessageMediaGeo.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageMediaGeoTypeID is TL type id of InputBotInlineMessageMediaGeo.
const InputBotInlineMessageMediaGeoTypeID = 0x96929a85

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageMediaGeo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageMediaGeo#96929a85 as nil")
	}
	b.PutID(InputBotInlineMessageMediaGeoTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaGeo#96929a85: field flags: %w", err)
	}
	if i.GeoPoint == nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaGeo#96929a85: field geo_point is nil")
	}
	if err := i.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaGeo#96929a85: field geo_point: %w", err)
	}
	if i.Flags.Has(0) {
		b.PutInt(i.Heading)
	}
	if i.Flags.Has(1) {
		b.PutInt(i.Period)
	}
	if i.Flags.Has(3) {
		b.PutInt(i.ProximityNotificationRadius)
	}
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaGeo#96929a85: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaGeo#96929a85: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetHeading sets value of Heading conditional field.
func (i *InputBotInlineMessageMediaGeo) SetHeading(value int) {
	i.Flags.Set(0)
	i.Heading = value
}

// GetHeading returns value of Heading conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaGeo) GetHeading() (value int, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.Heading, true
}

// SetPeriod sets value of Period conditional field.
func (i *InputBotInlineMessageMediaGeo) SetPeriod(value int) {
	i.Flags.Set(1)
	i.Period = value
}

// GetPeriod returns value of Period conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaGeo) GetPeriod() (value int, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Period, true
}

// SetProximityNotificationRadius sets value of ProximityNotificationRadius conditional field.
func (i *InputBotInlineMessageMediaGeo) SetProximityNotificationRadius(value int) {
	i.Flags.Set(3)
	i.ProximityNotificationRadius = value
}

// GetProximityNotificationRadius returns value of ProximityNotificationRadius conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaGeo) GetProximityNotificationRadius() (value int, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.ProximityNotificationRadius, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageMediaGeo) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaGeo) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageMediaGeo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageMediaGeo#96929a85 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageMediaGeoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field geo_point: %w", err)
		}
		i.GeoPoint = value
	}
	if i.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field heading: %w", err)
		}
		i.Heading = value
	}
	if i.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field period: %w", err)
		}
		i.Period = value
	}
	if i.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field proximity_notification_radius: %w", err)
		}
		i.ProximityNotificationRadius = value
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaGeo#96929a85: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageMediaGeo) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageMediaGeo.
var (
	_ bin.Encoder = &InputBotInlineMessageMediaGeo{}
	_ bin.Decoder = &InputBotInlineMessageMediaGeo{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageMediaGeo{}
)

// InputBotInlineMessageMediaVenue represents TL type `inputBotInlineMessageMediaVenue#417bbf11`.
type InputBotInlineMessageMediaVenue struct {
	// Flags field of InputBotInlineMessageMediaVenue.
	Flags bin.Fields
	// GeoPoint field of InputBotInlineMessageMediaVenue.
	GeoPoint InputGeoPointClass
	// Title field of InputBotInlineMessageMediaVenue.
	Title string
	// Address field of InputBotInlineMessageMediaVenue.
	Address string
	// Provider field of InputBotInlineMessageMediaVenue.
	Provider string
	// VenueID field of InputBotInlineMessageMediaVenue.
	VenueID string
	// VenueType field of InputBotInlineMessageMediaVenue.
	VenueType string
	// ReplyMarkup field of InputBotInlineMessageMediaVenue.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageMediaVenueTypeID is TL type id of InputBotInlineMessageMediaVenue.
const InputBotInlineMessageMediaVenueTypeID = 0x417bbf11

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageMediaVenue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageMediaVenue#417bbf11 as nil")
	}
	b.PutID(InputBotInlineMessageMediaVenueTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaVenue#417bbf11: field flags: %w", err)
	}
	if i.GeoPoint == nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaVenue#417bbf11: field geo_point is nil")
	}
	if err := i.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaVenue#417bbf11: field geo_point: %w", err)
	}
	b.PutString(i.Title)
	b.PutString(i.Address)
	b.PutString(i.Provider)
	b.PutString(i.VenueID)
	b.PutString(i.VenueType)
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaVenue#417bbf11: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaVenue#417bbf11: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageMediaVenue) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaVenue) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageMediaVenue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageMediaVenue#417bbf11 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageMediaVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field geo_point: %w", err)
		}
		i.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field title: %w", err)
		}
		i.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field address: %w", err)
		}
		i.Address = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field provider: %w", err)
		}
		i.Provider = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field venue_id: %w", err)
		}
		i.VenueID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field venue_type: %w", err)
		}
		i.VenueType = value
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaVenue#417bbf11: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageMediaVenue) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageMediaVenue.
var (
	_ bin.Encoder = &InputBotInlineMessageMediaVenue{}
	_ bin.Decoder = &InputBotInlineMessageMediaVenue{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageMediaVenue{}
)

// InputBotInlineMessageMediaContact represents TL type `inputBotInlineMessageMediaContact#a6edbffd`.
type InputBotInlineMessageMediaContact struct {
	// Flags field of InputBotInlineMessageMediaContact.
	Flags bin.Fields
	// PhoneNumber field of InputBotInlineMessageMediaContact.
	PhoneNumber string
	// FirstName field of InputBotInlineMessageMediaContact.
	FirstName string
	// LastName field of InputBotInlineMessageMediaContact.
	LastName string
	// Vcard field of InputBotInlineMessageMediaContact.
	Vcard string
	// ReplyMarkup field of InputBotInlineMessageMediaContact.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageMediaContactTypeID is TL type id of InputBotInlineMessageMediaContact.
const InputBotInlineMessageMediaContactTypeID = 0xa6edbffd

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageMediaContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageMediaContact#a6edbffd as nil")
	}
	b.PutID(InputBotInlineMessageMediaContactTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageMediaContact#a6edbffd: field flags: %w", err)
	}
	b.PutString(i.PhoneNumber)
	b.PutString(i.FirstName)
	b.PutString(i.LastName)
	b.PutString(i.Vcard)
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaContact#a6edbffd: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageMediaContact#a6edbffd: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageMediaContact) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageMediaContact) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageMediaContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageMediaContact#a6edbffd to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageMediaContactTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field phone_number: %w", err)
		}
		i.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field first_name: %w", err)
		}
		i.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field last_name: %w", err)
		}
		i.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field vcard: %w", err)
		}
		i.Vcard = value
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageMediaContact#a6edbffd: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageMediaContact) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageMediaContact.
var (
	_ bin.Encoder = &InputBotInlineMessageMediaContact{}
	_ bin.Decoder = &InputBotInlineMessageMediaContact{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageMediaContact{}
)

// InputBotInlineMessageGame represents TL type `inputBotInlineMessageGame#4b425864`.
type InputBotInlineMessageGame struct {
	// Flags field of InputBotInlineMessageGame.
	Flags bin.Fields
	// ReplyMarkup field of InputBotInlineMessageGame.
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// InputBotInlineMessageGameTypeID is TL type id of InputBotInlineMessageGame.
const InputBotInlineMessageGameTypeID = 0x4b425864

// Encode implements bin.Encoder.
func (i *InputBotInlineMessageGame) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineMessageGame#4b425864 as nil")
	}
	b.PutID(InputBotInlineMessageGameTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineMessageGame#4b425864: field flags: %w", err)
	}
	if i.Flags.Has(2) {
		if i.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageGame#4b425864: field reply_markup is nil")
		}
		if err := i.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineMessageGame#4b425864: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (i *InputBotInlineMessageGame) SetReplyMarkup(value ReplyMarkupClass) {
	i.Flags.Set(2)
	i.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineMessageGame) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineMessageGame) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineMessageGame#4b425864 to nil")
	}
	if err := b.ConsumeID(InputBotInlineMessageGameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineMessageGame#4b425864: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageGame#4b425864: field flags: %w", err)
		}
	}
	if i.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineMessageGame#4b425864: field reply_markup: %w", err)
		}
		i.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of InputBotInlineMessageClass.
func (i InputBotInlineMessageGame) construct() InputBotInlineMessageClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineMessageGame.
var (
	_ bin.Encoder = &InputBotInlineMessageGame{}
	_ bin.Decoder = &InputBotInlineMessageGame{}

	_ InputBotInlineMessageClass = &InputBotInlineMessageGame{}
)

// InputBotInlineMessageClass represents InputBotInlineMessage generic type.
//
// Example:
//  g, err := DecodeInputBotInlineMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputBotInlineMessageMediaAuto: // inputBotInlineMessageMediaAuto#3380c786
//  case *InputBotInlineMessageText: // inputBotInlineMessageText#3dcd7a87
//  case *InputBotInlineMessageMediaGeo: // inputBotInlineMessageMediaGeo#96929a85
//  case *InputBotInlineMessageMediaVenue: // inputBotInlineMessageMediaVenue#417bbf11
//  case *InputBotInlineMessageMediaContact: // inputBotInlineMessageMediaContact#a6edbffd
//  case *InputBotInlineMessageGame: // inputBotInlineMessageGame#4b425864
//  default: panic(v)
//  }
type InputBotInlineMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputBotInlineMessageClass
}

// DecodeInputBotInlineMessage implements binary de-serialization for InputBotInlineMessageClass.
func DecodeInputBotInlineMessage(buf *bin.Buffer) (InputBotInlineMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputBotInlineMessageMediaAutoTypeID:
		// Decoding inputBotInlineMessageMediaAuto#3380c786.
		v := InputBotInlineMessageMediaAuto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	case InputBotInlineMessageTextTypeID:
		// Decoding inputBotInlineMessageText#3dcd7a87.
		v := InputBotInlineMessageText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	case InputBotInlineMessageMediaGeoTypeID:
		// Decoding inputBotInlineMessageMediaGeo#96929a85.
		v := InputBotInlineMessageMediaGeo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	case InputBotInlineMessageMediaVenueTypeID:
		// Decoding inputBotInlineMessageMediaVenue#417bbf11.
		v := InputBotInlineMessageMediaVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	case InputBotInlineMessageMediaContactTypeID:
		// Decoding inputBotInlineMessageMediaContact#a6edbffd.
		v := InputBotInlineMessageMediaContact{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	case InputBotInlineMessageGameTypeID:
		// Decoding inputBotInlineMessageGame#4b425864.
		v := InputBotInlineMessageGame{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputBotInlineMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputBotInlineMessage boxes the InputBotInlineMessageClass providing a helper.
type InputBotInlineMessageBox struct {
	InputBotInlineMessage InputBotInlineMessageClass
}

// Decode implements bin.Decoder for InputBotInlineMessageBox.
func (b *InputBotInlineMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputBotInlineMessageBox to nil")
	}
	v, err := DecodeInputBotInlineMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputBotInlineMessage = v
	return nil
}

// Encode implements bin.Encode for InputBotInlineMessageBox.
func (b *InputBotInlineMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputBotInlineMessage == nil {
		return fmt.Errorf("unable to encode InputBotInlineMessageClass as nil")
	}
	return b.InputBotInlineMessage.Encode(buf)
}
