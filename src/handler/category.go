package handler

import (
	"context"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/labstack/echo/v4"
)

// Response all categories
func (h *Handler) CategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	c, err := h.DB.Client.Category.Query().All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, c)
}

func (h *Handler) CreateCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	name := e.FormValue("name")
	nameJa := e.FormValue("name_ja")
	emoji := e.FormValue("emoji")

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name")
	}
	if nameJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name_ja")
	}
	// emoji is one char
	if utf8.RuneCountInString(emoji) == 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: emoji")
	}

	c, err := h.DB.Client.Category.Create().
		SetName(name).
		SetNameJa(nameJa).
		SetEmoji(emoji).
		Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, c)
}

func (h *Handler) UpdateCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.FormValue("category_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: category_id")
	}

	c := h.DB.Client.Category.
		Update().
		Where(category.ID(uint32(id)))
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
	emoji := e.FormValue("emoji")
	if utf8.RuneCountInString(emoji) == 1 {
		c = c.SetEmoji(emoji)
		changed = true
	}

	if !changed {
		return echo.NewHTTPError(http.StatusBadRequest, "no changes")
	}

	if err = c.Exec(ctx); err != nil {
		return err
	}

	categoryDB, err := h.DB.Client.Category.Get(ctx, uint32(id))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, categoryDB)
}

func (h *Handler) DeleteCategoryHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("category_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params: category_id")
	}

	_, err = h.DB.Client.Category.
		Delete().
		Where(category.ID(uint32(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
