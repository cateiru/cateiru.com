package main

import (
	"context"
	"flag"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cateiru/cateiru.com/src"
)

func init() {
	src.Init()
}

func main() {
	flag.Parse()
	option := flag.Arg(0)

	if len(option) == 0 {
		log.Fatalln("No option")
	}

	ctx := context.Background()

	switch option {
	case "export":
		export(ctx)
	case "migration":
		migration(ctx)
	default:
		log.Fatalln("invalid option")
	}
}

// Export SQL Schema.
func export(ctx context.Context) {
	db, err := src.NewEmptySQL()
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	f, err := os.Create("schema.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()

	if err := db.WriteSchema(ctx, f); err != nil {
		log.Fatalf("write sql failed. %v", err)
	}
}

func migration(ctx context.Context) {
	db, err := src.NewConnectMySQL()
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	if err := db.Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
