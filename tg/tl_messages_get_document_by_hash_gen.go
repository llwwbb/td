// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesGetDocumentByHashRequest represents TL type `messages.getDocumentByHash#338e2464`.
// Get a document by its SHA256 hash, mainly used for gifs
//
// See https://core.telegram.org/method/messages.getDocumentByHash for reference.
type MessagesGetDocumentByHashRequest struct {
	// SHA256 of file
	SHA256 []byte
	// Size of the file in bytes
	Size int
	// Mime type
	MimeType string
}

// MessagesGetDocumentByHashRequestTypeID is TL type id of MessagesGetDocumentByHashRequest.
const MessagesGetDocumentByHashRequestTypeID = 0x338e2464

func (g *MessagesGetDocumentByHashRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SHA256 == nil) {
		return false
	}
	if !(g.Size == 0) {
		return false
	}
	if !(g.MimeType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDocumentByHashRequest) String() string {
	if g == nil {
		return "MessagesGetDocumentByHashRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetDocumentByHashRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSHA256: ")
	sb.WriteString(fmt.Sprint(g.SHA256))
	sb.WriteString(",\n")
	sb.WriteString("\tSize: ")
	sb.WriteString(fmt.Sprint(g.Size))
	sb.WriteString(",\n")
	sb.WriteString("\tMimeType: ")
	sb.WriteString(fmt.Sprint(g.MimeType))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetDocumentByHashRequest) TypeID() uint32 {
	return MessagesGetDocumentByHashRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetDocumentByHashRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDocumentByHash#338e2464 as nil")
	}
	b.PutID(MessagesGetDocumentByHashRequestTypeID)
	b.PutBytes(g.SHA256)
	b.PutInt(g.Size)
	b.PutString(g.MimeType)
	return nil
}

// GetSHA256 returns value of SHA256 field.
func (g *MessagesGetDocumentByHashRequest) GetSHA256() (value []byte) {
	return g.SHA256
}

// GetSize returns value of Size field.
func (g *MessagesGetDocumentByHashRequest) GetSize() (value int) {
	return g.Size
}

// GetMimeType returns value of MimeType field.
func (g *MessagesGetDocumentByHashRequest) GetMimeType() (value string) {
	return g.MimeType
}

// Decode implements bin.Decoder.
func (g *MessagesGetDocumentByHashRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDocumentByHash#338e2464 to nil")
	}
	if err := b.ConsumeID(MessagesGetDocumentByHashRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field sha256: %w", err)
		}
		g.SHA256 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field size: %w", err)
		}
		g.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDocumentByHash#338e2464: field mime_type: %w", err)
		}
		g.MimeType = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDocumentByHashRequest.
var (
	_ bin.Encoder = &MessagesGetDocumentByHashRequest{}
	_ bin.Decoder = &MessagesGetDocumentByHashRequest{}
)

// MessagesGetDocumentByHash invokes method messages.getDocumentByHash#338e2464 returning error if any.
// Get a document by its SHA256 hash, mainly used for gifs
//
// Possible errors:
//  400 SHA256_HASH_INVALID: The provided SHA256 hash is invalid
//
// See https://core.telegram.org/method/messages.getDocumentByHash for reference.
// Can be used by bots.
func (c *Client) MessagesGetDocumentByHash(ctx context.Context, request *MessagesGetDocumentByHashRequest) (DocumentClass, error) {
	var result DocumentBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Document, nil
}
