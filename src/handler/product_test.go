package handler_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/labstack/echo/v4"
)

func TestProductHandler(t *testing.T) {
	test.Init()

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
