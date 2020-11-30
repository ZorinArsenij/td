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

// AccountTmpPassword represents TL type `account.tmpPassword#db64fd34`.
type AccountTmpPassword struct {
	// TmpPassword field of AccountTmpPassword.
	TmpPassword []byte
	// ValidUntil field of AccountTmpPassword.
	ValidUntil int
}

// AccountTmpPasswordTypeID is TL type id of AccountTmpPassword.
const AccountTmpPasswordTypeID = 0xdb64fd34

// Encode implements bin.Encoder.
func (t *AccountTmpPassword) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.tmpPassword#db64fd34 as nil")
	}
	b.PutID(AccountTmpPasswordTypeID)
	b.PutBytes(t.TmpPassword)
	b.PutInt(t.ValidUntil)
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountTmpPassword) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.tmpPassword#db64fd34 to nil")
	}
	if err := b.ConsumeID(AccountTmpPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field tmp_password: %w", err)
		}
		t.TmpPassword = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field valid_until: %w", err)
		}
		t.ValidUntil = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountTmpPassword.
var (
	_ bin.Encoder = &AccountTmpPassword{}
	_ bin.Decoder = &AccountTmpPassword{}
)
