package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/contactdefault"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestPublicContactDefaultHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		_, err = tool.DB.Client.ContactDefault.Delete().Exec(ctx)
		require.NoError(t, err)

		c, err := tool.DB.Client.ContactDefault.Create().
			SetName("hoge").
			Save(ctx)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", c.ID), http.MethodGet, "")
		require.NoError(t, err)

		e := m.Echo()

		err = h.PublicContactDefaultHandler(e)
		require.NoError(t, err)

		categories := ent.ContactDefault{}
		err = m.Json(&categories)
		require.NoError(t, err)

		require.Equal(t, categories.Name, c.Name)
	})
}

func TestContactDefaultHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		_, err = tool.DB.Client.ContactDefault.Delete().Exec(ctx)
		require.NoError(t, err)

		c, err := tool.DB.Client.ContactDefault.Create().
			SetName("hoge").
			Save(ctx)
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.ContactDefaultHandler(e)
		require.NoError(t, err)

		categories := []ent.ContactDefault{}
		err = m.Json(&categories)
		require.NoError(t, err)

		require.Len(t, categories, 1)
		require.Equal(t, categories[0].Name, c.Name)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactDefaultHandler(e)
	})
}

func TestCreateContactDefaultHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("name", "hoge")
		form.Insert("email", "hoge@examile.com")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.CreateContactDefaultHandler(e)
		require.NoError(t, err)

		c := new(ent.ContactDefault)
		err = m.Json(c)
		require.NoError(t, err)

		require.Equal(t, *c.Name, "hoge")
		require.Equal(t, *c.Email, "hoge@examile.com")
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactDefaultHandler(e)
	})
}

func TestUpdateContactDefaultHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		c, err := tool.DB.Client.ContactDefault.Create().
			SetName("hoge").
			Save(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("id", fmt.Sprint(c.ID))
		form.Insert("name", "nyaaa")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateContactDefaultHandler(e)
		require.NoError(t, err)

		contactDefault, err := tool.DB.Client.ContactDefault.Get(ctx, c.ID)
		require.NoError(t, err)

		require.Equal(t, *contactDefault.Name, "nyaaa")
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactDefaultHandler(e)
	})
}

func TestDeleteContactDefaultHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		c, err := tool.DB.Client.ContactDefault.Create().
			SetName("hoge").
			Save(ctx)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%v", c.ID), http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.DeleteContactDefaultHandler(e)
		require.NoError(t, err)

		exist, err := tool.DB.Client.ContactDefault.
			Query().
			Where(contactdefault.ID(c.ID)).
			Exist(ctx)
		require.NoError(t, err)

		require.False(t, exist)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ContactDefaultHandler(e)
	})
}
