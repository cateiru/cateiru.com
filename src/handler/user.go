package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h Handler) MeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Base.Session(ctx, e); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, h.Base.User)
}

func (h Handler) UpdateUserHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Base.Session(ctx, e); err != nil {
		return err
	}

	familyName := e.FormValue("family_name")
	givenName := e.FormValue("given_name")
	familyNameJa := e.FormValue("family_name_ja")
	givenNameJa := e.FormValue("given_name_ja")

	birthDateStr := e.FormValue("birth_date")

	location := e.FormValue("location")
	locationJa := e.FormValue("location_ja")

	u := h.DB.Client.User.Update()
	changeFlag := false

	if familyName != "" {
		u = u.SetFamilyName(familyName)
		changeFlag = true
	}
	if givenName != "" {
		u = u.SetGivenName(givenName)
		changeFlag = true
	}
	if familyNameJa != "" {
		u = u.SetFamilyNameJa(familyNameJa)
		changeFlag = true
	}
	if givenNameJa != "" {
		u = u.SetGivenNameJa(givenNameJa)
		changeFlag = true
	}

	if birthDateStr != "" {
		// parse ISO 8601 types
		birthDate, err := time.Parse("2006-01-02T15:04:05-0700", birthDateStr)
		if err == nil {
			u = u.SetBirthDate(birthDate)
			changeFlag = true
		}
	}

	if location != "" {
		u = u.SetLocation(location)
		changeFlag = true
	}
	if locationJa != "" {
		u = u.SetLocationJa(locationJa)
		changeFlag = true
	}

	if !changeFlag {
		return echo.ErrBadRequest
	}

	err := u.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
