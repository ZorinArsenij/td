package mtproto

import (
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
)

func (c *Conn) handleResult(b *bin.Buffer) error {
	// Response to an RPC query.
	var res proto.Result
	if err := res.Decode(b); err != nil {
		return xerrors.Errorf("decode: %x", err)
	}
	c.logWithType(b).Debug("Handle result",
		zap.Int64("msg_id", res.RequestMessageID),
	)

	// Handling gzipped results.
	id, err := b.PeekID()
	if err != nil {
		return err
	}
	if id == proto.GZIPTypeID {
		content, err := gzip(b)
		if err != nil {
			return xerrors.Errorf("decompress: %w", err)
		}

		// Replacing buffer so callback will deal with uncompressed data.
		b = content
		c.logWithType(b).Debug("Decompressed",
			zap.Int64("msg_id", res.RequestMessageID),
		)

		// Replacing id with inner id if error is compressed for any reason.
		if id, err = b.PeekID(); err != nil {
			return xerrors.Errorf("peek id: %w", err)
		}
	}

	if id == mt.RPCErrorTypeID {
		var rpcErr mt.RPCError
		if err := rpcErr.Decode(b); err != nil {
			return xerrors.Errorf("error decode: %w", err)
		}
		c.rpc.NotifyError(
			res.RequestMessageID,
			NewError(rpcErr.ErrorCode, rpcErr.ErrorMessage),
		)

		return nil
	}
	if id == mt.PongTypeID {
		return c.handlePong(b)
	}

	return c.rpc.NotifyResult(res.RequestMessageID, b)
}
