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

// InputBotInlineResult represents TL type `inputBotInlineResult#88bf9319`.
type InputBotInlineResult struct {
	// Flags field of InputBotInlineResult.
	Flags bin.Fields
	// ID field of InputBotInlineResult.
	ID string
	// Type field of InputBotInlineResult.
	Type string
	// Title field of InputBotInlineResult.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description field of InputBotInlineResult.
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// URL field of InputBotInlineResult.
	//
	// Use SetURL and GetURL helpers.
	URL string
	// Thumb field of InputBotInlineResult.
	//
	// Use SetThumb and GetThumb helpers.
	Thumb InputWebDocument
	// Content field of InputBotInlineResult.
	//
	// Use SetContent and GetContent helpers.
	Content InputWebDocument
	// SendMessage field of InputBotInlineResult.
	SendMessage InputBotInlineMessageClass
}

// InputBotInlineResultTypeID is TL type id of InputBotInlineResult.
const InputBotInlineResultTypeID = 0x88bf9319

// Encode implements bin.Encoder.
func (i *InputBotInlineResult) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineResult#88bf9319 as nil")
	}
	b.PutID(InputBotInlineResultTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResult#88bf9319: field flags: %w", err)
	}
	b.PutString(i.ID)
	b.PutString(i.Type)
	if i.Flags.Has(1) {
		b.PutString(i.Title)
	}
	if i.Flags.Has(2) {
		b.PutString(i.Description)
	}
	if i.Flags.Has(3) {
		b.PutString(i.URL)
	}
	if i.Flags.Has(4) {
		if err := i.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineResult#88bf9319: field thumb: %w", err)
		}
	}
	if i.Flags.Has(5) {
		if err := i.Content.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputBotInlineResult#88bf9319: field content: %w", err)
		}
	}
	if i.SendMessage == nil {
		return fmt.Errorf("unable to encode inputBotInlineResult#88bf9319: field send_message is nil")
	}
	if err := i.SendMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResult#88bf9319: field send_message: %w", err)
	}
	return nil
}

// SetTitle sets value of Title conditional field.
func (i *InputBotInlineResult) SetTitle(value string) {
	i.Flags.Set(1)
	i.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResult) GetTitle() (value string, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Title, true
}

// SetDescription sets value of Description conditional field.
func (i *InputBotInlineResult) SetDescription(value string) {
	i.Flags.Set(2)
	i.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResult) GetDescription() (value string, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.Description, true
}

// SetURL sets value of URL conditional field.
func (i *InputBotInlineResult) SetURL(value string) {
	i.Flags.Set(3)
	i.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResult) GetURL() (value string, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.URL, true
}

// SetThumb sets value of Thumb conditional field.
func (i *InputBotInlineResult) SetThumb(value InputWebDocument) {
	i.Flags.Set(4)
	i.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResult) GetThumb() (value InputWebDocument, ok bool) {
	if !i.Flags.Has(4) {
		return value, false
	}
	return i.Thumb, true
}

// SetContent sets value of Content conditional field.
func (i *InputBotInlineResult) SetContent(value InputWebDocument) {
	i.Flags.Set(5)
	i.Content = value
}

// GetContent returns value of Content conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResult) GetContent() (value InputWebDocument, ok bool) {
	if !i.Flags.Has(5) {
		return value, false
	}
	return i.Content, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineResult) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineResult#88bf9319 to nil")
	}
	if err := b.ConsumeID(InputBotInlineResultTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field type: %w", err)
		}
		i.Type = value
	}
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field title: %w", err)
		}
		i.Title = value
	}
	if i.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field description: %w", err)
		}
		i.Description = value
	}
	if i.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field url: %w", err)
		}
		i.URL = value
	}
	if i.Flags.Has(4) {
		if err := i.Thumb.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field thumb: %w", err)
		}
	}
	if i.Flags.Has(5) {
		if err := i.Content.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field content: %w", err)
		}
	}
	{
		value, err := DecodeInputBotInlineMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResult#88bf9319: field send_message: %w", err)
		}
		i.SendMessage = value
	}
	return nil
}

// construct implements constructor of InputBotInlineResultClass.
func (i InputBotInlineResult) construct() InputBotInlineResultClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineResult.
var (
	_ bin.Encoder = &InputBotInlineResult{}
	_ bin.Decoder = &InputBotInlineResult{}

	_ InputBotInlineResultClass = &InputBotInlineResult{}
)

// InputBotInlineResultPhoto represents TL type `inputBotInlineResultPhoto#a8d864a7`.
type InputBotInlineResultPhoto struct {
	// ID field of InputBotInlineResultPhoto.
	ID string
	// Type field of InputBotInlineResultPhoto.
	Type string
	// Photo field of InputBotInlineResultPhoto.
	Photo InputPhotoClass
	// SendMessage field of InputBotInlineResultPhoto.
	SendMessage InputBotInlineMessageClass
}

// InputBotInlineResultPhotoTypeID is TL type id of InputBotInlineResultPhoto.
const InputBotInlineResultPhotoTypeID = 0xa8d864a7

// Encode implements bin.Encoder.
func (i *InputBotInlineResultPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineResultPhoto#a8d864a7 as nil")
	}
	b.PutID(InputBotInlineResultPhotoTypeID)
	b.PutString(i.ID)
	b.PutString(i.Type)
	if i.Photo == nil {
		return fmt.Errorf("unable to encode inputBotInlineResultPhoto#a8d864a7: field photo is nil")
	}
	if err := i.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultPhoto#a8d864a7: field photo: %w", err)
	}
	if i.SendMessage == nil {
		return fmt.Errorf("unable to encode inputBotInlineResultPhoto#a8d864a7: field send_message is nil")
	}
	if err := i.SendMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultPhoto#a8d864a7: field send_message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBotInlineResultPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineResultPhoto#a8d864a7 to nil")
	}
	if err := b.ConsumeID(InputBotInlineResultPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineResultPhoto#a8d864a7: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultPhoto#a8d864a7: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultPhoto#a8d864a7: field type: %w", err)
		}
		i.Type = value
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultPhoto#a8d864a7: field photo: %w", err)
		}
		i.Photo = value
	}
	{
		value, err := DecodeInputBotInlineMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultPhoto#a8d864a7: field send_message: %w", err)
		}
		i.SendMessage = value
	}
	return nil
}

// construct implements constructor of InputBotInlineResultClass.
func (i InputBotInlineResultPhoto) construct() InputBotInlineResultClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineResultPhoto.
var (
	_ bin.Encoder = &InputBotInlineResultPhoto{}
	_ bin.Decoder = &InputBotInlineResultPhoto{}

	_ InputBotInlineResultClass = &InputBotInlineResultPhoto{}
)

// InputBotInlineResultDocument represents TL type `inputBotInlineResultDocument#fff8fdc4`.
type InputBotInlineResultDocument struct {
	// Flags field of InputBotInlineResultDocument.
	Flags bin.Fields
	// ID field of InputBotInlineResultDocument.
	ID string
	// Type field of InputBotInlineResultDocument.
	Type string
	// Title field of InputBotInlineResultDocument.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description field of InputBotInlineResultDocument.
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// Document field of InputBotInlineResultDocument.
	Document InputDocumentClass
	// SendMessage field of InputBotInlineResultDocument.
	SendMessage InputBotInlineMessageClass
}

// InputBotInlineResultDocumentTypeID is TL type id of InputBotInlineResultDocument.
const InputBotInlineResultDocumentTypeID = 0xfff8fdc4

// Encode implements bin.Encoder.
func (i *InputBotInlineResultDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineResultDocument#fff8fdc4 as nil")
	}
	b.PutID(InputBotInlineResultDocumentTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultDocument#fff8fdc4: field flags: %w", err)
	}
	b.PutString(i.ID)
	b.PutString(i.Type)
	if i.Flags.Has(1) {
		b.PutString(i.Title)
	}
	if i.Flags.Has(2) {
		b.PutString(i.Description)
	}
	if i.Document == nil {
		return fmt.Errorf("unable to encode inputBotInlineResultDocument#fff8fdc4: field document is nil")
	}
	if err := i.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultDocument#fff8fdc4: field document: %w", err)
	}
	if i.SendMessage == nil {
		return fmt.Errorf("unable to encode inputBotInlineResultDocument#fff8fdc4: field send_message is nil")
	}
	if err := i.SendMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultDocument#fff8fdc4: field send_message: %w", err)
	}
	return nil
}

// SetTitle sets value of Title conditional field.
func (i *InputBotInlineResultDocument) SetTitle(value string) {
	i.Flags.Set(1)
	i.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResultDocument) GetTitle() (value string, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Title, true
}

// SetDescription sets value of Description conditional field.
func (i *InputBotInlineResultDocument) SetDescription(value string) {
	i.Flags.Set(2)
	i.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (i *InputBotInlineResultDocument) GetDescription() (value string, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.Description, true
}

// Decode implements bin.Decoder.
func (i *InputBotInlineResultDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineResultDocument#fff8fdc4 to nil")
	}
	if err := b.ConsumeID(InputBotInlineResultDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field type: %w", err)
		}
		i.Type = value
	}
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field title: %w", err)
		}
		i.Title = value
	}
	if i.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field description: %w", err)
		}
		i.Description = value
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field document: %w", err)
		}
		i.Document = value
	}
	{
		value, err := DecodeInputBotInlineMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultDocument#fff8fdc4: field send_message: %w", err)
		}
		i.SendMessage = value
	}
	return nil
}

// construct implements constructor of InputBotInlineResultClass.
func (i InputBotInlineResultDocument) construct() InputBotInlineResultClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineResultDocument.
var (
	_ bin.Encoder = &InputBotInlineResultDocument{}
	_ bin.Decoder = &InputBotInlineResultDocument{}

	_ InputBotInlineResultClass = &InputBotInlineResultDocument{}
)

// InputBotInlineResultGame represents TL type `inputBotInlineResultGame#4fa417f2`.
type InputBotInlineResultGame struct {
	// ID field of InputBotInlineResultGame.
	ID string
	// ShortName field of InputBotInlineResultGame.
	ShortName string
	// SendMessage field of InputBotInlineResultGame.
	SendMessage InputBotInlineMessageClass
}

// InputBotInlineResultGameTypeID is TL type id of InputBotInlineResultGame.
const InputBotInlineResultGameTypeID = 0x4fa417f2

// Encode implements bin.Encoder.
func (i *InputBotInlineResultGame) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotInlineResultGame#4fa417f2 as nil")
	}
	b.PutID(InputBotInlineResultGameTypeID)
	b.PutString(i.ID)
	b.PutString(i.ShortName)
	if i.SendMessage == nil {
		return fmt.Errorf("unable to encode inputBotInlineResultGame#4fa417f2: field send_message is nil")
	}
	if err := i.SendMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotInlineResultGame#4fa417f2: field send_message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBotInlineResultGame) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotInlineResultGame#4fa417f2 to nil")
	}
	if err := b.ConsumeID(InputBotInlineResultGameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotInlineResultGame#4fa417f2: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultGame#4fa417f2: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultGame#4fa417f2: field short_name: %w", err)
		}
		i.ShortName = value
	}
	{
		value, err := DecodeInputBotInlineMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotInlineResultGame#4fa417f2: field send_message: %w", err)
		}
		i.SendMessage = value
	}
	return nil
}

// construct implements constructor of InputBotInlineResultClass.
func (i InputBotInlineResultGame) construct() InputBotInlineResultClass { return &i }

// Ensuring interfaces in compile-time for InputBotInlineResultGame.
var (
	_ bin.Encoder = &InputBotInlineResultGame{}
	_ bin.Decoder = &InputBotInlineResultGame{}

	_ InputBotInlineResultClass = &InputBotInlineResultGame{}
)

// InputBotInlineResultClass represents InputBotInlineResult generic type.
//
// Example:
//  g, err := DecodeInputBotInlineResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputBotInlineResult: // inputBotInlineResult#88bf9319
//  case *InputBotInlineResultPhoto: // inputBotInlineResultPhoto#a8d864a7
//  case *InputBotInlineResultDocument: // inputBotInlineResultDocument#fff8fdc4
//  case *InputBotInlineResultGame: // inputBotInlineResultGame#4fa417f2
//  default: panic(v)
//  }
type InputBotInlineResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputBotInlineResultClass
}

// DecodeInputBotInlineResult implements binary de-serialization for InputBotInlineResultClass.
func DecodeInputBotInlineResult(buf *bin.Buffer) (InputBotInlineResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputBotInlineResultTypeID:
		// Decoding inputBotInlineResult#88bf9319.
		v := InputBotInlineResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineResultClass: %w", err)
		}
		return &v, nil
	case InputBotInlineResultPhotoTypeID:
		// Decoding inputBotInlineResultPhoto#a8d864a7.
		v := InputBotInlineResultPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineResultClass: %w", err)
		}
		return &v, nil
	case InputBotInlineResultDocumentTypeID:
		// Decoding inputBotInlineResultDocument#fff8fdc4.
		v := InputBotInlineResultDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineResultClass: %w", err)
		}
		return &v, nil
	case InputBotInlineResultGameTypeID:
		// Decoding inputBotInlineResultGame#4fa417f2.
		v := InputBotInlineResultGame{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotInlineResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputBotInlineResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputBotInlineResult boxes the InputBotInlineResultClass providing a helper.
type InputBotInlineResultBox struct {
	InputBotInlineResult InputBotInlineResultClass
}

// Decode implements bin.Decoder for InputBotInlineResultBox.
func (b *InputBotInlineResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputBotInlineResultBox to nil")
	}
	v, err := DecodeInputBotInlineResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputBotInlineResult = v
	return nil
}

// Encode implements bin.Encode for InputBotInlineResultBox.
func (b *InputBotInlineResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputBotInlineResult == nil {
		return fmt.Errorf("unable to encode InputBotInlineResultClass as nil")
	}
	return b.InputBotInlineResult.Encode(buf)
}
