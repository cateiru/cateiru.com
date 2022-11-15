package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/product"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestProductHandler(t *testing.T) {
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
		prod, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = prod.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		// other user and products
		secondU, err := tool.NewUser(ctx)
		require.NoError(t, err)
		secondProd, err := secondU.CreateProduct()
		require.NoError(t, err)
		_, err = secondProd.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.ProductHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		res := new([]ent.Product)
		err = m.Json(res)
		require.NoError(t, err)

		require.Len(t, *res, 1)

		require.EqualValues(t, (*res)[0].Name, prod.Name)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ProductHandler(e)
	})
}

func TestCreateProductHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("name", "cateiru.com")
		form.Insert("name_ja", "cateiru.com")
		form.Insert("detail", "my portfolio")
		form.Insert("detail_ja", "ポートフォリオ")
		form.Insert("site_url", "https://cateiru.com")
		form.Insert("dev_time", "2022-11-15T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateProductHandler(e)
		require.NoError(t, err)

		m.Status(t, http.StatusCreated)

		res := new(ent.Product)
		err = m.Json(res)
		require.NoError(t, err)
		prodInDB, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(res.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, prodInDB.Name, "cateiru.com")
		require.Equal(t, prodInDB.DetailJa, "ポートフォリオ")
		require.Equal(t, prodInDB.SiteURL, "https://cateiru.com")
		require.Equal(t, prodInDB.GithubURL, "")
		require.Equal(t, prodInDB.Thumbnail, "")
	})

	t.Run("all change", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("name", "cateiru.com")
		form.Insert("name_ja", "cateiru.com")
		form.Insert("detail", "my portfolio")
		form.Insert("detail_ja", "ポートフォリオ")
		form.Insert("site_url", "https://cateiru.com")
		form.Insert("github_url", "https://github.com/cateiru/cateiru.com")
		form.Insert("thumbnail", "https://caiteiru.com/ogp.png")
		form.Insert("dev_time", "2022-11-15T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateProductHandler(e)
		require.NoError(t, err)

		m.Status(t, http.StatusCreated)

		res := new(ent.Product)
		err = m.Json(res)
		require.NoError(t, err)
		prodInDB, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(res.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, prodInDB.Name, "cateiru.com")
		require.Equal(t, prodInDB.DetailJa, "ポートフォリオ")
		require.Equal(t, prodInDB.GithubURL, "https://github.com/cateiru/cateiru.com")
		require.Equal(t, prodInDB.Thumbnail, "https://caiteiru.com/ogp.png")
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateProductHandler(e)
	})
}

func TestUpdateProductHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		p, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = p.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("product_id", strconv.Itoa(int(p.Product.ID)))
		form.Insert("name", "cateiru--aaaa")

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateProductHandler(e)
		require.NoError(t, err)

		prod, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(p.Product.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, prod.Name, "cateiru--aaaa")
	})

	t.Run("all change", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		p, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = p.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("product_id", strconv.Itoa(int(p.Product.ID)))
		form.Insert("name", "cateiru.com")
		form.Insert("name_ja", "cateiru.com")
		form.Insert("detail", "my portfolio")
		form.Insert("detail_ja", "ポートフォリオ")
		form.Insert("site_url", "https://cateiru.com")
		form.Insert("github_url", "https://github.com/cateiru/cateiru.com")
		form.Insert("thumbnail", "https://caiteiru.com/ogp.png")
		form.Insert("dev_time", "2022-11-25T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateProductHandler(e)
		require.NoError(t, err)

		prod, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(p.Product.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, prod.Name, "cateiru.com")
		require.Equal(t, prod.NameJa, "cateiru.com")
		require.Equal(t, prod.Detail, "my portfolio")
		require.Equal(t, prod.DetailJa, "ポートフォリオ")
		require.Equal(t, prod.SiteURL, "https://cateiru.com")
		require.Equal(t, prod.GithubURL, "https://github.com/cateiru/cateiru.com")
		require.Equal(t, prod.Thumbnail, "https://caiteiru.com/ogp.png")
		require.Equal(t, prod.DevTime.Day(), 25)
	})

	t.Run("no changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		p, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = p.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		form := contents.NewMultipart()
		form.Insert("product_id", strconv.Itoa(int(p.Product.ID)))

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateProductHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateProductHandler(e)
	})
}

func TestDeleteProductHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		p, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = p.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodDelete, fmt.Sprintf("/?product_id=%v", p.Product.ID))
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.DeleteProductHandler(e)
		require.NoError(t, err)

		exist, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(p.Product.ID)).
			Exist(ctx)
		require.NoError(t, err)
		require.False(t, exist)
	})

	t.Run("other users", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		u2, err := tool.NewUser(ctx)
		require.NoError(t, err)
		p2, err := u2.CreateProduct()
		require.NoError(t, err)
		_, err = p2.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodDelete, fmt.Sprintf("/?product_id=%v", p2.Product.ID))
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.DeleteProductHandler(e)
		require.NoError(t, err)

		exist, err := tool.DB.Client.Product.
			Query().
			Where(product.ID(p2.Product.ID)).
			Exist(ctx)
		require.NoError(t, err)
		require.True(t, exist)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteProductHandler(e)
	})
}
