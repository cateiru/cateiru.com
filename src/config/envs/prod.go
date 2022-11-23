package envs

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ProdConfig = ConfigDefs{
	Mode: "prod",

	ApiDomain: url.URL{
		Host:   "api.cateiru.com",
		Scheme: "https",
	},
	PageDomain: url.URL{
		Host:   "cateiru.com",
		Scheme: "https",
	},

	Cors: middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"cateiru.com"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}),

	DBConfig: mysql.Config{
		DBName:               "cateirucom",
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 fmt.Sprintf("/cloudsql/%s", os.Getenv("INSTANCE_CONNECTION_NAME")),
		Net:                  "unix",
		ParseTime:            true,
		AllowNativePasswords: true,
	},

	SessionCookieName: "cateirucom-session",
	DBSessionTime:     time.Hour * 5,
	SessionCookieConfig: http.Cookie{
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	},
	SessionConfirmationCookieName: "cateirusso-issession",
	SessionConfirmationCookieConfig: http.Cookie{
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	},

	SSOTokenSecret: os.Getenv("SSO_TOKEN_SECRET"),
	SSORedirectURI: url.URL{
		Host:   "api.cateiru.com",
		Scheme: "https",
		Path:   "/login",
	},
	SSOClientID: os.Getenv("SSO_CLIENT_ID"),

	MailFromDomain:    "m.cateiru.com",
	MailgunAPIKey:     os.Getenv("MAILGUN_APIKEY"),
	SenderMailAddress: "Cateiru <info@m.cateiru.com>",
}
