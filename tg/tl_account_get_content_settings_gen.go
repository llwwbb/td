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

// AccountGetContentSettingsRequest represents TL type `account.getContentSettings#8b9b4dae`.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
type AccountGetContentSettingsRequest struct {
}

// AccountGetContentSettingsRequestTypeID is TL type id of AccountGetContentSettingsRequest.
const AccountGetContentSettingsRequestTypeID = 0x8b9b4dae

// Encode implements bin.Encoder.
func (g *AccountGetContentSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContentSettings#8b9b4dae as nil")
	}
	b.PutID(AccountGetContentSettingsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetContentSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContentSettings#8b9b4dae to nil")
	}
	if err := b.ConsumeID(AccountGetContentSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getContentSettings#8b9b4dae: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetContentSettingsRequest.
var (
	_ bin.Encoder = &AccountGetContentSettingsRequest{}
	_ bin.Decoder = &AccountGetContentSettingsRequest{}
)

// AccountGetContentSettings invokes method account.getContentSettings#8b9b4dae returning error if any.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
func (c *Client) AccountGetContentSettings(ctx context.Context, request *AccountGetContentSettingsRequest) (*AccountContentSettings, error) {
	var result AccountContentSettings
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
