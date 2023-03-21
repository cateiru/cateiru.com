package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestCategoryHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		l, err := u.CreateLink()
		require.NoError(t, err)
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.CategoryHandler(e)
		require.NoError(t, err)

		categories := []ent.Category{}
		err = m.Json(&categories)
		require.NoError(t, err)

		require.Len(t, categories, 1)
		require.Equal(t, categories[0].Name, l.Category.Name)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CategoryHandler(e)
	})
}

func TestCreateCategoryHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("name", "hoge")
		form.Insert("name_ja", "あああ")
		form.Insert("emoji", "🍙")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.CreateCategoryHandler(e)
		require.NoError(t, err)

		c := new(ent.Category)
		err = m.Json(c)
		require.NoError(t, err)

		require.Equal(t, c.Name, "hoge")
		require.Equal(t, c.NameJa, "あああ")
		require.Equal(t, c.Emoji, "🍙")
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateCategoryHandler(e)
	})
}

func TestUpdateCategoryHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		l, err := u.CreateLink()
		require.NoError(t, err)
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))
		form.Insert("name", "nyaaa")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateCategoryHandler(e)
		require.NoError(t, err)

		categoryDB, err := tool.DB.Client.Category.Get(ctx, l.Category.ID)
		require.NoError(t, err)

		require.Equal(t, categoryDB.Name, "nyaaa")
	})

	t.Run("all changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		l, err := u.CreateLink()
		require.NoError(t, err)
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))
		form.Insert("name", "nyaaa")
		form.Insert("name_ja", "ああああ")
		form.Insert("emoji", "🥑")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateCategoryHandler(e)
		require.NoError(t, err)

		categoryDB, err := tool.DB.Client.Category.Get(ctx, l.Category.ID)
		require.NoError(t, err)

		require.Equal(t, categoryDB.Name, "nyaaa")
		require.Equal(t, categoryDB.NameJa, "ああああ")
		require.Equal(t, categoryDB.Emoji, "🥑")
	})

	t.Run("no change", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		l, err := u.CreateLink()
		require.NoError(t, err)
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateCategoryHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateCategoryHandler(e)
	})
}

func TestDeleteCategoryHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearLink(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		l, err := u.CreateLink()
		require.NoError(t, err)
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?category_id=%v", l.Category.ID), http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.DeleteCategoryHandler(e)
		require.NoError(t, err)

		exist, err := tool.DB.Client.Category.
			Query().
			Where(category.ID(l.Category.ID)).
			Exist(ctx)
		require.NoError(t, err)

		require.False(t, exist)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteCategoryHandler(e)
	})
}
