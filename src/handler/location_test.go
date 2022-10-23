package handler_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/labstack/echo/v4"
)

func TestLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.LocationHandler(e)
	})
}

func TestCreateLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateLocationHandler(e)
	})
}

func TestUpdateLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateLocationHandler(e)
	})
}

func TestDeleteLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteLocationHandler(e)
	})
}
