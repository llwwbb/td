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

// InputWallPaper represents TL type `inputWallPaper#e630b979`.
// Wallpaper
//
// See https://core.telegram.org/constructor/inputWallPaper for reference.
type InputWallPaper struct {
	// Wallpaper ID
	ID int64
	// Access hash
	AccessHash int64
}

// InputWallPaperTypeID is TL type id of InputWallPaper.
const InputWallPaperTypeID = 0xe630b979

// Encode implements bin.Encoder.
func (i *InputWallPaper) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaper#e630b979 as nil")
	}
	b.PutID(InputWallPaperTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaper) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaper#e630b979 to nil")
	}
	if err := b.ConsumeID(InputWallPaperTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaper#e630b979: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaper#e630b979: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaper#e630b979: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaper) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaper.
var (
	_ bin.Encoder = &InputWallPaper{}
	_ bin.Decoder = &InputWallPaper{}

	_ InputWallPaperClass = &InputWallPaper{}
)

// InputWallPaperSlug represents TL type `inputWallPaperSlug#72091c80`.
// Wallpaper by slug (a unique ID)
//
// See https://core.telegram.org/constructor/inputWallPaperSlug for reference.
type InputWallPaperSlug struct {
	// Unique wallpaper ID
	Slug string
}

// InputWallPaperSlugTypeID is TL type id of InputWallPaperSlug.
const InputWallPaperSlugTypeID = 0x72091c80

// Encode implements bin.Encoder.
func (i *InputWallPaperSlug) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperSlug#72091c80 as nil")
	}
	b.PutID(InputWallPaperSlugTypeID)
	b.PutString(i.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaperSlug) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperSlug#72091c80 to nil")
	}
	if err := b.ConsumeID(InputWallPaperSlugTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaperSlug#72091c80: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaperSlug#72091c80: field slug: %w", err)
		}
		i.Slug = value
	}
	return nil
}

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaperSlug) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaperSlug.
var (
	_ bin.Encoder = &InputWallPaperSlug{}
	_ bin.Decoder = &InputWallPaperSlug{}

	_ InputWallPaperClass = &InputWallPaperSlug{}
)

// InputWallPaperNoFile represents TL type `inputWallPaperNoFile#8427bbac`.
// Wallpaper with no file
//
// See https://core.telegram.org/constructor/inputWallPaperNoFile for reference.
type InputWallPaperNoFile struct {
}

// InputWallPaperNoFileTypeID is TL type id of InputWallPaperNoFile.
const InputWallPaperNoFileTypeID = 0x8427bbac

// Encode implements bin.Encoder.
func (i *InputWallPaperNoFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperNoFile#8427bbac as nil")
	}
	b.PutID(InputWallPaperNoFileTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaperNoFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperNoFile#8427bbac to nil")
	}
	if err := b.ConsumeID(InputWallPaperNoFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaperNoFile#8427bbac: %w", err)
	}
	return nil
}

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaperNoFile) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaperNoFile.
var (
	_ bin.Encoder = &InputWallPaperNoFile{}
	_ bin.Decoder = &InputWallPaperNoFile{}

	_ InputWallPaperClass = &InputWallPaperNoFile{}
)

// InputWallPaperClass represents InputWallPaper generic type.
//
// See https://core.telegram.org/type/InputWallPaper for reference.
//
// Example:
//  g, err := DecodeInputWallPaper(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputWallPaper: // inputWallPaper#e630b979
//  case *InputWallPaperSlug: // inputWallPaperSlug#72091c80
//  case *InputWallPaperNoFile: // inputWallPaperNoFile#8427bbac
//  default: panic(v)
//  }
type InputWallPaperClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputWallPaperClass
}

// DecodeInputWallPaper implements binary de-serialization for InputWallPaperClass.
func DecodeInputWallPaper(buf *bin.Buffer) (InputWallPaperClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputWallPaperTypeID:
		// Decoding inputWallPaper#e630b979.
		v := InputWallPaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	case InputWallPaperSlugTypeID:
		// Decoding inputWallPaperSlug#72091c80.
		v := InputWallPaperSlug{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	case InputWallPaperNoFileTypeID:
		// Decoding inputWallPaperNoFile#8427bbac.
		v := InputWallPaperNoFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputWallPaper boxes the InputWallPaperClass providing a helper.
type InputWallPaperBox struct {
	InputWallPaper InputWallPaperClass
}

// Decode implements bin.Decoder for InputWallPaperBox.
func (b *InputWallPaperBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputWallPaperBox to nil")
	}
	v, err := DecodeInputWallPaper(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputWallPaper = v
	return nil
}

// Encode implements bin.Encode for InputWallPaperBox.
func (b *InputWallPaperBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputWallPaper == nil {
		return fmt.Errorf("unable to encode InputWallPaperClass as nil")
	}
	return b.InputWallPaper.Encode(buf)
}
