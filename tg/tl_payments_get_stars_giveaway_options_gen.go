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

// PaymentsGetStarsGiveawayOptionsRequest represents TL type `payments.getStarsGiveawayOptions#bd1efd3e`.
// Fetch a list of star giveaway options »¹.
//
// Links:
//  1. https://core.telegram.org/api/giveaways#star-giveaways
//
// See https://core.telegram.org/method/payments.getStarsGiveawayOptions for reference.
type PaymentsGetStarsGiveawayOptionsRequest struct {
}

// PaymentsGetStarsGiveawayOptionsRequestTypeID is TL type id of PaymentsGetStarsGiveawayOptionsRequest.
const PaymentsGetStarsGiveawayOptionsRequestTypeID = 0xbd1efd3e

// Ensuring interfaces in compile-time for PaymentsGetStarsGiveawayOptionsRequest.
var (
	_ bin.Encoder     = &PaymentsGetStarsGiveawayOptionsRequest{}
	_ bin.Decoder     = &PaymentsGetStarsGiveawayOptionsRequest{}
	_ bin.BareEncoder = &PaymentsGetStarsGiveawayOptionsRequest{}
	_ bin.BareDecoder = &PaymentsGetStarsGiveawayOptionsRequest{}
)

func (g *PaymentsGetStarsGiveawayOptionsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetStarsGiveawayOptionsRequest) String() string {
	if g == nil {
		return "PaymentsGetStarsGiveawayOptionsRequest(nil)"
	}
	type Alias PaymentsGetStarsGiveawayOptionsRequest
	return fmt.Sprintf("PaymentsGetStarsGiveawayOptionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetStarsGiveawayOptionsRequest) TypeID() uint32 {
	return PaymentsGetStarsGiveawayOptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetStarsGiveawayOptionsRequest) TypeName() string {
	return "payments.getStarsGiveawayOptions"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetStarsGiveawayOptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getStarsGiveawayOptions",
		ID:   PaymentsGetStarsGiveawayOptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetStarsGiveawayOptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsGiveawayOptions#bd1efd3e as nil")
	}
	b.PutID(PaymentsGetStarsGiveawayOptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetStarsGiveawayOptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsGiveawayOptions#bd1efd3e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetStarsGiveawayOptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsGiveawayOptions#bd1efd3e to nil")
	}
	if err := b.ConsumeID(PaymentsGetStarsGiveawayOptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getStarsGiveawayOptions#bd1efd3e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetStarsGiveawayOptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsGiveawayOptions#bd1efd3e to nil")
	}
	return nil
}

// PaymentsGetStarsGiveawayOptions invokes method payments.getStarsGiveawayOptions#bd1efd3e returning error if any.
// Fetch a list of star giveaway options »¹.
//
// Links:
//  1. https://core.telegram.org/api/giveaways#star-giveaways
//
// See https://core.telegram.org/method/payments.getStarsGiveawayOptions for reference.
// Can be used by bots.
func (c *Client) PaymentsGetStarsGiveawayOptions(ctx context.Context) ([]StarsGiveawayOption, error) {
	var result StarsGiveawayOptionVector

	request := &PaymentsGetStarsGiveawayOptionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []StarsGiveawayOption(result.Elems), nil
}
