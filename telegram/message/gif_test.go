package message

import (
	"context"
	"testing"
	"time"

	"github.com/gotd/td/tg"
)

func TestGIF(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)
	file := &tg.InputFile{
		ID: 10,
	}

	expectSendMedia(&tg.InputMediaUploadedDocument{
		File:     file,
		MimeType: DefaultGifMIME,
		Attributes: []tg.DocumentAttributeClass{
			&tg.DocumentAttributeAnimated{},
		},
	}, mock)
	expectSendMedia(&tg.InputMediaUploadedDocument{
		File:     file,
		MimeType: DefaultGifMIME,
		Attributes: []tg.DocumentAttributeClass{
			&tg.DocumentAttributeAnimated{},
		},
		TTLSeconds: 10,
	}, mock)

	mock.NoError(sender.Self().GIF(ctx, file))
	mock.NoError(sender.Self().Media(ctx, GIF(file).TTL(10*time.Second)))
}
