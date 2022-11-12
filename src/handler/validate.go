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

var MailRegexp = regexp.MustCompile(
	`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`,
)

func ValidateURL(target string) error {
	ok := UrlRegexp.MatchString(target)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid args: %s", target))
	}
	// _, err := url.Parse(target)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid args: %s", target))
	// }
	return nil
}

func ValidateMail(target string) error {
	ok := MailRegexp.MatchString(target)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid args: %s", target))
	}
	return nil
}
