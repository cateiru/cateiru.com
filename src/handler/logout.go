package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func (h Handler) LogoutHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	h.Base.Logout(ctx, e)

	return nil
}
