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

	Cors: nil,

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
	SessionConfirmationCookieName: "cateirucom-issession",
	SessionConfirmationCookieConfig: http.Cookie{
		HttpOnly: false,
		Secure:   false, // ローカルであるためfalse
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	},

	SSOTokenSecret: "12345",
	SSORedirectURI: url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
		Path:   "/login",
	},
	SSOClientID: "12345",

	MailFromDomain:    "m.cateiru.com",
	MailgunAPIKey:     "123456789abcd",
	SenderMailAddress: "Cateiru <info@m.cateiru.com>",
}
