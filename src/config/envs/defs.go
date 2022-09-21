package envs

type ConfigDefs struct {
	Mode string // Run mode. `local`, `test` or `prod`

	DBConfig string // DB Connect config
}
