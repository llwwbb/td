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

// ThemeParameters represents TL type `themeParameters#ef8395af`.
type ThemeParameters struct {
	// A color of the background in the RGB format
	BackgroundColor int32
	// A secondary color for the background in the RGB format
	SecondaryBackgroundColor int32
	// A color of the header background in the RGB format
	HeaderBackgroundColor int32
	// A color of the bottom bar background in the RGB format
	BottomBarBackgroundColor int32
	// A color of the section background in the RGB format
	SectionBackgroundColor int32
	// A color of the section separator in the RGB format
	SectionSeparatorColor int32
	// A color of text in the RGB format
	TextColor int32
	// An accent color of the text in the RGB format
	AccentTextColor int32
	// A color of text on the section headers in the RGB format
	SectionHeaderTextColor int32
	// A color of the subtitle text in the RGB format
	SubtitleTextColor int32
	// A color of the text for destructive actions in the RGB format
	DestructiveTextColor int32
	// A color of hints in the RGB format
	HintColor int32
	// A color of links in the RGB format
	LinkColor int32
	// A color of the buttons in the RGB format
	ButtonColor int32
	// A color of text on the buttons in the RGB format
	ButtonTextColor int32
}

// ThemeParametersTypeID is TL type id of ThemeParameters.
const ThemeParametersTypeID = 0xef8395af

// Ensuring interfaces in compile-time for ThemeParameters.
var (
	_ bin.Encoder     = &ThemeParameters{}
	_ bin.Decoder     = &ThemeParameters{}
	_ bin.BareEncoder = &ThemeParameters{}
	_ bin.BareDecoder = &ThemeParameters{}
)

func (t *ThemeParameters) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.BackgroundColor == 0) {
		return false
	}
	if !(t.SecondaryBackgroundColor == 0) {
		return false
	}
	if !(t.HeaderBackgroundColor == 0) {
		return false
	}
	if !(t.BottomBarBackgroundColor == 0) {
		return false
	}
	if !(t.SectionBackgroundColor == 0) {
		return false
	}
	if !(t.SectionSeparatorColor == 0) {
		return false
	}
	if !(t.TextColor == 0) {
		return false
	}
	if !(t.AccentTextColor == 0) {
		return false
	}
	if !(t.SectionHeaderTextColor == 0) {
		return false
	}
	if !(t.SubtitleTextColor == 0) {
		return false
	}
	if !(t.DestructiveTextColor == 0) {
		return false
	}
	if !(t.HintColor == 0) {
		return false
	}
	if !(t.LinkColor == 0) {
		return false
	}
	if !(t.ButtonColor == 0) {
		return false
	}
	if !(t.ButtonTextColor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThemeParameters) String() string {
	if t == nil {
		return "ThemeParameters(nil)"
	}
	type Alias ThemeParameters
	return fmt.Sprintf("ThemeParameters%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThemeParameters) TypeID() uint32 {
	return ThemeParametersTypeID
}

// TypeName returns name of type in TL schema.
func (*ThemeParameters) TypeName() string {
	return "themeParameters"
}

// TypeInfo returns info about TL type.
func (t *ThemeParameters) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "themeParameters",
		ID:   ThemeParametersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BackgroundColor",
			SchemaName: "background_color",
		},
		{
			Name:       "SecondaryBackgroundColor",
			SchemaName: "secondary_background_color",
		},
		{
			Name:       "HeaderBackgroundColor",
			SchemaName: "header_background_color",
		},
		{
			Name:       "BottomBarBackgroundColor",
			SchemaName: "bottom_bar_background_color",
		},
		{
			Name:       "SectionBackgroundColor",
			SchemaName: "section_background_color",
		},
		{
			Name:       "SectionSeparatorColor",
			SchemaName: "section_separator_color",
		},
		{
			Name:       "TextColor",
			SchemaName: "text_color",
		},
		{
			Name:       "AccentTextColor",
			SchemaName: "accent_text_color",
		},
		{
			Name:       "SectionHeaderTextColor",
			SchemaName: "section_header_text_color",
		},
		{
			Name:       "SubtitleTextColor",
			SchemaName: "subtitle_text_color",
		},
		{
			Name:       "DestructiveTextColor",
			SchemaName: "destructive_text_color",
		},
		{
			Name:       "HintColor",
			SchemaName: "hint_color",
		},
		{
			Name:       "LinkColor",
			SchemaName: "link_color",
		},
		{
			Name:       "ButtonColor",
			SchemaName: "button_color",
		},
		{
			Name:       "ButtonTextColor",
			SchemaName: "button_text_color",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThemeParameters) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeParameters#ef8395af as nil")
	}
	b.PutID(ThemeParametersTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThemeParameters) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeParameters#ef8395af as nil")
	}
	b.PutInt32(t.BackgroundColor)
	b.PutInt32(t.SecondaryBackgroundColor)
	b.PutInt32(t.HeaderBackgroundColor)
	b.PutInt32(t.BottomBarBackgroundColor)
	b.PutInt32(t.SectionBackgroundColor)
	b.PutInt32(t.SectionSeparatorColor)
	b.PutInt32(t.TextColor)
	b.PutInt32(t.AccentTextColor)
	b.PutInt32(t.SectionHeaderTextColor)
	b.PutInt32(t.SubtitleTextColor)
	b.PutInt32(t.DestructiveTextColor)
	b.PutInt32(t.HintColor)
	b.PutInt32(t.LinkColor)
	b.PutInt32(t.ButtonColor)
	b.PutInt32(t.ButtonTextColor)
	return nil
}

// Decode implements bin.Decoder.
func (t *ThemeParameters) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeParameters#ef8395af to nil")
	}
	if err := b.ConsumeID(ThemeParametersTypeID); err != nil {
		return fmt.Errorf("unable to decode themeParameters#ef8395af: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThemeParameters) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeParameters#ef8395af to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field background_color: %w", err)
		}
		t.BackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field secondary_background_color: %w", err)
		}
		t.SecondaryBackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field header_background_color: %w", err)
		}
		t.HeaderBackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field bottom_bar_background_color: %w", err)
		}
		t.BottomBarBackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_background_color: %w", err)
		}
		t.SectionBackgroundColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_separator_color: %w", err)
		}
		t.SectionSeparatorColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field text_color: %w", err)
		}
		t.TextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field accent_text_color: %w", err)
		}
		t.AccentTextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_header_text_color: %w", err)
		}
		t.SectionHeaderTextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field subtitle_text_color: %w", err)
		}
		t.SubtitleTextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field destructive_text_color: %w", err)
		}
		t.DestructiveTextColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field hint_color: %w", err)
		}
		t.HintColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field link_color: %w", err)
		}
		t.LinkColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field button_color: %w", err)
		}
		t.ButtonColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode themeParameters#ef8395af: field button_text_color: %w", err)
		}
		t.ButtonTextColor = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThemeParameters) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode themeParameters#ef8395af as nil")
	}
	b.ObjStart()
	b.PutID("themeParameters")
	b.Comma()
	b.FieldStart("background_color")
	b.PutInt32(t.BackgroundColor)
	b.Comma()
	b.FieldStart("secondary_background_color")
	b.PutInt32(t.SecondaryBackgroundColor)
	b.Comma()
	b.FieldStart("header_background_color")
	b.PutInt32(t.HeaderBackgroundColor)
	b.Comma()
	b.FieldStart("bottom_bar_background_color")
	b.PutInt32(t.BottomBarBackgroundColor)
	b.Comma()
	b.FieldStart("section_background_color")
	b.PutInt32(t.SectionBackgroundColor)
	b.Comma()
	b.FieldStart("section_separator_color")
	b.PutInt32(t.SectionSeparatorColor)
	b.Comma()
	b.FieldStart("text_color")
	b.PutInt32(t.TextColor)
	b.Comma()
	b.FieldStart("accent_text_color")
	b.PutInt32(t.AccentTextColor)
	b.Comma()
	b.FieldStart("section_header_text_color")
	b.PutInt32(t.SectionHeaderTextColor)
	b.Comma()
	b.FieldStart("subtitle_text_color")
	b.PutInt32(t.SubtitleTextColor)
	b.Comma()
	b.FieldStart("destructive_text_color")
	b.PutInt32(t.DestructiveTextColor)
	b.Comma()
	b.FieldStart("hint_color")
	b.PutInt32(t.HintColor)
	b.Comma()
	b.FieldStart("link_color")
	b.PutInt32(t.LinkColor)
	b.Comma()
	b.FieldStart("button_color")
	b.PutInt32(t.ButtonColor)
	b.Comma()
	b.FieldStart("button_text_color")
	b.PutInt32(t.ButtonTextColor)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThemeParameters) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode themeParameters#ef8395af to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("themeParameters"); err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: %w", err)
			}
		case "background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field background_color: %w", err)
			}
			t.BackgroundColor = value
		case "secondary_background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field secondary_background_color: %w", err)
			}
			t.SecondaryBackgroundColor = value
		case "header_background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field header_background_color: %w", err)
			}
			t.HeaderBackgroundColor = value
		case "bottom_bar_background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field bottom_bar_background_color: %w", err)
			}
			t.BottomBarBackgroundColor = value
		case "section_background_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_background_color: %w", err)
			}
			t.SectionBackgroundColor = value
		case "section_separator_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_separator_color: %w", err)
			}
			t.SectionSeparatorColor = value
		case "text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field text_color: %w", err)
			}
			t.TextColor = value
		case "accent_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field accent_text_color: %w", err)
			}
			t.AccentTextColor = value
		case "section_header_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field section_header_text_color: %w", err)
			}
			t.SectionHeaderTextColor = value
		case "subtitle_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field subtitle_text_color: %w", err)
			}
			t.SubtitleTextColor = value
		case "destructive_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field destructive_text_color: %w", err)
			}
			t.DestructiveTextColor = value
		case "hint_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field hint_color: %w", err)
			}
			t.HintColor = value
		case "link_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field link_color: %w", err)
			}
			t.LinkColor = value
		case "button_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field button_color: %w", err)
			}
			t.ButtonColor = value
		case "button_text_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode themeParameters#ef8395af: field button_text_color: %w", err)
			}
			t.ButtonTextColor = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackgroundColor returns value of BackgroundColor field.
func (t *ThemeParameters) GetBackgroundColor() (value int32) {
	if t == nil {
		return
	}
	return t.BackgroundColor
}

// GetSecondaryBackgroundColor returns value of SecondaryBackgroundColor field.
func (t *ThemeParameters) GetSecondaryBackgroundColor() (value int32) {
	if t == nil {
		return
	}
	return t.SecondaryBackgroundColor
}

// GetHeaderBackgroundColor returns value of HeaderBackgroundColor field.
func (t *ThemeParameters) GetHeaderBackgroundColor() (value int32) {
	if t == nil {
		return
	}
	return t.HeaderBackgroundColor
}

// GetBottomBarBackgroundColor returns value of BottomBarBackgroundColor field.
func (t *ThemeParameters) GetBottomBarBackgroundColor() (value int32) {
	if t == nil {
		return
	}
	return t.BottomBarBackgroundColor
}

// GetSectionBackgroundColor returns value of SectionBackgroundColor field.
func (t *ThemeParameters) GetSectionBackgroundColor() (value int32) {
	if t == nil {
		return
	}
	return t.SectionBackgroundColor
}

// GetSectionSeparatorColor returns value of SectionSeparatorColor field.
func (t *ThemeParameters) GetSectionSeparatorColor() (value int32) {
	if t == nil {
		return
	}
	return t.SectionSeparatorColor
}

// GetTextColor returns value of TextColor field.
func (t *ThemeParameters) GetTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.TextColor
}

// GetAccentTextColor returns value of AccentTextColor field.
func (t *ThemeParameters) GetAccentTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.AccentTextColor
}

// GetSectionHeaderTextColor returns value of SectionHeaderTextColor field.
func (t *ThemeParameters) GetSectionHeaderTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.SectionHeaderTextColor
}

// GetSubtitleTextColor returns value of SubtitleTextColor field.
func (t *ThemeParameters) GetSubtitleTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.SubtitleTextColor
}

// GetDestructiveTextColor returns value of DestructiveTextColor field.
func (t *ThemeParameters) GetDestructiveTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.DestructiveTextColor
}

// GetHintColor returns value of HintColor field.
func (t *ThemeParameters) GetHintColor() (value int32) {
	if t == nil {
		return
	}
	return t.HintColor
}

// GetLinkColor returns value of LinkColor field.
func (t *ThemeParameters) GetLinkColor() (value int32) {
	if t == nil {
		return
	}
	return t.LinkColor
}

// GetButtonColor returns value of ButtonColor field.
func (t *ThemeParameters) GetButtonColor() (value int32) {
	if t == nil {
		return
	}
	return t.ButtonColor
}

// GetButtonTextColor returns value of ButtonTextColor field.
func (t *ThemeParameters) GetButtonTextColor() (value int32) {
	if t == nil {
		return
	}
	return t.ButtonTextColor
}
