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

// InputMessageID represents TL type `inputMessageID#a676a322`.
type InputMessageID struct {
	// ID field of InputMessageID.
	ID int
}

// InputMessageIDTypeID is TL type id of InputMessageID.
const InputMessageIDTypeID = 0xa676a322

// Encode implements bin.Encoder.
func (i *InputMessageID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageID#a676a322 as nil")
	}
	b.PutID(InputMessageIDTypeID)
	b.PutInt(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageID#a676a322 to nil")
	}
	if err := b.ConsumeID(InputMessageIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageID#a676a322: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageID#a676a322: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageID) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageID.
var (
	_ bin.Encoder = &InputMessageID{}
	_ bin.Decoder = &InputMessageID{}

	_ InputMessageClass = &InputMessageID{}
)

// InputMessageReplyTo represents TL type `inputMessageReplyTo#bad88395`.
type InputMessageReplyTo struct {
	// ID field of InputMessageReplyTo.
	ID int
}

// InputMessageReplyToTypeID is TL type id of InputMessageReplyTo.
const InputMessageReplyToTypeID = 0xbad88395

// Encode implements bin.Encoder.
func (i *InputMessageReplyTo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyTo#bad88395 as nil")
	}
	b.PutID(InputMessageReplyToTypeID)
	b.PutInt(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageReplyTo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyTo#bad88395 to nil")
	}
	if err := b.ConsumeID(InputMessageReplyToTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageReplyTo#bad88395: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyTo#bad88395: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageReplyTo) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageReplyTo.
var (
	_ bin.Encoder = &InputMessageReplyTo{}
	_ bin.Decoder = &InputMessageReplyTo{}

	_ InputMessageClass = &InputMessageReplyTo{}
)

// InputMessagePinned represents TL type `inputMessagePinned#86872538`.
type InputMessagePinned struct {
}

// InputMessagePinnedTypeID is TL type id of InputMessagePinned.
const InputMessagePinnedTypeID = 0x86872538

// Encode implements bin.Encoder.
func (i *InputMessagePinned) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagePinned#86872538 as nil")
	}
	b.PutID(InputMessagePinnedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagePinned) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagePinned#86872538 to nil")
	}
	if err := b.ConsumeID(InputMessagePinnedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagePinned#86872538: %w", err)
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessagePinned) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessagePinned.
var (
	_ bin.Encoder = &InputMessagePinned{}
	_ bin.Decoder = &InputMessagePinned{}

	_ InputMessageClass = &InputMessagePinned{}
)

// InputMessageCallbackQuery represents TL type `inputMessageCallbackQuery#acfa1a7e`.
type InputMessageCallbackQuery struct {
	// ID field of InputMessageCallbackQuery.
	ID int
	// QueryID field of InputMessageCallbackQuery.
	QueryID int64
}

// InputMessageCallbackQueryTypeID is TL type id of InputMessageCallbackQuery.
const InputMessageCallbackQueryTypeID = 0xacfa1a7e

// Encode implements bin.Encoder.
func (i *InputMessageCallbackQuery) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageCallbackQuery#acfa1a7e as nil")
	}
	b.PutID(InputMessageCallbackQueryTypeID)
	b.PutInt(i.ID)
	b.PutLong(i.QueryID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageCallbackQuery) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageCallbackQuery#acfa1a7e to nil")
	}
	if err := b.ConsumeID(InputMessageCallbackQueryTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: field query_id: %w", err)
		}
		i.QueryID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageCallbackQuery) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageCallbackQuery.
var (
	_ bin.Encoder = &InputMessageCallbackQuery{}
	_ bin.Decoder = &InputMessageCallbackQuery{}

	_ InputMessageClass = &InputMessageCallbackQuery{}
)

// InputMessageClass represents InputMessage generic type.
//
// Example:
//  g, err := DecodeInputMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputMessageID: // inputMessageID#a676a322
//  case *InputMessageReplyTo: // inputMessageReplyTo#bad88395
//  case *InputMessagePinned: // inputMessagePinned#86872538
//  case *InputMessageCallbackQuery: // inputMessageCallbackQuery#acfa1a7e
//  default: panic(v)
//  }
type InputMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputMessageClass
}

// DecodeInputMessage implements binary de-serialization for InputMessageClass.
func DecodeInputMessage(buf *bin.Buffer) (InputMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputMessageIDTypeID:
		// Decoding inputMessageID#a676a322.
		v := InputMessageID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessageReplyToTypeID:
		// Decoding inputMessageReplyTo#bad88395.
		v := InputMessageReplyTo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessagePinnedTypeID:
		// Decoding inputMessagePinned#86872538.
		v := InputMessagePinned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessageCallbackQueryTypeID:
		// Decoding inputMessageCallbackQuery#acfa1a7e.
		v := InputMessageCallbackQuery{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputMessage boxes the InputMessageClass providing a helper.
type InputMessageBox struct {
	InputMessage InputMessageClass
}

// Decode implements bin.Decoder for InputMessageBox.
func (b *InputMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputMessageBox to nil")
	}
	v, err := DecodeInputMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputMessage = v
	return nil
}

// Encode implements bin.Encode for InputMessageBox.
func (b *InputMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputMessage == nil {
		return fmt.Errorf("unable to encode InputMessageClass as nil")
	}
	return b.InputMessage.Encode(buf)
}
