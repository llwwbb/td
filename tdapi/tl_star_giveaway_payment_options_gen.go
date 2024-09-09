// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// StarGiveawayPaymentOptions represents TL type `starGiveawayPaymentOptions#909fbf18`.
type StarGiveawayPaymentOptions struct {
	// The list of options
	Options []StarGiveawayPaymentOption
}

// StarGiveawayPaymentOptionsTypeID is TL type id of StarGiveawayPaymentOptions.
const StarGiveawayPaymentOptionsTypeID = 0x909fbf18

// Ensuring interfaces in compile-time for StarGiveawayPaymentOptions.
var (
	_ bin.Encoder     = &StarGiveawayPaymentOptions{}
	_ bin.Decoder     = &StarGiveawayPaymentOptions{}
	_ bin.BareEncoder = &StarGiveawayPaymentOptions{}
	_ bin.BareDecoder = &StarGiveawayPaymentOptions{}
)

func (s *StarGiveawayPaymentOptions) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Options == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGiveawayPaymentOptions) String() string {
	if s == nil {
		return "StarGiveawayPaymentOptions(nil)"
	}
	type Alias StarGiveawayPaymentOptions
	return fmt.Sprintf("StarGiveawayPaymentOptions%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGiveawayPaymentOptions) TypeID() uint32 {
	return StarGiveawayPaymentOptionsTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGiveawayPaymentOptions) TypeName() string {
	return "starGiveawayPaymentOptions"
}

// TypeInfo returns info about TL type.
func (s *StarGiveawayPaymentOptions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGiveawayPaymentOptions",
		ID:   StarGiveawayPaymentOptionsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Options",
			SchemaName: "options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarGiveawayPaymentOptions) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiveawayPaymentOptions#909fbf18 as nil")
	}
	b.PutID(StarGiveawayPaymentOptionsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGiveawayPaymentOptions) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiveawayPaymentOptions#909fbf18 as nil")
	}
	b.PutInt(len(s.Options))
	for idx, v := range s.Options {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare starGiveawayPaymentOptions#909fbf18: field options element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGiveawayPaymentOptions) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiveawayPaymentOptions#909fbf18 to nil")
	}
	if err := b.ConsumeID(StarGiveawayPaymentOptionsTypeID); err != nil {
		return fmt.Errorf("unable to decode starGiveawayPaymentOptions#909fbf18: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGiveawayPaymentOptions) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiveawayPaymentOptions#909fbf18 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiveawayPaymentOptions#909fbf18: field options: %w", err)
		}

		if headerLen > 0 {
			s.Options = make([]StarGiveawayPaymentOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StarGiveawayPaymentOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare starGiveawayPaymentOptions#909fbf18: field options: %w", err)
			}
			s.Options = append(s.Options, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarGiveawayPaymentOptions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiveawayPaymentOptions#909fbf18 as nil")
	}
	b.ObjStart()
	b.PutID("starGiveawayPaymentOptions")
	b.Comma()
	b.FieldStart("options")
	b.ArrStart()
	for idx, v := range s.Options {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode starGiveawayPaymentOptions#909fbf18: field options element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarGiveawayPaymentOptions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiveawayPaymentOptions#909fbf18 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starGiveawayPaymentOptions"); err != nil {
				return fmt.Errorf("unable to decode starGiveawayPaymentOptions#909fbf18: %w", err)
			}
		case "options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value StarGiveawayPaymentOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode starGiveawayPaymentOptions#909fbf18: field options: %w", err)
				}
				s.Options = append(s.Options, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode starGiveawayPaymentOptions#909fbf18: field options: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOptions returns value of Options field.
func (s *StarGiveawayPaymentOptions) GetOptions() (value []StarGiveawayPaymentOption) {
	if s == nil {
		return
	}
	return s.Options
}
