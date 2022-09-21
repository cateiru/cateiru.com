package envs

type ConfigDefs struct {
	Mode string // Run mode. `local`, `test` or `prod`

	// DB Connect config
	// like to `<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True`
	DBConfig string
}
