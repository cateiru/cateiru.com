package handler

import (
	"context"

	"github.com/cateiru/cateiru.com/src/base"
	"github.com/labstack/echo/v4"
)

func LogoutHandler(e echo.Context) error {
	ctx := context.Background()
	base, err := base.NewBase(e)
	if err != nil {
		return err
	}
	defer base.Close()

	return base.Logout(ctx)
}
