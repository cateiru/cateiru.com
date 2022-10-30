package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}
	return nil
}
