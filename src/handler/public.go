package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/labstack/echo/v4"
)

type PublicUser struct {
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// GivenNameJa holds the value of the "given_name_ja" field.
	GivenNameJa string `json:"given_name_ja,omitempty"`
	// FamilyNameJa holds the value of the "family_name_ja" field.
	FamilyNameJa string `json:"family_name_ja,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// LocationJa holds the value of the "location_ja" field.
	LocationJa string `json:"location_ja,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`

	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
}

func (h Handler) PublicProfileHandler(e echo.Context) error {
	ctx := context.Background()

	u, err := h.DB.Client.User.Query().Where(user.Selected(true)).First(ctx)
	if _, ok := err.(*ent.NotFoundError); ok {
		return echo.ErrNotFound
	}

	publicUser := PublicUser{
		GivenName:    u.GivenName,
		FamilyName:   u.FamilyName,
		GivenNameJa:  u.GivenNameJa,
		FamilyNameJa: u.FamilyNameJa,
		UserID:       u.UserID,
		BirthDate:    u.BirthDate,
		Location:     u.Location,
		LocationJa:   u.LocationJa,
		AvatarURL:    u.AvatarURL,
		Created:      u.Created,
		Modified:     u.Modified,
	}

	return e.JSON(http.StatusOK, publicUser)
}
