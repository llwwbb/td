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

// PageListOrderedItemText represents TL type `pageListOrderedItemText#5e068047`.
type PageListOrderedItemText struct {
	// Num field of PageListOrderedItemText.
	Num string
	// Text field of PageListOrderedItemText.
	Text RichTextClass
}

// PageListOrderedItemTextTypeID is TL type id of PageListOrderedItemText.
const PageListOrderedItemTextTypeID = 0x5e068047

// Encode implements bin.Encoder.
func (p *PageListOrderedItemText) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListOrderedItemText#5e068047 as nil")
	}
	b.PutID(PageListOrderedItemTextTypeID)
	b.PutString(p.Num)
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageListOrderedItemText#5e068047: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageListOrderedItemText#5e068047: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListOrderedItemText) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListOrderedItemText#5e068047 to nil")
	}
	if err := b.ConsumeID(PageListOrderedItemTextTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: field num: %w", err)
		}
		p.Num = value
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: field text: %w", err)
		}
		p.Text = value
	}
	return nil
}

// construct implements constructor of PageListOrderedItemClass.
func (p PageListOrderedItemText) construct() PageListOrderedItemClass { return &p }

// Ensuring interfaces in compile-time for PageListOrderedItemText.
var (
	_ bin.Encoder = &PageListOrderedItemText{}
	_ bin.Decoder = &PageListOrderedItemText{}

	_ PageListOrderedItemClass = &PageListOrderedItemText{}
)

// PageListOrderedItemBlocks represents TL type `pageListOrderedItemBlocks#98dd8936`.
type PageListOrderedItemBlocks struct {
	// Num field of PageListOrderedItemBlocks.
	Num string
	// Blocks field of PageListOrderedItemBlocks.
	Blocks []PageBlockClass
}

// PageListOrderedItemBlocksTypeID is TL type id of PageListOrderedItemBlocks.
const PageListOrderedItemBlocksTypeID = 0x98dd8936

// Encode implements bin.Encoder.
func (p *PageListOrderedItemBlocks) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListOrderedItemBlocks#98dd8936 as nil")
	}
	b.PutID(PageListOrderedItemBlocksTypeID)
	b.PutString(p.Num)
	b.PutVectorHeader(len(p.Blocks))
	for idx, v := range p.Blocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageListOrderedItemBlocks#98dd8936: field blocks element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageListOrderedItemBlocks#98dd8936: field blocks element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListOrderedItemBlocks) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListOrderedItemBlocks#98dd8936 to nil")
	}
	if err := b.ConsumeID(PageListOrderedItemBlocksTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field num: %w", err)
		}
		p.Num = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field blocks: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field blocks: %w", err)
			}
			p.Blocks = append(p.Blocks, value)
		}
	}
	return nil
}

// construct implements constructor of PageListOrderedItemClass.
func (p PageListOrderedItemBlocks) construct() PageListOrderedItemClass { return &p }

// Ensuring interfaces in compile-time for PageListOrderedItemBlocks.
var (
	_ bin.Encoder = &PageListOrderedItemBlocks{}
	_ bin.Decoder = &PageListOrderedItemBlocks{}

	_ PageListOrderedItemClass = &PageListOrderedItemBlocks{}
)

// PageListOrderedItemClass represents PageListOrderedItem generic type.
//
// Example:
//  g, err := DecodePageListOrderedItem(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PageListOrderedItemText: // pageListOrderedItemText#5e068047
//  case *PageListOrderedItemBlocks: // pageListOrderedItemBlocks#98dd8936
//  default: panic(v)
//  }
type PageListOrderedItemClass interface {
	bin.Encoder
	bin.Decoder
	construct() PageListOrderedItemClass
}

// DecodePageListOrderedItem implements binary de-serialization for PageListOrderedItemClass.
func DecodePageListOrderedItem(buf *bin.Buffer) (PageListOrderedItemClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PageListOrderedItemTextTypeID:
		// Decoding pageListOrderedItemText#5e068047.
		v := PageListOrderedItemText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", err)
		}
		return &v, nil
	case PageListOrderedItemBlocksTypeID:
		// Decoding pageListOrderedItemBlocks#98dd8936.
		v := PageListOrderedItemBlocks{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", bin.NewUnexpectedID(id))
	}
}

// PageListOrderedItem boxes the PageListOrderedItemClass providing a helper.
type PageListOrderedItemBox struct {
	PageListOrderedItem PageListOrderedItemClass
}

// Decode implements bin.Decoder for PageListOrderedItemBox.
func (b *PageListOrderedItemBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageListOrderedItemBox to nil")
	}
	v, err := DecodePageListOrderedItem(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageListOrderedItem = v
	return nil
}

// Encode implements bin.Encode for PageListOrderedItemBox.
func (b *PageListOrderedItemBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PageListOrderedItem == nil {
		return fmt.Errorf("unable to encode PageListOrderedItemClass as nil")
	}
	return b.PageListOrderedItem.Encode(buf)
}
