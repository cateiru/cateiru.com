package handler

import (
	"net/http"
	"strconv"

	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/labstack/echo/v4"
)

func (h *Handler) LocationHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	locations, err := h.DB.Client.Location.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, locations)
}

// create a new location
//
// - type: `univ` or `corp`
// - name: string
// - name_ja: string
// - address: string
// - address_ja: string
func (h *Handler) CreateLocationHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	t := e.FormValue("type")
	name := e.FormValue("name")
	nameJa := e.FormValue("name_ja")
	address := e.FormValue("address")
	addressJa := e.FormValue("address_ja")

	if t == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: type")
	}
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: name")
	}
	if nameJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: nameJa")
	}
	if address == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: address")
	}
	if addressJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: addressJa")
	}

	// t must be `univ` or `corp`
	if t != "univ" && t != "corp" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: type")
	}

	location, err := h.DB.Client.Location.
		Create().
		SetType(location.Type(t)).
		SetName(name).
		SetNameJa(nameJa).
		SetAddress(address).
		SetAddressJa(addressJa).
		Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, location)
}

// Update location
//
// - type: `univ` or `corp`
// - name: string
// - name_ja: string
// - address: string
// - address_ja: string
func (h *Handler) UpdateLocationHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.FormValue("location_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: location_id")
	}
	changed := false

	t := e.FormValue("type")
	name := e.FormValue("name")
	nameJa := e.FormValue("name_ja")
	address := e.FormValue("address")
	addressJa := e.FormValue("address_ja")

	u := h.DB.Client.Location.
		Update().
		Where(location.ID(uint32(id)))

	if t != "" {
		// t must be `univ` or `corp`
		if t != "univ" && t != "corp" {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: type")
		}

		u = u.SetType(location.Type(t))
		changed = true
	}
	if name != "" {
		u = u.SetName(name)
		changed = true
	}
	if nameJa != "" {
		u = u.SetNameJa(nameJa)
		changed = true
	}
	if address != "" {
		u = u.SetAddress(address)
		changed = true
	}
	if addressJa != "" {
		u = u.SetAddressJa(addressJa)
		changed = true
	}

	if !changed {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: no changes")
	}

	if err := u.Exec(ctx); err != nil {
		return err
	}

	locationDB, err := h.DB.Client.Location.
		Query().
		Where(location.ID(uint32(id))).
		First(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, locationDB)
}

func (h *Handler) DeleteLocationHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("location_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: location_id")
	}

	exist, err := h.DB.Client.Biography.
		Query().
		Where(biography.LocationID(uint32(id))).
		Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		return echo.NewHTTPError(http.StatusBadRequest, "This bio is already used")
	}

	_, err = h.DB.Client.Location.
		Delete().
		Where(location.ID(uint32(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
