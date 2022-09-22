package envs

var LocalConfig = ConfigDefs{
	Mode: "local",

	// This config is docker-cmpose MySQL connection.
	DBConfig: "docker:docker@tcp(localhost:3306)/cateiru?parseTime=True",

	SessionCookieName: "cateirucom-session",
}
