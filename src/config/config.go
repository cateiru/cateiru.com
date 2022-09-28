package config

import (
	"github.com/cateiru/cateiru.com/src/config/envs"
)

var Config *envs.ConfigDefs = &envs.TestConfig

func Init(mode string) {
	switch mode {
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
