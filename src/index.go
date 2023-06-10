package src

import (
	"time"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
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
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogHost:   true,
		LogError:  true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logging.Logger.Error("request",
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
					zap.String("error_message", v.Error.Error()),
				)
			} else {
				logging.Logger.Info("request",
					zap.String("URI", v.URI),
					zap.String("method", v.Method),
					zap.Int("status", v.Status),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
				)
			}

			return nil
		},
	}))

	if config.Config.Cors != nil {
		e.Use(config.Config.Cors)
	}

	db, err := db.NewConnectMySQL()
	if err != nil {
		logging.Sugar.Fatal(err)
	}
	defer db.Close()

	handler, err := handler.NewHandler(db)
	if err != nil {
		logging.Sugar.Fatal(err)
		return
	}

	// setting routes
	Routes(e, handler)

	s := http2.Server{}

	// Start a server
	// connection port is `8080`
	//
	// and, `http://localhist:8080` to access.
	e.Logger.Fatal(e.StartH2CServer(":8080", &s))
}
