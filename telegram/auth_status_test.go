package telegram

import (
	"context"
	"testing"

	"github.com/gotd/td/mtproto"
	"github.com/gotd/td/telegram/internal/rpcmock"
	"github.com/gotd/td/tg"
)

func TestClient_AuthStatus(t *testing.T) {
	ctx := context.Background()

	t.Run("Authorized", mockClient(func(a *rpcmock.Mock, client *Client) {
		user := &tg.User{
			Username: "user",
		}
		a.Expect().ThenResult(&tg.UserClassVector{Elems: []tg.UserClass{user}})

		status, err := client.AuthStatus(ctx)
		a.NoError(err)
		a.True(status.Authorized)
		a.Equal(user, status.User)
	}))

	t.Run("Unauthorized", mockClient(func(a *rpcmock.Mock, client *Client) {
		a.Expect().ThenUnregistered()

		status, err := client.AuthStatus(ctx)
		a.NoError(err)
		a.False(status.Authorized)
	}))

	t.Run("Error", mockClient(func(a *rpcmock.Mock, client *Client) {
		a.Expect().ThenRPCErr(&mtproto.Error{
			Code:    500,
			Message: "BRUH",
			Type:    "BRUH",
		})

		_, err := client.AuthStatus(ctx)
		a.Error(err)
	}))
}

func TestClient_AuthIfNecessary(t *testing.T) {
	ctx := context.Background()

	t.Run("Authorized", mockClient(func(a *rpcmock.Mock, client *Client) {
		user := &tg.User{
			Username: "user",
		}
		a.Expect().ThenResult(&tg.UserClassVector{Elems: []tg.UserClass{user}})

		// Pass empty AuthFlow because it should not be called anyway.
		err := client.AuthIfNecessary(ctx, AuthFlow{})
		a.NoError(err)
	}))
}
