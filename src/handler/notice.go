package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteNoticeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}
