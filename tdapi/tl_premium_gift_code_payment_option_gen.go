// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PremiumGiftCodePaymentOption represents TL type `premiumGiftCodePaymentOption#d89959ed`.
type PremiumGiftCodePaymentOption struct {
	// ISO 4217 currency code for Telegram Premium gift code payment
	Currency string
	// The amount to pay, in the smallest units of the currency
	Amount int64
	// The discount associated with this option, as a percentage
	DiscountPercentage int32
	// Number of users which will be able to activate the gift codes
	WinnerCount int32
	// Number of months the Telegram Premium subscription will be active
	MonthCount int32
	// Identifier of the store product associated with the option; may be empty if none
	StoreProductID string
	// Number of times the store product must be paid
	StoreProductQuantity int32
	// A sticker to be shown along with the gift code; may be null if unknown
	Sticker Sticker
}

// PremiumGiftCodePaymentOptionTypeID is TL type id of PremiumGiftCodePaymentOption.
const PremiumGiftCodePaymentOptionTypeID = 0xd89959ed

// Ensuring interfaces in compile-time for PremiumGiftCodePaymentOption.
var (
	_ bin.Encoder     = &PremiumGiftCodePaymentOption{}
	_ bin.Decoder     = &PremiumGiftCodePaymentOption{}
	_ bin.BareEncoder = &PremiumGiftCodePaymentOption{}
	_ bin.BareDecoder = &PremiumGiftCodePaymentOption{}
)

func (p *PremiumGiftCodePaymentOption) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Currency == "") {
		return false
	}
	if !(p.Amount == 0) {
		return false
	}
	if !(p.DiscountPercentage == 0) {
		return false
	}
	if !(p.WinnerCount == 0) {
		return false
	}
	if !(p.MonthCount == 0) {
		return false
	}
	if !(p.StoreProductID == "") {
		return false
	}
	if !(p.StoreProductQuantity == 0) {
		return false
	}
	if !(p.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumGiftCodePaymentOption) String() string {
	if p == nil {
		return "PremiumGiftCodePaymentOption(nil)"
	}
	type Alias PremiumGiftCodePaymentOption
	return fmt.Sprintf("PremiumGiftCodePaymentOption%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGiftCodePaymentOption) TypeID() uint32 {
	return PremiumGiftCodePaymentOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGiftCodePaymentOption) TypeName() string {
	return "premiumGiftCodePaymentOption"
}

// TypeInfo returns info about TL type.
func (p *PremiumGiftCodePaymentOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumGiftCodePaymentOption",
		ID:   PremiumGiftCodePaymentOptionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "DiscountPercentage",
			SchemaName: "discount_percentage",
		},
		{
			Name:       "WinnerCount",
			SchemaName: "winner_count",
		},
		{
			Name:       "MonthCount",
			SchemaName: "month_count",
		},
		{
			Name:       "StoreProductID",
			SchemaName: "store_product_id",
		},
		{
			Name:       "StoreProductQuantity",
			SchemaName: "store_product_quantity",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumGiftCodePaymentOption) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiftCodePaymentOption#d89959ed as nil")
	}
	b.PutID(PremiumGiftCodePaymentOptionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumGiftCodePaymentOption) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiftCodePaymentOption#d89959ed as nil")
	}
	b.PutString(p.Currency)
	b.PutInt53(p.Amount)
	b.PutInt32(p.DiscountPercentage)
	b.PutInt32(p.WinnerCount)
	b.PutInt32(p.MonthCount)
	b.PutString(p.StoreProductID)
	b.PutInt32(p.StoreProductQuantity)
	if err := p.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumGiftCodePaymentOption#d89959ed: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumGiftCodePaymentOption) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiftCodePaymentOption#d89959ed to nil")
	}
	if err := b.ConsumeID(PremiumGiftCodePaymentOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumGiftCodePaymentOption) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiftCodePaymentOption#d89959ed to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field currency: %w", err)
		}
		p.Currency = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field amount: %w", err)
		}
		p.Amount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field discount_percentage: %w", err)
		}
		p.DiscountPercentage = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field winner_count: %w", err)
		}
		p.WinnerCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field month_count: %w", err)
		}
		p.MonthCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field store_product_id: %w", err)
		}
		p.StoreProductID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field store_product_quantity: %w", err)
		}
		p.StoreProductQuantity = value
	}
	{
		if err := p.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field sticker: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumGiftCodePaymentOption) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiftCodePaymentOption#d89959ed as nil")
	}
	b.ObjStart()
	b.PutID("premiumGiftCodePaymentOption")
	b.Comma()
	b.FieldStart("currency")
	b.PutString(p.Currency)
	b.Comma()
	b.FieldStart("amount")
	b.PutInt53(p.Amount)
	b.Comma()
	b.FieldStart("discount_percentage")
	b.PutInt32(p.DiscountPercentage)
	b.Comma()
	b.FieldStart("winner_count")
	b.PutInt32(p.WinnerCount)
	b.Comma()
	b.FieldStart("month_count")
	b.PutInt32(p.MonthCount)
	b.Comma()
	b.FieldStart("store_product_id")
	b.PutString(p.StoreProductID)
	b.Comma()
	b.FieldStart("store_product_quantity")
	b.PutInt32(p.StoreProductQuantity)
	b.Comma()
	b.FieldStart("sticker")
	if err := p.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumGiftCodePaymentOption#d89959ed: field sticker: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumGiftCodePaymentOption) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiftCodePaymentOption#d89959ed to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumGiftCodePaymentOption"); err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: %w", err)
			}
		case "currency":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field currency: %w", err)
			}
			p.Currency = value
		case "amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field amount: %w", err)
			}
			p.Amount = value
		case "discount_percentage":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field discount_percentage: %w", err)
			}
			p.DiscountPercentage = value
		case "winner_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field winner_count: %w", err)
			}
			p.WinnerCount = value
		case "month_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field month_count: %w", err)
			}
			p.MonthCount = value
		case "store_product_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field store_product_id: %w", err)
			}
			p.StoreProductID = value
		case "store_product_quantity":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field store_product_quantity: %w", err)
			}
			p.StoreProductQuantity = value
		case "sticker":
			if err := p.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode premiumGiftCodePaymentOption#d89959ed: field sticker: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCurrency returns value of Currency field.
func (p *PremiumGiftCodePaymentOption) GetCurrency() (value string) {
	if p == nil {
		return
	}
	return p.Currency
}

// GetAmount returns value of Amount field.
func (p *PremiumGiftCodePaymentOption) GetAmount() (value int64) {
	if p == nil {
		return
	}
	return p.Amount
}

// GetDiscountPercentage returns value of DiscountPercentage field.
func (p *PremiumGiftCodePaymentOption) GetDiscountPercentage() (value int32) {
	if p == nil {
		return
	}
	return p.DiscountPercentage
}

// GetWinnerCount returns value of WinnerCount field.
func (p *PremiumGiftCodePaymentOption) GetWinnerCount() (value int32) {
	if p == nil {
		return
	}
	return p.WinnerCount
}

// GetMonthCount returns value of MonthCount field.
func (p *PremiumGiftCodePaymentOption) GetMonthCount() (value int32) {
	if p == nil {
		return
	}
	return p.MonthCount
}

// GetStoreProductID returns value of StoreProductID field.
func (p *PremiumGiftCodePaymentOption) GetStoreProductID() (value string) {
	if p == nil {
		return
	}
	return p.StoreProductID
}

// GetStoreProductQuantity returns value of StoreProductQuantity field.
func (p *PremiumGiftCodePaymentOption) GetStoreProductQuantity() (value int32) {
	if p == nil {
		return
	}
	return p.StoreProductQuantity
}

// GetSticker returns value of Sticker field.
func (p *PremiumGiftCodePaymentOption) GetSticker() (value Sticker) {
	if p == nil {
		return
	}
	return p.Sticker
}
