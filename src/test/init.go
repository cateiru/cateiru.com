package test

import (
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/logging"
)

func Init() {
	config.Init("test")
	logging.InitLogging("test")
}
