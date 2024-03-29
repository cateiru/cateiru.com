package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestLinkHandler(t *testing.T) {
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
		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.LinkHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		r := new([]handler.LinkResponse)
		err = m.Json(r)
		require.NoError(t, err)

		require.Len(t, *r, 1)

		require.Equal(t, (*r)[0].Link.ID, l.Link.ID)
		require.Equal(t, (*r)[0].Category.ID, l.Category.ID)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.LinkHandler(e)
	})
}

func TestCreateLinkHandler(t *testing.T) {
	test.Init()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	f, err := os.ReadFile("./favicon_short_url.html")
	require.NoError(t, err)

	httpmock.RegisterResponder(
		"GET",
		"https://cateiru.com",
		httpmock.NewBytesResponder(http.StatusOK, f),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://cateiru.com/aaa/favicon.ico",
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

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

		// INSERT DB only category
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("name", "hoge")
		form.Insert("name_ja", "あああ")
		form.Insert("site_url", "https://cateiru.com")
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.CreateLinkHandler(e)
		require.NoError(t, err)

		links, err := tool.DB.Client.Link.
			Query().
			Where(link.UserID(u.User.ID)).
			All(ctx)
		require.NoError(t, err)

		require.Len(t, links, 1)
		require.Equal(t, links[0].Name, "hoge")
		require.Equal(t, links[0].NameJa, "あああ")
		require.Equal(t, links[0].SiteURL, "https://cateiru.com")
		require.Equal(t, links[0].CategoryID, l.Category.ID)
	})

	t.Run("failed form", func(t *testing.T) {
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

		// INSERT DB only category
		_, err = l.CreateCategoryDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("name", "hoge")
		form.Insert("name_ja", "あああ")
		// form.Insert("site_url", "https://cateiru.com")
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.CreateLinkHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateLinkHandler(e)
	})
}

func TestUpdateLinkHandler(t *testing.T) {
	test.Init()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	f, err := os.ReadFile("./favicon_short_url.html")
	require.NoError(t, err)

	httpmock.RegisterResponder(
		"GET",
		"https://cateiru.com",
		httpmock.NewBytesResponder(http.StatusOK, f),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://cateiru.com/aaa/favicon.ico",
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

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

		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("link_id", strconv.Itoa(int(l.Link.ID)))
		form.Insert("name", "nyancat")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateLinkHandler(e)
		require.NoError(t, err)

		link, err := tool.DB.Client.Link.
			Query().
			Where(link.ID(l.Link.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, link.Name, "nyancat")
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

		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("link_id", strconv.Itoa(int(l.Link.ID)))
		form.Insert("name", "hoge")
		form.Insert("name_ja", "あああ")
		form.Insert("site_url", "https://cateiru.com")
		form.Insert("category_id", strconv.Itoa(int(l.Category.ID)))

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateLinkHandler(e)
		require.NoError(t, err)

		link, err := tool.DB.Client.Link.
			Query().
			Where(link.ID(l.Link.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, link.Name, "hoge")
		require.Equal(t, link.NameJa, "あああ")
		require.Equal(t, link.SiteURL, "https://cateiru.com")
		require.Equal(t, link.CategoryID, l.Category.ID)
	})

	t.Run("update favicon", func(t *testing.T) {
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
		l.SiteURL = "https://cateiru.com"

		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("link_id", strconv.Itoa(int(l.Link.ID)))
		form.Insert("update_favicon", "true")

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateLinkHandler(e)
		require.NoError(t, err)

	})

	t.Run("no changes", func(t *testing.T) {
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

		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("link_id", strconv.Itoa(int(l.Link.ID)))

		m, err := easy.NewFormData("/", http.MethodPut, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.UpdateLinkHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateLinkHandler(e)
	})
}

func TestDeleteLinkHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteLinkHandler(e)
	})

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

		_, err = l.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?link_id=%v", l.Link.ID), http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.DeleteLinkHandler(e)
		require.NoError(t, err)

		exist, err := tool.DB.Client.Link.
			Query().
			Where(link.ID(l.Link.ID)).
			Exist(ctx)
		require.NoError(t, err)
		require.False(t, exist)
	})
}

func TestGetFavicon(t *testing.T) {
	test.Init()

	t.Run("get favicon from HTML", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/html/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)

		fav, err := handler.GetFavicon("https://example.com")
		require.NoError(t, err)

		require.Equal(t, fav, "https://example.com/html/favicon.ico")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 0)
		require.Equal(t, info["GET https://example.com/html/favicon.ico"], 1)
	})

	t.Run("no defined favicon in HTML", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon_no_defined.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/html/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)

		fav, err := handler.GetFavicon("https://example.com")
		require.NoError(t, err)

		require.Equal(t, fav, "https://example.com/favicon.ico")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 1)
		require.Equal(t, info["GET https://example.com/html/favicon.ico"], 0)
	})

	t.Run("favicon short url", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon_short_url.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://cateiru.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://cateiru.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://cateiru.com/aaa/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)

		fav, err := handler.GetFavicon("https://cateiru.com")
		require.NoError(t, err)

		require.Equal(t, fav, "https://cateiru.com/aaa/favicon.ico")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://cateiru.com"], 1)
		require.Equal(t, info["GET https://cateiru.com/favicon.ico"], 0)
		require.Equal(t, info["GET https://cateiru.com/aaa/favicon.ico"], 1)
	})

	t.Run("favicon not found", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon_no_defined.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusNotFound, ""),
		)

		fav, err := handler.GetFavicon("https://example.com")
		require.NoError(t, err)
		require.Equal(t, fav, "")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 1)
	})
}
