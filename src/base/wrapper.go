package base

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Wrapper(h func(base *Base) error) echo.HandlerFunc {
	hand := func(c echo.Context) error {
		base, err := NewBase(c)
		if err != nil {
			return err
		}

		if err := h(base); err != nil {
			return err
		}

		return nil
	}

	return hand
}

func SessionWrapper(h func(base *SessionBase) error) echo.HandlerFunc {
	hand := func(c echo.Context) error {
		ctx := context.Background()

		base, err := NewSession(ctx, c)
		if err != nil {
			return err
		}

		if err := h(base); err != nil {
			return err
		}

		return nil
	}

	return hand
}
