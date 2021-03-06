package message

import (
	"context"
	"testing"

	"github.com/gotd/td/tg"
)

func TestGeoPoint(t *testing.T) {
	ctx := context.Background()
	sender, mock := testSender(t)
	point := tg.InputGeoPoint{
		Lat:            11,
		Long:           12,
		AccuracyRadius: 10,
	}

	expectSendMedia(&tg.InputMediaGeoPoint{
		GeoPoint: &point,
	}, mock)

	mock.NoError(sender.Self().Media(ctx, GeoPoint(11, 12, 10)))
}
