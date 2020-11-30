package telegram

import (
	"fmt"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/crypto"
	"github.com/gotd/td/internal/proto"
)

func (c *Client) newEncryptedMessage(id crypto.MessageID, seq int32, payload bin.Encoder, b *bin.Buffer) error {
	if err := payload.Encode(b); err != nil {
		return err
	}
	{
		typeID, err := b.PeekID()
		if err == nil {
			c.log.With(
				zap.Int64("message_id", int64(id)),
				zap.String("message_type", fmt.Sprintf("0x%x", typeID)),
				zap.String("message_type_str", c.types.Get(typeID)),
			).Debug("Request")
		}
	}
	d := proto.EncryptedMessageData{
		SessionID:              atomic.LoadInt64(&c.session),
		Salt:                   atomic.LoadInt64(&c.salt),
		MessageID:              id,
		SeqNo:                  seq,
		MessageDataLen:         int32(b.Len()),
		MessageDataWithPadding: b.Copy(),
	}

	b.Reset()
	if err := d.Encode(b); err != nil {
		return err
	}
	msg, err := c.encrypt(b.Copy())
	if err != nil {
		return err
	}

	b.Reset()
	if err := msg.Encode(b); err != nil {
		return err
	}

	c.log.With(zap.Int64("request_id", int64(d.MessageID))).Debug("Request")

	return nil
}
