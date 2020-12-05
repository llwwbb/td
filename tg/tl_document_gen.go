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

// DocumentEmpty represents TL type `documentEmpty#36f8c871`.
// Empty constructor, document doesn't exist.
//
// See https://core.telegram.org/constructor/documentEmpty for reference.
type DocumentEmpty struct {
	// Document ID or 0
	ID int64
}

// DocumentEmptyTypeID is TL type id of DocumentEmpty.
const DocumentEmptyTypeID = 0x36f8c871

// Encode implements bin.Encoder.
func (d *DocumentEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentEmpty#36f8c871 as nil")
	}
	b.PutID(DocumentEmptyTypeID)
	b.PutLong(d.ID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentEmpty#36f8c871 to nil")
	}
	if err := b.ConsumeID(DocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode documentEmpty#36f8c871: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode documentEmpty#36f8c871: field id: %w", err)
		}
		d.ID = value
	}
	return nil
}

// construct implements constructor of DocumentClass.
func (d DocumentEmpty) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for DocumentEmpty.
var (
	_ bin.Encoder = &DocumentEmpty{}
	_ bin.Decoder = &DocumentEmpty{}

	_ DocumentClass = &DocumentEmpty{}
)

// Document represents TL type `document#1e87342b`.
// Document
//
// See https://core.telegram.org/constructor/document for reference.
type Document struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Document ID
	ID int64
	// Check sum, dependant on document ID
	AccessHash int64
	// File reference
	FileReference []byte
	// Creation date
	Date int
	// MIME type
	MimeType string
	// Size
	Size int
	// Thumbnails
	//
	// Use SetThumbs and GetThumbs helpers.
	Thumbs []PhotoSizeClass
	// Video thumbnails
	//
	// Use SetVideoThumbs and GetVideoThumbs helpers.
	VideoThumbs []VideoSize
	// DC ID
	DCID int
	// Attributes
	Attributes []DocumentAttributeClass
}

// DocumentTypeID is TL type id of Document.
const DocumentTypeID = 0x1e87342b

// Encode implements bin.Encoder.
func (d *Document) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#1e87342b as nil")
	}
	b.PutID(DocumentTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#1e87342b: field flags: %w", err)
	}
	b.PutLong(d.ID)
	b.PutLong(d.AccessHash)
	b.PutBytes(d.FileReference)
	b.PutInt(d.Date)
	b.PutString(d.MimeType)
	b.PutInt(d.Size)
	if d.Flags.Has(0) {
		b.PutVectorHeader(len(d.Thumbs))
		for idx, v := range d.Thumbs {
			if v == nil {
				return fmt.Errorf("unable to encode document#1e87342b: field thumbs element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#1e87342b: field thumbs element with index %d: %w", idx, err)
			}
		}
	}
	if d.Flags.Has(1) {
		b.PutVectorHeader(len(d.VideoThumbs))
		for idx, v := range d.VideoThumbs {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode document#1e87342b: field video_thumbs element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(d.DCID)
	b.PutVectorHeader(len(d.Attributes))
	for idx, v := range d.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode document#1e87342b: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode document#1e87342b: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetThumbs sets value of Thumbs conditional field.
func (d *Document) SetThumbs(value []PhotoSizeClass) {
	d.Flags.Set(0)
	d.Thumbs = value
}

// GetThumbs returns value of Thumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetThumbs() (value []PhotoSizeClass, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Thumbs, true
}

// SetVideoThumbs sets value of VideoThumbs conditional field.
func (d *Document) SetVideoThumbs(value []VideoSize) {
	d.Flags.Set(1)
	d.VideoThumbs = value
}

// GetVideoThumbs returns value of VideoThumbs conditional field and
// boolean which is true if field was set.
func (d *Document) GetVideoThumbs() (value []VideoSize, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.VideoThumbs, true
}

// Decode implements bin.Decoder.
func (d *Document) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#1e87342b to nil")
	}
	if err := b.ConsumeID(DocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode document#1e87342b: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field access_hash: %w", err)
		}
		d.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field file_reference: %w", err)
		}
		d.FileReference = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field mime_type: %w", err)
		}
		d.MimeType = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field size: %w", err)
		}
		d.Size = value
	}
	if d.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field thumbs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field thumbs: %w", err)
			}
			d.Thumbs = append(d.Thumbs, value)
		}
	}
	if d.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field video_thumbs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value VideoSize
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field video_thumbs: %w", err)
			}
			d.VideoThumbs = append(d.VideoThumbs, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field dc_id: %w", err)
		}
		d.DCID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode document#1e87342b: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode document#1e87342b: field attributes: %w", err)
			}
			d.Attributes = append(d.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of DocumentClass.
func (d Document) construct() DocumentClass { return &d }

// Ensuring interfaces in compile-time for Document.
var (
	_ bin.Encoder = &Document{}
	_ bin.Decoder = &Document{}

	_ DocumentClass = &Document{}
)

// DocumentClass represents Document generic type.
//
// See https://core.telegram.org/type/Document for reference.
//
// Example:
//  g, err := DecodeDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *DocumentEmpty: // documentEmpty#36f8c871
//  case *Document: // document#1e87342b
//  default: panic(v)
//  }
type DocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() DocumentClass
}

// DecodeDocument implements binary de-serialization for DocumentClass.
func DecodeDocument(buf *bin.Buffer) (DocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DocumentEmptyTypeID:
		// Decoding documentEmpty#36f8c871.
		v := DocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	case DocumentTypeID:
		// Decoding document#1e87342b.
		v := Document{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// Document boxes the DocumentClass providing a helper.
type DocumentBox struct {
	Document DocumentClass
}

// Decode implements bin.Decoder for DocumentBox.
func (b *DocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DocumentBox to nil")
	}
	v, err := DecodeDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Document = v
	return nil
}

// Encode implements bin.Encode for DocumentBox.
func (b *DocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Document == nil {
		return fmt.Errorf("unable to encode DocumentClass as nil")
	}
	return b.Document.Encode(buf)
}
