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

// TogglePaidMessageReactionIsAnonymousRequest represents TL type `togglePaidMessageReactionIsAnonymous#9774db11`.
type TogglePaidMessageReactionIsAnonymousRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// Pass true to make paid reaction of the user on the message anonymous; pass false to
	// make the user's profile visible among top reactors
	IsAnonymous bool
}

// TogglePaidMessageReactionIsAnonymousRequestTypeID is TL type id of TogglePaidMessageReactionIsAnonymousRequest.
const TogglePaidMessageReactionIsAnonymousRequestTypeID = 0x9774db11

// Ensuring interfaces in compile-time for TogglePaidMessageReactionIsAnonymousRequest.
var (
	_ bin.Encoder     = &TogglePaidMessageReactionIsAnonymousRequest{}
	_ bin.Decoder     = &TogglePaidMessageReactionIsAnonymousRequest{}
	_ bin.BareEncoder = &TogglePaidMessageReactionIsAnonymousRequest{}
	_ bin.BareDecoder = &TogglePaidMessageReactionIsAnonymousRequest{}
)

func (t *TogglePaidMessageReactionIsAnonymousRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.MessageID == 0) {
		return false
	}
	if !(t.IsAnonymous == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TogglePaidMessageReactionIsAnonymousRequest) String() string {
	if t == nil {
		return "TogglePaidMessageReactionIsAnonymousRequest(nil)"
	}
	type Alias TogglePaidMessageReactionIsAnonymousRequest
	return fmt.Sprintf("TogglePaidMessageReactionIsAnonymousRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TogglePaidMessageReactionIsAnonymousRequest) TypeID() uint32 {
	return TogglePaidMessageReactionIsAnonymousRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TogglePaidMessageReactionIsAnonymousRequest) TypeName() string {
	return "togglePaidMessageReactionIsAnonymous"
}

// TypeInfo returns info about TL type.
func (t *TogglePaidMessageReactionIsAnonymousRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "togglePaidMessageReactionIsAnonymous",
		ID:   TogglePaidMessageReactionIsAnonymousRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "IsAnonymous",
			SchemaName: "is_anonymous",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode togglePaidMessageReactionIsAnonymous#9774db11 as nil")
	}
	b.PutID(TogglePaidMessageReactionIsAnonymousRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode togglePaidMessageReactionIsAnonymous#9774db11 as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutInt53(t.MessageID)
	b.PutBool(t.IsAnonymous)
	return nil
}

// Decode implements bin.Decoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode togglePaidMessageReactionIsAnonymous#9774db11 to nil")
	}
	if err := b.ConsumeID(TogglePaidMessageReactionIsAnonymousRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode togglePaidMessageReactionIsAnonymous#9774db11 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field message_id: %w", err)
		}
		t.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field is_anonymous: %w", err)
		}
		t.IsAnonymous = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode togglePaidMessageReactionIsAnonymous#9774db11 as nil")
	}
	b.ObjStart()
	b.PutID("togglePaidMessageReactionIsAnonymous")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(t.MessageID)
	b.Comma()
	b.FieldStart("is_anonymous")
	b.PutBool(t.IsAnonymous)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TogglePaidMessageReactionIsAnonymousRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode togglePaidMessageReactionIsAnonymous#9774db11 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("togglePaidMessageReactionIsAnonymous"); err != nil {
				return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field chat_id: %w", err)
			}
			t.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field message_id: %w", err)
			}
			t.MessageID = value
		case "is_anonymous":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode togglePaidMessageReactionIsAnonymous#9774db11: field is_anonymous: %w", err)
			}
			t.IsAnonymous = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *TogglePaidMessageReactionIsAnonymousRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetMessageID returns value of MessageID field.
func (t *TogglePaidMessageReactionIsAnonymousRequest) GetMessageID() (value int64) {
	if t == nil {
		return
	}
	return t.MessageID
}

// GetIsAnonymous returns value of IsAnonymous field.
func (t *TogglePaidMessageReactionIsAnonymousRequest) GetIsAnonymous() (value bool) {
	if t == nil {
		return
	}
	return t.IsAnonymous
}

// TogglePaidMessageReactionIsAnonymous invokes method togglePaidMessageReactionIsAnonymous#9774db11 returning error if any.
func (c *Client) TogglePaidMessageReactionIsAnonymous(ctx context.Context, request *TogglePaidMessageReactionIsAnonymousRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
