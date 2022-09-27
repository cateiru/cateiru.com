package src

import (
	"github.com/cateiru/cateiru.com/src/base"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/labstack/echo/v4"
)

// API Routes
func Routes(e *echo.Echo) {
	e.GET("/", base.Wrapper(handler.RootHandler))
}