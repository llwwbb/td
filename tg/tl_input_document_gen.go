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

// InputDocumentEmpty represents TL type `inputDocumentEmpty#72f0eaae`.
type InputDocumentEmpty struct {
}

// InputDocumentEmptyTypeID is TL type id of InputDocumentEmpty.
const InputDocumentEmptyTypeID = 0x72f0eaae

// Encode implements bin.Encoder.
func (i *InputDocumentEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocumentEmpty#72f0eaae as nil")
	}
	b.PutID(InputDocumentEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDocumentEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocumentEmpty#72f0eaae to nil")
	}
	if err := b.ConsumeID(InputDocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocumentEmpty#72f0eaae: %w", err)
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocumentEmpty) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocumentEmpty.
var (
	_ bin.Encoder = &InputDocumentEmpty{}
	_ bin.Decoder = &InputDocumentEmpty{}

	_ InputDocumentClass = &InputDocumentEmpty{}
)

// InputDocument represents TL type `inputDocument#1abfb575`.
type InputDocument struct {
	// ID field of InputDocument.
	ID int64
	// AccessHash field of InputDocument.
	AccessHash int64
	// FileReference field of InputDocument.
	FileReference []byte
}

// InputDocumentTypeID is TL type id of InputDocument.
const InputDocumentTypeID = 0x1abfb575

// Encode implements bin.Encoder.
func (i *InputDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocument#1abfb575 as nil")
	}
	b.PutID(InputDocumentTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocument#1abfb575 to nil")
	}
	if err := b.ConsumeID(InputDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocument#1abfb575: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocument) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocument.
var (
	_ bin.Encoder = &InputDocument{}
	_ bin.Decoder = &InputDocument{}

	_ InputDocumentClass = &InputDocument{}
)

// InputDocumentClass represents InputDocument generic type.
//
// Example:
//  g, err := DecodeInputDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputDocumentEmpty: // inputDocumentEmpty#72f0eaae
//  case *InputDocument: // inputDocument#1abfb575
//  default: panic(v)
//  }
type InputDocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputDocumentClass
}

// DecodeInputDocument implements binary de-serialization for InputDocumentClass.
func DecodeInputDocument(buf *bin.Buffer) (InputDocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputDocumentEmptyTypeID:
		// Decoding inputDocumentEmpty#72f0eaae.
		v := InputDocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	case InputDocumentTypeID:
		// Decoding inputDocument#1abfb575.
		v := InputDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputDocument boxes the InputDocumentClass providing a helper.
type InputDocumentBox struct {
	InputDocument InputDocumentClass
}

// Decode implements bin.Decoder for InputDocumentBox.
func (b *InputDocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputDocumentBox to nil")
	}
	v, err := DecodeInputDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputDocument = v
	return nil
}

// Encode implements bin.Encode for InputDocumentBox.
func (b *InputDocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputDocument == nil {
		return fmt.Errorf("unable to encode InputDocumentClass as nil")
	}
	return b.InputDocument.Encode(buf)
}
