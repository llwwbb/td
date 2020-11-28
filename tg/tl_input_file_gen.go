// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// InputFile represents TL type `inputFile#f52ff27f`.
type InputFile struct {
	// ID field of InputFile.
	ID int64
	// Parts field of InputFile.
	Parts int
	// Name field of InputFile.
	Name string
	// Md5Checksum field of InputFile.
	Md5Checksum string
}

// InputFileTypeID is TL type id of InputFile.
const InputFileTypeID = 0xf52ff27f

// Encode implements bin.Encoder.
func (i *InputFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFile#f52ff27f as nil")
	}
	b.PutID(InputFileTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Name)
	b.PutString(i.Md5Checksum)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFile#f52ff27f to nil")
	}
	if err := b.ConsumeID(InputFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFile#f52ff27f: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field name: %w", err)
		}
		i.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFile#f52ff27f: field md5_checksum: %w", err)
		}
		i.Md5Checksum = value
	}
	return nil
}

// construct implements constructor of InputFileClass.
func (i InputFile) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFile.
var (
	_ bin.Encoder = &InputFile{}
	_ bin.Decoder = &InputFile{}

	_ InputFileClass = &InputFile{}
)

// InputFileBig represents TL type `inputFileBig#fa4f0bb5`.
type InputFileBig struct {
	// ID field of InputFileBig.
	ID int64
	// Parts field of InputFileBig.
	Parts int
	// Name field of InputFileBig.
	Name string
}

// InputFileBigTypeID is TL type id of InputFileBig.
const InputFileBigTypeID = 0xfa4f0bb5

// Encode implements bin.Encoder.
func (i *InputFileBig) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileBig#fa4f0bb5 as nil")
	}
	b.PutID(InputFileBigTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Name)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileBig) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileBig#fa4f0bb5 to nil")
	}
	if err := b.ConsumeID(InputFileBigTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileBig#fa4f0bb5: field name: %w", err)
		}
		i.Name = value
	}
	return nil
}

// construct implements constructor of InputFileClass.
func (i InputFileBig) construct() InputFileClass { return &i }

// Ensuring interfaces in compile-time for InputFileBig.
var (
	_ bin.Encoder = &InputFileBig{}
	_ bin.Decoder = &InputFileBig{}

	_ InputFileClass = &InputFileBig{}
)

// InputFileClass represents InputFile generic type.
//
// Example:
//  g, err := DecodeInputFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputFile: // inputFile#f52ff27f
//  case *InputFileBig: // inputFileBig#fa4f0bb5
//  default: panic(v)
//  }
type InputFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputFileClass
}

// DecodeInputFile implements binary de-serialization for InputFileClass.
func DecodeInputFile(buf *bin.Buffer) (InputFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputFileTypeID:
		// Decoding inputFile#f52ff27f.
		v := InputFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	case InputFileBigTypeID:
		// Decoding inputFileBig#fa4f0bb5.
		v := InputFileBig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputFile boxes the InputFileClass providing a helper.
type InputFileBox struct {
	InputFile InputFileClass
}

// Decode implements bin.Decoder for InputFileBox.
func (b *InputFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputFileBox to nil")
	}
	v, err := DecodeInputFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputFile = v
	return nil
}

// Encode implements bin.Encode for InputFileBox.
func (b *InputFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputFile == nil {
		return fmt.Errorf("unable to encode InputFileClass as nil")
	}
	return b.InputFile.Encode(buf)
}
