package handler

import (
	"net/http"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/labstack/echo/v4"
)

func RootHandler(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, config.Config.PageDomain.String())
}
