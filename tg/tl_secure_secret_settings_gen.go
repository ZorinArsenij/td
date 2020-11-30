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

// SecureSecretSettings represents TL type `secureSecretSettings#1527bcac`.
type SecureSecretSettings struct {
	// SecureAlgo field of SecureSecretSettings.
	SecureAlgo SecurePasswordKdfAlgoClass
	// SecureSecret field of SecureSecretSettings.
	SecureSecret []byte
	// SecureSecretID field of SecureSecretSettings.
	SecureSecretID int64
}

// SecureSecretSettingsTypeID is TL type id of SecureSecretSettings.
const SecureSecretSettingsTypeID = 0x1527bcac

// Encode implements bin.Encoder.
func (s *SecureSecretSettings) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureSecretSettings#1527bcac as nil")
	}
	b.PutID(SecureSecretSettingsTypeID)
	if s.SecureAlgo == nil {
		return fmt.Errorf("unable to encode secureSecretSettings#1527bcac: field secure_algo is nil")
	}
	if err := s.SecureAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureSecretSettings#1527bcac: field secure_algo: %w", err)
	}
	b.PutBytes(s.SecureSecret)
	b.PutLong(s.SecureSecretID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureSecretSettings) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureSecretSettings#1527bcac to nil")
	}
	if err := b.ConsumeID(SecureSecretSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: %w", err)
	}
	{
		value, err := DecodeSecurePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_algo: %w", err)
		}
		s.SecureAlgo = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_secret: %w", err)
		}
		s.SecureSecret = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_secret_id: %w", err)
		}
		s.SecureSecretID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureSecretSettings.
var (
	_ bin.Encoder = &SecureSecretSettings{}
	_ bin.Decoder = &SecureSecretSettings{}
)
