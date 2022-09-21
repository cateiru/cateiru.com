package envs

var TestConfig = ConfigDefs{
	Mode: "test",

	DBConfig: "docker:docker@tcp(localhost:3306)/test?parseTime=True",
}
