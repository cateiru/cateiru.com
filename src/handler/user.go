package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/labstack/echo/v4"
)

// Response login user data
func (h *Handler) MeHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Base.Session(ctx, e); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, h.Base.User)
}

// Update user profiles
//
// changeable profiles
// - family_name
// - given_name
// - family_name_ja
// - given_name_ja
// - birth_date
// - location
// - location_ja
func (h *Handler) UpdateUserHandler(e echo.Context) error {
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

// Response all users in users db
func (h *Handler) AllUsersHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	users, err := h.DB.Client.User.Query().All(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, users)
}

func (h *Handler) ChangeSelect(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	newSelectUserIdStr := e.FormValue("id")
	newSelectedUserId, err := strconv.Atoi(newSelectUserIdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id query is number type")
	}

	// clear select all users
	_, err = h.DB.Client.User.
		Update().
		Where(user.Selected(true)).
		SetSelected(false).
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = h.DB.Client.User.Update().
		Where(user.ID(uint32(newSelectedUserId))).
		SetSelected(true).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
