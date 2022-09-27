package envs

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-sql-driver/mysql"
)

var LocalConfig = ConfigDefs{
	Mode: "local",

	ApiDomain: url.URL{
		Host: "localhost:8080",
	},
	PageDomain: url.URL{
		Host: "localhost:3000",
	},

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
}
