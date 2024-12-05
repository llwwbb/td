// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// GetAvailableGiftsRequest represents TL type `getAvailableGifts#f6d565eb`.
type GetAvailableGiftsRequest struct {
}

// GetAvailableGiftsRequestTypeID is TL type id of GetAvailableGiftsRequest.
const GetAvailableGiftsRequestTypeID = 0xf6d565eb

// Ensuring interfaces in compile-time for GetAvailableGiftsRequest.
var (
	_ bin.Encoder     = &GetAvailableGiftsRequest{}
	_ bin.Decoder     = &GetAvailableGiftsRequest{}
	_ bin.BareEncoder = &GetAvailableGiftsRequest{}
	_ bin.BareDecoder = &GetAvailableGiftsRequest{}
)

func (g *GetAvailableGiftsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetAvailableGiftsRequest) String() string {
	if g == nil {
		return "GetAvailableGiftsRequest(nil)"
	}
	type Alias GetAvailableGiftsRequest
	return fmt.Sprintf("GetAvailableGiftsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetAvailableGiftsRequest) TypeID() uint32 {
	return GetAvailableGiftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetAvailableGiftsRequest) TypeName() string {
	return "getAvailableGifts"
}

// TypeInfo returns info about TL type.
func (g *GetAvailableGiftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getAvailableGifts",
		ID:   GetAvailableGiftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetAvailableGiftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableGifts#f6d565eb as nil")
	}
	b.PutID(GetAvailableGiftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetAvailableGiftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableGifts#f6d565eb as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetAvailableGiftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableGifts#f6d565eb to nil")
	}
	if err := b.ConsumeID(GetAvailableGiftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getAvailableGifts#f6d565eb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetAvailableGiftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableGifts#f6d565eb to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetAvailableGiftsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getAvailableGifts#f6d565eb as nil")
	}
	b.ObjStart()
	b.PutID("getAvailableGifts")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetAvailableGiftsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getAvailableGifts#f6d565eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getAvailableGifts"); err != nil {
				return fmt.Errorf("unable to decode getAvailableGifts#f6d565eb: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAvailableGifts invokes method getAvailableGifts#f6d565eb returning error if any.
func (c *Client) GetAvailableGifts(ctx context.Context) (*Gifts, error) {
	var result Gifts

	request := &GetAvailableGiftsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}