// Code generated by gotdgen, DO NOT EDIT.

package td

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

// SendMultipleSMSRequest represents TL type `sendMultipleSMS#df18e5ca`.
//
// See https://localhost:80/doc/constructor/sendMultipleSMS for reference.
type SendMultipleSMSRequest struct {
	// Messages field of SendMultipleSMSRequest.
	Messages []SMS
}

// SendMultipleSMSRequestTypeID is TL type id of SendMultipleSMSRequest.
const SendMultipleSMSRequestTypeID = 0xdf18e5ca

func (s *SendMultipleSMSRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMultipleSMSRequest) String() string {
	if s == nil {
		return "SendMultipleSMSRequest(nil)"
	}
	type Alias SendMultipleSMSRequest
	return fmt.Sprintf("SendMultipleSMSRequest%+v", Alias(*s))
}

// FillFrom fills SendMultipleSMSRequest from given interface.
func (s *SendMultipleSMSRequest) FillFrom(from interface {
	GetMessages() (value []SMS)
}) {
	s.Messages = from.GetMessages()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendMultipleSMSRequest) TypeID() uint32 {
	return SendMultipleSMSRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendMultipleSMSRequest) TypeName() string {
	return "sendMultipleSMS"
}

// TypeInfo returns info about TL type.
func (s *SendMultipleSMSRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendMultipleSMS",
		ID:   SendMultipleSMSRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendMultipleSMSRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMultipleSMS#df18e5ca as nil")
	}
	b.PutID(SendMultipleSMSRequestTypeID)
	b.PutVectorHeader(len(s.Messages))
	for idx, v := range s.Messages {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sendMultipleSMS#df18e5ca: field messages element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetMessages returns value of Messages field.
func (s *SendMultipleSMSRequest) GetMessages() (value []SMS) {
	return s.Messages
}

// Decode implements bin.Decoder.
func (s *SendMultipleSMSRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMultipleSMS#df18e5ca to nil")
	}
	if err := b.ConsumeID(SendMultipleSMSRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SMS
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode sendMultipleSMS#df18e5ca: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SendMultipleSMSRequest.
var (
	_ bin.Encoder = &SendMultipleSMSRequest{}
	_ bin.Decoder = &SendMultipleSMSRequest{}
)

// SendMultipleSMS invokes method sendMultipleSMS#df18e5ca returning error if any.
//
// See https://localhost:80/doc/constructor/sendMultipleSMS for reference.
func (c *Client) SendMultipleSMS(ctx context.Context, messages []SMS) error {
	var ok Ok

	request := &SendMultipleSMSRequest{
		Messages: messages,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
