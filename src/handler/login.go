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

// Login from OIDC
// Require `?code=` param.
// - code param is Open ID Connect code.
func (h *Handler) LoginHandler(c echo.Context) error {
	ctx := context.Background()

	code := c.QueryParams().Get("code")
	if code == "" {
		return echo.ErrBadRequest
	}

	claims, err := GetClaims(code)
	if err != nil {
		return err
	}

	// Check role
	if err := CheckAdminRole(claims.Role); err != nil {
		return err
	}

	existUser, err := h.DB.Client.User.Query().Where(user.SSOToken(claims.ID)).Exist(ctx)
	if err != nil {
		return err
	}

	var u *ent.User

	if existUser {
		u, err = h.DB.Client.User.Query().Where(user.SSOToken(claims.ID)).First(ctx)
		if err != nil {
			return err
		}
	} else {
		u, err = CreateUser(ctx, h.Base, claims)
		if err != nil {
			return err
		}
	}

	if err := h.Base.Login(ctx, c, u); err != nil {
		return err
	}

	// TODO
	redirectURL := url.URL{
		Host:   config.Config.PageDomain.Host,
		Scheme: config.Config.PageDomain.Scheme,
		Path:   "/admin",
	}

	return c.Redirect(http.StatusMovedPermanently, redirectURL.String())
}

// Get Claims from sso code via CateiruSSO
func GetClaims(code string) (*sso.Claims, error) {
	token, err := sso.GetToken(code, config.Config.SSORedirectURI.String(), config.Config.SSOTokenSecret)
	if err != nil {
		logging.Sugar.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed get token from CateiruSSO")
	}

	claims, err := sso.ValidateIDToken(token.IDToken)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "failed valid JWT token")
	}

	return claims, nil
}

// Check exists `admin` role in CateiruSSO received claims.
func CheckAdminRole(rolesStr string) error {
	roles := strings.Split(rolesStr, " ")
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
	return nil
}

// Create New User from SSO Claims
func CreateUser(ctx context.Context, base *base.Base, claims *sso.Claims) (*ent.User, error) {
	userConf := base.DB.Client.User.Create().
		SetGivenName(claims.GivenName).
		SetFamilyName(claims.FamilyName).
		SetUserID(claims.NickName).
		SetMail(claims.Email).
		SetBirthDate(time.Now()).
		SetGivenNameJa(claims.GivenName).
		SetFamilyNameJa(claims.FamilyName).
		SetLocation("todo").
		SetSSOToken(claims.ID).
		SetLocationJa("todo")

	if claims.Picture != "" {
		userConf = userConf.SetAvatarURL(claims.Picture)
	}

	u, err := userConf.Save(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed insert form user db")
	}

	_, err = base.DB.Client.Notice.Create().
		SetUserID(u.ID).
		Save(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed insert form notice db")
	}

	return u, nil
}
