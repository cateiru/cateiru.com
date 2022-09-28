package base

import (
	"context"
	"net/http"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src"
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
	token, err := c.getSessionToken(c.E)
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

// Get Session token from Cookies
// Cookie session token name is defined `SessionCookieName` in config
//
// If don't get values or invalid value, return 403 error.
func (c *Base) getSessionToken(e echo.Context) (uuid.UUID, error) {
	tokenCookie, err := e.Cookie(config.Config.SessionCookieName)
	if err != nil {
		return uuid.UUID{}, err
	}
	token := tokenCookie.Value

	u, err := uuid.Parse(token)
	if err != nil {
		src.Sugar.Error(err)
		return uuid.UUID{}, echo.ErrForbidden
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
		src.Sugar.Error(err)
		return nil, echo.ErrForbidden
	}

	user, err := client.User.Get(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
