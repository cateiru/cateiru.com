package handler

import (
	"context"
	"net/http"
	"strconv"
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

	if err := ValidateURL(siteUrl); err != nil {
		return err
	}
	if err := ValidateURL(githubUrl); err != nil {
		return err
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
// - github_url: string
// - dev_time: ISO 8601 type date
// - thumbnail: string
func (h *Handler) UpdateProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.FormValue("product_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	c := h.DB.Client.Product.
		Update().
		Where(product.ID(uint32(id)))

	changed := false

	name := e.FormValue("name")
	if name != "" {
		c = c.SetName(name)
		changed = true
	}
	nameJa := e.FormValue("name_ja")
	if nameJa != "" {
		c = c.SetNameJa(nameJa)
		changed = true
	}
	detail := e.FormValue("detail")
	if detail != "" {
		c = c.SetDetail(detail)
		changed = true
	}
	detailJa := e.FormValue("detail_ja")
	if detail != "" {
		c = c.SetDetailJa(detailJa)
		changed = true
	}
	siteUrl := e.FormValue("site_url")
	if siteUrl != "" {
		if err := ValidateURL(siteUrl); err != nil {
			return err
		}
		c = c.SetSiteURL(siteUrl)
		changed = true
	}
	githubUrl := e.FormValue("github_url")
	if githubUrl != "" {
		if err := ValidateURL(githubUrl); err != nil {
			return err
		}
		c = c.SetGithubURL(githubUrl)
		changed = true
	}
	devTimeStr := e.FormValue("dev_time")
	if devTimeStr != "" {
		devTime, err := time.Parse("2006-01-02T15:04:05-0700", devTimeStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: dev_time")
		}
		c = c.SetDevTime(devTime)
		changed = true
	}
	thumbnail := e.FormValue("thumbnail")
	if thumbnail != "" {
		c = c.SetThumbnail(thumbnail)
		changed = true
	}

	if !changed {
		return echo.NewHTTPError(http.StatusBadRequest, "no changes")
	}

	if err := c.Exec(ctx); err != nil {
		return err
	}

	prod, err := h.DB.Client.Product.
		Query().
		Where(product.ID(uint32(id))).
		First(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, prod)
}

func (h *Handler) DeleteProductHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("product_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: product_id")
	}

	_, err = h.DB.Client.Product.
		Delete().
		Where(product.ID(uint32(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
