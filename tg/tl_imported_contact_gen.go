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

// ImportedContact represents TL type `importedContact#d0028438`.
// Successfully imported contact.
//
// See https://core.telegram.org/constructor/importedContact for reference.
type ImportedContact struct {
	// User identifier
	UserID int
	// The contact's client identifier (passed to one of the InputContact constructors)
	ClientID int64
}

// ImportedContactTypeID is TL type id of ImportedContact.
const ImportedContactTypeID = 0xd0028438

// Encode implements bin.Encoder.
func (i *ImportedContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContact#d0028438 as nil")
	}
	b.PutID(ImportedContactTypeID)
	b.PutInt(i.UserID)
	b.PutLong(i.ClientID)
	return nil
}

// Decode implements bin.Decoder.
func (i *ImportedContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContact#d0028438 to nil")
	}
	if err := b.ConsumeID(ImportedContactTypeID); err != nil {
		return fmt.Errorf("unable to decode importedContact#d0028438: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#d0028438: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#d0028438: field client_id: %w", err)
		}
		i.ClientID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ImportedContact.
var (
	_ bin.Encoder = &ImportedContact{}
	_ bin.Decoder = &ImportedContact{}
)
