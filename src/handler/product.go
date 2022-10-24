package handler

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru.com/ent/product"
	"github.com/labstack/echo/v4"
)

// Response to all products of me.
func (h *Handler) ProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	products, err := h.DB.Client.Product.
		Query().
		Where(product.UserID(h.User.ID)).
		All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, products)
}

// Create a new products
//
// form values:
// - name: string
// - name_ja: string
// - detail: string
// - detail_ja: string
// - site_url: string
// - github_url: Optional string
// - dev_time: ISO 8601 type date
// - thumbnail: Optional string
func (h *Handler) CreateProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}
