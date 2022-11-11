package handler_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/contact"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestContactHandler(t *testing.T) {
	test.Init()

}

func TestContactGetHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactGetHandler(e)
	})

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, true)
		require.NoError(t, err)

		form, err := tool.NewForm()
		require.NoError(t, err)

		_, err = form.InsertDB(ctx, tool.DB, u.User.ID)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()

		err = h.ContactGetHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		resp := []ent.Category{}

		err = m.Json(&resp)
		require.NoError(t, err)

		require.Len(t, resp, 1)
		require.Equal(t, resp[0].Name, form.Name)
	})
}

func TestContactDeleteHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, true)
		require.NoError(t, err)

		form, err := tool.NewForm()
		require.NoError(t, err)

		c, err := form.InsertDB(ctx, tool.DB, u.User.ID)
		require.NoError(t, err)

		m, err := mock.NewGet("", fmt.Sprintf("/?contact_id=%v", c.ID))
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()

		err = h.ContactDeleteHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		exists, err := tool.DB.Client.Contact.
			Query().
			Where(contact.ID(c.ID)).
			Exist(ctx)
		require.NoError(t, err)
		require.False(t, exists)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactDeleteHandler(e)
	})
}
