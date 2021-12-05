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

// BackgroundTypeWallpaper represents TL type `backgroundTypeWallpaper#758c4c7b`.
type BackgroundTypeWallpaper struct {
	// True, if the wallpaper must be downscaled to fit in 450x450 square and then
	// box-blurred with radius 12
	IsBlurred bool
	// True, if the background needs to be slightly moved when device is tilted
	IsMoving bool
}

// BackgroundTypeWallpaperTypeID is TL type id of BackgroundTypeWallpaper.
const BackgroundTypeWallpaperTypeID = 0x758c4c7b

// construct implements constructor of BackgroundTypeClass.
func (b BackgroundTypeWallpaper) construct() BackgroundTypeClass { return &b }

// Ensuring interfaces in compile-time for BackgroundTypeWallpaper.
var (
	_ bin.Encoder     = &BackgroundTypeWallpaper{}
	_ bin.Decoder     = &BackgroundTypeWallpaper{}
	_ bin.BareEncoder = &BackgroundTypeWallpaper{}
	_ bin.BareDecoder = &BackgroundTypeWallpaper{}

	_ BackgroundTypeClass = &BackgroundTypeWallpaper{}
)

func (b *BackgroundTypeWallpaper) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.IsBlurred == false) {
		return false
	}
	if !(b.IsMoving == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundTypeWallpaper) String() string {
	if b == nil {
		return "BackgroundTypeWallpaper(nil)"
	}
	type Alias BackgroundTypeWallpaper
	return fmt.Sprintf("BackgroundTypeWallpaper%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundTypeWallpaper) TypeID() uint32 {
	return BackgroundTypeWallpaperTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundTypeWallpaper) TypeName() string {
	return "backgroundTypeWallpaper"
}

// TypeInfo returns info about TL type.
func (b *BackgroundTypeWallpaper) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundTypeWallpaper",
		ID:   BackgroundTypeWallpaperTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsBlurred",
			SchemaName: "is_blurred",
		},
		{
			Name:       "IsMoving",
			SchemaName: "is_moving",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundTypeWallpaper) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeWallpaper#758c4c7b as nil")
	}
	buf.PutID(BackgroundTypeWallpaperTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundTypeWallpaper) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeWallpaper#758c4c7b as nil")
	}
	buf.PutBool(b.IsBlurred)
	buf.PutBool(b.IsMoving)
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundTypeWallpaper) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeWallpaper#758c4c7b to nil")
	}
	if err := buf.ConsumeID(BackgroundTypeWallpaperTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundTypeWallpaper) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeWallpaper#758c4c7b to nil")
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: field is_blurred: %w", err)
		}
		b.IsBlurred = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: field is_moving: %w", err)
		}
		b.IsMoving = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BackgroundTypeWallpaper) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeWallpaper#758c4c7b as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundTypeWallpaper")
	buf.FieldStart("is_blurred")
	buf.PutBool(b.IsBlurred)
	buf.FieldStart("is_moving")
	buf.PutBool(b.IsMoving)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BackgroundTypeWallpaper) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeWallpaper#758c4c7b to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("backgroundTypeWallpaper"); err != nil {
				return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: %w", err)
			}
		case "is_blurred":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: field is_blurred: %w", err)
			}
			b.IsBlurred = value
		case "is_moving":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypeWallpaper#758c4c7b: field is_moving: %w", err)
			}
			b.IsMoving = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetIsBlurred returns value of IsBlurred field.
func (b *BackgroundTypeWallpaper) GetIsBlurred() (value bool) {
	return b.IsBlurred
}

// GetIsMoving returns value of IsMoving field.
func (b *BackgroundTypeWallpaper) GetIsMoving() (value bool) {
	return b.IsMoving
}

// BackgroundTypePattern represents TL type `backgroundTypePattern#4ce716fd`.
type BackgroundTypePattern struct {
	// Fill of the background
	Fill BackgroundFillClass
	// Intensity of the pattern when it is shown above the filled background; 0-100.
	Intensity int32
	// True, if the background fill must be applied only to the pattern itself. All other
	// pixels are black in this case. For dark themes only
	IsInverted bool
	// True, if the background needs to be slightly moved when device is tilted
	IsMoving bool
}

// BackgroundTypePatternTypeID is TL type id of BackgroundTypePattern.
const BackgroundTypePatternTypeID = 0x4ce716fd

// construct implements constructor of BackgroundTypeClass.
func (b BackgroundTypePattern) construct() BackgroundTypeClass { return &b }

// Ensuring interfaces in compile-time for BackgroundTypePattern.
var (
	_ bin.Encoder     = &BackgroundTypePattern{}
	_ bin.Decoder     = &BackgroundTypePattern{}
	_ bin.BareEncoder = &BackgroundTypePattern{}
	_ bin.BareDecoder = &BackgroundTypePattern{}

	_ BackgroundTypeClass = &BackgroundTypePattern{}
)

func (b *BackgroundTypePattern) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Fill == nil) {
		return false
	}
	if !(b.Intensity == 0) {
		return false
	}
	if !(b.IsInverted == false) {
		return false
	}
	if !(b.IsMoving == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundTypePattern) String() string {
	if b == nil {
		return "BackgroundTypePattern(nil)"
	}
	type Alias BackgroundTypePattern
	return fmt.Sprintf("BackgroundTypePattern%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundTypePattern) TypeID() uint32 {
	return BackgroundTypePatternTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundTypePattern) TypeName() string {
	return "backgroundTypePattern"
}

// TypeInfo returns info about TL type.
func (b *BackgroundTypePattern) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundTypePattern",
		ID:   BackgroundTypePatternTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Fill",
			SchemaName: "fill",
		},
		{
			Name:       "Intensity",
			SchemaName: "intensity",
		},
		{
			Name:       "IsInverted",
			SchemaName: "is_inverted",
		},
		{
			Name:       "IsMoving",
			SchemaName: "is_moving",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundTypePattern) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypePattern#4ce716fd as nil")
	}
	buf.PutID(BackgroundTypePatternTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundTypePattern) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypePattern#4ce716fd as nil")
	}
	if b.Fill == nil {
		return fmt.Errorf("unable to encode backgroundTypePattern#4ce716fd: field fill is nil")
	}
	if err := b.Fill.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode backgroundTypePattern#4ce716fd: field fill: %w", err)
	}
	buf.PutInt32(b.Intensity)
	buf.PutBool(b.IsInverted)
	buf.PutBool(b.IsMoving)
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundTypePattern) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypePattern#4ce716fd to nil")
	}
	if err := buf.ConsumeID(BackgroundTypePatternTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundTypePattern) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypePattern#4ce716fd to nil")
	}
	{
		value, err := DecodeBackgroundFill(buf)
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field fill: %w", err)
		}
		b.Fill = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field intensity: %w", err)
		}
		b.Intensity = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field is_inverted: %w", err)
		}
		b.IsInverted = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field is_moving: %w", err)
		}
		b.IsMoving = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BackgroundTypePattern) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypePattern#4ce716fd as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundTypePattern")
	buf.FieldStart("fill")
	if b.Fill == nil {
		return fmt.Errorf("unable to encode backgroundTypePattern#4ce716fd: field fill is nil")
	}
	if err := b.Fill.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode backgroundTypePattern#4ce716fd: field fill: %w", err)
	}
	buf.FieldStart("intensity")
	buf.PutInt32(b.Intensity)
	buf.FieldStart("is_inverted")
	buf.PutBool(b.IsInverted)
	buf.FieldStart("is_moving")
	buf.PutBool(b.IsMoving)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BackgroundTypePattern) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypePattern#4ce716fd to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("backgroundTypePattern"); err != nil {
				return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: %w", err)
			}
		case "fill":
			value, err := DecodeTDLibJSONBackgroundFill(buf)
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field fill: %w", err)
			}
			b.Fill = value
		case "intensity":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field intensity: %w", err)
			}
			b.Intensity = value
		case "is_inverted":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field is_inverted: %w", err)
			}
			b.IsInverted = value
		case "is_moving":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypePattern#4ce716fd: field is_moving: %w", err)
			}
			b.IsMoving = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetFill returns value of Fill field.
func (b *BackgroundTypePattern) GetFill() (value BackgroundFillClass) {
	return b.Fill
}

// GetIntensity returns value of Intensity field.
func (b *BackgroundTypePattern) GetIntensity() (value int32) {
	return b.Intensity
}

// GetIsInverted returns value of IsInverted field.
func (b *BackgroundTypePattern) GetIsInverted() (value bool) {
	return b.IsInverted
}

// GetIsMoving returns value of IsMoving field.
func (b *BackgroundTypePattern) GetIsMoving() (value bool) {
	return b.IsMoving
}

// BackgroundTypeFill represents TL type `backgroundTypeFill#3b301c2c`.
type BackgroundTypeFill struct {
	// The background fill
	Fill BackgroundFillClass
}

// BackgroundTypeFillTypeID is TL type id of BackgroundTypeFill.
const BackgroundTypeFillTypeID = 0x3b301c2c

// construct implements constructor of BackgroundTypeClass.
func (b BackgroundTypeFill) construct() BackgroundTypeClass { return &b }

// Ensuring interfaces in compile-time for BackgroundTypeFill.
var (
	_ bin.Encoder     = &BackgroundTypeFill{}
	_ bin.Decoder     = &BackgroundTypeFill{}
	_ bin.BareEncoder = &BackgroundTypeFill{}
	_ bin.BareDecoder = &BackgroundTypeFill{}

	_ BackgroundTypeClass = &BackgroundTypeFill{}
)

func (b *BackgroundTypeFill) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Fill == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BackgroundTypeFill) String() string {
	if b == nil {
		return "BackgroundTypeFill(nil)"
	}
	type Alias BackgroundTypeFill
	return fmt.Sprintf("BackgroundTypeFill%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BackgroundTypeFill) TypeID() uint32 {
	return BackgroundTypeFillTypeID
}

// TypeName returns name of type in TL schema.
func (*BackgroundTypeFill) TypeName() string {
	return "backgroundTypeFill"
}

// TypeInfo returns info about TL type.
func (b *BackgroundTypeFill) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "backgroundTypeFill",
		ID:   BackgroundTypeFillTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Fill",
			SchemaName: "fill",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BackgroundTypeFill) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeFill#3b301c2c as nil")
	}
	buf.PutID(BackgroundTypeFillTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BackgroundTypeFill) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeFill#3b301c2c as nil")
	}
	if b.Fill == nil {
		return fmt.Errorf("unable to encode backgroundTypeFill#3b301c2c: field fill is nil")
	}
	if err := b.Fill.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode backgroundTypeFill#3b301c2c: field fill: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BackgroundTypeFill) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeFill#3b301c2c to nil")
	}
	if err := buf.ConsumeID(BackgroundTypeFillTypeID); err != nil {
		return fmt.Errorf("unable to decode backgroundTypeFill#3b301c2c: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BackgroundTypeFill) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeFill#3b301c2c to nil")
	}
	{
		value, err := DecodeBackgroundFill(buf)
		if err != nil {
			return fmt.Errorf("unable to decode backgroundTypeFill#3b301c2c: field fill: %w", err)
		}
		b.Fill = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BackgroundTypeFill) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode backgroundTypeFill#3b301c2c as nil")
	}
	buf.ObjStart()
	buf.PutID("backgroundTypeFill")
	buf.FieldStart("fill")
	if b.Fill == nil {
		return fmt.Errorf("unable to encode backgroundTypeFill#3b301c2c: field fill is nil")
	}
	if err := b.Fill.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode backgroundTypeFill#3b301c2c: field fill: %w", err)
	}
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BackgroundTypeFill) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode backgroundTypeFill#3b301c2c to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("backgroundTypeFill"); err != nil {
				return fmt.Errorf("unable to decode backgroundTypeFill#3b301c2c: %w", err)
			}
		case "fill":
			value, err := DecodeTDLibJSONBackgroundFill(buf)
			if err != nil {
				return fmt.Errorf("unable to decode backgroundTypeFill#3b301c2c: field fill: %w", err)
			}
			b.Fill = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetFill returns value of Fill field.
func (b *BackgroundTypeFill) GetFill() (value BackgroundFillClass) {
	return b.Fill
}

// BackgroundTypeClassName is schema name of BackgroundTypeClass.
const BackgroundTypeClassName = "BackgroundType"

// BackgroundTypeClass represents BackgroundType generic type.
//
// Example:
//  g, err := tdapi.DecodeBackgroundType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.BackgroundTypeWallpaper: // backgroundTypeWallpaper#758c4c7b
//  case *tdapi.BackgroundTypePattern: // backgroundTypePattern#4ce716fd
//  case *tdapi.BackgroundTypeFill: // backgroundTypeFill#3b301c2c
//  default: panic(v)
//  }
type BackgroundTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BackgroundTypeClass

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
}

// DecodeBackgroundType implements binary de-serialization for BackgroundTypeClass.
func DecodeBackgroundType(buf *bin.Buffer) (BackgroundTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BackgroundTypeWallpaperTypeID:
		// Decoding backgroundTypeWallpaper#758c4c7b.
		v := BackgroundTypeWallpaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	case BackgroundTypePatternTypeID:
		// Decoding backgroundTypePattern#4ce716fd.
		v := BackgroundTypePattern{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	case BackgroundTypeFillTypeID:
		// Decoding backgroundTypeFill#3b301c2c.
		v := BackgroundTypeFill{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONBackgroundType implements binary de-serialization for BackgroundTypeClass.
func DecodeTDLibJSONBackgroundType(buf tdjson.Decoder) (BackgroundTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "backgroundTypeWallpaper":
		// Decoding backgroundTypeWallpaper#758c4c7b.
		v := BackgroundTypeWallpaper{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	case "backgroundTypePattern":
		// Decoding backgroundTypePattern#4ce716fd.
		v := BackgroundTypePattern{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	case "backgroundTypeFill":
		// Decoding backgroundTypeFill#3b301c2c.
		v := BackgroundTypeFill{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BackgroundTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// BackgroundType boxes the BackgroundTypeClass providing a helper.
type BackgroundTypeBox struct {
	BackgroundType BackgroundTypeClass
}

// Decode implements bin.Decoder for BackgroundTypeBox.
func (b *BackgroundTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BackgroundTypeBox to nil")
	}
	v, err := DecodeBackgroundType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BackgroundType = v
	return nil
}

// Encode implements bin.Encode for BackgroundTypeBox.
func (b *BackgroundTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BackgroundType == nil {
		return fmt.Errorf("unable to encode BackgroundTypeClass as nil")
	}
	return b.BackgroundType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for BackgroundTypeBox.
func (b *BackgroundTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode BackgroundTypeBox to nil")
	}
	v, err := DecodeTDLibJSONBackgroundType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BackgroundType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for BackgroundTypeBox.
func (b *BackgroundTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.BackgroundType == nil {
		return fmt.Errorf("unable to encode BackgroundTypeClass as nil")
	}
	return b.BackgroundType.EncodeTDLibJSON(buf)
}
