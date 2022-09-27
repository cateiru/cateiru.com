package base

import (
	"context"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SessionBase struct {
	User ent.User
	Base
}

// Create new user session.
func NewSession(ctx context.Context, e echo.Context) (*SessionBase, error) {
	base, err := NewBase(e)
	if err != nil {
		return nil, err
	}

	token, err := getSessionToken(e)
	if err != nil {
		return nil, err
	}
	u, err := sessionLogin(ctx, *base.DB.Client, token)
	if err != nil {
		return nil, err
	}

	return &SessionBase{
		User: *u,
		Base: *base,
	}, nil
}

// Get Session token from Cookies
// Cookie session token name is defined `SessionCookieName` in config
//
// If don't get values or invalid value, return 403 error.
func getSessionToken(e echo.Context) (uuid.UUID, error) {
	tokenCookie, err := e.Cookie(config.Config.SessionCookieName)
	if err != nil {
		return uuid.UUID{}, err
	}
	token := tokenCookie.Value

	u, err := uuid.FromBytes([]byte(token))
	if err != nil {
		return uuid.UUID{}, echo.ErrForbidden
	}

	return u, nil
}

// Login from session token
func sessionLogin(ctx context.Context, client ent.Client, sessionToken uuid.UUID) (*ent.User, error) {
	session, err := client.Session.Get(ctx, sessionToken)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, echo.ErrForbidden
	}

	user, err := client.User.Get(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
