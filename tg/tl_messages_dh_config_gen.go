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

// MessagesDhConfigNotModified represents TL type `messages.dhConfigNotModified#c0e24635`.
type MessagesDhConfigNotModified struct {
	// Random field of MessagesDhConfigNotModified.
	Random []byte
}

// MessagesDhConfigNotModifiedTypeID is TL type id of MessagesDhConfigNotModified.
const MessagesDhConfigNotModifiedTypeID = 0xc0e24635

// Encode implements bin.Encoder.
func (d *MessagesDhConfigNotModified) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfigNotModified#c0e24635 as nil")
	}
	b.PutID(MessagesDhConfigNotModifiedTypeID)
	b.PutBytes(d.Random)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDhConfigNotModified) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfigNotModified#c0e24635 to nil")
	}
	if err := b.ConsumeID(MessagesDhConfigNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dhConfigNotModified#c0e24635: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfigNotModified#c0e24635: field random: %w", err)
		}
		d.Random = value
	}
	return nil
}

// construct implements constructor of MessagesDhConfigClass.
func (d MessagesDhConfigNotModified) construct() MessagesDhConfigClass { return &d }

// Ensuring interfaces in compile-time for MessagesDhConfigNotModified.
var (
	_ bin.Encoder = &MessagesDhConfigNotModified{}
	_ bin.Decoder = &MessagesDhConfigNotModified{}

	_ MessagesDhConfigClass = &MessagesDhConfigNotModified{}
)

// MessagesDhConfig represents TL type `messages.dhConfig#2c221edd`.
type MessagesDhConfig struct {
	// G field of MessagesDhConfig.
	G int
	// P field of MessagesDhConfig.
	P []byte
	// Version field of MessagesDhConfig.
	Version int
	// Random field of MessagesDhConfig.
	Random []byte
}

// MessagesDhConfigTypeID is TL type id of MessagesDhConfig.
const MessagesDhConfigTypeID = 0x2c221edd

// Encode implements bin.Encoder.
func (d *MessagesDhConfig) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfig#2c221edd as nil")
	}
	b.PutID(MessagesDhConfigTypeID)
	b.PutInt(d.G)
	b.PutBytes(d.P)
	b.PutInt(d.Version)
	b.PutBytes(d.Random)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDhConfig) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfig#2c221edd to nil")
	}
	if err := b.ConsumeID(MessagesDhConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field g: %w", err)
		}
		d.G = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field p: %w", err)
		}
		d.P = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field version: %w", err)
		}
		d.Version = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field random: %w", err)
		}
		d.Random = value
	}
	return nil
}

// construct implements constructor of MessagesDhConfigClass.
func (d MessagesDhConfig) construct() MessagesDhConfigClass { return &d }

// Ensuring interfaces in compile-time for MessagesDhConfig.
var (
	_ bin.Encoder = &MessagesDhConfig{}
	_ bin.Decoder = &MessagesDhConfig{}

	_ MessagesDhConfigClass = &MessagesDhConfig{}
)

// MessagesDhConfigClass represents messages.DhConfig generic type.
//
// Example:
//  g, err := DecodeMessagesDhConfig(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesDhConfigNotModified: // messages.dhConfigNotModified#c0e24635
//  case *MessagesDhConfig: // messages.dhConfig#2c221edd
//  default: panic(v)
//  }
type MessagesDhConfigClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesDhConfigClass
}

// DecodeMessagesDhConfig implements binary de-serialization for MessagesDhConfigClass.
func DecodeMessagesDhConfig(buf *bin.Buffer) (MessagesDhConfigClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesDhConfigNotModifiedTypeID:
		// Decoding messages.dhConfigNotModified#c0e24635.
		v := MessagesDhConfigNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", err)
		}
		return &v, nil
	case MessagesDhConfigTypeID:
		// Decoding messages.dhConfig#2c221edd.
		v := MessagesDhConfig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesDhConfig boxes the MessagesDhConfigClass providing a helper.
type MessagesDhConfigBox struct {
	MessagesDhConfig MessagesDhConfigClass
}

// Decode implements bin.Decoder for MessagesDhConfigBox.
func (b *MessagesDhConfigBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesDhConfigBox to nil")
	}
	v, err := DecodeMessagesDhConfig(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessagesDhConfig = v
	return nil
}

// Encode implements bin.Encode for MessagesDhConfigBox.
func (b *MessagesDhConfigBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessagesDhConfig == nil {
		return fmt.Errorf("unable to encode MessagesDhConfigClass as nil")
	}
	return b.MessagesDhConfig.Encode(buf)
}
