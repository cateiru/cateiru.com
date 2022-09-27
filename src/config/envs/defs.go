package envs

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-sql-driver/mysql"
)

type ConfigDefs struct {
	Mode string // Run mode. `local`, `test` or `prod`

	ApiDomain  url.URL
	PageDomain url.URL

	// DB Connect config
	DBConfig mysql.Config

	SessionCookieName             string
	SessionConfirmationCookieName string
	DBSessionTime                 time.Duration
	// Cookie config
	// Name and domain is overwrite other config.
	SessionCookieConfig             http.Cookie
	SessionConfirmationCookieConfig http.Cookie
}
