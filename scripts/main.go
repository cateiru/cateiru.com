package main

import (
	"context"
	"errors"
	"os"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src"
	"github.com/cateiru/cateiru.com/src/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
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

	config.Config.DBConfig = mysql.Config{
		DBName:    "em",
		User:      "docker",
		Passwd:    "docker",
		Addr:      "127.0.0.1:3306",
		Net:       "tcp",
		ParseTime: true,
	}

	client, err := ent.Open("mysql", config.Config.DBConfig.FormatDSN())
	if err != nil {
		return err
	}
	defer client.Close()

	f, err := os.Create("schema.sql")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := client.Schema.WriteTo(ctx, f); err != nil {
		return err
	}

	return nil
}

// migration
func (cmd *Migration) Execute(args []string) error {
	ctx := context.Background()

	mode := cmd.Mode

	dbName := ""
	switch mode {
	case "local":
		dbName = "cateiru"
	case "test":
		dbName = "cateiru_test"
	default:
		return errors.New("invalid mode")
	}

	config.Config.DBConfig = mysql.Config{
		DBName:    dbName,
		User:      "docker",
		Passwd:    "docker",
		Addr:      "127.0.0.1:3306",
		Net:       "tcp",
		ParseTime: true,
	}

	client, err := ent.Open("mysql", config.Config.DBConfig.FormatDSN())
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		return err
	}

	return nil
}
