package message

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

// RequestBuilder is a intermediate builder to make different RPC calls using Sender.
type RequestBuilder struct {
	Builder
}

// ScreenshotNotify sends notification about screenshot.
// Parameter msgID is a ID of message that was screenshotted, can be 0.
func (b *RequestBuilder) ScreenshotNotify(ctx context.Context, msgID int) error {
	p, err := b.peer(ctx)
	if err != nil {
		return xerrors.Errorf("peer: %w", err)
	}

	if err := b.sender.sendScreenshotNotification(ctx, &tg.MessagesSendScreenshotNotificationRequest{
		Peer:         p,
		ReplyToMsgID: msgID,
	}); err != nil {
		return xerrors.Errorf("send screenshot notify: %w", err)
	}

	return nil
}

// PeerSettings returns peer settings.
func (b *RequestBuilder) PeerSettings(ctx context.Context) (*tg.PeerSettings, error) {
	p, err := b.peer(ctx)
	if err != nil {
		return nil, xerrors.Errorf("peer: %w", err)
	}

	settings, err := b.sender.getPeerSettings(ctx, p)
	if err != nil {
		return nil, xerrors.Errorf("get peer settings: %w", err)
	}

	return settings, nil
}

type startBotBuilder struct {
	bot   tg.InputUserClass
	param string
}

// StartBotOption is a option for StartBot method.
type StartBotOption func(s *startBotBuilder)

// StartBotInputUser sets InputUserClass to start bot.
func StartBotInputUser(bot tg.InputUserClass) func(s *startBotBuilder) {
	return func(s *startBotBuilder) {
		s.bot = bot
	}
}

// StartBotParam sets deeplink param to start bot.
func StartBotParam(param string) func(s *startBotBuilder) {
	return func(s *startBotBuilder) {
		s.param = param
	}
}

// StartBot starts a conversation with a bot using a deep linking parameter.
func (b *RequestBuilder) StartBot(ctx context.Context, opts ...StartBotOption) error {
	p, err := b.peer(ctx)
	if err != nil {
		return xerrors.Errorf("peer: %w", err)
	}

	sb := startBotBuilder{}
	for _, opt := range opts {
		opt(&sb)
	}

	if sb.bot == nil {
		switch u := p.(type) {
		case *tg.InputPeerUser:
			v := new(tg.InputUser)
			v.FillFrom(u)
			sb.bot = v
		case *tg.InputPeerUserFromMessage:
			v := new(tg.InputUserFromMessage)
			v.FillFrom(u)
			sb.bot = v
		default:
			return xerrors.Errorf("unexpected peer type %T, try to pass input user manually", p)
		}
	}

	if err := b.sender.startBot(ctx, &tg.MessagesStartBotRequest{
		Peer:       p,
		Bot:        sb.bot,
		StartParam: sb.param,
	}); err != nil {
		return xerrors.Errorf("start bot: %w", err)
	}

	return nil
}
