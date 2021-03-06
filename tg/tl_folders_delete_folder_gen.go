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

// FoldersDeleteFolderRequest represents TL type `folders.deleteFolder#1c295881`.
// Delete a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/method/folders.deleteFolder for reference.
type FoldersDeleteFolderRequest struct {
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// FoldersDeleteFolderRequestTypeID is TL type id of FoldersDeleteFolderRequest.
const FoldersDeleteFolderRequestTypeID = 0x1c295881

func (d *FoldersDeleteFolderRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *FoldersDeleteFolderRequest) String() string {
	if d == nil {
		return "FoldersDeleteFolderRequest(nil)"
	}
	type Alias FoldersDeleteFolderRequest
	return fmt.Sprintf("FoldersDeleteFolderRequest%+v", Alias(*d))
}

// FillFrom fills FoldersDeleteFolderRequest from given interface.
func (d *FoldersDeleteFolderRequest) FillFrom(from interface {
	GetFolderID() (value int)
}) {
	d.FolderID = from.GetFolderID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FoldersDeleteFolderRequest) TypeID() uint32 {
	return FoldersDeleteFolderRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*FoldersDeleteFolderRequest) TypeName() string {
	return "folders.deleteFolder"
}

// TypeInfo returns info about TL type.
func (d *FoldersDeleteFolderRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "folders.deleteFolder",
		ID:   FoldersDeleteFolderRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *FoldersDeleteFolderRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode folders.deleteFolder#1c295881 as nil")
	}
	b.PutID(FoldersDeleteFolderRequestTypeID)
	b.PutInt(d.FolderID)
	return nil
}

// GetFolderID returns value of FolderID field.
func (d *FoldersDeleteFolderRequest) GetFolderID() (value int) {
	return d.FolderID
}

// Decode implements bin.Decoder.
func (d *FoldersDeleteFolderRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode folders.deleteFolder#1c295881 to nil")
	}
	if err := b.ConsumeID(FoldersDeleteFolderRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode folders.deleteFolder#1c295881: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode folders.deleteFolder#1c295881: field folder_id: %w", err)
		}
		d.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FoldersDeleteFolderRequest.
var (
	_ bin.Encoder = &FoldersDeleteFolderRequest{}
	_ bin.Decoder = &FoldersDeleteFolderRequest{}
)

// FoldersDeleteFolder invokes method folders.deleteFolder#1c295881 returning error if any.
// Delete a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/method/folders.deleteFolder for reference.
func (c *Client) FoldersDeleteFolder(ctx context.Context, folderid int) (UpdatesClass, error) {
	var result UpdatesBox

	request := &FoldersDeleteFolderRequest{
		FolderID: folderid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
