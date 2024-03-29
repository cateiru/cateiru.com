package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cateiru/cateiru-sso/pkg/go/sso"
	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/cateiru/cateiru.com/ent/session"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/golang-jwt/jwt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const JWT_PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIG5AIBAAKCAYEA1dWsixCX0Bky4PK43Y5OMsQ32wQY+FsV2iXpnrfdITeFqH3H
BB+wQ2EeV5QVAnVM3o4j52ClEgmLm97+yIERd+UOLOrCC0eFmMXcxTH/DYt9tg4k
87piFXrDcfGtNTVCaMIqbHBtEC9dPwtNewxjiCYlkQDOW226+1AciFZX2DPDDjQB
Q9ruH8TJcQpj3je3NQKmBr2JIXaDAc2L94gKK6nuTIMFiEWjXHuKFmk3NI51xOyZ
nFfmbkNOFQsjxynfFoxgZE40DEJBdIMpNgdmSDCNR52L0LCJOs3fynNL8NnVU77C
rNxywhwlcPKnzCaU3czf5jQJmf6BaNmZi6bkye39oqYOBiZIGoa6uZsCMTBpOp4H
uL519QQOzzi7z8UF/9DfGfX6pO3IL0hy9lkxj/RsYKIUpuPFBzJLx47RQGFIALXn
yRsbp36Ffsbd6bTjh7ZyXZjhKKCkSDFM37ULF1W/mLnMfgQaKg8Xbuyps0ySTWO8
olwZwbojzfKq7nB9AgMBAAECggGAO+o6BWEp2HInEmaQK+witxDJwcFdKcGD1vMG
iaVk/VisR7CuYdZrMgX1VX61gFHTrwxOvRcUSYCJMKyD4TRg66IvnaBNrJFmuo6+
NDa1C1uJZsiBzBTuRKx4NOYNrTn2GJpG/slllcJfszl8hTzMaqqdngqGx5FwcdhZ
lk3zIS4nukQSEqw2Sop+EwFhSDN+Twkl8f+Le+fT/15TmJDyoJdloOSZWTyB8Hjh
ETR79mrwTrQoumjatQyAxV6wsK5nQ0fWXFAsm/Z5fnkGH41iy+LC00K5CiA93cpf
0EcM+RUGCkim7Hy2ykPF7k8q+z73D3ijXmAIjYK6OD4B1BIFoM/5w/Z504Apu4IF
cgFf8IFP95M9wz1Ox7HPJqwViDmmClfqXQu1uf0Ss6LWg79ylZtGp8AM8ekBeUsb
hbNQobdSh6FKkfmmDs4sv75HG3HSsSVc2FDWiXK5WhzySFIM+i0dWxLySZMlyRZf
dJJvoZiHMsdbQGohLzm7R9HGWg+hAoHBAP1NOnS2viKnb1WWkP/5MzewB0dIJqjM
9otRCvn77P5iogmvlpyR+1G5RBPYXGnlLCnlTs4ZU8Q9FJz95k/GM1crs1/io7c4
Jz38PzNeCT300K9OXn/rf7FQMjQ5hESoYntw07/YfI+zX6DBtYTmd6ugydofjnV7
guQEIiaT7gR9InWuu9MCeAXPGs5sqA4GKPEbFIcewiAuyJQ7ycLNxAYJMYcb1pih
cy13TKXo1M0uHDWIjtAzED3wmYpP8avxeQKBwQDYHND6FPIACOg31/eoXX3QDwFa
BdfZ5g363RmXK5yQmWyKf8yX2+Vg7oLs+ScInXi6mbjXCSVCKFBbYODVhj9mDNpU
KI+LloT2h9IHQ+4JM07sRLn8g5TJ1d297hSd38EdNjmmn3DCa8OfuNsoSTDn956/
PX4CmAPfVb5wRa1IE4r6oPsYGHGMhPjuSs69Z0r5+mcrOrD9+KvU4jnkAVfgmgEQ
7gFSCTqs2QdfuDLOOMVfL1dBy43L9h3Zx7+wWiUCgcEA7On7v8hgY+c48dNnAZOi
PPauiBRK3F0AR+PgIVppGrFEH00U0JKVfqWKsAkQvpbzzKfGImx79bKvgfrUHE8k
G+cq2YcQW0DxKz+wGSpd+I3vVdg1+O1aPIzmuAQH23Om0RABbZFdR8acra4gShKJ
YYR85z/Wrbl8/imDi5OW7qnfvjRVRpIrkjtYNjRYpu1KA+CxPYQeS52WS1b0afis
3Hxiz5zRBWcUZDUOAquXxeTXy5ZTBRaNnXFZGo2VW535AoHBALR/tOb6sXjn95QU
tEuR8m/g6H8Y3ESpCcoZ+rKxAS4ghnBS9z1qsvU3oHBuVHS0mU448BHiGJVd+Gep
zX9phfAlEgEyc1nh2KHTpM7epKRYjzV03WYTPCUrk+17OQr3BGLylEofZt1rhlMT
4S+PmeJUhekyYdmmv607/zfkaisF75EO/UkNuWDk+siEAJJfAFczFhIcDhHBLTs7
y6Asown7dqJ+NuoevTv6dFc1EHH/JpIuhxF0ArV0lU/8rpSLZQKBwAlZ0vUOe+VM
i/cq3TFNnFh0sJWMMVLTKhzJISw9n1XCje3+htoJDCZ38uHCkeLgA5v8v/vQ+ytC
7hdwLPHSMkPVyHeSrMSpOXgjB7ehZIca+OsGQLVuusS3TYPMRoKgEv6FYTpYsXVK
jyWd+eLWVT3ziFk2w6H2eM9/WJPaNDUQwnNdcYjLJBZ3H4HhfvL+uNTeRNzK5PQe
4uheKTqPvxrpyvUOEYipS9FX8FjoZQUo4hn5tHHVIQjL6w0thKwi7Q==
-----END RSA PRIVATE KEY-----`

const JWT_PUBLIC_PKCS8_KEY = `-----BEGIN PUBLIC KEY-----
MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA1dWsixCX0Bky4PK43Y5O
MsQ32wQY+FsV2iXpnrfdITeFqH3HBB+wQ2EeV5QVAnVM3o4j52ClEgmLm97+yIER
d+UOLOrCC0eFmMXcxTH/DYt9tg4k87piFXrDcfGtNTVCaMIqbHBtEC9dPwtNewxj
iCYlkQDOW226+1AciFZX2DPDDjQBQ9ruH8TJcQpj3je3NQKmBr2JIXaDAc2L94gK
K6nuTIMFiEWjXHuKFmk3NI51xOyZnFfmbkNOFQsjxynfFoxgZE40DEJBdIMpNgdm
SDCNR52L0LCJOs3fynNL8NnVU77CrNxywhwlcPKnzCaU3czf5jQJmf6BaNmZi6bk
ye39oqYOBiZIGoa6uZsCMTBpOp4HuL519QQOzzi7z8UF/9DfGfX6pO3IL0hy9lkx
j/RsYKIUpuPFBzJLx47RQGFIALXnyRsbp36Ffsbd6bTjh7ZyXZjhKKCkSDFM37UL
F1W/mLnMfgQaKg8Xbuyps0ySTWO8olwZwbojzfKq7nB9AgMBAAE=
-----END PUBLIC KEY-----`

func TestLoginHandler(t *testing.T) {
	test.Init()

	code := "abc123"
	redirect := config.Config.SSORedirectURI.String()

	testUser, err := test.NewUser()
	require.NoError(t, err)

	claims := sso.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test",
			Subject:   "hoge",
			Audience:  "nya",
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},

		Name:       testUser.FamilyName,
		GivenName:  testUser.GivenName,
		FamilyName: testUser.FamilyName,
		NickName:   testUser.UserId,
		Email:      testUser.Mail,
		Picture:    testUser.AvatarURL,

		ID:   testUser.SSOToken,
		Role: "admin",

		Iat:      time.Now().Unix(),
		AuthTime: time.Now().Unix(),

		PreferredUserName: "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(JWT_PRIVATE_KEY))
	require.NoError(t, err)

	idToken, err := token.SignedString(signKey)
	require.NoError(t, err)

	// mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	key := sso.PublicKey{
		Pkcs8: JWT_PUBLIC_PKCS8_KEY,
	}
	body, err := json.Marshal(key)
	require.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://api.sso.cateiru.com/v1/oauth/jwt/key",
		httpmock.NewStringResponder(200, string(body)),
	)

	successResp := sso.TokenResponse{
		AccessToken:  "hogehoge",
		TokenType:    "Bearer",
		RefreshToken: "hugahuga",
		ExpiresIn:    3600,
		IDToken:      idToken,
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf(
		"https://api.sso.cateiru.com/v1/oauth/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
		url.QueryEscape(code), url.QueryEscape(redirect),
	),
		func(req *http.Request) (*http.Response, error) {
			secret := strings.Split(req.Header.Get("Authorization"), " ")
			if secret[1] != config.Config.SSOTokenSecret {
				return httpmock.NewStringResponse(403, ""), nil
			}

			resp, err := httpmock.NewJsonResponse(200, successResp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		base, err := test.NewTestTool()
		require.NoError(t, err)
		defer base.Close()

		m, err := easy.NewMock(fmt.Sprintf("/login?code=%s", code), http.MethodGet, "")
		require.NoError(t, err)

		h, err := base.Handler()
		require.NoError(t, err)

		e := m.Echo()

		err = h.LoginHandler(e)
		require.NoError(t, err)

		cookies := m.SetCookies()

		sessionToken := ""
		for _, cookie := range cookies {
			if cookie.Name == config.Config.SessionCookieName {
				sessionToken = cookie.Value
			}
		}
		require.NotEqual(t, sessionToken, "")

		// SELECT * FROM `users` WHERE users.id = (SELECT user_id FROM `sessions` WHERE session_token = ?)
		users, err := base.DB.Client.User.Query().Where(func(s *sql.Selector) {
			t := sql.Table(session.Table)
			s.Where(
				sql.EQ(
					s.C(user.FieldID),
					sql.Select(
						t.C(session.FieldUserID),
					).From(t).Where(sql.EQ(
						t.C(session.FieldID),
						sessionToken,
					)),
				),
			)
		}).All(ctx)
		require.NoError(t, err)
		require.Len(t, users, 1)
		dbInsertedUser := users[0]

		require.Equal(t, dbInsertedUser.FamilyName, testUser.FamilyName)
		require.Equal(t, dbInsertedUser.UserID, testUser.UserId)

		nExist, err := base.DB.Client.Notice.Query().Where(notice.UserID(users[0].ID)).Exist(ctx)
		require.NoError(t, err)
		require.True(t, nExist)
	})
}

func TestLoginAlreadyExistUser(t *testing.T) {
	test.Init()

	code := "abc123"
	redirect := config.Config.SSORedirectURI.String()

	testUser, err := test.NewUser()
	require.NoError(t, err)

	claims := sso.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test",
			Subject:   "hoge",
			Audience:  "nya",
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},

		Name:       testUser.FamilyName,
		GivenName:  testUser.GivenName,
		FamilyName: testUser.FamilyName,
		NickName:   testUser.UserId,
		Email:      testUser.Mail,
		Picture:    testUser.AvatarURL,

		ID:   testUser.SSOToken,
		Role: "admin",

		Iat:      time.Now().Unix(),
		AuthTime: time.Now().Unix(),

		PreferredUserName: "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(JWT_PRIVATE_KEY))
	require.NoError(t, err)

	idToken, err := token.SignedString(signKey)
	require.NoError(t, err)

	// mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	key := sso.PublicKey{
		Pkcs8: JWT_PUBLIC_PKCS8_KEY,
	}
	body, err := json.Marshal(key)
	require.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://api.sso.cateiru.com/v1/oauth/jwt/key",
		httpmock.NewStringResponder(200, string(body)),
	)

	successResp := sso.TokenResponse{
		AccessToken:  "hogehoge",
		TokenType:    "Bearer",
		RefreshToken: "hugahuga",
		ExpiresIn:    3600,
		IDToken:      idToken,
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf(
		"https://api.sso.cateiru.com/v1/oauth/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
		url.QueryEscape(code), url.QueryEscape(redirect),
	),
		func(req *http.Request) (*http.Response, error) {
			secret := strings.Split(req.Header.Get("Authorization"), " ")
			if secret[1] != config.Config.SSOTokenSecret {
				return httpmock.NewStringResponse(403, ""), nil
			}

			resp, err := httpmock.NewJsonResponse(200, successResp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		base, err := test.NewTestTool()
		require.NoError(t, err)
		defer base.Close()

		// Insert DB
		testUser.CreateDB(ctx, base.DB)

		m, err := easy.NewMock(fmt.Sprintf("/login?code=%s", code), http.MethodGet, "")
		require.NoError(t, err)

		h, err := base.Handler()
		require.NoError(t, err)

		e := m.Echo()

		err = h.LoginHandler(e)
		require.NoError(t, err)

		cookies := m.SetCookies()

		sessionToken := ""
		for _, cookie := range cookies {
			if cookie.Name == config.Config.SessionCookieName {
				sessionToken = cookie.Value
			}
		}
		require.NotEqual(t, sessionToken, "")

		// SELECT * FROM `users` WHERE users.id = (SELECT user_id FROM `sessions` WHERE session_token = ?)
		users, err := base.DB.Client.User.Query().Where(func(s *sql.Selector) {
			t := sql.Table(session.Table)
			s.Where(
				sql.EQ(
					s.C(user.FieldID),
					sql.Select(
						t.C(session.FieldUserID),
					).From(t).Where(sql.EQ(
						t.C(session.FieldID),
						sessionToken,
					)),
				),
			)
		}).All(ctx)
		require.NoError(t, err)
		require.Len(t, users, 1)
		dbInsertedUser := users[0]

		require.Equal(t, dbInsertedUser.FamilyName, testUser.FamilyName)
		require.Equal(t, dbInsertedUser.UserID, testUser.UserId)

		user2, err := base.DB.Client.User.Query().Where(user.SSOToken(testUser.SSOToken)).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, user2, 1)
	})
}

func TestLoginNoAdminUser(t *testing.T) {
	test.Init()

	code := "abc123"
	redirect := config.Config.SSORedirectURI.String()

	testUser, err := test.NewUser()
	require.NoError(t, err)

	claims := sso.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test",
			Subject:   "hoge",
			Audience:  "nya",
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},

		Name:       testUser.FamilyName,
		GivenName:  testUser.GivenName,
		FamilyName: testUser.FamilyName,
		NickName:   testUser.UserId,
		Email:      testUser.Mail,
		Picture:    testUser.AvatarURL,

		ID:   testUser.SSOToken,
		Role: "pro user", // no admin user

		Iat:      time.Now().Unix(),
		AuthTime: time.Now().Unix(),

		PreferredUserName: "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(JWT_PRIVATE_KEY))
	require.NoError(t, err)

	idToken, err := token.SignedString(signKey)
	require.NoError(t, err)

	// mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	key := sso.PublicKey{
		Pkcs8: JWT_PUBLIC_PKCS8_KEY,
	}
	body, err := json.Marshal(key)
	require.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://api.sso.cateiru.com/v1/oauth/jwt/key",
		httpmock.NewStringResponder(200, string(body)),
	)

	successResp := sso.TokenResponse{
		AccessToken:  "hogehoge",
		TokenType:    "Bearer",
		RefreshToken: "hugahuga",
		ExpiresIn:    3600,
		IDToken:      idToken,
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf(
		"https://api.sso.cateiru.com/v1/oauth/token?grant_type=authorization_code&code=%s&redirect_uri=%s",
		url.QueryEscape(code), url.QueryEscape(redirect),
	),
		func(req *http.Request) (*http.Response, error) {
			secret := strings.Split(req.Header.Get("Authorization"), " ")
			if secret[1] != config.Config.SSOTokenSecret {
				return httpmock.NewStringResponse(403, ""), nil
			}

			resp, err := httpmock.NewJsonResponse(200, successResp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	t.Run("success", func(t *testing.T) {
		base, err := test.NewTestTool()
		require.NoError(t, err)
		defer base.Close()

		m, err := easy.NewMock(fmt.Sprintf("/login?code=%s", code), http.MethodGet, "")
		require.NoError(t, err)

		h, err := base.Handler()
		require.NoError(t, err)

		e := m.Echo()

		err = h.LoginHandler(e)
		require.Error(t, err)
	})
}

func TestCheckAdminRole(t *testing.T) {
	codes := map[string]bool{
		"admin":                         true,
		"admin test":                    true,
		"local test":                    false,
		"local":                         false,
		"local hogehoge user admin nya": true,
	}

	for roles, check := range codes {
		err := handler.CheckAdminRole(roles)
		if check {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}
