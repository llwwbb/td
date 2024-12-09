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

// GetChatAffiliateProgramsRequest represents TL type `getChatAffiliatePrograms#f9df5251`.
type GetChatAffiliateProgramsRequest struct {
	// Identifier of the chat for which the affiliate programs were connected. Can be an
	// identifier of the Saved Messages chat, of a chat with an owned bot, or of a channel
	// chat with can_post_messages administrator right
	ChatID int64
	// Offset of the first affiliate program to return as received from the previous request;
	// use empty string to get the first chunk of results
	Offset string
	// The maximum number of affiliate programs to return
	Limit int32
}

// GetChatAffiliateProgramsRequestTypeID is TL type id of GetChatAffiliateProgramsRequest.
const GetChatAffiliateProgramsRequestTypeID = 0xf9df5251

// Ensuring interfaces in compile-time for GetChatAffiliateProgramsRequest.
var (
	_ bin.Encoder     = &GetChatAffiliateProgramsRequest{}
	_ bin.Decoder     = &GetChatAffiliateProgramsRequest{}
	_ bin.BareEncoder = &GetChatAffiliateProgramsRequest{}
	_ bin.BareDecoder = &GetChatAffiliateProgramsRequest{}
)

func (g *GetChatAffiliateProgramsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatAffiliateProgramsRequest) String() string {
	if g == nil {
		return "GetChatAffiliateProgramsRequest(nil)"
	}
	type Alias GetChatAffiliateProgramsRequest
	return fmt.Sprintf("GetChatAffiliateProgramsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatAffiliateProgramsRequest) TypeID() uint32 {
	return GetChatAffiliateProgramsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatAffiliateProgramsRequest) TypeName() string {
	return "getChatAffiliatePrograms"
}

// TypeInfo returns info about TL type.
func (g *GetChatAffiliateProgramsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatAffiliatePrograms",
		ID:   GetChatAffiliateProgramsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatAffiliateProgramsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAffiliatePrograms#f9df5251 as nil")
	}
	b.PutID(GetChatAffiliateProgramsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatAffiliateProgramsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAffiliatePrograms#f9df5251 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatAffiliateProgramsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAffiliatePrograms#f9df5251 to nil")
	}
	if err := b.ConsumeID(GetChatAffiliateProgramsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatAffiliateProgramsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAffiliatePrograms#f9df5251 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatAffiliateProgramsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatAffiliatePrograms#f9df5251 as nil")
	}
	b.ObjStart()
	b.PutID("getChatAffiliatePrograms")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatAffiliateProgramsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatAffiliatePrograms#f9df5251 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatAffiliatePrograms"); err != nil {
				return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field chat_id: %w", err)
			}
			g.ChatID = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatAffiliatePrograms#f9df5251: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatAffiliateProgramsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetOffset returns value of Offset field.
func (g *GetChatAffiliateProgramsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetChatAffiliateProgramsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatAffiliatePrograms invokes method getChatAffiliatePrograms#f9df5251 returning error if any.
func (c *Client) GetChatAffiliatePrograms(ctx context.Context, request *GetChatAffiliateProgramsRequest) (*ChatAffiliatePrograms, error) {
	var result ChatAffiliatePrograms

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}