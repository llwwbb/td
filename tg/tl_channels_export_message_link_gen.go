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

// ChannelsExportMessageLinkRequest represents TL type `channels.exportMessageLink#e63fadeb`.
// Get link and embed info of a message in a channel/supergroup
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
type ChannelsExportMessageLinkRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to include other grouped media (for albums)
	Grouped bool
	// Whether to also include a thread ID, if available, inside of the link
	Thread bool
	// Channel
	Channel InputChannelClass
	// Message ID
	ID int
}

// ChannelsExportMessageLinkRequestTypeID is TL type id of ChannelsExportMessageLinkRequest.
const ChannelsExportMessageLinkRequestTypeID = 0xe63fadeb

// Encode implements bin.Encoder.
func (e *ChannelsExportMessageLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.exportMessageLink#e63fadeb as nil")
	}
	b.PutID(ChannelsExportMessageLinkRequestTypeID)
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field flags: %w", err)
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel: %w", err)
	}
	b.PutInt(e.ID)
	return nil
}

// SetGrouped sets value of Grouped conditional field.
func (e *ChannelsExportMessageLinkRequest) SetGrouped(value bool) {
	if value {
		e.Flags.Set(0)
	} else {
		e.Flags.Unset(0)
	}
}

// SetThread sets value of Thread conditional field.
func (e *ChannelsExportMessageLinkRequest) SetThread(value bool) {
	if value {
		e.Flags.Set(1)
	} else {
		e.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (e *ChannelsExportMessageLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.exportMessageLink#e63fadeb to nil")
	}
	if err := b.ConsumeID(ChannelsExportMessageLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field flags: %w", err)
		}
	}
	e.Grouped = e.Flags.Has(0)
	e.Thread = e.Flags.Has(1)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsExportMessageLinkRequest.
var (
	_ bin.Encoder = &ChannelsExportMessageLinkRequest{}
	_ bin.Decoder = &ChannelsExportMessageLinkRequest{}
)

// ChannelsExportMessageLink invokes method channels.exportMessageLink#e63fadeb returning error if any.
// Get link and embed info of a message in a channel/supergroup
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
func (c *Client) ChannelsExportMessageLink(ctx context.Context, request *ChannelsExportMessageLinkRequest) (*ExportedMessageLink, error) {
	var result ExportedMessageLink
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
