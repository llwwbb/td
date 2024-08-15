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

// GetChatInviteLinkMembersRequest represents TL type `getChatInviteLinkMembers#6704ed3c`.
type GetChatInviteLinkMembersRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link for which to return chat members
	InviteLink string
	// Pass true if the link is a subscription link and only members with expired
	// subscription must be returned
	OnlyWithExpiredSubscription bool
	// A chat member from which to return next chat members; pass null to get results from
	// the beginning
	OffsetMember ChatInviteLinkMember
	// The maximum number of chat members to return; up to 100
	Limit int32
}

// GetChatInviteLinkMembersRequestTypeID is TL type id of GetChatInviteLinkMembersRequest.
const GetChatInviteLinkMembersRequestTypeID = 0x6704ed3c

// Ensuring interfaces in compile-time for GetChatInviteLinkMembersRequest.
var (
	_ bin.Encoder     = &GetChatInviteLinkMembersRequest{}
	_ bin.Decoder     = &GetChatInviteLinkMembersRequest{}
	_ bin.BareEncoder = &GetChatInviteLinkMembersRequest{}
	_ bin.BareDecoder = &GetChatInviteLinkMembersRequest{}
)

func (g *GetChatInviteLinkMembersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.InviteLink == "") {
		return false
	}
	if !(g.OnlyWithExpiredSubscription == false) {
		return false
	}
	if !(g.OffsetMember.Zero()) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatInviteLinkMembersRequest) String() string {
	if g == nil {
		return "GetChatInviteLinkMembersRequest(nil)"
	}
	type Alias GetChatInviteLinkMembersRequest
	return fmt.Sprintf("GetChatInviteLinkMembersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatInviteLinkMembersRequest) TypeID() uint32 {
	return GetChatInviteLinkMembersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatInviteLinkMembersRequest) TypeName() string {
	return "getChatInviteLinkMembers"
}

// TypeInfo returns info about TL type.
func (g *GetChatInviteLinkMembersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatInviteLinkMembers",
		ID:   GetChatInviteLinkMembersRequestTypeID,
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
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "OnlyWithExpiredSubscription",
			SchemaName: "only_with_expired_subscription",
		},
		{
			Name:       "OffsetMember",
			SchemaName: "offset_member",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatInviteLinkMembersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#6704ed3c as nil")
	}
	b.PutID(GetChatInviteLinkMembersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatInviteLinkMembersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#6704ed3c as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.InviteLink)
	b.PutBool(g.OnlyWithExpiredSubscription)
	if err := g.OffsetMember.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatInviteLinkMembers#6704ed3c: field offset_member: %w", err)
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatInviteLinkMembersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#6704ed3c to nil")
	}
	if err := b.ConsumeID(GetChatInviteLinkMembersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatInviteLinkMembersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#6704ed3c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field invite_link: %w", err)
		}
		g.InviteLink = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field only_with_expired_subscription: %w", err)
		}
		g.OnlyWithExpiredSubscription = value
	}
	{
		if err := g.OffsetMember.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field offset_member: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatInviteLinkMembersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#6704ed3c as nil")
	}
	b.ObjStart()
	b.PutID("getChatInviteLinkMembers")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(g.InviteLink)
	b.Comma()
	b.FieldStart("only_with_expired_subscription")
	b.PutBool(g.OnlyWithExpiredSubscription)
	b.Comma()
	b.FieldStart("offset_member")
	if err := g.OffsetMember.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatInviteLinkMembers#6704ed3c: field offset_member: %w", err)
	}
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatInviteLinkMembersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#6704ed3c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatInviteLinkMembers"); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field chat_id: %w", err)
			}
			g.ChatID = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field invite_link: %w", err)
			}
			g.InviteLink = value
		case "only_with_expired_subscription":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field only_with_expired_subscription: %w", err)
			}
			g.OnlyWithExpiredSubscription = value
		case "offset_member":
			if err := g.OffsetMember.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field offset_member: %w", err)
			}
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#6704ed3c: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatInviteLinkMembersRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (g *GetChatInviteLinkMembersRequest) GetInviteLink() (value string) {
	if g == nil {
		return
	}
	return g.InviteLink
}

// GetOnlyWithExpiredSubscription returns value of OnlyWithExpiredSubscription field.
func (g *GetChatInviteLinkMembersRequest) GetOnlyWithExpiredSubscription() (value bool) {
	if g == nil {
		return
	}
	return g.OnlyWithExpiredSubscription
}

// GetOffsetMember returns value of OffsetMember field.
func (g *GetChatInviteLinkMembersRequest) GetOffsetMember() (value ChatInviteLinkMember) {
	if g == nil {
		return
	}
	return g.OffsetMember
}

// GetLimit returns value of Limit field.
func (g *GetChatInviteLinkMembersRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatInviteLinkMembers invokes method getChatInviteLinkMembers#6704ed3c returning error if any.
func (c *Client) GetChatInviteLinkMembers(ctx context.Context, request *GetChatInviteLinkMembersRequest) (*ChatInviteLinkMembers, error) {
	var result ChatInviteLinkMembers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
