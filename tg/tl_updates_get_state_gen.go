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

// UpdatesGetStateRequest represents TL type `updates.getState#edd4882a`.
// Returns a current state of updates.
//
// See https://core.telegram.org/method/updates.getState for reference.
type UpdatesGetStateRequest struct {
}

// UpdatesGetStateRequestTypeID is TL type id of UpdatesGetStateRequest.
const UpdatesGetStateRequestTypeID = 0xedd4882a

func (g *UpdatesGetStateRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *UpdatesGetStateRequest) String() string {
	if g == nil {
		return "UpdatesGetStateRequest(nil)"
	}
	type Alias UpdatesGetStateRequest
	return fmt.Sprintf("UpdatesGetStateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesGetStateRequest) TypeID() uint32 {
	return UpdatesGetStateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesGetStateRequest) TypeName() string {
	return "updates.getState"
}

// TypeInfo returns info about TL type.
func (g *UpdatesGetStateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.getState",
		ID:   UpdatesGetStateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *UpdatesGetStateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getState#edd4882a as nil")
	}
	b.PutID(UpdatesGetStateRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *UpdatesGetStateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getState#edd4882a to nil")
	}
	if err := b.ConsumeID(UpdatesGetStateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getState#edd4882a: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesGetStateRequest.
var (
	_ bin.Encoder = &UpdatesGetStateRequest{}
	_ bin.Decoder = &UpdatesGetStateRequest{}
)

// UpdatesGetState invokes method updates.getState#edd4882a returning error if any.
// Returns a current state of updates.
//
// See https://core.telegram.org/method/updates.getState for reference.
// Can be used by bots.
func (c *Client) UpdatesGetState(ctx context.Context) (*UpdatesState, error) {
	var result UpdatesState

	request := &UpdatesGetStateRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
