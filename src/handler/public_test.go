package handler_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestPublicProfileHandler(t *testing.T) {
	tool, err := test.NewTestTool()
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, true)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.NoError(t, err)

		m.Ok(t)
	})

	t.Run("all status false", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)

		_, err = tool.NewUser(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.ErrorIs(t, err, echo.ErrNotFound)
	})

	t.Run("user empty", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.ErrorIs(t, err, echo.ErrNotFound)
	})
}
