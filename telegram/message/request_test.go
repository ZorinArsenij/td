package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
)

func TestPure(t *testing.T) {
	s := Sender{}
	dialog := s.Self()

	b1 := dialog.Reply(1).Reply(2).Reply(1)
	b2 := dialog.Reply(2)
	b3 := dialog

	require.Equal(t, 1, b1.replyToMsgID)
	require.Equal(t, 2, b2.replyToMsgID)
	require.Equal(t, 0, b3.replyToMsgID)
}

func TestRequestBuilder_ScreenshotNotify(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)

	mock.ExpectFunc(func(b bin.Encoder) {
		req, ok := b.(*tg.MessagesSendScreenshotNotificationRequest)
		mock.True(ok)
		mock.Equal(&tg.InputPeerSelf{}, req.Peer)
		mock.Equal(10, req.ReplyToMsgID)
	}).ThenResult(&tg.Updates{})
	mock.NoError(sender.Self().ScreenshotNotify(ctx, 10))

	mock.ExpectFunc(func(b bin.Encoder) {
		req, ok := b.(*tg.MessagesSendScreenshotNotificationRequest)
		mock.True(ok)
		mock.Equal(&tg.InputPeerSelf{}, req.Peer)
		mock.Equal(10, req.ReplyToMsgID)
	}).ThenRPCErr(testRPCError())
	mock.Error(sender.Self().ScreenshotNotify(ctx, 10))
}

func TestRequestBuilder_PeerSettings(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)

	expected := &tg.PeerSettings{
		ReportSpam: true,
	}
	mock.ExpectCall(&tg.MessagesGetPeerSettingsRequest{
		Peer: &tg.InputPeerSelf{},
	}).ThenResult(expected)

	s, err := sender.Self().PeerSettings(ctx)
	mock.NoError(err)
	mock.Equal(expected, s)

	mock.ExpectCall(&tg.MessagesGetPeerSettingsRequest{
		Peer: &tg.InputPeerSelf{},
	}).ThenRPCErr(testRPCError())

	_, err = sender.Self().PeerSettings(ctx)
	mock.Error(err)
}

func TestRequestBuilder_StartBot(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)
	var peer tg.InputPeerClass = &tg.InputPeerUser{
		UserID:     10,
		AccessHash: 10,
	}

	mock.ExpectFunc(func(b bin.Encoder) {
		req, ok := b.(*tg.MessagesStartBotRequest)
		mock.True(ok)
		mock.Equal(peer, req.Peer)
		mock.Equal(&tg.InputUser{
			UserID:     10,
			AccessHash: 10,
		}, req.Bot)
		mock.Equal("abc", req.StartParam)
		mock.NotZero(req.RandomID)
	}).ThenResult(&tg.Updates{})
	mock.NoError(sender.Peer(peer).StartBot(ctx, StartBotParam("abc")))

	inputUser := &tg.InputUserEmpty{}
	mock.ExpectFunc(func(b bin.Encoder) {
		req, ok := b.(*tg.MessagesStartBotRequest)
		mock.True(ok)
		mock.Equal(peer, req.Peer)
		mock.Equal(inputUser, req.Bot)
		mock.NotZero(req.RandomID)
	}).ThenResult(&tg.Updates{})
	mock.NoError(sender.Peer(peer).StartBot(ctx, StartBotInputUser(inputUser)))

	// Should not make RPC calls.
	mock.Error(sender.Self().StartBot(ctx))

	peerFromMsg := &tg.InputPeerUserFromMessage{
		Peer:   peer,
		UserID: 10,
		MsgID:  10,
	}
	mock.ExpectFunc(func(b bin.Encoder) {
		req, ok := b.(*tg.MessagesStartBotRequest)
		mock.True(ok)
		mock.Equal(peerFromMsg, req.Peer)
		mock.Equal(&tg.InputUserFromMessage{
			Peer:   peer,
			MsgID:  10,
			UserID: 10,
		}, req.Bot)
		mock.NotZero(req.RandomID)
	}).ThenRPCErr(testRPCError())

	mock.Error(sender.Peer(peerFromMsg).StartBot(ctx))
}
