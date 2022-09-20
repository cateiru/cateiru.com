package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cateiru/cateir.com/ent"
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
	client, err := ent.Open("mysql", "docker:docker@tcp(localhost:3306)/cateiru")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	f, err := os.Create("schema.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()
	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}

func migration() {
	fmt.Println("OK")
}
