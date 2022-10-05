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
	E  echo.Context
	DB *db.DB

	User *ent.User
}

// Create a Base objects
func NewBase(e echo.Context) (*Base, error) {
	db, err := db.NewConnectMySQL()
	if err != nil {
		return nil, err
	}

	return &Base{
		E:    e,
		DB:   db,
		User: nil,
	}, nil
}

func (c *Base) Close() {
	c.DB.Close()
}

// Require session
func (c *Base) Session(ctx context.Context) error {
	token, err := c.getSessionToken()
	if err != nil {
		return err
	}
	u, err := c.sessionLogin(ctx, *c.DB.Client, token)
	if err != nil {
		return err
	}

	c.User = u

	return nil
}

func (c *Base) Login(ctx context.Context, u *ent.User) error {
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
		Domain: config.Config.PageDomain.Path,

		// FIXME: want to deepcopy and use
		HttpOnly: config.Config.SessionCookieConfig.HttpOnly,
		Secure:   config.Config.SessionCookieConfig.Secure,
		SameSite: config.Config.SessionCookieConfig.SameSite,
		Path:     config.Config.SessionCookieConfig.Path,
	}
	c.E.SetCookie(sessionCookie)

	confirmCookie := &http.Cookie{
		Name:   config.Config.SessionConfirmationCookieName,
		Value:  "true",
		Domain: config.Config.PageDomain.Path,

		// FIXME: want to deepcopy and use
		HttpOnly: config.Config.SessionConfirmationCookieConfig.HttpOnly,
		Secure:   config.Config.SessionConfirmationCookieConfig.Secure,
		SameSite: config.Config.SessionConfirmationCookieConfig.SameSite,
		Path:     config.Config.SessionConfirmationCookieConfig.Path,
	}
	c.E.SetCookie(confirmCookie)

	c.User = u

	return nil
}

// Logout
func (c *Base) Logout(ctx context.Context) error {
	tokenCookie, err := c.E.Cookie(config.Config.SessionCookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed login: %v", err))
	}
	tokenCookie.Expires = time.Now()
	tokenCookie.MaxAge = 0
	c.E.SetCookie(tokenCookie)

	checkCookie, err := c.E.Cookie(config.Config.SessionConfirmationCookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed login: %v", err))
	}
	checkCookie.Expires = time.Now()
	checkCookie.MaxAge = 0
	c.E.SetCookie(checkCookie)

	return nil
}

// Get Session token from Cookies
// Cookie session token name is defined `SessionCookieName` in config
//
// If don't get values or invalid value, return 403 error.
func (c *Base) getSessionToken() (uuid.UUID, error) {
	tokenCookie, err := c.E.Cookie(config.Config.SessionCookieName)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(http.StatusBadRequest, "cookie is not found")
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
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, echo.NewHTTPError(http.StatusForbidden, "session token is empty")
	}

	user, err := client.User.Get(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
