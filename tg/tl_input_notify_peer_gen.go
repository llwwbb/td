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

// InputNotifyPeer represents TL type `inputNotifyPeer#b8bc5b0c`.
type InputNotifyPeer struct {
	// Peer field of InputNotifyPeer.
	Peer InputPeerClass
}

// InputNotifyPeerTypeID is TL type id of InputNotifyPeer.
const InputNotifyPeerTypeID = 0xb8bc5b0c

// Encode implements bin.Encoder.
func (i *InputNotifyPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyPeer#b8bc5b0c as nil")
	}
	b.PutID(InputNotifyPeerTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputNotifyPeer#b8bc5b0c: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputNotifyPeer#b8bc5b0c: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyPeer#b8bc5b0c to nil")
	}
	if err := b.ConsumeID(InputNotifyPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyPeer#b8bc5b0c: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputNotifyPeer#b8bc5b0c: field peer: %w", err)
		}
		i.Peer = value
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyPeer) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyPeer.
var (
	_ bin.Encoder = &InputNotifyPeer{}
	_ bin.Decoder = &InputNotifyPeer{}

	_ InputNotifyPeerClass = &InputNotifyPeer{}
)

// InputNotifyUsers represents TL type `inputNotifyUsers#193b4417`.
type InputNotifyUsers struct {
}

// InputNotifyUsersTypeID is TL type id of InputNotifyUsers.
const InputNotifyUsersTypeID = 0x193b4417

// Encode implements bin.Encoder.
func (i *InputNotifyUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyUsers#193b4417 as nil")
	}
	b.PutID(InputNotifyUsersTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyUsers#193b4417 to nil")
	}
	if err := b.ConsumeID(InputNotifyUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyUsers#193b4417: %w", err)
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyUsers) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyUsers.
var (
	_ bin.Encoder = &InputNotifyUsers{}
	_ bin.Decoder = &InputNotifyUsers{}

	_ InputNotifyPeerClass = &InputNotifyUsers{}
)

// InputNotifyChats represents TL type `inputNotifyChats#4a95e84e`.
type InputNotifyChats struct {
}

// InputNotifyChatsTypeID is TL type id of InputNotifyChats.
const InputNotifyChatsTypeID = 0x4a95e84e

// Encode implements bin.Encoder.
func (i *InputNotifyChats) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyChats#4a95e84e as nil")
	}
	b.PutID(InputNotifyChatsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyChats) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyChats#4a95e84e to nil")
	}
	if err := b.ConsumeID(InputNotifyChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyChats#4a95e84e: %w", err)
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyChats) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyChats.
var (
	_ bin.Encoder = &InputNotifyChats{}
	_ bin.Decoder = &InputNotifyChats{}

	_ InputNotifyPeerClass = &InputNotifyChats{}
)

// InputNotifyBroadcasts represents TL type `inputNotifyBroadcasts#b1db7c7e`.
type InputNotifyBroadcasts struct {
}

// InputNotifyBroadcastsTypeID is TL type id of InputNotifyBroadcasts.
const InputNotifyBroadcastsTypeID = 0xb1db7c7e

// Encode implements bin.Encoder.
func (i *InputNotifyBroadcasts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyBroadcasts#b1db7c7e as nil")
	}
	b.PutID(InputNotifyBroadcastsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyBroadcasts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyBroadcasts#b1db7c7e to nil")
	}
	if err := b.ConsumeID(InputNotifyBroadcastsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyBroadcasts#b1db7c7e: %w", err)
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyBroadcasts) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyBroadcasts.
var (
	_ bin.Encoder = &InputNotifyBroadcasts{}
	_ bin.Decoder = &InputNotifyBroadcasts{}

	_ InputNotifyPeerClass = &InputNotifyBroadcasts{}
)

// InputNotifyPeerClass represents InputNotifyPeer generic type.
//
// Example:
//  g, err := DecodeInputNotifyPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputNotifyPeer: // inputNotifyPeer#b8bc5b0c
//  case *InputNotifyUsers: // inputNotifyUsers#193b4417
//  case *InputNotifyChats: // inputNotifyChats#4a95e84e
//  case *InputNotifyBroadcasts: // inputNotifyBroadcasts#b1db7c7e
//  default: panic(v)
//  }
type InputNotifyPeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputNotifyPeerClass
}

// DecodeInputNotifyPeer implements binary de-serialization for InputNotifyPeerClass.
func DecodeInputNotifyPeer(buf *bin.Buffer) (InputNotifyPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputNotifyPeerTypeID:
		// Decoding inputNotifyPeer#b8bc5b0c.
		v := InputNotifyPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyUsersTypeID:
		// Decoding inputNotifyUsers#193b4417.
		v := InputNotifyUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyChatsTypeID:
		// Decoding inputNotifyChats#4a95e84e.
		v := InputNotifyChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyBroadcastsTypeID:
		// Decoding inputNotifyBroadcasts#b1db7c7e.
		v := InputNotifyBroadcasts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputNotifyPeer boxes the InputNotifyPeerClass providing a helper.
type InputNotifyPeerBox struct {
	InputNotifyPeer InputNotifyPeerClass
}

// Decode implements bin.Decoder for InputNotifyPeerBox.
func (b *InputNotifyPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputNotifyPeerBox to nil")
	}
	v, err := DecodeInputNotifyPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputNotifyPeer = v
	return nil
}

// Encode implements bin.Encode for InputNotifyPeerBox.
func (b *InputNotifyPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputNotifyPeer == nil {
		return fmt.Errorf("unable to encode InputNotifyPeerClass as nil")
	}
	return b.InputNotifyPeer.Encode(buf)
}
