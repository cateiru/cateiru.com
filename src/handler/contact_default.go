package handler

import (
	"net/http"
	"strconv"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/contactdefault"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PublicContactDefaultHandler(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.QueryParam("id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is empty")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	defaultContact, err := h.DB.Client.ContactDefault.Get(ctx, uint32(id))
	if _, ok := err.(*ent.NotFoundError); ok {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, defaultContact)
}

func (h *Handler) ContactDefaultHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.Session(ctx, c); err != nil {
		return err
	}

	idStr := c.QueryParam("id")
	if idStr == "" {
		defaultContacts, err := h.DB.Client.ContactDefault.Query().All(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, defaultContacts)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params: category_id")
	}

	defaultContact, err := h.DB.Client.ContactDefault.Get(ctx, uint32(id))
	if _, ok := err.(*ent.NotFoundError); ok {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, defaultContact)
}

func (h *Handler) CreateContactDefaultHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.Session(ctx, c); err != nil {
		return err
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	url := c.FormValue("url")
	category := c.FormValue("category")
	customTitle := c.FormValue("custom_title")
	description := c.FormValue("description")

	d := h.DB.Client.ContactDefault.Create()

	if name != "" {
		d = d.SetName(name)
	}
	if email != "" {
		d = d.SetEmail(email)
	}
	if url != "" {
		d = d.SetURL(url)
	}
	if category != "" {
		d = d.SetCategory(category)
	}
	if customTitle != "" {
		d = d.SetCustomTitle(customTitle)
	}
	if description != "" {
		d = d.SetDescription(description)
	}

	contactDefault, err := d.Save(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, contactDefault)
}

func (h *Handler) UpdateContactDefaultHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.Session(ctx, c); err != nil {
		return err
	}

	idStr := c.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: category_id")
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	url := c.FormValue("url")
	category := c.FormValue("category")
	customTitle := c.FormValue("custom_title")
	description := c.FormValue("description")

	d := h.DB.Client.ContactDefault.Update().Where(contactdefault.ID(uint32(id)))

	if name != "" {
		d = d.SetName(name)
	}
	if email != "" {
		d = d.SetEmail(email)
	}
	if url != "" {
		d = d.SetURL(url)
	}
	if category != "" {
		d = d.SetCategory(category)
	}
	if customTitle != "" {
		d = d.SetCustomTitle(customTitle)
	}
	if description != "" {
		d = d.SetDescription(description)
	}

	contactDefault, err := d.Save(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, contactDefault)
}

func (h *Handler) DeleteContactDefaultHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := h.Session(ctx, c); err != nil {
		return err
	}

	idStr := c.QueryParam("id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is empty")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params: category_id")
	}

	_, err = h.DB.Client.ContactDefault.
		Delete().
		Where(contactdefault.ID(uint32(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
