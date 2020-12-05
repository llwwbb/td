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

// MessagesGetDhConfigRequest represents TL type `messages.getDhConfig#26cf8950`.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
type MessagesGetDhConfigRequest struct {
	// Value of the version parameter from messages.dhConfig, avialable at the client
	Version int
	// Length of the required random sequence
	RandomLength int
}

// MessagesGetDhConfigRequestTypeID is TL type id of MessagesGetDhConfigRequest.
const MessagesGetDhConfigRequestTypeID = 0x26cf8950

// Encode implements bin.Encoder.
func (g *MessagesGetDhConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDhConfig#26cf8950 as nil")
	}
	b.PutID(MessagesGetDhConfigRequestTypeID)
	b.PutInt(g.Version)
	b.PutInt(g.RandomLength)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDhConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDhConfig#26cf8950 to nil")
	}
	if err := b.ConsumeID(MessagesGetDhConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field version: %w", err)
		}
		g.Version = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field random_length: %w", err)
		}
		g.RandomLength = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDhConfigRequest.
var (
	_ bin.Encoder = &MessagesGetDhConfigRequest{}
	_ bin.Decoder = &MessagesGetDhConfigRequest{}
)

// MessagesGetDhConfig invokes method messages.getDhConfig#26cf8950 returning error if any.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
func (c *Client) MessagesGetDhConfig(ctx context.Context, request *MessagesGetDhConfigRequest) (MessagesDhConfigClass, error) {
	var result MessagesDhConfigBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DhConfig, nil
}
