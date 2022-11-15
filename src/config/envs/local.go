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

var HOST = GetURL()

func GetURL() string {
	u := os.Getenv("URL")
	fmt.Println("URL: " + u)
	if u == "" {
		return "localhost"
	}
	return u
}

var LocalConfig = ConfigDefs{
	Mode: "local",

	ApiDomain: url.URL{
		Host:   fmt.Sprintf("%s:8080", HOST),
		Scheme: "http",
	},
	PageDomain: url.URL{
		Host:   fmt.Sprintf("%s:3000", HOST),
		Scheme: "http",
	},

	Cors: middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{fmt.Sprintf("http://%s:3000", HOST)},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}),

	// This config is docker-cmpose MySQL connection.
	DBConfig: mysql.Config{
		DBName:    "cateiru",
		User:      "docker",
		Passwd:    "docker",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
	},

	SessionCookieName: "cateirucom-session",
	DBSessionTime:     time.Hour * 5,
	SessionCookieConfig: http.Cookie{
		HttpOnly: true,
		Secure:   false, // ローカルであるためfalse
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	},
	SessionConfirmationCookieName: "cateirusso-issession",
	SessionConfirmationCookieConfig: http.Cookie{
		HttpOnly: false,
		Secure:   false, // ローカルであるためfalse
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	},

	SSOTokenSecret: "2974d92793c53756ec347fe2a8246fd9f91a2dde291147f081292907cc20b385",
	SSORedirectURI: url.URL{
		Host:   fmt.Sprintf("%s:8080", HOST),
		Scheme: "http",
		Path:   "/login",
	},
	SSOClientID: "e962e6d0db161d59ae5110736ae8ac",

	MailFromDomain:    "m.cateiru.com",
	MailgunAPIKey:     "123456789abcd",
	SenderMailAddress: "Cateiru <info@m.cateiru.com>",
}
