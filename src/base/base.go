package base

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Base struct {
	DB   *db.DB
	User *ent.User
}

// Create a Base objects
func NewBase(db *db.DB) (*Base, error) {
	return &Base{
		DB:   db,
		User: nil,
	}, nil
}

// Require session
func (c *Base) Session(ctx context.Context, e echo.Context) error {
	token, err := c.getSessionToken(e)
	if err != nil {
		c.Logout(e)
		return err
	}
	u, err := c.sessionLogin(ctx, *c.DB.Client, token)
	if err != nil {
		c.Logout(e)
		return err
	}

	c.User = u

	return nil
}

func (c *Base) Login(ctx context.Context, e echo.Context, u *ent.User) error {
	session, err := c.DB.Client.Session.
		Create().
		SetUserID(u.ID).
		SetPeriod(time.Now().Add(config.Config.DBSessionTime)).
		Save(ctx)
	if err != nil {
		return err
	}

	sessionCookie := &http.Cookie{
		Name:   config.Config.SessionCookieName,
		Value:  session.ID.String(),
		Domain: config.Config.PageDomain.Host,

		// FIXME: want to deepcopy and use
		HttpOnly: config.Config.SessionCookieConfig.HttpOnly,
		Secure:   config.Config.SessionCookieConfig.Secure,
		SameSite: config.Config.SessionCookieConfig.SameSite,
		Path:     config.Config.SessionCookieConfig.Path,
	}
	e.SetCookie(sessionCookie)

	confirmCookie := &http.Cookie{
		Name:   config.Config.SessionConfirmationCookieName,
		Value:  "true",
		Domain: config.Config.PageDomain.Host,

		// FIXME: want to deepcopy and use
		HttpOnly: config.Config.SessionConfirmationCookieConfig.HttpOnly,
		Secure:   config.Config.SessionConfirmationCookieConfig.Secure,
		SameSite: config.Config.SessionConfirmationCookieConfig.SameSite,
		Path:     config.Config.SessionConfirmationCookieConfig.Path,
	}
	e.SetCookie(confirmCookie)

	c.User = u

	return nil
}

// Logout
func (c *Base) Logout(e echo.Context) {
	tokenCookie, err := e.Cookie(config.Config.SessionCookieName)
	if err == nil {
		tokenCookie.Domain = config.Config.PageDomain.Host
		tokenCookie.HttpOnly = config.Config.SessionCookieConfig.HttpOnly
		tokenCookie.Secure = config.Config.SessionCookieConfig.Secure
		tokenCookie.SameSite = config.Config.SessionCookieConfig.SameSite
		tokenCookie.Path = config.Config.SessionCookieConfig.Path

		tokenCookie.Expires = time.Now().Add(-1 * time.Hour)
		tokenCookie.MaxAge = -1
		e.SetCookie(tokenCookie)
	}

	checkCookie, err := e.Cookie(config.Config.SessionConfirmationCookieName)
	if err == nil {
		checkCookie.Domain = config.Config.PageDomain.Host
		checkCookie.HttpOnly = config.Config.SessionConfirmationCookieConfig.HttpOnly
		checkCookie.Secure = config.Config.SessionConfirmationCookieConfig.Secure
		checkCookie.SameSite = config.Config.SessionConfirmationCookieConfig.SameSite
		checkCookie.Path = config.Config.SessionConfirmationCookieConfig.Path

		checkCookie.Expires = time.Now().Add(-1 * time.Hour)
		checkCookie.MaxAge = -1
		e.SetCookie(checkCookie)
	}
}

// Get Session token from Cookies
// Cookie session token name is defined `SessionCookieName` in config
//
// If don't get values or invalid value, return 403 error.
func (c *Base) getSessionToken(e echo.Context) (uuid.UUID, error) {
	tokenCookie, err := e.Cookie(config.Config.SessionCookieName)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("cookie is not found: %v", err))
	}
	token := tokenCookie.Value

	u, err := uuid.Parse(token)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusForbidden, "failed parse session token uuid")
	}

	return u, nil
}

// Login from session token
func (c *Base) sessionLogin(ctx context.Context, client ent.Client, sessionToken uuid.UUID) (*ent.User, error) {
	session, err := client.Session.Get(ctx, sessionToken)
	if _, ok := err.(*ent.NotFoundError); ok {
		return nil, echo.NewHTTPError(http.StatusForbidden, "session token invalid")
	}
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, echo.NewHTTPError(http.StatusForbidden, "session token is empty")
	}

	if time.Until(session.Period) < 0 {
		return nil, echo.NewHTTPError(http.StatusForbidden, "session token is empty")
	}

	user, err := client.User.Get(ctx, session.UserID)
	if _, ok := err.(*ent.NotFoundError); ok {
		return nil, echo.NewHTTPError(http.StatusForbidden, "no user")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
