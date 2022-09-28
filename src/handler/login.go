package handler

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cateiru/cateiru-sso/pkg/go/sso"
	"github.com/cateiru/cateiru.com/src/base"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	ctx := context.Background()
	base, err := base.NewBase(c)
	if err != nil {
		return err
	}

	code := c.QueryParams().Get("code")
	if code == "" {
		return echo.ErrBadRequest
	}

	token, err := sso.GetToken(code, config.Config.SSORedirectURI.String(), config.Config.SSOTokenSecret)
	if err != nil {
		logging.Sugar.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed get token from CateiruSSO")
	}

	claims, err := sso.ValidateIDToken(token.IDToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed valid JWT token")
	}

	user, err := base.DB.Client.User.Create().
		SetGivenName(claims.GivenName).
		SetFamilyName(claims.FamilyName).
		SetUserID(claims.Name).
		SetMail(claims.Email).
		SetBirthDate(time.Now()).
		SetGivenNameJa(claims.GivenName).
		SetFamilyNameJa(claims.FamilyName).
		SetLocation("").
		SetLocationJa("").
		Save(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed insert form user db")
	}

	if err := base.Login(ctx, user); err != nil {
		return err
	}

	redirectURL := url.URL{}

	return c.Redirect(http.StatusMovedPermanently, redirectURL.String())
}
