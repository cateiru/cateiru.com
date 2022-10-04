package handler

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/pkg/go/sso"
	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/user"
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
	defer base.Close()

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

	// Check role
	roles := strings.Split(claims.Role, " ")
	findAdmin := false
	for _, role := range roles {
		if role == "admin" {
			findAdmin = true
			break
		}
	}
	if !findAdmin {
		return echo.NewHTTPError(http.StatusForbidden, "you do not have access")
	}

	existUser, err := base.DB.Client.User.Query().Where(user.SSOToken(claims.ID)).Exist(ctx)
	if err != nil {
		return err
	}

	var u *ent.User

	if existUser {
		u, err = base.DB.Client.User.Query().Where(user.SSOToken(claims.ID)).First(ctx)
		if err != nil {
			return err
		}
	} else {
		userConf := base.DB.Client.User.Create().
			SetGivenName(claims.GivenName).
			SetFamilyName(claims.FamilyName).
			SetUserID(claims.NickName).
			SetMail(claims.Email).
			SetBirthDate(time.Now()).
			SetGivenNameJa(claims.GivenName).
			SetFamilyNameJa(claims.FamilyName).
			SetLocation("").
			SetSSOToken(claims.ID).
			SetLocationJa("")

		if claims.Picture != "" {
			userConf = userConf.SetAvatarURL(claims.Picture)
		}

		u, err = userConf.Save(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed insert form user db")
		}
	}

	if err := base.Login(ctx, u); err != nil {
		return err
	}

	redirectURL := url.URL{}

	return c.Redirect(http.StatusMovedPermanently, redirectURL.String())
}
