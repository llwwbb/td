// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// UserAuth represents TL type `user.auth#f4815592`.
type UserAuth struct {
	// Foo field of UserAuth.
	Foo string
}

// UserAuthTypeID is TL type id of UserAuth.
const UserAuthTypeID = 0xf4815592

// Encode implements bin.Encoder.
func (a *UserAuth) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.auth#f4815592 as nil")
	}
	b.PutID(UserAuthTypeID)
	b.PutString(a.Foo)
	return nil
}

// Decode implements bin.Decoder.
func (a *UserAuth) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.auth#f4815592 to nil")
	}
	if err := b.ConsumeID(UserAuthTypeID); err != nil {
		return fmt.Errorf("unable to decode user.auth#f4815592: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user.auth#f4815592: field foo: %w", err)
		}
		a.Foo = value
	}
	return nil
}

// construct implements constructor of UserAuthClass.
func (a UserAuth) construct() UserAuthClass { return &a }

// Ensuring interfaces in compile-time for UserAuth.
var (
	_ bin.Encoder = &UserAuth{}
	_ bin.Decoder = &UserAuth{}

	_ UserAuthClass = &UserAuth{}
)

// UserAuthPassword represents TL type `user.authPassword#5981e317`.
type UserAuthPassword struct {
	// Pwd field of UserAuthPassword.
	Pwd string
}

// UserAuthPasswordTypeID is TL type id of UserAuthPassword.
const UserAuthPasswordTypeID = 0x5981e317

// Encode implements bin.Encoder.
func (a *UserAuthPassword) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.authPassword#5981e317 as nil")
	}
	b.PutID(UserAuthPasswordTypeID)
	b.PutString(a.Pwd)
	return nil
}

// Decode implements bin.Decoder.
func (a *UserAuthPassword) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.authPassword#5981e317 to nil")
	}
	if err := b.ConsumeID(UserAuthPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode user.authPassword#5981e317: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user.authPassword#5981e317: field pwd: %w", err)
		}
		a.Pwd = value
	}
	return nil
}

// construct implements constructor of UserAuthClass.
func (a UserAuthPassword) construct() UserAuthClass { return &a }

// Ensuring interfaces in compile-time for UserAuthPassword.
var (
	_ bin.Encoder = &UserAuthPassword{}
	_ bin.Decoder = &UserAuthPassword{}

	_ UserAuthClass = &UserAuthPassword{}
)

// UserAuthClass represents user.Auth generic type.
//
// Example:
//  g, err := DecodeUserAuth(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserAuth: // user.auth#f4815592
//  case *UserAuthPassword: // user.authPassword#5981e317
//  default: panic(v)
//  }
type UserAuthClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserAuthClass
}

// DecodeUserAuth implements binary de-serialization for UserAuthClass.
func DecodeUserAuth(buf *bin.Buffer) (UserAuthClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserAuthTypeID:
		// Decoding user.auth#f4815592.
		v := UserAuth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserAuthClass: %w", err)
		}
		return &v, nil
	case UserAuthPasswordTypeID:
		// Decoding user.authPassword#5981e317.
		v := UserAuthPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserAuthClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserAuthClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserAuth boxes the UserAuthClass providing a helper.
type UserAuthBox struct {
	UserAuth UserAuthClass
}

// Decode implements bin.Decoder for UserAuthBox.
func (b *UserAuthBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserAuthBox to nil")
	}
	v, err := DecodeUserAuth(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserAuth = v
	return nil
}

// Encode implements bin.Encode for UserAuthBox.
func (b *UserAuthBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserAuth == nil {
		return fmt.Errorf("unable to encode UserAuthClass as nil")
	}
	return b.UserAuth.Encode(buf)
}
