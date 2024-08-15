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

// StarSubscription represents TL type `starSubscription#6caf0ade`.
type StarSubscription struct {
	// Unique identifier of the subscription
	ID string
	// Identifier of the channel chat that is subscribed
	ChatID int64
	// Point in time (Unix timestamp) when the subscription will expire or expired
	ExpirationDate int32
	// True, if the subscription is active and the user can use the method
	// reuseStarSubscription to join the subscribed chat again
	CanReuse bool
	// True, if the subscription was canceled
	IsCanceled bool
	// True, if the subscription expires soon and there are no enough Telegram Stars on the
	// user's balance to extend it
	IsExpiring bool
	// The invite link that can be used to renew the subscription if it has been expired; may
	// be empty, if the link isn't available anymore
	InviteLink string
	// The subscription plan
	Pricing StarSubscriptionPricing
}

// StarSubscriptionTypeID is TL type id of StarSubscription.
const StarSubscriptionTypeID = 0x6caf0ade

// Ensuring interfaces in compile-time for StarSubscription.
var (
	_ bin.Encoder     = &StarSubscription{}
	_ bin.Decoder     = &StarSubscription{}
	_ bin.BareEncoder = &StarSubscription{}
	_ bin.BareDecoder = &StarSubscription{}
)

func (s *StarSubscription) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.ExpirationDate == 0) {
		return false
	}
	if !(s.CanReuse == false) {
		return false
	}
	if !(s.IsCanceled == false) {
		return false
	}
	if !(s.IsExpiring == false) {
		return false
	}
	if !(s.InviteLink == "") {
		return false
	}
	if !(s.Pricing.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarSubscription) String() string {
	if s == nil {
		return "StarSubscription(nil)"
	}
	type Alias StarSubscription
	return fmt.Sprintf("StarSubscription%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarSubscription) TypeID() uint32 {
	return StarSubscriptionTypeID
}

// TypeName returns name of type in TL schema.
func (*StarSubscription) TypeName() string {
	return "starSubscription"
}

// TypeInfo returns info about TL type.
func (s *StarSubscription) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starSubscription",
		ID:   StarSubscriptionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "ExpirationDate",
			SchemaName: "expiration_date",
		},
		{
			Name:       "CanReuse",
			SchemaName: "can_reuse",
		},
		{
			Name:       "IsCanceled",
			SchemaName: "is_canceled",
		},
		{
			Name:       "IsExpiring",
			SchemaName: "is_expiring",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "Pricing",
			SchemaName: "pricing",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarSubscription) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starSubscription#6caf0ade as nil")
	}
	b.PutID(StarSubscriptionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarSubscription) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starSubscription#6caf0ade as nil")
	}
	b.PutString(s.ID)
	b.PutInt53(s.ChatID)
	b.PutInt32(s.ExpirationDate)
	b.PutBool(s.CanReuse)
	b.PutBool(s.IsCanceled)
	b.PutBool(s.IsExpiring)
	b.PutString(s.InviteLink)
	if err := s.Pricing.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starSubscription#6caf0ade: field pricing: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarSubscription) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starSubscription#6caf0ade to nil")
	}
	if err := b.ConsumeID(StarSubscriptionTypeID); err != nil {
		return fmt.Errorf("unable to decode starSubscription#6caf0ade: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarSubscription) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starSubscription#6caf0ade to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field expiration_date: %w", err)
		}
		s.ExpirationDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field can_reuse: %w", err)
		}
		s.CanReuse = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field is_canceled: %w", err)
		}
		s.IsCanceled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field is_expiring: %w", err)
		}
		s.IsExpiring = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field invite_link: %w", err)
		}
		s.InviteLink = value
	}
	{
		if err := s.Pricing.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starSubscription#6caf0ade: field pricing: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarSubscription) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starSubscription#6caf0ade as nil")
	}
	b.ObjStart()
	b.PutID("starSubscription")
	b.Comma()
	b.FieldStart("id")
	b.PutString(s.ID)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("expiration_date")
	b.PutInt32(s.ExpirationDate)
	b.Comma()
	b.FieldStart("can_reuse")
	b.PutBool(s.CanReuse)
	b.Comma()
	b.FieldStart("is_canceled")
	b.PutBool(s.IsCanceled)
	b.Comma()
	b.FieldStart("is_expiring")
	b.PutBool(s.IsExpiring)
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(s.InviteLink)
	b.Comma()
	b.FieldStart("pricing")
	if err := s.Pricing.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode starSubscription#6caf0ade: field pricing: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarSubscription) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starSubscription#6caf0ade to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starSubscription"); err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field id: %w", err)
			}
			s.ID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field chat_id: %w", err)
			}
			s.ChatID = value
		case "expiration_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field expiration_date: %w", err)
			}
			s.ExpirationDate = value
		case "can_reuse":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field can_reuse: %w", err)
			}
			s.CanReuse = value
		case "is_canceled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field is_canceled: %w", err)
			}
			s.IsCanceled = value
		case "is_expiring":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field is_expiring: %w", err)
			}
			s.IsExpiring = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field invite_link: %w", err)
			}
			s.InviteLink = value
		case "pricing":
			if err := s.Pricing.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode starSubscription#6caf0ade: field pricing: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *StarSubscription) GetID() (value string) {
	if s == nil {
		return
	}
	return s.ID
}

// GetChatID returns value of ChatID field.
func (s *StarSubscription) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetExpirationDate returns value of ExpirationDate field.
func (s *StarSubscription) GetExpirationDate() (value int32) {
	if s == nil {
		return
	}
	return s.ExpirationDate
}

// GetCanReuse returns value of CanReuse field.
func (s *StarSubscription) GetCanReuse() (value bool) {
	if s == nil {
		return
	}
	return s.CanReuse
}

// GetIsCanceled returns value of IsCanceled field.
func (s *StarSubscription) GetIsCanceled() (value bool) {
	if s == nil {
		return
	}
	return s.IsCanceled
}

// GetIsExpiring returns value of IsExpiring field.
func (s *StarSubscription) GetIsExpiring() (value bool) {
	if s == nil {
		return
	}
	return s.IsExpiring
}

// GetInviteLink returns value of InviteLink field.
func (s *StarSubscription) GetInviteLink() (value string) {
	if s == nil {
		return
	}
	return s.InviteLink
}

// GetPricing returns value of Pricing field.
func (s *StarSubscription) GetPricing() (value StarSubscriptionPricing) {
	if s == nil {
		return
	}
	return s.Pricing
}
