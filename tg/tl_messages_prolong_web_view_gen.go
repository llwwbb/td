// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesProlongWebViewRequest represents TL type `messages.prolongWebView#b0d81a83`.
// Indicate to the server (from the user side) that the user is still using a web app.
// If the method returns a QUERY_ID_INVALID error, the webview must be closed.
//
// See https://core.telegram.org/method/messages.prolongWebView for reference.
type MessagesProlongWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the inline message that will be sent by the bot on behalf of the user once the
	// web app interaction is terminated¹ should be sent silently (no notifications for the
	// receivers).
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendWebViewResultMessage
	Silent bool
	// Dialog where the web app was opened.
	Peer InputPeerClass
	// Bot that owns the web app¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps
	Bot InputUserClass
	// Web app interaction ID obtained from messages.requestWebView¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.requestWebView
	QueryID int64
	// If set, indicates that the inline message that will be sent by the bot on behalf of
	// the user once the web app interaction is terminated¹ should be sent in reply to the
	// specified message or story.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendWebViewResultMessage
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// Open the web app as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesProlongWebViewRequestTypeID is TL type id of MessagesProlongWebViewRequest.
const MessagesProlongWebViewRequestTypeID = 0xb0d81a83

// Ensuring interfaces in compile-time for MessagesProlongWebViewRequest.
var (
	_ bin.Encoder     = &MessagesProlongWebViewRequest{}
	_ bin.Decoder     = &MessagesProlongWebViewRequest{}
	_ bin.BareEncoder = &MessagesProlongWebViewRequest{}
	_ bin.BareDecoder = &MessagesProlongWebViewRequest{}
)

func (p *MessagesProlongWebViewRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Silent == false) {
		return false
	}
	if !(p.Peer == nil) {
		return false
	}
	if !(p.Bot == nil) {
		return false
	}
	if !(p.QueryID == 0) {
		return false
	}
	if !(p.ReplyTo == nil) {
		return false
	}
	if !(p.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *MessagesProlongWebViewRequest) String() string {
	if p == nil {
		return "MessagesProlongWebViewRequest(nil)"
	}
	type Alias MessagesProlongWebViewRequest
	return fmt.Sprintf("MessagesProlongWebViewRequest%+v", Alias(*p))
}

// FillFrom fills MessagesProlongWebViewRequest from given interface.
func (p *MessagesProlongWebViewRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetPeer() (value InputPeerClass)
	GetBot() (value InputUserClass)
	GetQueryID() (value int64)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	p.Silent = from.GetSilent()
	p.Peer = from.GetPeer()
	p.Bot = from.GetBot()
	p.QueryID = from.GetQueryID()
	if val, ok := from.GetReplyTo(); ok {
		p.ReplyTo = val
	}

	if val, ok := from.GetSendAs(); ok {
		p.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesProlongWebViewRequest) TypeID() uint32 {
	return MessagesProlongWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesProlongWebViewRequest) TypeName() string {
	return "messages.prolongWebView"
}

// TypeInfo returns info about TL type.
func (p *MessagesProlongWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.prolongWebView",
		ID:   MessagesProlongWebViewRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !p.Flags.Has(5),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !p.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *MessagesProlongWebViewRequest) SetFlags() {
	if !(p.Silent == false) {
		p.Flags.Set(5)
	}
	if !(p.ReplyTo == nil) {
		p.Flags.Set(0)
	}
	if !(p.SendAs == nil) {
		p.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (p *MessagesProlongWebViewRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.prolongWebView#b0d81a83 as nil")
	}
	b.PutID(MessagesProlongWebViewRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *MessagesProlongWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.prolongWebView#b0d81a83 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field flags: %w", err)
	}
	if p.Peer == nil {
		return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field peer is nil")
	}
	if err := p.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field peer: %w", err)
	}
	if p.Bot == nil {
		return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field bot is nil")
	}
	if err := p.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field bot: %w", err)
	}
	b.PutLong(p.QueryID)
	if p.Flags.Has(0) {
		if p.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field reply_to is nil")
		}
		if err := p.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field reply_to: %w", err)
		}
	}
	if p.Flags.Has(13) {
		if p.SendAs == nil {
			return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field send_as is nil")
		}
		if err := p.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.prolongWebView#b0d81a83: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *MessagesProlongWebViewRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.prolongWebView#b0d81a83 to nil")
	}
	if err := b.ConsumeID(MessagesProlongWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *MessagesProlongWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.prolongWebView#b0d81a83 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field flags: %w", err)
		}
	}
	p.Silent = p.Flags.Has(5)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field peer: %w", err)
		}
		p.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field bot: %w", err)
		}
		p.Bot = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field query_id: %w", err)
		}
		p.QueryID = value
	}
	if p.Flags.Has(0) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field reply_to: %w", err)
		}
		p.ReplyTo = value
	}
	if p.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.prolongWebView#b0d81a83: field send_as: %w", err)
		}
		p.SendAs = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (p *MessagesProlongWebViewRequest) SetSilent(value bool) {
	if value {
		p.Flags.Set(5)
		p.Silent = true
	} else {
		p.Flags.Unset(5)
		p.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (p *MessagesProlongWebViewRequest) GetSilent() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(5)
}

// GetPeer returns value of Peer field.
func (p *MessagesProlongWebViewRequest) GetPeer() (value InputPeerClass) {
	if p == nil {
		return
	}
	return p.Peer
}

// GetBot returns value of Bot field.
func (p *MessagesProlongWebViewRequest) GetBot() (value InputUserClass) {
	if p == nil {
		return
	}
	return p.Bot
}

// GetQueryID returns value of QueryID field.
func (p *MessagesProlongWebViewRequest) GetQueryID() (value int64) {
	if p == nil {
		return
	}
	return p.QueryID
}

// SetReplyTo sets value of ReplyTo conditional field.
func (p *MessagesProlongWebViewRequest) SetReplyTo(value InputReplyToClass) {
	p.Flags.Set(0)
	p.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (p *MessagesProlongWebViewRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.ReplyTo, true
}

// SetSendAs sets value of SendAs conditional field.
func (p *MessagesProlongWebViewRequest) SetSendAs(value InputPeerClass) {
	p.Flags.Set(13)
	p.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (p *MessagesProlongWebViewRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(13) {
		return value, false
	}
	return p.SendAs, true
}

// MessagesProlongWebView invokes method messages.prolongWebView#b0d81a83 returning error if any.
// Indicate to the server (from the user side) that the user is still using a web app.
// If the method returns a QUERY_ID_INVALID error, the webview must be closed.
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//
// See https://core.telegram.org/method/messages.prolongWebView for reference.
func (c *Client) MessagesProlongWebView(ctx context.Context, request *MessagesProlongWebViewRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
