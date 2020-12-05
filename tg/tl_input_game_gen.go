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

// InputGameID represents TL type `inputGameID#32c3e77`.
// Indicates an already sent game
//
// See https://core.telegram.org/constructor/inputGameID for reference.
type InputGameID struct {
	// game ID from Game constructor
	ID int64
	// access hash from Game constructor
	AccessHash int64
}

// InputGameIDTypeID is TL type id of InputGameID.
const InputGameIDTypeID = 0x32c3e77

// Encode implements bin.Encoder.
func (i *InputGameID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameID#32c3e77 as nil")
	}
	b.PutID(InputGameIDTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGameID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameID#32c3e77 to nil")
	}
	if err := b.ConsumeID(InputGameIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGameID#32c3e77: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameID#32c3e77: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameID#32c3e77: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputGameClass.
func (i InputGameID) construct() InputGameClass { return &i }

// Ensuring interfaces in compile-time for InputGameID.
var (
	_ bin.Encoder = &InputGameID{}
	_ bin.Decoder = &InputGameID{}

	_ InputGameClass = &InputGameID{}
)

// InputGameShortName represents TL type `inputGameShortName#c331e80a`.
// Game by short name
//
// See https://core.telegram.org/constructor/inputGameShortName for reference.
type InputGameShortName struct {
	// The bot that provides the game
	BotID InputUserClass
	// The game's short name
	ShortName string
}

// InputGameShortNameTypeID is TL type id of InputGameShortName.
const InputGameShortNameTypeID = 0xc331e80a

// Encode implements bin.Encoder.
func (i *InputGameShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameShortName#c331e80a as nil")
	}
	b.PutID(InputGameShortNameTypeID)
	if i.BotID == nil {
		return fmt.Errorf("unable to encode inputGameShortName#c331e80a: field bot_id is nil")
	}
	if err := i.BotID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputGameShortName#c331e80a: field bot_id: %w", err)
	}
	b.PutString(i.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGameShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameShortName#c331e80a to nil")
	}
	if err := b.ConsumeID(InputGameShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGameShortName#c331e80a: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputGameShortName#c331e80a: field bot_id: %w", err)
		}
		i.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameShortName#c331e80a: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// construct implements constructor of InputGameClass.
func (i InputGameShortName) construct() InputGameClass { return &i }

// Ensuring interfaces in compile-time for InputGameShortName.
var (
	_ bin.Encoder = &InputGameShortName{}
	_ bin.Decoder = &InputGameShortName{}

	_ InputGameClass = &InputGameShortName{}
)

// InputGameClass represents InputGame generic type.
//
// See https://core.telegram.org/type/InputGame for reference.
//
// Example:
//  g, err := DecodeInputGame(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputGameID: // inputGameID#32c3e77
//  case *InputGameShortName: // inputGameShortName#c331e80a
//  default: panic(v)
//  }
type InputGameClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputGameClass
}

// DecodeInputGame implements binary de-serialization for InputGameClass.
func DecodeInputGame(buf *bin.Buffer) (InputGameClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputGameIDTypeID:
		// Decoding inputGameID#32c3e77.
		v := InputGameID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGameClass: %w", err)
		}
		return &v, nil
	case InputGameShortNameTypeID:
		// Decoding inputGameShortName#c331e80a.
		v := InputGameShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGameClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputGameClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputGame boxes the InputGameClass providing a helper.
type InputGameBox struct {
	InputGame InputGameClass
}

// Decode implements bin.Decoder for InputGameBox.
func (b *InputGameBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputGameBox to nil")
	}
	v, err := DecodeInputGame(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputGame = v
	return nil
}

// Encode implements bin.Encode for InputGameBox.
func (b *InputGameBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputGame == nil {
		return fmt.Errorf("unable to encode InputGameClass as nil")
	}
	return b.InputGame.Encode(buf)
}
