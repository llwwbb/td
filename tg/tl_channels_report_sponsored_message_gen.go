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

// ChannelsReportSponsoredMessageRequest represents TL type `channels.reportSponsoredMessage#af8ff6b9`.
// Report a sponsored message »¹, see here »² for more info on the full flow.
//
// Links:
//  1. https://core.telegram.org/api/sponsored-messages
//  2. https://core.telegram.org/api/sponsored-messages#reporting-sponsored-messages
//
// See https://core.telegram.org/method/channels.reportSponsoredMessage for reference.
type ChannelsReportSponsoredMessageRequest struct {
	// The channel where the sponsored message can be seen.
	Channel InputChannelClass
	// ID of the sponsored message.
	RandomID []byte
	// Chosen report option, initially an empty string, see here »¹ for more info on the
	// full flow.
	//
	// Links:
	//  1) https://core.telegram.org/api/sponsored-messages#reporting-sponsored-messages
	Option []byte
}

// ChannelsReportSponsoredMessageRequestTypeID is TL type id of ChannelsReportSponsoredMessageRequest.
const ChannelsReportSponsoredMessageRequestTypeID = 0xaf8ff6b9

// Ensuring interfaces in compile-time for ChannelsReportSponsoredMessageRequest.
var (
	_ bin.Encoder     = &ChannelsReportSponsoredMessageRequest{}
	_ bin.Decoder     = &ChannelsReportSponsoredMessageRequest{}
	_ bin.BareEncoder = &ChannelsReportSponsoredMessageRequest{}
	_ bin.BareDecoder = &ChannelsReportSponsoredMessageRequest{}
)

func (r *ChannelsReportSponsoredMessageRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.RandomID == nil) {
		return false
	}
	if !(r.Option == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReportSponsoredMessageRequest) String() string {
	if r == nil {
		return "ChannelsReportSponsoredMessageRequest(nil)"
	}
	type Alias ChannelsReportSponsoredMessageRequest
	return fmt.Sprintf("ChannelsReportSponsoredMessageRequest%+v", Alias(*r))
}

// FillFrom fills ChannelsReportSponsoredMessageRequest from given interface.
func (r *ChannelsReportSponsoredMessageRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetRandomID() (value []byte)
	GetOption() (value []byte)
}) {
	r.Channel = from.GetChannel()
	r.RandomID = from.GetRandomID()
	r.Option = from.GetOption()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsReportSponsoredMessageRequest) TypeID() uint32 {
	return ChannelsReportSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsReportSponsoredMessageRequest) TypeName() string {
	return "channels.reportSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (r *ChannelsReportSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.reportSponsoredMessage",
		ID:   ChannelsReportSponsoredMessageRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ChannelsReportSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportSponsoredMessage#af8ff6b9 as nil")
	}
	b.PutID(ChannelsReportSponsoredMessageRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ChannelsReportSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportSponsoredMessage#af8ff6b9 as nil")
	}
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.reportSponsoredMessage#af8ff6b9: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportSponsoredMessage#af8ff6b9: field channel: %w", err)
	}
	b.PutBytes(r.RandomID)
	b.PutBytes(r.Option)
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReportSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportSponsoredMessage#af8ff6b9 to nil")
	}
	if err := b.ConsumeID(ChannelsReportSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.reportSponsoredMessage#af8ff6b9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ChannelsReportSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportSponsoredMessage#af8ff6b9 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSponsoredMessage#af8ff6b9: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSponsoredMessage#af8ff6b9: field random_id: %w", err)
		}
		r.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSponsoredMessage#af8ff6b9: field option: %w", err)
		}
		r.Option = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReportSponsoredMessageRequest) GetChannel() (value InputChannelClass) {
	if r == nil {
		return
	}
	return r.Channel
}

// GetRandomID returns value of RandomID field.
func (r *ChannelsReportSponsoredMessageRequest) GetRandomID() (value []byte) {
	if r == nil {
		return
	}
	return r.RandomID
}

// GetOption returns value of Option field.
func (r *ChannelsReportSponsoredMessageRequest) GetOption() (value []byte) {
	if r == nil {
		return
	}
	return r.Option
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (r *ChannelsReportSponsoredMessageRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return r.Channel.AsNotEmpty()
}

// ChannelsReportSponsoredMessage invokes method channels.reportSponsoredMessage#af8ff6b9 returning error if any.
// Report a sponsored message »¹, see here »² for more info on the full flow.
//
// Links:
//  1. https://core.telegram.org/api/sponsored-messages
//  2. https://core.telegram.org/api/sponsored-messages#reporting-sponsored-messages
//
// Possible errors:
//
//	400 AD_EXPIRED: The ad has expired (too old or not found).
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 PREMIUM_ACCOUNT_REQUIRED: A premium account is required to execute this action.
//
// See https://core.telegram.org/method/channels.reportSponsoredMessage for reference.
func (c *Client) ChannelsReportSponsoredMessage(ctx context.Context, request *ChannelsReportSponsoredMessageRequest) (ChannelsSponsoredMessageReportResultClass, error) {
	var result ChannelsSponsoredMessageReportResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SponsoredMessageReportResult, nil
}
