// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesHidePeerSettingsBarRequest represents TL type `messages.hidePeerSettingsBar#4facb138`.
// Should be called after the user hides the report spam/add as contact bar of a new chat, effectively prevents the user from executing the actions specified in the peer's settings.
//
// See https://core.telegram.org/method/messages.hidePeerSettingsBar for reference.
type MessagesHidePeerSettingsBarRequest struct {
	// Peer
	Peer InputPeerClass
}

// MessagesHidePeerSettingsBarRequestTypeID is TL type id of MessagesHidePeerSettingsBarRequest.
const MessagesHidePeerSettingsBarRequestTypeID = 0x4facb138

// Encode implements bin.Encoder.
func (h *MessagesHidePeerSettingsBarRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hidePeerSettingsBar#4facb138 as nil")
	}
	b.PutID(MessagesHidePeerSettingsBarRequestTypeID)
	if h.Peer == nil {
		return fmt.Errorf("unable to encode messages.hidePeerSettingsBar#4facb138: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hidePeerSettingsBar#4facb138: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHidePeerSettingsBarRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hidePeerSettingsBar#4facb138 to nil")
	}
	if err := b.ConsumeID(MessagesHidePeerSettingsBarRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.hidePeerSettingsBar#4facb138: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.hidePeerSettingsBar#4facb138: field peer: %w", err)
		}
		h.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesHidePeerSettingsBarRequest.
var (
	_ bin.Encoder = &MessagesHidePeerSettingsBarRequest{}
	_ bin.Decoder = &MessagesHidePeerSettingsBarRequest{}
)

// MessagesHidePeerSettingsBar invokes method messages.hidePeerSettingsBar#4facb138 returning error if any.
// Should be called after the user hides the report spam/add as contact bar of a new chat, effectively prevents the user from executing the actions specified in the peer's settings.
//
// See https://core.telegram.org/method/messages.hidePeerSettingsBar for reference.
func (c *Client) MessagesHidePeerSettingsBar(ctx context.Context, request *MessagesHidePeerSettingsBarRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
