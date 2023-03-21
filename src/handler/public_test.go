package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
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
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, tool.DB, true)
		require.NoError(t, err)

		// bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		bio.IsPublic = true
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)
		// product
		product, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = product.CreateDB(ctx, tool.DB)
		require.NoError(t, err)
		// link
		link, err := u.CreateLink()
		require.NoError(t, err)
		_, err = link.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		response := new(handler.Public)
		err = m.Json(response)
		require.NoError(t, err)

		require.Equal(t, response.GivenName, u.GivenName)
		require.Len(t, response.Biographies, 1)
		require.Len(t, response.Products, 1)
		require.Len(t, response.Links, 1)
	})

	t.Run("all status false", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		_, err = tool.NewUser(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.ErrorIs(t, err, echo.ErrNotFound)
	})

	t.Run("user empty", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.ErrorIs(t, err, echo.ErrNotFound)
	})

	t.Run("no bio", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, tool.DB, true)
		require.NoError(t, err)

		// product
		product, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = product.CreateDB(ctx, tool.DB)
		require.NoError(t, err)
		// link
		link, err := u.CreateLink()
		require.NoError(t, err)
		_, err = link.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		response := new(handler.Public)
		err = m.Json(response)
		require.NoError(t, err)

		require.Equal(t, response.GivenName, u.GivenName)
		require.Len(t, response.Biographies, 0)
		require.Len(t, response.Products, 1)
		require.Len(t, response.Links, 1)
	})

	t.Run("no product", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, tool.DB, true)
		require.NoError(t, err)

		// bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		bio.IsPublic = true
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)
		// link
		link, err := u.CreateLink()
		require.NoError(t, err)
		_, err = link.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		response := new(handler.Public)
		err = m.Json(response)
		require.NoError(t, err)

		require.Equal(t, response.GivenName, u.GivenName)
		require.Len(t, response.Biographies, 1)
		require.Len(t, response.Products, 0)
		require.Len(t, response.Links, 1)
	})

	t.Run("no links", func(t *testing.T) {
		ctx := context.Background()
		err := tool.ClearUser(ctx)
		require.NoError(t, err)
		err = tool.ClearBio(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, tool.DB, true)
		require.NoError(t, err)

		// bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		bio.IsPublic = true
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)
		// product
		product, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = product.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProfileHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		response := new(handler.Public)
		err = m.Json(response)
		require.NoError(t, err)

		require.Equal(t, response.GivenName, u.GivenName)
		require.Len(t, response.Biographies, 1)
		require.Len(t, response.Products, 1)
		require.Len(t, response.Links, 0)
	})
}

func TestPublicProductHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		p1, err := u.CreateProduct()
		require.NoError(t, err)
		_, err = p1.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?product_id=%v", p1.Product.ID), http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProductsHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		responseProduct := new(handler.PublicProduct)
		err = m.Json(responseProduct)
		require.NoError(t, err)

		require.Equal(t, p1.Name, responseProduct.Name)
	})

	t.Run("no exist product id", func(t *testing.T) {
		tool, err := test.NewTestTool()
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := easy.NewMock("/?product_id=123456", http.MethodGet, "")
		require.NoError(t, err)
		e := m.Echo()

		err = h.PublicProductsHandler(e)
		require.ErrorIs(t, err, echo.ErrNotFound)
	})
}
