package handler

import (
	"context"
	"net/http"

	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/labstack/echo/v4"
)

func (h *Handler) LocationHandler(e echo.Context) error {
	ctx := context.Background()

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
	ctx := context.Background()

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

func (h *Handler) UpdateLocationHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteLocationHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	return nil
}
