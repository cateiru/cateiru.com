package handler

import (
	"context"
	"net/http"
	"time"

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

// Create a new product
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

	name := e.FormValue("name")
	nameJa := e.FormValue("name_ja")
	detail := e.FormValue("detail")
	detailJa := e.FormValue("detail_ja")
	siteUrl := e.FormValue("site_url")
	githubUrl := e.FormValue("github_url")
	devTimeStr := e.FormValue("dev_time")
	thumbnail := e.FormValue("thumbnail")

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name")
	}
	if nameJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name_ja")
	}
	if detail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: detail")
	}
	if detailJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: detail_ja")
	}
	if siteUrl == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: site_url")
	}
	if devTimeStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: dev_time")
	}

	devTime, err := time.Parse("2006-01-02T15:04:05-0700", devTimeStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: dev_time")
	}

	c := h.DB.Client.Product.
		Create().
		SetName(name).
		SetNameJa(nameJa).
		SetDetail(detail).
		SetDetailJa(detailJa).
		SetDevTime(devTime)

	if githubUrl != "" {
		c = c.SetGithubURL(githubUrl)
	}
	if thumbnail != "" {
		c = c.SetThumbnail(thumbnail)
	}

	prod, err := c.Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, prod)
}

// Update a product
//
// changeable form values:
// - name: string
// - name_ja: string
// - detail: string
// - detail_ja: string
// - site_url: string
// - github_url: Optional string
// - dev_time: ISO 8601 type date
// - thumbnail: Optional string
func (h *Handler) UpdateProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	// idStr := e.FormValue("product_id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (h *Handler) DeleteProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}
