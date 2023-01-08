package handler

import (
	"github.com/labstack/echo/v4"
)

// Logout session
// no delete users.
func (h *Handler) LogoutHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	h.Base.Logout(e)

	return nil
}
