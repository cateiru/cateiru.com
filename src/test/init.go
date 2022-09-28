package test

import (
	"github.com/cateiru/cateiru.com/src"
	"github.com/cateiru/cateiru.com/src/config"
)

func Init() {
	config.Init("test")
	src.InitLogging("test")
}
