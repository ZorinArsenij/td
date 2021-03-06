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

// AuthExportLoginTokenRequest represents TL type `auth.exportLoginToken#b1b41517`.
// Generate a login token, for login via QR code¹.
// The generated login token should be encoded using base64url, then shown as a tg://login?token=base64encodedtoken URL in the QR code.
// For more info, see login via QR code¹.
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//  2) https://core.telegram.org/api/qr-login
//
// See https://core.telegram.org/method/auth.exportLoginToken for reference.
type AuthExportLoginTokenRequest struct {
	// Application identifier (see. App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIID int
	// Application identifier hash (see. App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIHash string
	// List of already logged-in user IDs, to prevent logging in twice with the same user
	ExceptIds []int
}

// AuthExportLoginTokenRequestTypeID is TL type id of AuthExportLoginTokenRequest.
const AuthExportLoginTokenRequestTypeID = 0xb1b41517

func (e *AuthExportLoginTokenRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.APIID == 0) {
		return false
	}
	if !(e.APIHash == "") {
		return false
	}
	if !(e.ExceptIds == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AuthExportLoginTokenRequest) String() string {
	if e == nil {
		return "AuthExportLoginTokenRequest(nil)"
	}
	type Alias AuthExportLoginTokenRequest
	return fmt.Sprintf("AuthExportLoginTokenRequest%+v", Alias(*e))
}

// FillFrom fills AuthExportLoginTokenRequest from given interface.
func (e *AuthExportLoginTokenRequest) FillFrom(from interface {
	GetAPIID() (value int)
	GetAPIHash() (value string)
	GetExceptIds() (value []int)
}) {
	e.APIID = from.GetAPIID()
	e.APIHash = from.GetAPIHash()
	e.ExceptIds = from.GetExceptIds()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthExportLoginTokenRequest) TypeID() uint32 {
	return AuthExportLoginTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthExportLoginTokenRequest) TypeName() string {
	return "auth.exportLoginToken"
}

// TypeInfo returns info about TL type.
func (e *AuthExportLoginTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.exportLoginToken",
		ID:   AuthExportLoginTokenRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "APIHash",
			SchemaName: "api_hash",
		},
		{
			Name:       "ExceptIds",
			SchemaName: "except_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AuthExportLoginTokenRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportLoginToken#b1b41517 as nil")
	}
	b.PutID(AuthExportLoginTokenRequestTypeID)
	b.PutInt(e.APIID)
	b.PutString(e.APIHash)
	b.PutVectorHeader(len(e.ExceptIds))
	for _, v := range e.ExceptIds {
		b.PutInt(v)
	}
	return nil
}

// GetAPIID returns value of APIID field.
func (e *AuthExportLoginTokenRequest) GetAPIID() (value int) {
	return e.APIID
}

// GetAPIHash returns value of APIHash field.
func (e *AuthExportLoginTokenRequest) GetAPIHash() (value string) {
	return e.APIHash
}

// GetExceptIds returns value of ExceptIds field.
func (e *AuthExportLoginTokenRequest) GetExceptIds() (value []int) {
	return e.ExceptIds
}

// Decode implements bin.Decoder.
func (e *AuthExportLoginTokenRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportLoginToken#b1b41517 to nil")
	}
	if err := b.ConsumeID(AuthExportLoginTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportLoginToken#b1b41517: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b1b41517: field api_id: %w", err)
		}
		e.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b1b41517: field api_hash: %w", err)
		}
		e.APIHash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportLoginToken#b1b41517: field except_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode auth.exportLoginToken#b1b41517: field except_ids: %w", err)
			}
			e.ExceptIds = append(e.ExceptIds, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthExportLoginTokenRequest.
var (
	_ bin.Encoder = &AuthExportLoginTokenRequest{}
	_ bin.Decoder = &AuthExportLoginTokenRequest{}
)

// AuthExportLoginToken invokes method auth.exportLoginToken#b1b41517 returning error if any.
// Generate a login token, for login via QR code¹.
// The generated login token should be encoded using base64url, then shown as a tg://login?token=base64encodedtoken URL in the QR code.
// For more info, see login via QR code¹.
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//  2) https://core.telegram.org/api/qr-login
//
// See https://core.telegram.org/method/auth.exportLoginToken for reference.
func (c *Client) AuthExportLoginToken(ctx context.Context, request *AuthExportLoginTokenRequest) (AuthLoginTokenClass, error) {
	var result AuthLoginTokenBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LoginToken, nil
}
