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

// AuthSendCodeRequest represents TL type `auth.sendCode#a677244f`.
type AuthSendCodeRequest struct {
	// PhoneNumber field of AuthSendCodeRequest.
	PhoneNumber string
	// APIID field of AuthSendCodeRequest.
	APIID int
	// APIHash field of AuthSendCodeRequest.
	APIHash string
	// Settings field of AuthSendCodeRequest.
	Settings CodeSettings
}

// AuthSendCodeRequestTypeID is TL type id of AuthSendCodeRequest.
const AuthSendCodeRequestTypeID = 0xa677244f

// Encode implements bin.Encoder.
func (s *AuthSendCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sendCode#a677244f as nil")
	}
	b.PutID(AuthSendCodeRequestTypeID)
	b.PutString(s.PhoneNumber)
	b.PutInt(s.APIID)
	b.PutString(s.APIHash)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.sendCode#a677244f: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSendCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sendCode#a677244f to nil")
	}
	if err := b.ConsumeID(AuthSendCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sendCode#a677244f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field api_id: %w", err)
		}
		s.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field api_hash: %w", err)
		}
		s.APIHash = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthSendCodeRequest.
var (
	_ bin.Encoder = &AuthSendCodeRequest{}
	_ bin.Decoder = &AuthSendCodeRequest{}
)
