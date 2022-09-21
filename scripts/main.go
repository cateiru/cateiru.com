package main

import (
	"context"
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cateiru/cateiru.com/src"
	"github.com/cateiru/cateiru.com/src/config"
)

func init() {
	src.Init()
}

func main() {
	flag.Parse()
	option := flag.Arg(0)

	if len(option) == 0 {
		src.Sugar.Fatalln("no options err")
	}

	ctx := context.Background()

	switch option {
	case "export":
		export(ctx)
	case "migration":
		migration(ctx)
	default:
		src.Sugar.Fatalln("invalid option")
	}
}

// Export SQL Schema.
func export(ctx context.Context) {
	db, err := src.NewEmptySQL()
	if err != nil {
		src.Sugar.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	f, err := os.Create("schema.sql")
	if err != nil {
		src.Sugar.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()

	if err := db.WriteSchema(ctx, f); err != nil {
		src.Sugar.Fatalf("write sql failed. %v", err)
	}
}

// Migrations
//
// Usage:
// go run ./scripts/main migration [mode]
//
// mode: Config mode. `test`, `local` or `prod`
// If no mode is set, it will migrate the `test` table by default.
func migration(ctx context.Context) {
	mode := flag.Arg(1)
	if len(mode) != 0 {
		// Overwrite config
		config.Mode = mode
		config.Init()
	}

	db, err := src.NewConnectMySQL()
	if err != nil {
		src.Sugar.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	if err := db.Client.Schema.Create(ctx); err != nil {
		src.Sugar.Fatalf("failed creating schema resources: %v", err)
	}
}
