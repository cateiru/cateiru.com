package handler

import (
	"github.com/cateiru/cateiru.com/src/base"
	"github.com/cateiru/cateiru.com/src/db"
)

type Handler struct {
	*base.Base
}

func NewHandler(db *db.DB) (*Handler, error) {
	base, err := base.NewBase(db)
	if err != nil {
		return nil, err
	}

	return &Handler{
		Base: base,
	}, nil
}
