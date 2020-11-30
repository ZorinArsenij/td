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

// WebPageEmpty represents TL type `webPageEmpty#eb1477e8`.
type WebPageEmpty struct {
	// ID field of WebPageEmpty.
	ID int64
}

// WebPageEmptyTypeID is TL type id of WebPageEmpty.
const WebPageEmptyTypeID = 0xeb1477e8

// Encode implements bin.Encoder.
func (w *WebPageEmpty) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageEmpty#eb1477e8 as nil")
	}
	b.PutID(WebPageEmptyTypeID)
	b.PutLong(w.ID)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageEmpty) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageEmpty#eb1477e8 to nil")
	}
	if err := b.ConsumeID(WebPageEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageEmpty#eb1477e8: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPageEmpty#eb1477e8: field id: %w", err)
		}
		w.ID = value
	}
	return nil
}

// construct implements constructor of WebPageClass.
func (w WebPageEmpty) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPageEmpty.
var (
	_ bin.Encoder = &WebPageEmpty{}
	_ bin.Decoder = &WebPageEmpty{}

	_ WebPageClass = &WebPageEmpty{}
)

// WebPagePending represents TL type `webPagePending#c586da1c`.
type WebPagePending struct {
	// ID field of WebPagePending.
	ID int64
	// Date field of WebPagePending.
	Date int
}

// WebPagePendingTypeID is TL type id of WebPagePending.
const WebPagePendingTypeID = 0xc586da1c

// Encode implements bin.Encoder.
func (w *WebPagePending) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPagePending#c586da1c as nil")
	}
	b.PutID(WebPagePendingTypeID)
	b.PutLong(w.ID)
	b.PutInt(w.Date)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPagePending) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPagePending#c586da1c to nil")
	}
	if err := b.ConsumeID(WebPagePendingTypeID); err != nil {
		return fmt.Errorf("unable to decode webPagePending#c586da1c: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPagePending#c586da1c: field id: %w", err)
		}
		w.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPagePending#c586da1c: field date: %w", err)
		}
		w.Date = value
	}
	return nil
}

// construct implements constructor of WebPageClass.
func (w WebPagePending) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPagePending.
var (
	_ bin.Encoder = &WebPagePending{}
	_ bin.Decoder = &WebPagePending{}

	_ WebPageClass = &WebPagePending{}
)

// WebPage represents TL type `webPage#e89c45b2`.
type WebPage struct {
	// Flags field of WebPage.
	Flags bin.Fields
	// ID field of WebPage.
	ID int64
	// URL field of WebPage.
	URL string
	// DisplayURL field of WebPage.
	DisplayURL string
	// Hash field of WebPage.
	Hash int
	// Type field of WebPage.
	//
	// Use SetType and GetType helpers.
	Type string
	// SiteName field of WebPage.
	//
	// Use SetSiteName and GetSiteName helpers.
	SiteName string
	// Title field of WebPage.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description field of WebPage.
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// Photo field of WebPage.
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo PhotoClass
	// EmbedURL field of WebPage.
	//
	// Use SetEmbedURL and GetEmbedURL helpers.
	EmbedURL string
	// EmbedType field of WebPage.
	//
	// Use SetEmbedType and GetEmbedType helpers.
	EmbedType string
	// EmbedWidth field of WebPage.
	//
	// Use SetEmbedWidth and GetEmbedWidth helpers.
	EmbedWidth int
	// EmbedHeight field of WebPage.
	//
	// Use SetEmbedHeight and GetEmbedHeight helpers.
	EmbedHeight int
	// Duration field of WebPage.
	//
	// Use SetDuration and GetDuration helpers.
	Duration int
	// Author field of WebPage.
	//
	// Use SetAuthor and GetAuthor helpers.
	Author string
	// Document field of WebPage.
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// CachedPage field of WebPage.
	//
	// Use SetCachedPage and GetCachedPage helpers.
	CachedPage Page
	// Attributes field of WebPage.
	//
	// Use SetAttributes and GetAttributes helpers.
	Attributes []WebPageAttributeTheme
}

// WebPageTypeID is TL type id of WebPage.
const WebPageTypeID = 0xe89c45b2

// Encode implements bin.Encoder.
func (w *WebPage) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#e89c45b2 as nil")
	}
	b.PutID(WebPageTypeID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#e89c45b2: field flags: %w", err)
	}
	b.PutLong(w.ID)
	b.PutString(w.URL)
	b.PutString(w.DisplayURL)
	b.PutInt(w.Hash)
	if w.Flags.Has(0) {
		b.PutString(w.Type)
	}
	if w.Flags.Has(1) {
		b.PutString(w.SiteName)
	}
	if w.Flags.Has(2) {
		b.PutString(w.Title)
	}
	if w.Flags.Has(3) {
		b.PutString(w.Description)
	}
	if w.Flags.Has(4) {
		if w.Photo == nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field photo is nil")
		}
		if err := w.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field photo: %w", err)
		}
	}
	if w.Flags.Has(5) {
		b.PutString(w.EmbedURL)
	}
	if w.Flags.Has(5) {
		b.PutString(w.EmbedType)
	}
	if w.Flags.Has(6) {
		b.PutInt(w.EmbedWidth)
	}
	if w.Flags.Has(6) {
		b.PutInt(w.EmbedHeight)
	}
	if w.Flags.Has(7) {
		b.PutInt(w.Duration)
	}
	if w.Flags.Has(8) {
		b.PutString(w.Author)
	}
	if w.Flags.Has(9) {
		if w.Document == nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field document is nil")
		}
		if err := w.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field document: %w", err)
		}
	}
	if w.Flags.Has(10) {
		if err := w.CachedPage.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPage#e89c45b2: field cached_page: %w", err)
		}
	}
	if w.Flags.Has(12) {
		b.PutVectorHeader(len(w.Attributes))
		for idx, v := range w.Attributes {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode webPage#e89c45b2: field attributes element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetType sets value of Type conditional field.
func (w *WebPage) SetType(value string) {
	w.Flags.Set(0)
	w.Type = value
}

// GetType returns value of Type conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetType() (value string, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Type, true
}

// SetSiteName sets value of SiteName conditional field.
func (w *WebPage) SetSiteName(value string) {
	w.Flags.Set(1)
	w.SiteName = value
}

// GetSiteName returns value of SiteName conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetSiteName() (value string, ok bool) {
	if !w.Flags.Has(1) {
		return value, false
	}
	return w.SiteName, true
}

// SetTitle sets value of Title conditional field.
func (w *WebPage) SetTitle(value string) {
	w.Flags.Set(2)
	w.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetTitle() (value string, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Title, true
}

// SetDescription sets value of Description conditional field.
func (w *WebPage) SetDescription(value string) {
	w.Flags.Set(3)
	w.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDescription() (value string, ok bool) {
	if !w.Flags.Has(3) {
		return value, false
	}
	return w.Description, true
}

// SetPhoto sets value of Photo conditional field.
func (w *WebPage) SetPhoto(value PhotoClass) {
	w.Flags.Set(4)
	w.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetPhoto() (value PhotoClass, ok bool) {
	if !w.Flags.Has(4) {
		return value, false
	}
	return w.Photo, true
}

// SetEmbedURL sets value of EmbedURL conditional field.
func (w *WebPage) SetEmbedURL(value string) {
	w.Flags.Set(5)
	w.EmbedURL = value
}

// GetEmbedURL returns value of EmbedURL conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedURL() (value string, ok bool) {
	if !w.Flags.Has(5) {
		return value, false
	}
	return w.EmbedURL, true
}

// SetEmbedType sets value of EmbedType conditional field.
func (w *WebPage) SetEmbedType(value string) {
	w.Flags.Set(5)
	w.EmbedType = value
}

// GetEmbedType returns value of EmbedType conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedType() (value string, ok bool) {
	if !w.Flags.Has(5) {
		return value, false
	}
	return w.EmbedType, true
}

// SetEmbedWidth sets value of EmbedWidth conditional field.
func (w *WebPage) SetEmbedWidth(value int) {
	w.Flags.Set(6)
	w.EmbedWidth = value
}

// GetEmbedWidth returns value of EmbedWidth conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedWidth() (value int, ok bool) {
	if !w.Flags.Has(6) {
		return value, false
	}
	return w.EmbedWidth, true
}

// SetEmbedHeight sets value of EmbedHeight conditional field.
func (w *WebPage) SetEmbedHeight(value int) {
	w.Flags.Set(6)
	w.EmbedHeight = value
}

// GetEmbedHeight returns value of EmbedHeight conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetEmbedHeight() (value int, ok bool) {
	if !w.Flags.Has(6) {
		return value, false
	}
	return w.EmbedHeight, true
}

// SetDuration sets value of Duration conditional field.
func (w *WebPage) SetDuration(value int) {
	w.Flags.Set(7)
	w.Duration = value
}

// GetDuration returns value of Duration conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDuration() (value int, ok bool) {
	if !w.Flags.Has(7) {
		return value, false
	}
	return w.Duration, true
}

// SetAuthor sets value of Author conditional field.
func (w *WebPage) SetAuthor(value string) {
	w.Flags.Set(8)
	w.Author = value
}

// GetAuthor returns value of Author conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetAuthor() (value string, ok bool) {
	if !w.Flags.Has(8) {
		return value, false
	}
	return w.Author, true
}

// SetDocument sets value of Document conditional field.
func (w *WebPage) SetDocument(value DocumentClass) {
	w.Flags.Set(9)
	w.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetDocument() (value DocumentClass, ok bool) {
	if !w.Flags.Has(9) {
		return value, false
	}
	return w.Document, true
}

// SetCachedPage sets value of CachedPage conditional field.
func (w *WebPage) SetCachedPage(value Page) {
	w.Flags.Set(10)
	w.CachedPage = value
}

// GetCachedPage returns value of CachedPage conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetCachedPage() (value Page, ok bool) {
	if !w.Flags.Has(10) {
		return value, false
	}
	return w.CachedPage, true
}

// SetAttributes sets value of Attributes conditional field.
func (w *WebPage) SetAttributes(value []WebPageAttributeTheme) {
	w.Flags.Set(12)
	w.Attributes = value
}

// GetAttributes returns value of Attributes conditional field and
// boolean which is true if field was set.
func (w *WebPage) GetAttributes() (value []WebPageAttributeTheme, ok bool) {
	if !w.Flags.Has(12) {
		return value, false
	}
	return w.Attributes, true
}

// Decode implements bin.Decoder.
func (w *WebPage) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#e89c45b2 to nil")
	}
	if err := b.ConsumeID(WebPageTypeID); err != nil {
		return fmt.Errorf("unable to decode webPage#e89c45b2: %w", err)
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field id: %w", err)
		}
		w.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field display_url: %w", err)
		}
		w.DisplayURL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field hash: %w", err)
		}
		w.Hash = value
	}
	if w.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field type: %w", err)
		}
		w.Type = value
	}
	if w.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field site_name: %w", err)
		}
		w.SiteName = value
	}
	if w.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field title: %w", err)
		}
		w.Title = value
	}
	if w.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field description: %w", err)
		}
		w.Description = value
	}
	if w.Flags.Has(4) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field photo: %w", err)
		}
		w.Photo = value
	}
	if w.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_url: %w", err)
		}
		w.EmbedURL = value
	}
	if w.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_type: %w", err)
		}
		w.EmbedType = value
	}
	if w.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_width: %w", err)
		}
		w.EmbedWidth = value
	}
	if w.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field embed_height: %w", err)
		}
		w.EmbedHeight = value
	}
	if w.Flags.Has(7) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field duration: %w", err)
		}
		w.Duration = value
	}
	if w.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field author: %w", err)
		}
		w.Author = value
	}
	if w.Flags.Has(9) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field document: %w", err)
		}
		w.Document = value
	}
	if w.Flags.Has(10) {
		if err := w.CachedPage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field cached_page: %w", err)
		}
	}
	if w.Flags.Has(12) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#e89c45b2: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value WebPageAttributeTheme
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode webPage#e89c45b2: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of WebPageClass.
func (w WebPage) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPage.
var (
	_ bin.Encoder = &WebPage{}
	_ bin.Decoder = &WebPage{}

	_ WebPageClass = &WebPage{}
)

// WebPageNotModified represents TL type `webPageNotModified#7311ca11`.
type WebPageNotModified struct {
	// Flags field of WebPageNotModified.
	Flags bin.Fields
	// CachedPageViews field of WebPageNotModified.
	//
	// Use SetCachedPageViews and GetCachedPageViews helpers.
	CachedPageViews int
}

// WebPageNotModifiedTypeID is TL type id of WebPageNotModified.
const WebPageNotModifiedTypeID = 0x7311ca11

// Encode implements bin.Encoder.
func (w *WebPageNotModified) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageNotModified#7311ca11 as nil")
	}
	b.PutID(WebPageNotModifiedTypeID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageNotModified#7311ca11: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutInt(w.CachedPageViews)
	}
	return nil
}

// SetCachedPageViews sets value of CachedPageViews conditional field.
func (w *WebPageNotModified) SetCachedPageViews(value int) {
	w.Flags.Set(0)
	w.CachedPageViews = value
}

// GetCachedPageViews returns value of CachedPageViews conditional field and
// boolean which is true if field was set.
func (w *WebPageNotModified) GetCachedPageViews() (value int, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.CachedPageViews, true
}

// Decode implements bin.Decoder.
func (w *WebPageNotModified) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageNotModified#7311ca11 to nil")
	}
	if err := b.ConsumeID(WebPageNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageNotModified#7311ca11: %w", err)
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageNotModified#7311ca11: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPageNotModified#7311ca11: field cached_page_views: %w", err)
		}
		w.CachedPageViews = value
	}
	return nil
}

// construct implements constructor of WebPageClass.
func (w WebPageNotModified) construct() WebPageClass { return &w }

// Ensuring interfaces in compile-time for WebPageNotModified.
var (
	_ bin.Encoder = &WebPageNotModified{}
	_ bin.Decoder = &WebPageNotModified{}

	_ WebPageClass = &WebPageNotModified{}
)

// WebPageClass represents WebPage generic type.
//
// Example:
//  g, err := DecodeWebPage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *WebPageEmpty: // webPageEmpty#eb1477e8
//  case *WebPagePending: // webPagePending#c586da1c
//  case *WebPage: // webPage#e89c45b2
//  case *WebPageNotModified: // webPageNotModified#7311ca11
//  default: panic(v)
//  }
type WebPageClass interface {
	bin.Encoder
	bin.Decoder
	construct() WebPageClass
}

// DecodeWebPage implements binary de-serialization for WebPageClass.
func DecodeWebPage(buf *bin.Buffer) (WebPageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WebPageEmptyTypeID:
		// Decoding webPageEmpty#eb1477e8.
		v := WebPageEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPagePendingTypeID:
		// Decoding webPagePending#c586da1c.
		v := WebPagePending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPageTypeID:
		// Decoding webPage#e89c45b2.
		v := WebPage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	case WebPageNotModifiedTypeID:
		// Decoding webPageNotModified#7311ca11.
		v := WebPageNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WebPageClass: %w", bin.NewUnexpectedID(id))
	}
}

// WebPage boxes the WebPageClass providing a helper.
type WebPageBox struct {
	WebPage WebPageClass
}

// Decode implements bin.Decoder for WebPageBox.
func (b *WebPageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WebPageBox to nil")
	}
	v, err := DecodeWebPage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WebPage = v
	return nil
}

// Encode implements bin.Encode for WebPageBox.
func (b *WebPageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WebPage == nil {
		return fmt.Errorf("unable to encode WebPageClass as nil")
	}
	return b.WebPage.Encode(buf)
}
