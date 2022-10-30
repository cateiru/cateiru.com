package handler_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestNoticeHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearNotice(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		n, err := u.CreateNotice()
		require.NoError(t, err)
		_, err = n.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.NoticeHandler(e)
		require.NoError(t, err)

		responseNotice := []ent.Notice{}
		err = m.Json(&responseNotice)
		require.NoError(t, err)

		require.Len(t, responseNotice, 1)
		require.Equal(t, responseNotice[0].DiscordWebhook, n.DiscordWebhook)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.NoticeHandler(e)
	})
}

func TestCreateNoticeHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateNoticeHandler(e)
	})
}

func TestUpdateNoticeHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateNoticeHandler(e)
	})
}

func TestDeleteNoticeHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteNoticeHandler(e)
	})
}
