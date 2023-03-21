package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
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

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.NoticeHandler(e)
		require.NoError(t, err)

		responseNotice := ent.Notice{}
		err = m.Json(&responseNotice)
		require.NoError(t, err)

		require.Equal(t, responseNotice.DiscordWebhook, n.DiscordWebhook)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.NoticeHandler(e)
	})
}

func TestUpdateNoticeHandler(t *testing.T) {
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

		form := easy.NewMultipart()
		form.Insert("discord_webhook", "https://webhook.discord.com/hogehoge")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateNoticeHandler(e)
		require.NoError(t, err)

		nDB, err := tool.DB.Client.Notice.
			Query().
			Where(notice.UserID(u.User.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, nDB.Mail, "")
		require.Equal(t, nDB.DiscordWebhook, "https://webhook.discord.com/hogehoge")
		require.Equal(t, nDB.SlackWebhook, "")
	})

	t.Run("all changes", func(t *testing.T) {
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

		form := easy.NewMultipart()
		form.Insert("discord_webhook", "https://webhook.discord.com/hogehoge")
		form.Insert("slack_webhook", "https://webhook.slack.com/aaaaaaa")
		form.Insert("mail", "test@cateiru.com")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateNoticeHandler(e)
		require.NoError(t, err)

		nDB, err := tool.DB.Client.Notice.
			Query().
			Where(notice.UserID(u.User.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, nDB.Mail, "test@cateiru.com")
		require.Equal(t, nDB.DiscordWebhook, "https://webhook.discord.com/hogehoge")
		require.Equal(t, nDB.SlackWebhook, "https://webhook.slack.com/aaaaaaa")
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateNoticeHandler(e)
	})
}
