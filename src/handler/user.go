package handler

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru.com/src/base"
	"github.com/labstack/echo/v4"
)

func MeHandler(e echo.Context) error {
	ctx := context.Background()
	base, err := base.NewBase(e)
	if err != nil {
		return err
	}
	defer base.Close()

	if err := base.Session(ctx); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, base.User)
}

// func ChangeUserHandler(e echo.Context) error {
// 	ctx := context.Background()
// 	base, err := base.NewBase(e)
// 	if err != nil {
// 		return err
// 	}
// 	defer base.Close()

// 	base.Session(ctx)

// 	return nil
// }
