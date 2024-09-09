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

// GiveawayInfoOngoing represents TL type `giveawayInfoOngoing#624ee050`.
type GiveawayInfoOngoing struct {
	// Point in time (Unix timestamp) when the giveaway was created
	CreationDate int32
	// Status of the current user in the giveaway
	Status GiveawayParticipantStatusClass
	// True, if the giveaway has ended and results are being prepared
	IsEnded bool
}

// GiveawayInfoOngoingTypeID is TL type id of GiveawayInfoOngoing.
const GiveawayInfoOngoingTypeID = 0x624ee050

// construct implements constructor of GiveawayInfoClass.
func (g GiveawayInfoOngoing) construct() GiveawayInfoClass { return &g }

// Ensuring interfaces in compile-time for GiveawayInfoOngoing.
var (
	_ bin.Encoder     = &GiveawayInfoOngoing{}
	_ bin.Decoder     = &GiveawayInfoOngoing{}
	_ bin.BareEncoder = &GiveawayInfoOngoing{}
	_ bin.BareDecoder = &GiveawayInfoOngoing{}

	_ GiveawayInfoClass = &GiveawayInfoOngoing{}
)

func (g *GiveawayInfoOngoing) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.CreationDate == 0) {
		return false
	}
	if !(g.Status == nil) {
		return false
	}
	if !(g.IsEnded == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GiveawayInfoOngoing) String() string {
	if g == nil {
		return "GiveawayInfoOngoing(nil)"
	}
	type Alias GiveawayInfoOngoing
	return fmt.Sprintf("GiveawayInfoOngoing%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GiveawayInfoOngoing) TypeID() uint32 {
	return GiveawayInfoOngoingTypeID
}

// TypeName returns name of type in TL schema.
func (*GiveawayInfoOngoing) TypeName() string {
	return "giveawayInfoOngoing"
}

// TypeInfo returns info about TL type.
func (g *GiveawayInfoOngoing) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "giveawayInfoOngoing",
		ID:   GiveawayInfoOngoingTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CreationDate",
			SchemaName: "creation_date",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
		{
			Name:       "IsEnded",
			SchemaName: "is_ended",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GiveawayInfoOngoing) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoOngoing#624ee050 as nil")
	}
	b.PutID(GiveawayInfoOngoingTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GiveawayInfoOngoing) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoOngoing#624ee050 as nil")
	}
	b.PutInt32(g.CreationDate)
	if g.Status == nil {
		return fmt.Errorf("unable to encode giveawayInfoOngoing#624ee050: field status is nil")
	}
	if err := g.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode giveawayInfoOngoing#624ee050: field status: %w", err)
	}
	b.PutBool(g.IsEnded)
	return nil
}

// Decode implements bin.Decoder.
func (g *GiveawayInfoOngoing) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoOngoing#624ee050 to nil")
	}
	if err := b.ConsumeID(GiveawayInfoOngoingTypeID); err != nil {
		return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GiveawayInfoOngoing) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoOngoing#624ee050 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field creation_date: %w", err)
		}
		g.CreationDate = value
	}
	{
		value, err := DecodeGiveawayParticipantStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field status: %w", err)
		}
		g.Status = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field is_ended: %w", err)
		}
		g.IsEnded = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GiveawayInfoOngoing) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoOngoing#624ee050 as nil")
	}
	b.ObjStart()
	b.PutID("giveawayInfoOngoing")
	b.Comma()
	b.FieldStart("creation_date")
	b.PutInt32(g.CreationDate)
	b.Comma()
	b.FieldStart("status")
	if g.Status == nil {
		return fmt.Errorf("unable to encode giveawayInfoOngoing#624ee050: field status is nil")
	}
	if err := g.Status.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode giveawayInfoOngoing#624ee050: field status: %w", err)
	}
	b.Comma()
	b.FieldStart("is_ended")
	b.PutBool(g.IsEnded)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GiveawayInfoOngoing) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoOngoing#624ee050 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("giveawayInfoOngoing"); err != nil {
				return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: %w", err)
			}
		case "creation_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field creation_date: %w", err)
			}
			g.CreationDate = value
		case "status":
			value, err := DecodeTDLibJSONGiveawayParticipantStatus(b)
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field status: %w", err)
			}
			g.Status = value
		case "is_ended":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoOngoing#624ee050: field is_ended: %w", err)
			}
			g.IsEnded = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCreationDate returns value of CreationDate field.
func (g *GiveawayInfoOngoing) GetCreationDate() (value int32) {
	if g == nil {
		return
	}
	return g.CreationDate
}

// GetStatus returns value of Status field.
func (g *GiveawayInfoOngoing) GetStatus() (value GiveawayParticipantStatusClass) {
	if g == nil {
		return
	}
	return g.Status
}

// GetIsEnded returns value of IsEnded field.
func (g *GiveawayInfoOngoing) GetIsEnded() (value bool) {
	if g == nil {
		return
	}
	return g.IsEnded
}

// GiveawayInfoCompleted represents TL type `giveawayInfoCompleted#328cc35c`.
type GiveawayInfoCompleted struct {
	// Point in time (Unix timestamp) when the giveaway was created
	CreationDate int32
	// Point in time (Unix timestamp) when the winners were selected. May be bigger than
	// winners selection date specified in parameters of the giveaway
	ActualWinnersSelectionDate int32
	// True, if the giveaway was canceled and was fully refunded
	WasRefunded bool
	// True, if the cuurent user is a winner of the giveaway
	IsWinner bool
	// Number of winners in the giveaway
	WinnerCount int32
	// Number of winners, which activated their gift codes; for Telegram Premium giveaways
	// only
	ActivationCount int32
	// Telegram Premium gift code that was received by the current user; empty if the user
	// isn't a winner in the giveaway or the giveaway isn't a Telegram Premium giveaway
	GiftCode string
	// The amount of Telegram Stars won by the current user; 0 if the user isn't a winner in
	// the giveaway or the giveaway isn't a Telegram Star giveaway
	WonStarCount int64
}

// GiveawayInfoCompletedTypeID is TL type id of GiveawayInfoCompleted.
const GiveawayInfoCompletedTypeID = 0x328cc35c

// construct implements constructor of GiveawayInfoClass.
func (g GiveawayInfoCompleted) construct() GiveawayInfoClass { return &g }

// Ensuring interfaces in compile-time for GiveawayInfoCompleted.
var (
	_ bin.Encoder     = &GiveawayInfoCompleted{}
	_ bin.Decoder     = &GiveawayInfoCompleted{}
	_ bin.BareEncoder = &GiveawayInfoCompleted{}
	_ bin.BareDecoder = &GiveawayInfoCompleted{}

	_ GiveawayInfoClass = &GiveawayInfoCompleted{}
)

func (g *GiveawayInfoCompleted) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.CreationDate == 0) {
		return false
	}
	if !(g.ActualWinnersSelectionDate == 0) {
		return false
	}
	if !(g.WasRefunded == false) {
		return false
	}
	if !(g.IsWinner == false) {
		return false
	}
	if !(g.WinnerCount == 0) {
		return false
	}
	if !(g.ActivationCount == 0) {
		return false
	}
	if !(g.GiftCode == "") {
		return false
	}
	if !(g.WonStarCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GiveawayInfoCompleted) String() string {
	if g == nil {
		return "GiveawayInfoCompleted(nil)"
	}
	type Alias GiveawayInfoCompleted
	return fmt.Sprintf("GiveawayInfoCompleted%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GiveawayInfoCompleted) TypeID() uint32 {
	return GiveawayInfoCompletedTypeID
}

// TypeName returns name of type in TL schema.
func (*GiveawayInfoCompleted) TypeName() string {
	return "giveawayInfoCompleted"
}

// TypeInfo returns info about TL type.
func (g *GiveawayInfoCompleted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "giveawayInfoCompleted",
		ID:   GiveawayInfoCompletedTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CreationDate",
			SchemaName: "creation_date",
		},
		{
			Name:       "ActualWinnersSelectionDate",
			SchemaName: "actual_winners_selection_date",
		},
		{
			Name:       "WasRefunded",
			SchemaName: "was_refunded",
		},
		{
			Name:       "IsWinner",
			SchemaName: "is_winner",
		},
		{
			Name:       "WinnerCount",
			SchemaName: "winner_count",
		},
		{
			Name:       "ActivationCount",
			SchemaName: "activation_count",
		},
		{
			Name:       "GiftCode",
			SchemaName: "gift_code",
		},
		{
			Name:       "WonStarCount",
			SchemaName: "won_star_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GiveawayInfoCompleted) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoCompleted#328cc35c as nil")
	}
	b.PutID(GiveawayInfoCompletedTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GiveawayInfoCompleted) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoCompleted#328cc35c as nil")
	}
	b.PutInt32(g.CreationDate)
	b.PutInt32(g.ActualWinnersSelectionDate)
	b.PutBool(g.WasRefunded)
	b.PutBool(g.IsWinner)
	b.PutInt32(g.WinnerCount)
	b.PutInt32(g.ActivationCount)
	b.PutString(g.GiftCode)
	b.PutInt53(g.WonStarCount)
	return nil
}

// Decode implements bin.Decoder.
func (g *GiveawayInfoCompleted) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoCompleted#328cc35c to nil")
	}
	if err := b.ConsumeID(GiveawayInfoCompletedTypeID); err != nil {
		return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GiveawayInfoCompleted) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoCompleted#328cc35c to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field creation_date: %w", err)
		}
		g.CreationDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field actual_winners_selection_date: %w", err)
		}
		g.ActualWinnersSelectionDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field was_refunded: %w", err)
		}
		g.WasRefunded = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field is_winner: %w", err)
		}
		g.IsWinner = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field winner_count: %w", err)
		}
		g.WinnerCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field activation_count: %w", err)
		}
		g.ActivationCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field gift_code: %w", err)
		}
		g.GiftCode = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field won_star_count: %w", err)
		}
		g.WonStarCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GiveawayInfoCompleted) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode giveawayInfoCompleted#328cc35c as nil")
	}
	b.ObjStart()
	b.PutID("giveawayInfoCompleted")
	b.Comma()
	b.FieldStart("creation_date")
	b.PutInt32(g.CreationDate)
	b.Comma()
	b.FieldStart("actual_winners_selection_date")
	b.PutInt32(g.ActualWinnersSelectionDate)
	b.Comma()
	b.FieldStart("was_refunded")
	b.PutBool(g.WasRefunded)
	b.Comma()
	b.FieldStart("is_winner")
	b.PutBool(g.IsWinner)
	b.Comma()
	b.FieldStart("winner_count")
	b.PutInt32(g.WinnerCount)
	b.Comma()
	b.FieldStart("activation_count")
	b.PutInt32(g.ActivationCount)
	b.Comma()
	b.FieldStart("gift_code")
	b.PutString(g.GiftCode)
	b.Comma()
	b.FieldStart("won_star_count")
	b.PutInt53(g.WonStarCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GiveawayInfoCompleted) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode giveawayInfoCompleted#328cc35c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("giveawayInfoCompleted"); err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: %w", err)
			}
		case "creation_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field creation_date: %w", err)
			}
			g.CreationDate = value
		case "actual_winners_selection_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field actual_winners_selection_date: %w", err)
			}
			g.ActualWinnersSelectionDate = value
		case "was_refunded":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field was_refunded: %w", err)
			}
			g.WasRefunded = value
		case "is_winner":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field is_winner: %w", err)
			}
			g.IsWinner = value
		case "winner_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field winner_count: %w", err)
			}
			g.WinnerCount = value
		case "activation_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field activation_count: %w", err)
			}
			g.ActivationCount = value
		case "gift_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field gift_code: %w", err)
			}
			g.GiftCode = value
		case "won_star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode giveawayInfoCompleted#328cc35c: field won_star_count: %w", err)
			}
			g.WonStarCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCreationDate returns value of CreationDate field.
func (g *GiveawayInfoCompleted) GetCreationDate() (value int32) {
	if g == nil {
		return
	}
	return g.CreationDate
}

// GetActualWinnersSelectionDate returns value of ActualWinnersSelectionDate field.
func (g *GiveawayInfoCompleted) GetActualWinnersSelectionDate() (value int32) {
	if g == nil {
		return
	}
	return g.ActualWinnersSelectionDate
}

// GetWasRefunded returns value of WasRefunded field.
func (g *GiveawayInfoCompleted) GetWasRefunded() (value bool) {
	if g == nil {
		return
	}
	return g.WasRefunded
}

// GetIsWinner returns value of IsWinner field.
func (g *GiveawayInfoCompleted) GetIsWinner() (value bool) {
	if g == nil {
		return
	}
	return g.IsWinner
}

// GetWinnerCount returns value of WinnerCount field.
func (g *GiveawayInfoCompleted) GetWinnerCount() (value int32) {
	if g == nil {
		return
	}
	return g.WinnerCount
}

// GetActivationCount returns value of ActivationCount field.
func (g *GiveawayInfoCompleted) GetActivationCount() (value int32) {
	if g == nil {
		return
	}
	return g.ActivationCount
}

// GetGiftCode returns value of GiftCode field.
func (g *GiveawayInfoCompleted) GetGiftCode() (value string) {
	if g == nil {
		return
	}
	return g.GiftCode
}

// GetWonStarCount returns value of WonStarCount field.
func (g *GiveawayInfoCompleted) GetWonStarCount() (value int64) {
	if g == nil {
		return
	}
	return g.WonStarCount
}

// GiveawayInfoClassName is schema name of GiveawayInfoClass.
const GiveawayInfoClassName = "GiveawayInfo"

// GiveawayInfoClass represents GiveawayInfo generic type.
//
// Example:
//
//	g, err := tdapi.DecodeGiveawayInfo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.GiveawayInfoOngoing: // giveawayInfoOngoing#624ee050
//	case *tdapi.GiveawayInfoCompleted: // giveawayInfoCompleted#328cc35c
//	default: panic(v)
//	}
type GiveawayInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() GiveawayInfoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error

	// Point in time (Unix timestamp) when the giveaway was created
	GetCreationDate() (value int32)
}

// DecodeGiveawayInfo implements binary de-serialization for GiveawayInfoClass.
func DecodeGiveawayInfo(buf *bin.Buffer) (GiveawayInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case GiveawayInfoOngoingTypeID:
		// Decoding giveawayInfoOngoing#624ee050.
		v := GiveawayInfoOngoing{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", err)
		}
		return &v, nil
	case GiveawayInfoCompletedTypeID:
		// Decoding giveawayInfoCompleted#328cc35c.
		v := GiveawayInfoCompleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONGiveawayInfo implements binary de-serialization for GiveawayInfoClass.
func DecodeTDLibJSONGiveawayInfo(buf tdjson.Decoder) (GiveawayInfoClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "giveawayInfoOngoing":
		// Decoding giveawayInfoOngoing#624ee050.
		v := GiveawayInfoOngoing{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", err)
		}
		return &v, nil
	case "giveawayInfoCompleted":
		// Decoding giveawayInfoCompleted#328cc35c.
		v := GiveawayInfoCompleted{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode GiveawayInfoClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// GiveawayInfo boxes the GiveawayInfoClass providing a helper.
type GiveawayInfoBox struct {
	GiveawayInfo GiveawayInfoClass
}

// Decode implements bin.Decoder for GiveawayInfoBox.
func (b *GiveawayInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode GiveawayInfoBox to nil")
	}
	v, err := DecodeGiveawayInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.GiveawayInfo = v
	return nil
}

// Encode implements bin.Encode for GiveawayInfoBox.
func (b *GiveawayInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.GiveawayInfo == nil {
		return fmt.Errorf("unable to encode GiveawayInfoClass as nil")
	}
	return b.GiveawayInfo.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for GiveawayInfoBox.
func (b *GiveawayInfoBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode GiveawayInfoBox to nil")
	}
	v, err := DecodeTDLibJSONGiveawayInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.GiveawayInfo = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for GiveawayInfoBox.
func (b *GiveawayInfoBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.GiveawayInfo == nil {
		return fmt.Errorf("unable to encode GiveawayInfoClass as nil")
	}
	return b.GiveawayInfo.EncodeTDLibJSON(buf)
}
