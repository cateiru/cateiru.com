package handler

import (
	"net/http"

	"github.com/cateiru/cateiru.com/src/base"
)

func RootHandler(base *base.Base) error {
	return base.E.String(http.StatusOK, "Hello World")
}
