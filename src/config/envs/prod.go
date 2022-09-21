package envs

import "os"

var ProdConfig = ConfigDefs{
	Mode: "prod",

	DBConfig: os.Getenv("DB_CONFIG"),
}
