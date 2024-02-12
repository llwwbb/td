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

// MessagesRequestWebViewRequest represents TL type `messages.requestWebView#269dc2c1`.
// Open a bot mini app¹, sending over user information after user confirmation.
// After calling this method, until the user closes the webview, messages
// prolongWebView¹ must be called every 60 seconds.
//
// Links:
//  1. https://core.telegram.org/bots/webapps
//  2. https://core.telegram.org/method/messages.prolongWebView
//
// See https://core.telegram.org/method/messages.requestWebView for reference.
type MessagesRequestWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the webview was opened by clicking on the bot's menu button »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/menu
	FromBotMenu bool
	// Whether the inline message that will be sent by the bot on behalf of the user once the
	// web app interaction is terminated¹ should be sent silently (no notifications for the
	// receivers).
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendWebViewResultMessage
	Silent bool
	// Dialog where the web app is being opened, and where the resulting message will be sent
	// (see the docs for more info »¹).
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps
	Peer InputPeerClass
	// Bot that owns the web app¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps
	Bot InputUserClass
	// Web app URL¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps
	//
	// Use SetURL and GetURL helpers.
	URL string
	// If the web app was opened from the attachment menu using a attachment menu deep link¹
	// start_param should contain the data from the startattach parameter.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#bot-attachment-or-side-menu-links
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Theme parameters »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#theme-parameters
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
	// Short name of the application; 0-64 English letters, digits, and underscores
	Platform string
	// If set, indicates that the inline message that will be sent by the bot on behalf of
	// the user once the web app interaction is terminated¹ should be sent in reply to the
	// specified message or story.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendWebViewResultMessage
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// Open the web app as the specified peer, sending the resulting the message as the
	// specified peer.
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesRequestWebViewRequestTypeID is TL type id of MessagesRequestWebViewRequest.
const MessagesRequestWebViewRequestTypeID = 0x269dc2c1

// Ensuring interfaces in compile-time for MessagesRequestWebViewRequest.
var (
	_ bin.Encoder     = &MessagesRequestWebViewRequest{}
	_ bin.Decoder     = &MessagesRequestWebViewRequest{}
	_ bin.BareEncoder = &MessagesRequestWebViewRequest{}
	_ bin.BareDecoder = &MessagesRequestWebViewRequest{}
)

func (r *MessagesRequestWebViewRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.FromBotMenu == false) {
		return false
	}
	if !(r.Silent == false) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.StartParam == "") {
		return false
	}
	if !(r.ThemeParams.Zero()) {
		return false
	}
	if !(r.Platform == "") {
		return false
	}
	if !(r.ReplyTo == nil) {
		return false
	}
	if !(r.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestWebViewRequest) String() string {
	if r == nil {
		return "MessagesRequestWebViewRequest(nil)"
	}
	type Alias MessagesRequestWebViewRequest
	return fmt.Sprintf("MessagesRequestWebViewRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestWebViewRequest from given interface.
func (r *MessagesRequestWebViewRequest) FillFrom(from interface {
	GetFromBotMenu() (value bool)
	GetSilent() (value bool)
	GetPeer() (value InputPeerClass)
	GetBot() (value InputUserClass)
	GetURL() (value string, ok bool)
	GetStartParam() (value string, ok bool)
	GetThemeParams() (value DataJSON, ok bool)
	GetPlatform() (value string)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	r.FromBotMenu = from.GetFromBotMenu()
	r.Silent = from.GetSilent()
	r.Peer = from.GetPeer()
	r.Bot = from.GetBot()
	if val, ok := from.GetURL(); ok {
		r.URL = val
	}

	if val, ok := from.GetStartParam(); ok {
		r.StartParam = val
	}

	if val, ok := from.GetThemeParams(); ok {
		r.ThemeParams = val
	}

	r.Platform = from.GetPlatform()
	if val, ok := from.GetReplyTo(); ok {
		r.ReplyTo = val
	}

	if val, ok := from.GetSendAs(); ok {
		r.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRequestWebViewRequest) TypeID() uint32 {
	return MessagesRequestWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRequestWebViewRequest) TypeName() string {
	return "messages.requestWebView"
}

// TypeInfo returns info about TL type.
func (r *MessagesRequestWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.requestWebView",
		ID:   MessagesRequestWebViewRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FromBotMenu",
			SchemaName: "from_bot_menu",
			Null:       !r.Flags.Has(4),
		},
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !r.Flags.Has(5),
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
			Name:       "URL",
			SchemaName: "url",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !r.Flags.Has(3),
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !r.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesRequestWebViewRequest) SetFlags() {
	if !(r.FromBotMenu == false) {
		r.Flags.Set(4)
	}
	if !(r.Silent == false) {
		r.Flags.Set(5)
	}
	if !(r.URL == "") {
		r.Flags.Set(1)
	}
	if !(r.StartParam == "") {
		r.Flags.Set(3)
	}
	if !(r.ThemeParams.Zero()) {
		r.Flags.Set(2)
	}
	if !(r.ReplyTo == nil) {
		r.Flags.Set(0)
	}
	if !(r.SendAs == nil) {
		r.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesRequestWebViewRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestWebView#269dc2c1 as nil")
	}
	b.PutID(MessagesRequestWebViewRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRequestWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestWebView#269dc2c1 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field flags: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field peer: %w", err)
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field bot: %w", err)
	}
	if r.Flags.Has(1) {
		b.PutString(r.URL)
	}
	if r.Flags.Has(3) {
		b.PutString(r.StartParam)
	}
	if r.Flags.Has(2) {
		if err := r.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field theme_params: %w", err)
		}
	}
	b.PutString(r.Platform)
	if r.Flags.Has(0) {
		if r.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field reply_to is nil")
		}
		if err := r.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field reply_to: %w", err)
		}
	}
	if r.Flags.Has(13) {
		if r.SendAs == nil {
			return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field send_as is nil")
		}
		if err := r.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestWebView#269dc2c1: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestWebViewRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestWebView#269dc2c1 to nil")
	}
	if err := b.ConsumeID(MessagesRequestWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRequestWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestWebView#269dc2c1 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field flags: %w", err)
		}
	}
	r.FromBotMenu = r.Flags.Has(4)
	r.Silent = r.Flags.Has(5)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field bot: %w", err)
		}
		r.Bot = value
	}
	if r.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field url: %w", err)
		}
		r.URL = value
	}
	if r.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field start_param: %w", err)
		}
		r.StartParam = value
	}
	if r.Flags.Has(2) {
		if err := r.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field theme_params: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field platform: %w", err)
		}
		r.Platform = value
	}
	if r.Flags.Has(0) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field reply_to: %w", err)
		}
		r.ReplyTo = value
	}
	if r.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestWebView#269dc2c1: field send_as: %w", err)
		}
		r.SendAs = value
	}
	return nil
}

// SetFromBotMenu sets value of FromBotMenu conditional field.
func (r *MessagesRequestWebViewRequest) SetFromBotMenu(value bool) {
	if value {
		r.Flags.Set(4)
		r.FromBotMenu = true
	} else {
		r.Flags.Unset(4)
		r.FromBotMenu = false
	}
}

// GetFromBotMenu returns value of FromBotMenu conditional field.
func (r *MessagesRequestWebViewRequest) GetFromBotMenu() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(4)
}

// SetSilent sets value of Silent conditional field.
func (r *MessagesRequestWebViewRequest) SetSilent(value bool) {
	if value {
		r.Flags.Set(5)
		r.Silent = true
	} else {
		r.Flags.Unset(5)
		r.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (r *MessagesRequestWebViewRequest) GetSilent() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(5)
}

// GetPeer returns value of Peer field.
func (r *MessagesRequestWebViewRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetBot returns value of Bot field.
func (r *MessagesRequestWebViewRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// SetURL sets value of URL conditional field.
func (r *MessagesRequestWebViewRequest) SetURL(value string) {
	r.Flags.Set(1)
	r.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestWebViewRequest) GetURL() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.URL, true
}

// SetStartParam sets value of StartParam conditional field.
func (r *MessagesRequestWebViewRequest) SetStartParam(value string) {
	r.Flags.Set(3)
	r.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestWebViewRequest) GetStartParam() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.StartParam, true
}

// SetThemeParams sets value of ThemeParams conditional field.
func (r *MessagesRequestWebViewRequest) SetThemeParams(value DataJSON) {
	r.Flags.Set(2)
	r.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestWebViewRequest) GetThemeParams() (value DataJSON, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(2) {
		return value, false
	}
	return r.ThemeParams, true
}

// GetPlatform returns value of Platform field.
func (r *MessagesRequestWebViewRequest) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// SetReplyTo sets value of ReplyTo conditional field.
func (r *MessagesRequestWebViewRequest) SetReplyTo(value InputReplyToClass) {
	r.Flags.Set(0)
	r.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestWebViewRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.ReplyTo, true
}

// SetSendAs sets value of SendAs conditional field.
func (r *MessagesRequestWebViewRequest) SetSendAs(value InputPeerClass) {
	r.Flags.Set(13)
	r.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestWebViewRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(13) {
		return value, false
	}
	return r.SendAs, true
}

// MessagesRequestWebView invokes method messages.requestWebView#269dc2c1 returning error if any.
// Open a bot mini app¹, sending over user information after user confirmation.
// After calling this method, until the user closes the webview, messages
// prolongWebView¹ must be called every 60 seconds.
//
// Links:
//  1. https://core.telegram.org/bots/webapps
//  2. https://core.telegram.org/method/messages.prolongWebView
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//	400 BOT_WEBVIEW_DISABLED: A webview cannot be opened in the specified conditions: emitted for example if from_bot_menu or url are set and peer is not the chat with the bot.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.requestWebView for reference.
func (c *Client) MessagesRequestWebView(ctx context.Context, request *MessagesRequestWebViewRequest) (*WebViewResultURL, error) {
	var result WebViewResultURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
