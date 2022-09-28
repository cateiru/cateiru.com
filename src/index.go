package src

import (
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/labstack/echo/v4"
)

// Initialize Server
func Init(mode string) {
	config.Init(mode)
	logging.InitLogging(mode)
}

// Start API Server
//
// This API is run to hold information for `cateiru.com“.
func Server() {
	e := echo.New()

	// setting routes
	Routes(e)

	// Start a server
	// connection port is `8080`
	//
	// and, `http://localhist:8080` to access.
	e.Logger.Fatal(e.Start(":8080"))
}
