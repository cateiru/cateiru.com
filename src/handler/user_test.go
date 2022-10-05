package handler_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

// Response login user data
func TestMeHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		tool, err := test.NewTestToolDB()
		require.NoError(t, err)
		defer tool.Close()

		user, err := tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/user/me")
		require.NoError(t, err)

		user.HandlerSession(ctx, tool.DB, m)

		e := m.Echo()

		err = handler.MeHandler(e)
		require.NoError(t, err)

		m.Ok(t)
	})

	t.Run("no login", func(t *testing.T) {
		tool, err := test.NewTestToolDB()
		require.NoError(t, err)
		defer tool.Close()

		m, err := mock.NewGet("", "/user/me")
		require.NoError(t, err)

		e := m.Echo()

		err = handler.MeHandler(e)
		require.Error(t, err)
	})
}
