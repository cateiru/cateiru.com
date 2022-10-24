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

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateProductHandler(e)
	})
}

func TestUpdateProductHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateProductHandler(e)
	})
}

func TestDeleteProductHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteProductHandler(e)
	})
}
