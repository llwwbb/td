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

// LangpackGetLanguagesRequest represents TL type `langpack.getLanguages#42c6978f`.
// Get information about all languages in a localization pack
//
// See https://core.telegram.org/constructor/langpack.getLanguages for reference.
type LangpackGetLanguagesRequest struct {
	// Language pack
	LangPack string
}

// LangpackGetLanguagesRequestTypeID is TL type id of LangpackGetLanguagesRequest.
const LangpackGetLanguagesRequestTypeID = 0x42c6978f

// Encode implements bin.Encoder.
func (g *LangpackGetLanguagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLanguages#42c6978f as nil")
	}
	b.PutID(LangpackGetLanguagesRequestTypeID)
	b.PutString(g.LangPack)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetLanguagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLanguages#42c6978f to nil")
	}
	if err := b.ConsumeID(LangpackGetLanguagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getLanguages#42c6978f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLanguages#42c6978f: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetLanguagesRequest.
var (
	_ bin.Encoder = &LangpackGetLanguagesRequest{}
	_ bin.Decoder = &LangpackGetLanguagesRequest{}
)
