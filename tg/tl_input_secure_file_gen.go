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

// InputSecureFileUploaded represents TL type `inputSecureFileUploaded#3334b0f0`.
type InputSecureFileUploaded struct {
	// ID field of InputSecureFileUploaded.
	ID int64
	// Parts field of InputSecureFileUploaded.
	Parts int
	// Md5Checksum field of InputSecureFileUploaded.
	Md5Checksum string
	// FileHash field of InputSecureFileUploaded.
	FileHash []byte
	// Secret field of InputSecureFileUploaded.
	Secret []byte
}

// InputSecureFileUploadedTypeID is TL type id of InputSecureFileUploaded.
const InputSecureFileUploadedTypeID = 0x3334b0f0

// Encode implements bin.Encoder.
func (i *InputSecureFileUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFileUploaded#3334b0f0 as nil")
	}
	b.PutID(InputSecureFileUploadedTypeID)
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.Md5Checksum)
	b.PutBytes(i.FileHash)
	b.PutBytes(i.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSecureFileUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFileUploaded#3334b0f0 to nil")
	}
	if err := b.ConsumeID(InputSecureFileUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field md5_checksum: %w", err)
		}
		i.Md5Checksum = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field file_hash: %w", err)
		}
		i.FileHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field secret: %w", err)
		}
		i.Secret = value
	}
	return nil
}

// construct implements constructor of InputSecureFileClass.
func (i InputSecureFileUploaded) construct() InputSecureFileClass { return &i }

// Ensuring interfaces in compile-time for InputSecureFileUploaded.
var (
	_ bin.Encoder = &InputSecureFileUploaded{}
	_ bin.Decoder = &InputSecureFileUploaded{}

	_ InputSecureFileClass = &InputSecureFileUploaded{}
)

// InputSecureFile represents TL type `inputSecureFile#5367e5be`.
type InputSecureFile struct {
	// ID field of InputSecureFile.
	ID int64
	// AccessHash field of InputSecureFile.
	AccessHash int64
}

// InputSecureFileTypeID is TL type id of InputSecureFile.
const InputSecureFileTypeID = 0x5367e5be

// Encode implements bin.Encoder.
func (i *InputSecureFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFile#5367e5be as nil")
	}
	b.PutID(InputSecureFileTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSecureFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFile#5367e5be to nil")
	}
	if err := b.ConsumeID(InputSecureFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureFile#5367e5be: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFile#5367e5be: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFile#5367e5be: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputSecureFileClass.
func (i InputSecureFile) construct() InputSecureFileClass { return &i }

// Ensuring interfaces in compile-time for InputSecureFile.
var (
	_ bin.Encoder = &InputSecureFile{}
	_ bin.Decoder = &InputSecureFile{}

	_ InputSecureFileClass = &InputSecureFile{}
)

// InputSecureFileClass represents InputSecureFile generic type.
//
// Example:
//  g, err := DecodeInputSecureFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputSecureFileUploaded: // inputSecureFileUploaded#3334b0f0
//  case *InputSecureFile: // inputSecureFile#5367e5be
//  default: panic(v)
//  }
type InputSecureFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputSecureFileClass
}

// DecodeInputSecureFile implements binary de-serialization for InputSecureFileClass.
func DecodeInputSecureFile(buf *bin.Buffer) (InputSecureFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputSecureFileUploadedTypeID:
		// Decoding inputSecureFileUploaded#3334b0f0.
		v := InputSecureFileUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", err)
		}
		return &v, nil
	case InputSecureFileTypeID:
		// Decoding inputSecureFile#5367e5be.
		v := InputSecureFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputSecureFile boxes the InputSecureFileClass providing a helper.
type InputSecureFileBox struct {
	InputSecureFile InputSecureFileClass
}

// Decode implements bin.Decoder for InputSecureFileBox.
func (b *InputSecureFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputSecureFileBox to nil")
	}
	v, err := DecodeInputSecureFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputSecureFile = v
	return nil
}

// Encode implements bin.Encode for InputSecureFileBox.
func (b *InputSecureFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputSecureFile == nil {
		return fmt.Errorf("unable to encode InputSecureFileClass as nil")
	}
	return b.InputSecureFile.Encode(buf)
}
