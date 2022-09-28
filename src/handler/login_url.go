package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/pkg/go/sso"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/labstack/echo/v4"
)

// Redirect to CateiruSSO Login page.
func LoginURLHandler(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, sso.CreateURI(config.Config.SSOClientID, config.Config.SSORedirectURI.String()))
}
