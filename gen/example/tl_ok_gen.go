// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}

// can be returned by functions as result.
type Ok struct {
}

// OkTypeID is TL type id of Ok.
const OkTypeID = 0xd4edbe69

// Encode implements bin.Encoder.
func (o *Ok) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode ok#d4edbe69 as nil")
	}
	b.PutID(OkTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (o *Ok) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode ok#d4edbe69 to nil")
	}
	if err := b.ConsumeID(OkTypeID); err != nil {
		return fmt.Errorf("unable to decode ok#d4edbe69: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Ok.
var (
	_ bin.Encoder = &Ok{}
	_ bin.Decoder = &Ok{}
)