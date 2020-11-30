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

// WebAuthorization represents TL type `webAuthorization#cac943f2`.
type WebAuthorization struct {
	// Hash field of WebAuthorization.
	Hash int64
	// BotID field of WebAuthorization.
	BotID int
	// Domain field of WebAuthorization.
	Domain string
	// Browser field of WebAuthorization.
	Browser string
	// Platform field of WebAuthorization.
	Platform string
	// DateCreated field of WebAuthorization.
	DateCreated int
	// DateActive field of WebAuthorization.
	DateActive int
	// IP field of WebAuthorization.
	IP string
	// Region field of WebAuthorization.
	Region string
}

// WebAuthorizationTypeID is TL type id of WebAuthorization.
const WebAuthorizationTypeID = 0xcac943f2

// Encode implements bin.Encoder.
func (w *WebAuthorization) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webAuthorization#cac943f2 as nil")
	}
	b.PutID(WebAuthorizationTypeID)
	b.PutLong(w.Hash)
	b.PutInt(w.BotID)
	b.PutString(w.Domain)
	b.PutString(w.Browser)
	b.PutString(w.Platform)
	b.PutInt(w.DateCreated)
	b.PutInt(w.DateActive)
	b.PutString(w.IP)
	b.PutString(w.Region)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebAuthorization) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webAuthorization#cac943f2 to nil")
	}
	if err := b.ConsumeID(WebAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode webAuthorization#cac943f2: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field hash: %w", err)
		}
		w.Hash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field bot_id: %w", err)
		}
		w.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field domain: %w", err)
		}
		w.Domain = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field browser: %w", err)
		}
		w.Browser = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field platform: %w", err)
		}
		w.Platform = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field date_created: %w", err)
		}
		w.DateCreated = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field date_active: %w", err)
		}
		w.DateActive = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field ip: %w", err)
		}
		w.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field region: %w", err)
		}
		w.Region = value
	}
	return nil
}

// Ensuring interfaces in compile-time for WebAuthorization.
var (
	_ bin.Encoder = &WebAuthorization{}
	_ bin.Decoder = &WebAuthorization{}
)
