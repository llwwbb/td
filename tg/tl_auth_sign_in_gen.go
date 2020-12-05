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

// AuthSignInRequest represents TL type `auth.signIn#bcd51581`.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
type AuthSignInRequest struct {
	// Phone number in the international format
	PhoneNumber string
	// SMS-message ID, obtained from auth.sendCode
	PhoneCodeHash string
	// Valid numerical code from the SMS-message
	PhoneCode string
}

// AuthSignInRequestTypeID is TL type id of AuthSignInRequest.
const AuthSignInRequestTypeID = 0xbcd51581

// Encode implements bin.Encoder.
func (s *AuthSignInRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.signIn#bcd51581 as nil")
	}
	b.PutID(AuthSignInRequestTypeID)
	b.PutString(s.PhoneNumber)
	b.PutString(s.PhoneCodeHash)
	b.PutString(s.PhoneCode)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSignInRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.signIn#bcd51581 to nil")
	}
	if err := b.ConsumeID(AuthSignInRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.signIn#bcd51581: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_code_hash: %w", err)
		}
		s.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_code: %w", err)
		}
		s.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthSignInRequest.
var (
	_ bin.Encoder = &AuthSignInRequest{}
	_ bin.Decoder = &AuthSignInRequest{}
)

// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
