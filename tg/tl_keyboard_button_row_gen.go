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

// KeyboardButtonRow represents TL type `keyboardButtonRow#77608b83`.
// Inline keyboard row
//
// See https://core.telegram.org/constructor/keyboardButtonRow for reference.
type KeyboardButtonRow struct {
	// Bot or inline keyboard buttons
	Buttons []KeyboardButtonClass
}

// KeyboardButtonRowTypeID is TL type id of KeyboardButtonRow.
const KeyboardButtonRowTypeID = 0x77608b83

// Encode implements bin.Encoder.
func (k *KeyboardButtonRow) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonRow#77608b83 as nil")
	}
	b.PutID(KeyboardButtonRowTypeID)
	b.PutVectorHeader(len(k.Buttons))
	for idx, v := range k.Buttons {
		if v == nil {
			return fmt.Errorf("unable to encode keyboardButtonRow#77608b83: field buttons element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode keyboardButtonRow#77608b83: field buttons element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonRow) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonRow#77608b83 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonRowTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: field buttons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeKeyboardButton(b)
			if err != nil {
				return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: field buttons: %w", err)
			}
			k.Buttons = append(k.Buttons, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for KeyboardButtonRow.
var (
	_ bin.Encoder = &KeyboardButtonRow{}
	_ bin.Decoder = &KeyboardButtonRow{}
)
