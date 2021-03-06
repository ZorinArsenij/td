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

// BotCommand represents TL type `botCommand#c27ac8c7`.
// Describes a bot command that can be used in a chat
//
// See https://core.telegram.org/constructor/botCommand for reference.
type BotCommand struct {
	// /command name
	Command string
	// Description of the command
	Description string
}

// BotCommandTypeID is TL type id of BotCommand.
const BotCommandTypeID = 0xc27ac8c7

func (b *BotCommand) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Command == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommand) String() string {
	if b == nil {
		return "BotCommand(nil)"
	}
	type Alias BotCommand
	return fmt.Sprintf("BotCommand%+v", Alias(*b))
}

// FillFrom fills BotCommand from given interface.
func (b *BotCommand) FillFrom(from interface {
	GetCommand() (value string)
	GetDescription() (value string)
}) {
	b.Command = from.GetCommand()
	b.Description = from.GetDescription()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommand) TypeID() uint32 {
	return BotCommandTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommand) TypeName() string {
	return "botCommand"
}

// TypeInfo returns info about TL type.
func (b *BotCommand) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommand",
		ID:   BotCommandTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Command",
			SchemaName: "command",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommand) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommand#c27ac8c7 as nil")
	}
	buf.PutID(BotCommandTypeID)
	buf.PutString(b.Command)
	buf.PutString(b.Description)
	return nil
}

// GetCommand returns value of Command field.
func (b *BotCommand) GetCommand() (value string) {
	return b.Command
}

// GetDescription returns value of Description field.
func (b *BotCommand) GetDescription() (value string) {
	return b.Description
}

// Decode implements bin.Decoder.
func (b *BotCommand) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommand#c27ac8c7 to nil")
	}
	if err := buf.ConsumeID(BotCommandTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommand#c27ac8c7: %w", err)
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botCommand#c27ac8c7: field command: %w", err)
		}
		b.Command = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botCommand#c27ac8c7: field description: %w", err)
		}
		b.Description = value
	}
	return nil
}

// Ensuring interfaces in compile-time for BotCommand.
var (
	_ bin.Encoder = &BotCommand{}
	_ bin.Decoder = &BotCommand{}
)
