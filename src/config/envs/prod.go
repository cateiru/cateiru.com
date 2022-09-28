package envs

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var ProdConfig = ConfigDefs{
	Mode: "prod",

	ApiDomain: url.URL{
		Host:   "cateiru.com",
		Scheme: "https",
	},
	PageDomain: url.URL{
		Host:   "api.cateiru.com",
		Scheme: "https",
	},

	DBConfig: mysql.Config{
		DBName:    "cateirucom",
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
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
}
