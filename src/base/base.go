package base

import (
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/labstack/echo/v4"
)

type Base struct {
	E  echo.Context
	DB *db.DB
}

// Create a Base objects
func NewBase(e echo.Context) (*Base, error) {
	db, err := db.NewConnectMySQL()
	if err != nil {
		return nil, err
	}

	return &Base{
		E:  e,
		DB: db,
	}, nil
}
