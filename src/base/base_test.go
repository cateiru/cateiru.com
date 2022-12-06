package base_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru.com/src/base"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	test.Init()

	t.Run("create base", func(t *testing.T) {
		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		base, err := base.NewBase(tool.DB)
		require.NoError(t, err)

		db := base.DB
		require.NotEmpty(t, db)
	})

	t.Run("session", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
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
			b, err := base.NewBase(tool.DB)
			require.NoError(t, err)

			err = b.Session(ctx, c)
			require.NoError(t, err)
			require.NotNil(t, b.User)
		})

		t.Run("no exist", func(t *testing.T) {
			m, err := mock.NewMock("", http.MethodGet, "/")
			require.NoError(t, err)

			c := m.Echo()
			b, err := base.NewBase(tool.DB)
			require.NoError(t, err)

			err = b.Session(ctx, c)
			require.Error(t, err)
			require.Nil(t, b.User)
		})
	})

	t.Run("expired session token", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		// Create session token
		user, err := tool.NewUser(ctx)
		require.NoError(t, err)
		s, err := tool.DB.Client.Session.Create().
			SetUserID(user.User.ID).
			SetPeriod(time.Now().Add(-10 * time.Hour)).
			Save(ctx)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		sessionCookie := &http.Cookie{
			Name:  config.Config.SessionCookieName,
			Value: s.String(),
		}
		m.Cookie([]*http.Cookie{sessionCookie})

		c := m.Echo()
		b, err := base.NewBase(tool.DB)
		require.NoError(t, err)

		err = b.Session(ctx, c)
		require.Error(t, err)
		require.Nil(t, b.User)
	})

	t.Run("login", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		c := m.Echo()
		b, err := base.NewBase(tool.DB)
		require.NoError(t, err)

		err = b.Login(ctx, c, u.User)
		require.NoError(t, err)

		err = c.String(http.StatusOK, "")
		require.NoError(t, err)

		m.Ok(t)

		cookies := m.W.Result().Cookies()

		require.Len(t, cookies, 2)
	})

	t.Run("logout", func(t *testing.T) {
		ctx := context.Background()
		tool, err := test.NewTestTool()
		require.NoError(t, err)
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

		e := m.Echo()
		handler := func(e echo.Context) error {
			base, err := base.NewBase(tool.DB)
			if err != nil {
				return err
			}

			base.Logout(e)
			return nil
		}
		err = handler(e)
		require.NoError(t, err)

		m.Ok(t)
		cookies := m.W.Result().Cookies()
		require.Len(t, cookies, 2)

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
