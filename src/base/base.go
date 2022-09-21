package base

import "github.com/labstack/echo/v4"

type Base struct {
	e *echo.Echo
}

// Create a Base objects
func New(e *echo.Echo) *Base {
	return &Base{
		e: e,
	}
}
