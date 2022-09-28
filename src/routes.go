package src

import (
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/labstack/echo/v4"
)

// API Routes
func Routes(e *echo.Echo) {
	e.GET("/", handler.RootHandler)

	e.GET("/login", handler.LoginHandler)
	e.GET("/login/url", handler.LoginURLHandler)
}
