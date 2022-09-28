package base_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/base"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	test.Init()

	t.Run("create base", func(t *testing.T) {
		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		c := m.Echo()

		base, err := base.NewBase(c)
		require.NoError(t, err)

		db := base.DB
		require.NotEmpty(t, db)
		e := base.E
		require.NotEmpty(t, e)
	})

	t.Run("session", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestToolDB()
		require.NoError(t, err)
		defer tool.Close()

		// Create session token
		user, err := tool.NewUser(ctx)
		require.NoError(t, err)
		sessionToken, err := user.CreateSession(ctx, tool.DB)
		require.NoError(t, err)

		t.Run("exist token in cookies", func(t *testing.T) {
			m, err := mock.NewMock("", http.MethodGet, "/")
			require.NoError(t, err)

			sessionCookie := &http.Cookie{
				Name:  config.Config.SessionCookieName,
				Value: sessionToken.String(),
			}
			m.Cookie([]*http.Cookie{sessionCookie})

			c := m.Echo()
			b, err := base.NewBase(c)
			require.NoError(t, err)
			defer b.Close()

			err = b.Session(ctx)
			require.NoError(t, err)
			require.NotNil(t, b.User)
		})

		t.Run("no exist", func(t *testing.T) {
			m, err := mock.NewMock("", http.MethodGet, "/")
			require.NoError(t, err)

			c := m.Echo()
			b, err := base.NewBase(c)
			require.NoError(t, err)
			defer b.Close()

			err = b.Session(ctx)
			require.Error(t, err)
			require.Nil(t, b.User)
		})
	})

	t.Run("login", func(t *testing.T) {
		ctx := context.Background()

		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		c := m.Echo()
		b, err := base.NewBase(c)
		require.NoError(t, err)
		defer b.Close()

		tool := test.NewTestTool(b.DB)
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = b.Login(ctx, u.User)
		require.NoError(t, err)

		err = b.E.String(http.StatusOK, "")
		require.NoError(t, err)

		m.Ok(t)

		cookies := m.W.Result().Cookies()

		require.Len(t, cookies, 2)
	})
}
