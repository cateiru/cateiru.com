package envs

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-sql-driver/mysql"
)

var TestConfig = ConfigDefs{
	Mode: "test",

	ApiDomain: url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	},
	PageDomain: url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	},

	DBConfig: mysql.Config{
		DBName:    "test",
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
		Host:   "localhost:8080",
		Scheme: "http",
		Path:   "/login",
	},
	SSOClientID: "e962e6d0db161d59ae5110736ae8ac",
}
