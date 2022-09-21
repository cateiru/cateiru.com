package config

import (
	"github.com/cateiru/cateiru.com/src/config/envs"
)

// Set this variable at build time.
//
// Example:
//
//	go build  -ldflags="-X config.Mode=prod"
var Mode string = ""

var Config *envs.ConfigDefs = nil

func Init() {
	switch Mode {
	case "test":
		Config = &envs.TestConfig
	case "local":
		Config = &envs.LocalConfig
	case "prod":
		Config = &envs.ProdConfig
	default:
		Config = &envs.TestConfig
	}
}
