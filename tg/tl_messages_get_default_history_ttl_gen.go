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

// MessagesGetDefaultHistoryTTLRequest represents TL type `messages.getDefaultHistoryTTL#658b7188`.
// Gets the default value of the Time-To-Live setting, applied to all new chats.
//
// See https://core.telegram.org/method/messages.getDefaultHistoryTTL for reference.
type MessagesGetDefaultHistoryTTLRequest struct {
}

// MessagesGetDefaultHistoryTTLRequestTypeID is TL type id of MessagesGetDefaultHistoryTTLRequest.
const MessagesGetDefaultHistoryTTLRequestTypeID = 0x658b7188

// Ensuring interfaces in compile-time for MessagesGetDefaultHistoryTTLRequest.
var (
	_ bin.Encoder     = &MessagesGetDefaultHistoryTTLRequest{}
	_ bin.Decoder     = &MessagesGetDefaultHistoryTTLRequest{}
	_ bin.BareEncoder = &MessagesGetDefaultHistoryTTLRequest{}
	_ bin.BareDecoder = &MessagesGetDefaultHistoryTTLRequest{}
)

func (g *MessagesGetDefaultHistoryTTLRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDefaultHistoryTTLRequest) String() string {
	if g == nil {
		return "MessagesGetDefaultHistoryTTLRequest(nil)"
	}
	type Alias MessagesGetDefaultHistoryTTLRequest
	return fmt.Sprintf("MessagesGetDefaultHistoryTTLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDefaultHistoryTTLRequest) TypeID() uint32 {
	return MessagesGetDefaultHistoryTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDefaultHistoryTTLRequest) TypeName() string {
	return "messages.getDefaultHistoryTTL"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDefaultHistoryTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDefaultHistoryTTL",
		ID:   MessagesGetDefaultHistoryTTLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDefaultHistoryTTLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDefaultHistoryTTL#658b7188 as nil")
	}
	b.PutID(MessagesGetDefaultHistoryTTLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDefaultHistoryTTLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDefaultHistoryTTL#658b7188 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDefaultHistoryTTLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDefaultHistoryTTL#658b7188 to nil")
	}
	if err := b.ConsumeID(MessagesGetDefaultHistoryTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDefaultHistoryTTL#658b7188: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDefaultHistoryTTLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDefaultHistoryTTL#658b7188 to nil")
	}
	return nil
}

// MessagesGetDefaultHistoryTTL invokes method messages.getDefaultHistoryTTL#658b7188 returning error if any.
// Gets the default value of the Time-To-Live setting, applied to all new chats.
//
// See https://core.telegram.org/method/messages.getDefaultHistoryTTL for reference.
func (c *Client) MessagesGetDefaultHistoryTTL(ctx context.Context) (*DefaultHistoryTTL, error) {
	var result DefaultHistoryTTL

	request := &MessagesGetDefaultHistoryTTLRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
