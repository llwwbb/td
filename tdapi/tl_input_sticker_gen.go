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

// InputStickerStatic represents TL type `inputStickerStatic#540604db`.
type InputStickerStatic struct {
	// PNG image with the sticker; must be up to 512 KB in size and fit in a 512x512 square
	Sticker InputFileClass
	// Emojis corresponding to the sticker
	Emojis string
	// For masks, position where the mask is placed; pass null if unspecified
	MaskPosition MaskPosition
}

// InputStickerStaticTypeID is TL type id of InputStickerStatic.
const InputStickerStaticTypeID = 0x540604db

// construct implements constructor of InputStickerClass.
func (i InputStickerStatic) construct() InputStickerClass { return &i }

// Ensuring interfaces in compile-time for InputStickerStatic.
var (
	_ bin.Encoder     = &InputStickerStatic{}
	_ bin.Decoder     = &InputStickerStatic{}
	_ bin.BareEncoder = &InputStickerStatic{}
	_ bin.BareDecoder = &InputStickerStatic{}

	_ InputStickerClass = &InputStickerStatic{}
)

func (i *InputStickerStatic) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Sticker == nil) {
		return false
	}
	if !(i.Emojis == "") {
		return false
	}
	if !(i.MaskPosition.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerStatic) String() string {
	if i == nil {
		return "InputStickerStatic(nil)"
	}
	type Alias InputStickerStatic
	return fmt.Sprintf("InputStickerStatic%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerStatic) TypeID() uint32 {
	return InputStickerStaticTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerStatic) TypeName() string {
	return "inputStickerStatic"
}

// TypeInfo returns info about TL type.
func (i *InputStickerStatic) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerStatic",
		ID:   InputStickerStaticTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Emojis",
			SchemaName: "emojis",
		},
		{
			Name:       "MaskPosition",
			SchemaName: "mask_position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerStatic) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerStatic#540604db as nil")
	}
	b.PutID(InputStickerStaticTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerStatic) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerStatic#540604db as nil")
	}
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field sticker is nil")
	}
	if err := i.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field sticker: %w", err)
	}
	b.PutString(i.Emojis)
	if err := i.MaskPosition.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field mask_position: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerStatic) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerStatic#540604db to nil")
	}
	if err := b.ConsumeID(InputStickerStaticTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerStatic#540604db: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerStatic) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerStatic#540604db to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerStatic#540604db: field sticker: %w", err)
		}
		i.Sticker = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerStatic#540604db: field emojis: %w", err)
		}
		i.Emojis = value
	}
	{
		if err := i.MaskPosition.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerStatic#540604db: field mask_position: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStickerStatic) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerStatic#540604db as nil")
	}
	b.ObjStart()
	b.PutID("inputStickerStatic")
	b.FieldStart("sticker")
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field sticker is nil")
	}
	if err := i.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field sticker: %w", err)
	}
	b.FieldStart("emojis")
	b.PutString(i.Emojis)
	b.FieldStart("mask_position")
	if err := i.MaskPosition.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerStatic#540604db: field mask_position: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStickerStatic) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerStatic#540604db to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStickerStatic"); err != nil {
				return fmt.Errorf("unable to decode inputStickerStatic#540604db: %w", err)
			}
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputStickerStatic#540604db: field sticker: %w", err)
			}
			i.Sticker = value
		case "emojis":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputStickerStatic#540604db: field emojis: %w", err)
			}
			i.Emojis = value
		case "mask_position":
			if err := i.MaskPosition.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputStickerStatic#540604db: field mask_position: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (i *InputStickerStatic) GetSticker() (value InputFileClass) {
	return i.Sticker
}

// GetEmojis returns value of Emojis field.
func (i *InputStickerStatic) GetEmojis() (value string) {
	return i.Emojis
}

// GetMaskPosition returns value of MaskPosition field.
func (i *InputStickerStatic) GetMaskPosition() (value MaskPosition) {
	return i.MaskPosition
}

// InputStickerAnimated represents TL type `inputStickerAnimated#bccf4960`.
type InputStickerAnimated struct {
	// File with the animated sticker. Only local or uploaded within a week files are
	// supported. See https://core.telegram.org/animated_stickers#technical-requirements for
	// technical requirements
	Sticker InputFileClass
	// Emojis corresponding to the sticker
	Emojis string
}

// InputStickerAnimatedTypeID is TL type id of InputStickerAnimated.
const InputStickerAnimatedTypeID = 0xbccf4960

// construct implements constructor of InputStickerClass.
func (i InputStickerAnimated) construct() InputStickerClass { return &i }

// Ensuring interfaces in compile-time for InputStickerAnimated.
var (
	_ bin.Encoder     = &InputStickerAnimated{}
	_ bin.Decoder     = &InputStickerAnimated{}
	_ bin.BareEncoder = &InputStickerAnimated{}
	_ bin.BareDecoder = &InputStickerAnimated{}

	_ InputStickerClass = &InputStickerAnimated{}
)

func (i *InputStickerAnimated) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Sticker == nil) {
		return false
	}
	if !(i.Emojis == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerAnimated) String() string {
	if i == nil {
		return "InputStickerAnimated(nil)"
	}
	type Alias InputStickerAnimated
	return fmt.Sprintf("InputStickerAnimated%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerAnimated) TypeID() uint32 {
	return InputStickerAnimatedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerAnimated) TypeName() string {
	return "inputStickerAnimated"
}

// TypeInfo returns info about TL type.
func (i *InputStickerAnimated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerAnimated",
		ID:   InputStickerAnimatedTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Emojis",
			SchemaName: "emojis",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStickerAnimated) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerAnimated#bccf4960 as nil")
	}
	b.PutID(InputStickerAnimatedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerAnimated) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerAnimated#bccf4960 as nil")
	}
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputStickerAnimated#bccf4960: field sticker is nil")
	}
	if err := i.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerAnimated#bccf4960: field sticker: %w", err)
	}
	b.PutString(i.Emojis)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerAnimated) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerAnimated#bccf4960 to nil")
	}
	if err := b.ConsumeID(InputStickerAnimatedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerAnimated) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerAnimated#bccf4960 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: field sticker: %w", err)
		}
		i.Sticker = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: field emojis: %w", err)
		}
		i.Emojis = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStickerAnimated) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerAnimated#bccf4960 as nil")
	}
	b.ObjStart()
	b.PutID("inputStickerAnimated")
	b.FieldStart("sticker")
	if i.Sticker == nil {
		return fmt.Errorf("unable to encode inputStickerAnimated#bccf4960: field sticker is nil")
	}
	if err := i.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerAnimated#bccf4960: field sticker: %w", err)
	}
	b.FieldStart("emojis")
	b.PutString(i.Emojis)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStickerAnimated) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerAnimated#bccf4960 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStickerAnimated"); err != nil {
				return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: %w", err)
			}
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: field sticker: %w", err)
			}
			i.Sticker = value
		case "emojis":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputStickerAnimated#bccf4960: field emojis: %w", err)
			}
			i.Emojis = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (i *InputStickerAnimated) GetSticker() (value InputFileClass) {
	return i.Sticker
}

// GetEmojis returns value of Emojis field.
func (i *InputStickerAnimated) GetEmojis() (value string) {
	return i.Emojis
}

// InputStickerClassName is schema name of InputStickerClass.
const InputStickerClassName = "InputSticker"

// InputStickerClass represents InputSticker generic type.
//
// Example:
//  g, err := tdapi.DecodeInputSticker(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.InputStickerStatic: // inputStickerStatic#540604db
//  case *tdapi.InputStickerAnimated: // inputStickerAnimated#bccf4960
//  default: panic(v)
//  }
type InputStickerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStickerClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error

	// PNG image with the sticker; must be up to 512 KB in size and fit in a 512x512 square
	GetSticker() (value InputFileClass)
	// Emojis corresponding to the sticker
	GetEmojis() (value string)
}

// DecodeInputSticker implements binary de-serialization for InputStickerClass.
func DecodeInputSticker(buf *bin.Buffer) (InputStickerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStickerStaticTypeID:
		// Decoding inputStickerStatic#540604db.
		v := InputStickerStatic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerClass: %w", err)
		}
		return &v, nil
	case InputStickerAnimatedTypeID:
		// Decoding inputStickerAnimated#bccf4960.
		v := InputStickerAnimated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickerClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputSticker implements binary de-serialization for InputStickerClass.
func DecodeTDLibJSONInputSticker(buf tdjson.Decoder) (InputStickerClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputStickerStatic":
		// Decoding inputStickerStatic#540604db.
		v := InputStickerStatic{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerClass: %w", err)
		}
		return &v, nil
	case "inputStickerAnimated":
		// Decoding inputStickerAnimated#bccf4960.
		v := InputStickerAnimated{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStickerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStickerClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputSticker boxes the InputStickerClass providing a helper.
type InputStickerBox struct {
	InputSticker InputStickerClass
}

// Decode implements bin.Decoder for InputStickerBox.
func (b *InputStickerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickerBox to nil")
	}
	v, err := DecodeInputSticker(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputSticker = v
	return nil
}

// Encode implements bin.Encode for InputStickerBox.
func (b *InputStickerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputSticker == nil {
		return fmt.Errorf("unable to encode InputStickerClass as nil")
	}
	return b.InputSticker.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputStickerBox.
func (b *InputStickerBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStickerBox to nil")
	}
	v, err := DecodeTDLibJSONInputSticker(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputSticker = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputStickerBox.
func (b *InputStickerBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputSticker == nil {
		return fmt.Errorf("unable to encode InputStickerClass as nil")
	}
	return b.InputSticker.EncodeTDLibJSON(buf)
}
