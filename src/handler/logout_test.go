package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

func TestLogout(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()
		err = h.LogoutHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		cookies := m.W.Result().Cookies()
		c := new(http.Cookie)
		for _, cookie := range cookies {
			if cookie.Name == config.Config.SessionCookieName {
				c = cookie
			}
		}
		require.NotNil(t, c)
		require.Equal(t, c.MaxAge, 0)
	})
}
