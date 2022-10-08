package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/cateiru/cateiru.com/src"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/cateiru.com/src/logging"
	"github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

type Export struct {
}

type Migration struct {
	Mode string `short:"m" long:"mode" default:"local"`
	Env  string `short:"e" long:"env"`
}

type Options struct {
}

var opts Options
var export Export
var migration Migration

func init() {
	src.Init("local")
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "go run ./script/main.go"

	parser.AddCommand("export", "export sql schema", "", &export)
	parser.AddCommand("migration", "sql migration", "", &migration)

	_, err := parser.Parse()
	if err != nil {
		return
	}
}

// Export SQL Schema.
func (cmd *Export) Execute(args []string) error {
	ctx := context.Background()

	db, err := db.NewEmptySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	f, err := os.Create("schema.sql")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := db.WriteSchema(ctx, f); err != nil {
		return err
	}

	logging.Sugar.Infoln("Success export SQL schema.sql")

	return nil
}

// migration
func (cmd *Migration) Execute(args []string) error {
	ctx := context.Background()

	// Overwrite config
	config.Init(cmd.Mode)

	if cmd.Mode == "prod" && cmd.Env != "" {
		return cmd.migrationProd(ctx, cmd.Env)
	}

	db, err := db.NewConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Client.Schema.Create(ctx); err != nil {
		return err
	}

	logging.Sugar.Infof("Success migration! DATABASE: %s\n", config.Config.DBConfig)

	return nil
}

func (cmd *Migration) migrationProd(ctx context.Context, envPath string) error {
	if !filepath.IsAbs(envPath) {
		p, err := os.Getwd()
		if err != nil {
			return err
		}

		envAbsPath := filepath.Join(p, envPath)

		if _, err := os.Stat(envPath); err != nil {
			return err
		}

		envPath = envAbsPath
	}

	if err := godotenv.Load(envPath); err != nil {
		return err
	}

	config.Config.DBConfig = mysql.Config{
		DBName:    "cateirucom",
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      "127.0.0.1:3306",
		Net:       "tcp",
		ParseTime: true,
	}

	db, err := db.NewConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	f, err := os.Create("migration_diff.sql")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := db.WriteSchema(ctx, f); err != nil {
		return err
	}

	logging.Sugar.Infof("Success migration prod!\n")

	return nil
}
