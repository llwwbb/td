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

// ContactsTopPeersNotModified represents TL type `contacts.topPeersNotModified#de266ef5`.
// Top peer info hasn't changed
//
// See https://core.telegram.org/constructor/contacts.topPeersNotModified for reference.
type ContactsTopPeersNotModified struct {
}

// ContactsTopPeersNotModifiedTypeID is TL type id of ContactsTopPeersNotModified.
const ContactsTopPeersNotModifiedTypeID = 0xde266ef5

// Encode implements bin.Encoder.
func (t *ContactsTopPeersNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersNotModified#de266ef5 as nil")
	}
	b.PutID(ContactsTopPeersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersNotModified#de266ef5 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersNotModified#de266ef5: %w", err)
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersNotModified) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersNotModified.
var (
	_ bin.Encoder = &ContactsTopPeersNotModified{}
	_ bin.Decoder = &ContactsTopPeersNotModified{}

	_ ContactsTopPeersClass = &ContactsTopPeersNotModified{}
)

// ContactsTopPeers represents TL type `contacts.topPeers#70b772a8`.
// Top peers
//
// See https://core.telegram.org/constructor/contacts.topPeers for reference.
type ContactsTopPeers struct {
	// Top peers by top peer category
	Categories []TopPeerCategoryPeers
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
}

// ContactsTopPeersTypeID is TL type id of ContactsTopPeers.
const ContactsTopPeersTypeID = 0x70b772a8

// Encode implements bin.Encoder.
func (t *ContactsTopPeers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeers#70b772a8 as nil")
	}
	b.PutID(ContactsTopPeersTypeID)
	b.PutVectorHeader(len(t.Categories))
	for idx, v := range t.Categories {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field categories element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Chats))
	for idx, v := range t.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(t.Users))
	for idx, v := range t.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.topPeers#70b772a8: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeers#70b772a8 to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TopPeerCategoryPeers
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field categories: %w", err)
			}
			t.Categories = append(t.Categories, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field chats: %w", err)
			}
			t.Chats = append(t.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.topPeers#70b772a8: field users: %w", err)
			}
			t.Users = append(t.Users, value)
		}
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeers) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeers.
var (
	_ bin.Encoder = &ContactsTopPeers{}
	_ bin.Decoder = &ContactsTopPeers{}

	_ ContactsTopPeersClass = &ContactsTopPeers{}
)

// ContactsTopPeersDisabled represents TL type `contacts.topPeersDisabled#b52c939d`.
// Top peers disabled
//
// See https://core.telegram.org/constructor/contacts.topPeersDisabled for reference.
type ContactsTopPeersDisabled struct {
}

// ContactsTopPeersDisabledTypeID is TL type id of ContactsTopPeersDisabled.
const ContactsTopPeersDisabledTypeID = 0xb52c939d

// Encode implements bin.Encoder.
func (t *ContactsTopPeersDisabled) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.topPeersDisabled#b52c939d as nil")
	}
	b.PutID(ContactsTopPeersDisabledTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsTopPeersDisabled) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.topPeersDisabled#b52c939d to nil")
	}
	if err := b.ConsumeID(ContactsTopPeersDisabledTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.topPeersDisabled#b52c939d: %w", err)
	}
	return nil
}

// construct implements constructor of ContactsTopPeersClass.
func (t ContactsTopPeersDisabled) construct() ContactsTopPeersClass { return &t }

// Ensuring interfaces in compile-time for ContactsTopPeersDisabled.
var (
	_ bin.Encoder = &ContactsTopPeersDisabled{}
	_ bin.Decoder = &ContactsTopPeersDisabled{}

	_ ContactsTopPeersClass = &ContactsTopPeersDisabled{}
)

// ContactsTopPeersClass represents contacts.TopPeers generic type.
//
// See https://core.telegram.org/type/contacts.TopPeers for reference.
//
// Example:
//  g, err := DecodeContactsTopPeers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ContactsTopPeersNotModified: // contacts.topPeersNotModified#de266ef5
//  case *ContactsTopPeers: // contacts.topPeers#70b772a8
//  case *ContactsTopPeersDisabled: // contacts.topPeersDisabled#b52c939d
//  default: panic(v)
//  }
type ContactsTopPeersClass interface {
	bin.Encoder
	bin.Decoder
	construct() ContactsTopPeersClass
}

// DecodeContactsTopPeers implements binary de-serialization for ContactsTopPeersClass.
func DecodeContactsTopPeers(buf *bin.Buffer) (ContactsTopPeersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsTopPeersNotModifiedTypeID:
		// Decoding contacts.topPeersNotModified#de266ef5.
		v := ContactsTopPeersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersTypeID:
		// Decoding contacts.topPeers#70b772a8.
		v := ContactsTopPeers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	case ContactsTopPeersDisabledTypeID:
		// Decoding contacts.topPeersDisabled#b52c939d.
		v := ContactsTopPeersDisabled{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsTopPeersClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsTopPeers boxes the ContactsTopPeersClass providing a helper.
type ContactsTopPeersBox struct {
	TopPeers ContactsTopPeersClass
}

// Decode implements bin.Decoder for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsTopPeersBox to nil")
	}
	v, err := DecodeContactsTopPeers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopPeers = v
	return nil
}

// Encode implements bin.Encode for ContactsTopPeersBox.
func (b *ContactsTopPeersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopPeers == nil {
		return fmt.Errorf("unable to encode ContactsTopPeersClass as nil")
	}
	return b.TopPeers.Encode(buf)
}
