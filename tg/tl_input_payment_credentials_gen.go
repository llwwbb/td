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

// InputPaymentCredentialsSaved represents TL type `inputPaymentCredentialsSaved#c10eb2cf`.
type InputPaymentCredentialsSaved struct {
	// ID field of InputPaymentCredentialsSaved.
	ID string
	// TmpPassword field of InputPaymentCredentialsSaved.
	TmpPassword []byte
}

// InputPaymentCredentialsSavedTypeID is TL type id of InputPaymentCredentialsSaved.
const InputPaymentCredentialsSavedTypeID = 0xc10eb2cf

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsSaved) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsSaved#c10eb2cf as nil")
	}
	b.PutID(InputPaymentCredentialsSavedTypeID)
	b.PutString(i.ID)
	b.PutBytes(i.TmpPassword)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsSaved) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsSaved#c10eb2cf to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsSavedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field tmp_password: %w", err)
		}
		i.TmpPassword = value
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsSaved) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsSaved.
var (
	_ bin.Encoder = &InputPaymentCredentialsSaved{}
	_ bin.Decoder = &InputPaymentCredentialsSaved{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsSaved{}
)

// InputPaymentCredentials represents TL type `inputPaymentCredentials#3417d728`.
type InputPaymentCredentials struct {
	// Flags field of InputPaymentCredentials.
	Flags bin.Fields
	// Save field of InputPaymentCredentials.
	Save bool
	// Data field of InputPaymentCredentials.
	Data DataJSON
}

// InputPaymentCredentialsTypeID is TL type id of InputPaymentCredentials.
const InputPaymentCredentialsTypeID = 0x3417d728

// Encode implements bin.Encoder.
func (i *InputPaymentCredentials) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentials#3417d728 as nil")
	}
	b.PutID(InputPaymentCredentialsTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field flags: %w", err)
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field data: %w", err)
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (i *InputPaymentCredentials) SetSave(value bool) {
	if value {
		i.Flags.Set(0)
	} else {
		i.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentials) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentials#3417d728 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field flags: %w", err)
		}
	}
	i.Save = i.Flags.Has(0)
	{
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentials) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentials.
var (
	_ bin.Encoder = &InputPaymentCredentials{}
	_ bin.Decoder = &InputPaymentCredentials{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentials{}
)

// InputPaymentCredentialsApplePay represents TL type `inputPaymentCredentialsApplePay#aa1c39f`.
type InputPaymentCredentialsApplePay struct {
	// PaymentData field of InputPaymentCredentialsApplePay.
	PaymentData DataJSON
}

// InputPaymentCredentialsApplePayTypeID is TL type id of InputPaymentCredentialsApplePay.
const InputPaymentCredentialsApplePayTypeID = 0xaa1c39f

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsApplePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsApplePay#aa1c39f as nil")
	}
	b.PutID(InputPaymentCredentialsApplePayTypeID)
	if err := i.PaymentData.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsApplePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsApplePay#aa1c39f to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsApplePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: %w", err)
	}
	{
		if err := i.PaymentData.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsApplePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsApplePay.
var (
	_ bin.Encoder = &InputPaymentCredentialsApplePay{}
	_ bin.Decoder = &InputPaymentCredentialsApplePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsApplePay{}
)

// InputPaymentCredentialsAndroidPay represents TL type `inputPaymentCredentialsAndroidPay#ca05d50e`.
type InputPaymentCredentialsAndroidPay struct {
	// PaymentToken field of InputPaymentCredentialsAndroidPay.
	PaymentToken DataJSON
	// GoogleTransactionID field of InputPaymentCredentialsAndroidPay.
	GoogleTransactionID string
}

// InputPaymentCredentialsAndroidPayTypeID is TL type id of InputPaymentCredentialsAndroidPay.
const InputPaymentCredentialsAndroidPayTypeID = 0xca05d50e

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsAndroidPay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsAndroidPay#ca05d50e as nil")
	}
	b.PutID(InputPaymentCredentialsAndroidPayTypeID)
	if err := i.PaymentToken.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsAndroidPay#ca05d50e: field payment_token: %w", err)
	}
	b.PutString(i.GoogleTransactionID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsAndroidPay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsAndroidPay#ca05d50e to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsAndroidPayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: %w", err)
	}
	{
		if err := i.PaymentToken.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: field payment_token: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: field google_transaction_id: %w", err)
		}
		i.GoogleTransactionID = value
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsAndroidPay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsAndroidPay.
var (
	_ bin.Encoder = &InputPaymentCredentialsAndroidPay{}
	_ bin.Decoder = &InputPaymentCredentialsAndroidPay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsAndroidPay{}
)

// InputPaymentCredentialsClass represents InputPaymentCredentials generic type.
//
// Example:
//  g, err := DecodeInputPaymentCredentials(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputPaymentCredentialsSaved: // inputPaymentCredentialsSaved#c10eb2cf
//  case *InputPaymentCredentials: // inputPaymentCredentials#3417d728
//  case *InputPaymentCredentialsApplePay: // inputPaymentCredentialsApplePay#aa1c39f
//  case *InputPaymentCredentialsAndroidPay: // inputPaymentCredentialsAndroidPay#ca05d50e
//  default: panic(v)
//  }
type InputPaymentCredentialsClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPaymentCredentialsClass
}

// DecodeInputPaymentCredentials implements binary de-serialization for InputPaymentCredentialsClass.
func DecodeInputPaymentCredentials(buf *bin.Buffer) (InputPaymentCredentialsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPaymentCredentialsSavedTypeID:
		// Decoding inputPaymentCredentialsSaved#c10eb2cf.
		v := InputPaymentCredentialsSaved{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsTypeID:
		// Decoding inputPaymentCredentials#3417d728.
		v := InputPaymentCredentials{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsApplePayTypeID:
		// Decoding inputPaymentCredentialsApplePay#aa1c39f.
		v := InputPaymentCredentialsApplePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsAndroidPayTypeID:
		// Decoding inputPaymentCredentialsAndroidPay#ca05d50e.
		v := InputPaymentCredentialsAndroidPay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPaymentCredentials boxes the InputPaymentCredentialsClass providing a helper.
type InputPaymentCredentialsBox struct {
	InputPaymentCredentials InputPaymentCredentialsClass
}

// Decode implements bin.Decoder for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaymentCredentialsBox to nil")
	}
	v, err := DecodeInputPaymentCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaymentCredentials = v
	return nil
}

// Encode implements bin.Encode for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPaymentCredentials == nil {
		return fmt.Errorf("unable to encode InputPaymentCredentialsClass as nil")
	}
	return b.InputPaymentCredentials.Encode(buf)
}
