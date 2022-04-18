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

// AttachMenuBot represents TL type `attachMenuBot#e93cb772`.
//
// See https://core.telegram.org/constructor/attachMenuBot for reference.
type AttachMenuBot struct {
	// Flags field of AttachMenuBot.
	Flags bin.Fields
	// Inactive field of AttachMenuBot.
	Inactive bool
	// BotID field of AttachMenuBot.
	BotID int64
	// ShortName field of AttachMenuBot.
	ShortName string
	// Icons field of AttachMenuBot.
	Icons []AttachMenuBotIcon
}

// AttachMenuBotTypeID is TL type id of AttachMenuBot.
const AttachMenuBotTypeID = 0xe93cb772

// Ensuring interfaces in compile-time for AttachMenuBot.
var (
	_ bin.Encoder     = &AttachMenuBot{}
	_ bin.Decoder     = &AttachMenuBot{}
	_ bin.BareEncoder = &AttachMenuBot{}
	_ bin.BareDecoder = &AttachMenuBot{}
)

func (a *AttachMenuBot) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Inactive == false) {
		return false
	}
	if !(a.BotID == 0) {
		return false
	}
	if !(a.ShortName == "") {
		return false
	}
	if !(a.Icons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuBot) String() string {
	if a == nil {
		return "AttachMenuBot(nil)"
	}
	type Alias AttachMenuBot
	return fmt.Sprintf("AttachMenuBot%+v", Alias(*a))
}

// FillFrom fills AttachMenuBot from given interface.
func (a *AttachMenuBot) FillFrom(from interface {
	GetInactive() (value bool)
	GetBotID() (value int64)
	GetShortName() (value string)
	GetIcons() (value []AttachMenuBotIcon)
}) {
	a.Inactive = from.GetInactive()
	a.BotID = from.GetBotID()
	a.ShortName = from.GetShortName()
	a.Icons = from.GetIcons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuBot) TypeID() uint32 {
	return AttachMenuBotTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuBot) TypeName() string {
	return "attachMenuBot"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuBot) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuBot",
		ID:   AttachMenuBotTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inactive",
			SchemaName: "inactive",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "Icons",
			SchemaName: "icons",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AttachMenuBot) SetFlags() {
	if !(a.Inactive == false) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *AttachMenuBot) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBot#e93cb772 as nil")
	}
	b.PutID(AttachMenuBotTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuBot) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBot#e93cb772 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode attachMenuBot#e93cb772: field flags: %w", err)
	}
	b.PutLong(a.BotID)
	b.PutString(a.ShortName)
	b.PutVectorHeader(len(a.Icons))
	for idx, v := range a.Icons {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode attachMenuBot#e93cb772: field icons element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuBot) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBot#e93cb772 to nil")
	}
	if err := b.ConsumeID(AttachMenuBotTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuBot#e93cb772: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuBot) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBot#e93cb772 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#e93cb772: field flags: %w", err)
		}
	}
	a.Inactive = a.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#e93cb772: field bot_id: %w", err)
		}
		a.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#e93cb772: field short_name: %w", err)
		}
		a.ShortName = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#e93cb772: field icons: %w", err)
		}

		if headerLen > 0 {
			a.Icons = make([]AttachMenuBotIcon, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AttachMenuBotIcon
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode attachMenuBot#e93cb772: field icons: %w", err)
			}
			a.Icons = append(a.Icons, value)
		}
	}
	return nil
}

// SetInactive sets value of Inactive conditional field.
func (a *AttachMenuBot) SetInactive(value bool) {
	if value {
		a.Flags.Set(0)
		a.Inactive = true
	} else {
		a.Flags.Unset(0)
		a.Inactive = false
	}
}

// GetInactive returns value of Inactive conditional field.
func (a *AttachMenuBot) GetInactive() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// GetBotID returns value of BotID field.
func (a *AttachMenuBot) GetBotID() (value int64) {
	if a == nil {
		return
	}
	return a.BotID
}

// GetShortName returns value of ShortName field.
func (a *AttachMenuBot) GetShortName() (value string) {
	if a == nil {
		return
	}
	return a.ShortName
}

// GetIcons returns value of Icons field.
func (a *AttachMenuBot) GetIcons() (value []AttachMenuBotIcon) {
	if a == nil {
		return
	}
	return a.Icons
}