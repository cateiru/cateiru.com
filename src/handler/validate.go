package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

var UrlRegexp = regexp.MustCompile(
	`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`,
)

func ValidateURL(target string) error {
	ok := UrlRegexp.MatchString(target)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid args: %s", target))
	}
	return nil
}
