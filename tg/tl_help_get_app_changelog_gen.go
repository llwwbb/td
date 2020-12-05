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

// HelpGetAppChangelogRequest represents TL type `help.getAppChangelog#9010ef6f`.
// Get changelog of current app.
// Typically, an updates constructor will be returned, containing one or more updateServiceNotification updates with app-specific changelogs.
//
// See https://core.telegram.org/method/help.getAppChangelog for reference.
type HelpGetAppChangelogRequest struct {
	// Previous app version
	PrevAppVersion string
}

// HelpGetAppChangelogRequestTypeID is TL type id of HelpGetAppChangelogRequest.
const HelpGetAppChangelogRequestTypeID = 0x9010ef6f

// Encode implements bin.Encoder.
func (g *HelpGetAppChangelogRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppChangelog#9010ef6f as nil")
	}
	b.PutID(HelpGetAppChangelogRequestTypeID)
	b.PutString(g.PrevAppVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetAppChangelogRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppChangelog#9010ef6f to nil")
	}
	if err := b.ConsumeID(HelpGetAppChangelogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: field prev_app_version: %w", err)
		}
		g.PrevAppVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetAppChangelogRequest.
var (
	_ bin.Encoder = &HelpGetAppChangelogRequest{}
	_ bin.Decoder = &HelpGetAppChangelogRequest{}
)

// HelpGetAppChangelog invokes method help.getAppChangelog#9010ef6f returning error if any.
// Get changelog of current app.
// Typically, an updates constructor will be returned, containing one or more updateServiceNotification updates with app-specific changelogs.
//
// See https://core.telegram.org/method/help.getAppChangelog for reference.
func (c *Client) HelpGetAppChangelog(ctx context.Context, request *HelpGetAppChangelogRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
