// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// RPCDropAnswerRequest represents TL type `rpc_drop_answer#58e4a740`.
type RPCDropAnswerRequest struct {
	// ReqMsgID field of RPCDropAnswerRequest.
	ReqMsgID int64
}

// RPCDropAnswerRequestTypeID is TL type id of RPCDropAnswerRequest.
const RPCDropAnswerRequestTypeID = 0x58e4a740

func (r *RPCDropAnswerRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ReqMsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCDropAnswerRequest) String() string {
	if r == nil {
		return "RPCDropAnswerRequest(nil)"
	}
	type Alias RPCDropAnswerRequest
	return fmt.Sprintf("RPCDropAnswerRequest%+v", Alias(*r))
}

// FillFrom fills RPCDropAnswerRequest from given interface.
func (r *RPCDropAnswerRequest) FillFrom(from interface {
	GetReqMsgID() (value int64)
}) {
	r.ReqMsgID = from.GetReqMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCDropAnswerRequest) TypeID() uint32 {
	return RPCDropAnswerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCDropAnswerRequest) TypeName() string {
	return "rpc_drop_answer"
}

// TypeInfo returns info about TL type.
func (r *RPCDropAnswerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_drop_answer",
		ID:   RPCDropAnswerRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReqMsgID",
			SchemaName: "req_msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCDropAnswerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_drop_answer#58e4a740 as nil")
	}
	b.PutID(RPCDropAnswerRequestTypeID)
	b.PutLong(r.ReqMsgID)
	return nil
}

// GetReqMsgID returns value of ReqMsgID field.
func (r *RPCDropAnswerRequest) GetReqMsgID() (value int64) {
	return r.ReqMsgID
}

// Decode implements bin.Decoder.
func (r *RPCDropAnswerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_drop_answer#58e4a740 to nil")
	}
	if err := b.ConsumeID(RPCDropAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_drop_answer#58e4a740: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_drop_answer#58e4a740: field req_msg_id: %w", err)
		}
		r.ReqMsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RPCDropAnswerRequest.
var (
	_ bin.Encoder = &RPCDropAnswerRequest{}
	_ bin.Decoder = &RPCDropAnswerRequest{}
)

// RPCDropAnswer invokes method rpc_drop_answer#58e4a740 returning error if any.
func (c *Client) RPCDropAnswer(ctx context.Context, reqmsgid int64) (RpcDropAnswerClass, error) {
	var result RpcDropAnswerBox

	request := &RPCDropAnswerRequest{
		ReqMsgID: reqmsgid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.RpcDropAnswer, nil
}
