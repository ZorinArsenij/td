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

// HelpGetDeepLinkInfoRequest represents TL type `help.getDeepLinkInfo#3fedc75f`.
type HelpGetDeepLinkInfoRequest struct {
	// Path field of HelpGetDeepLinkInfoRequest.
	Path string
}

// HelpGetDeepLinkInfoRequestTypeID is TL type id of HelpGetDeepLinkInfoRequest.
const HelpGetDeepLinkInfoRequestTypeID = 0x3fedc75f

// Encode implements bin.Encoder.
func (g *HelpGetDeepLinkInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getDeepLinkInfo#3fedc75f as nil")
	}
	b.PutID(HelpGetDeepLinkInfoRequestTypeID)
	b.PutString(g.Path)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetDeepLinkInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getDeepLinkInfo#3fedc75f to nil")
	}
	if err := b.ConsumeID(HelpGetDeepLinkInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getDeepLinkInfo#3fedc75f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getDeepLinkInfo#3fedc75f: field path: %w", err)
		}
		g.Path = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetDeepLinkInfoRequest.
var (
	_ bin.Encoder = &HelpGetDeepLinkInfoRequest{}
	_ bin.Decoder = &HelpGetDeepLinkInfoRequest{}
)
