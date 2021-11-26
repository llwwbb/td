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

// GetJSONValueRequest represents TL type `getJsonValue#92fa5a05`.
type GetJSONValueRequest struct {
	// The JSON-serialized string
	JSON string
}

// GetJSONValueRequestTypeID is TL type id of GetJSONValueRequest.
const GetJSONValueRequestTypeID = 0x92fa5a05

// Ensuring interfaces in compile-time for GetJSONValueRequest.
var (
	_ bin.Encoder     = &GetJSONValueRequest{}
	_ bin.Decoder     = &GetJSONValueRequest{}
	_ bin.BareEncoder = &GetJSONValueRequest{}
	_ bin.BareDecoder = &GetJSONValueRequest{}
)

func (g *GetJSONValueRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.JSON == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetJSONValueRequest) String() string {
	if g == nil {
		return "GetJSONValueRequest(nil)"
	}
	type Alias GetJSONValueRequest
	return fmt.Sprintf("GetJSONValueRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetJSONValueRequest) TypeID() uint32 {
	return GetJSONValueRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetJSONValueRequest) TypeName() string {
	return "getJsonValue"
}

// TypeInfo returns info about TL type.
func (g *GetJSONValueRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getJsonValue",
		ID:   GetJSONValueRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JSON",
			SchemaName: "json",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetJSONValueRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonValue#92fa5a05 as nil")
	}
	b.PutID(GetJSONValueRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetJSONValueRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonValue#92fa5a05 as nil")
	}
	b.PutString(g.JSON)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetJSONValueRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonValue#92fa5a05 to nil")
	}
	if err := b.ConsumeID(GetJSONValueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getJsonValue#92fa5a05: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetJSONValueRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonValue#92fa5a05 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getJsonValue#92fa5a05: field json: %w", err)
		}
		g.JSON = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetJSONValueRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonValue#92fa5a05 as nil")
	}
	b.ObjStart()
	b.PutID("getJsonValue")
	b.FieldStart("json")
	b.PutString(g.JSON)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetJSONValueRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonValue#92fa5a05 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getJsonValue"); err != nil {
				return fmt.Errorf("unable to decode getJsonValue#92fa5a05: %w", err)
			}
		case "json":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getJsonValue#92fa5a05: field json: %w", err)
			}
			g.JSON = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetJSON returns value of JSON field.
func (g *GetJSONValueRequest) GetJSON() (value string) {
	return g.JSON
}

// GetJSONValue invokes method getJsonValue#92fa5a05 returning error if any.
func (c *Client) GetJSONValue(ctx context.Context, json string) (JSONValueClass, error) {
	var result JSONValueBox

	request := &GetJSONValueRequest{
		JSON: json,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.JsonValue, nil
}