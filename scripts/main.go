package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cateiru/cateiru.com/src"
)

func main() {
	flag.Parse()
	option := flag.Arg(0)

	if len(option) == 0 {
		log.Fatalln("No option")
	}

	switch option {
	case "export":
		export()
	case "migration":
		migration()
	default:
		log.Fatalln("invalid option")
	}
}

// Export SQL Schema.
func export() {
	db, err := src.NewEmptySQL()
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	f, err := os.Create("schema.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()

	if err := db.WriteSchema(ctx, f); err != nil {
		log.Fatalf("write sql failed. %v", err)
	}
}

func migration() {
	fmt.Println("OK")
}
