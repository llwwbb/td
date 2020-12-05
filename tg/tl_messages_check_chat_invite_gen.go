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

// MessagesCheckChatInviteRequest represents TL type `messages.checkChatInvite#3eadb1bb`.
// Check the validity of a chat invite link and get basic info about it
//
// See https://core.telegram.org/method/messages.checkChatInvite for reference.
type MessagesCheckChatInviteRequest struct {
	// Invite hash in t.me/joinchat/hash
	Hash string
}

// MessagesCheckChatInviteRequestTypeID is TL type id of MessagesCheckChatInviteRequest.
const MessagesCheckChatInviteRequestTypeID = 0x3eadb1bb

// Encode implements bin.Encoder.
func (c *MessagesCheckChatInviteRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkChatInvite#3eadb1bb as nil")
	}
	b.PutID(MessagesCheckChatInviteRequestTypeID)
	b.PutString(c.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesCheckChatInviteRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkChatInvite#3eadb1bb to nil")
	}
	if err := b.ConsumeID(MessagesCheckChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.checkChatInvite#3eadb1bb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.checkChatInvite#3eadb1bb: field hash: %w", err)
		}
		c.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCheckChatInviteRequest.
var (
	_ bin.Encoder = &MessagesCheckChatInviteRequest{}
	_ bin.Decoder = &MessagesCheckChatInviteRequest{}
)

// MessagesCheckChatInvite invokes method messages.checkChatInvite#3eadb1bb returning error if any.
// Check the validity of a chat invite link and get basic info about it
//
// See https://core.telegram.org/method/messages.checkChatInvite for reference.
func (c *Client) MessagesCheckChatInvite(ctx context.Context, request *MessagesCheckChatInviteRequest) (ChatInviteClass, error) {
	var result ChatInviteBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChatInvite, nil
}
