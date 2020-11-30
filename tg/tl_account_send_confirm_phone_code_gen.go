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

// AccountSendConfirmPhoneCodeRequest represents TL type `account.sendConfirmPhoneCode#1b3faa88`.
type AccountSendConfirmPhoneCodeRequest struct {
	// Hash field of AccountSendConfirmPhoneCodeRequest.
	Hash string
	// Settings field of AccountSendConfirmPhoneCodeRequest.
	Settings CodeSettings
}

// AccountSendConfirmPhoneCodeRequestTypeID is TL type id of AccountSendConfirmPhoneCodeRequest.
const AccountSendConfirmPhoneCodeRequestTypeID = 0x1b3faa88

// Encode implements bin.Encoder.
func (s *AccountSendConfirmPhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendConfirmPhoneCode#1b3faa88 as nil")
	}
	b.PutID(AccountSendConfirmPhoneCodeRequestTypeID)
	b.PutString(s.Hash)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendConfirmPhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendConfirmPhoneCode#1b3faa88 to nil")
	}
	if err := b.ConsumeID(AccountSendConfirmPhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendConfirmPhoneCode#1b3faa88: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendConfirmPhoneCodeRequest.
var (
	_ bin.Encoder = &AccountSendConfirmPhoneCodeRequest{}
	_ bin.Decoder = &AccountSendConfirmPhoneCodeRequest{}
)
