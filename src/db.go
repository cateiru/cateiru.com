package src

import (
	"context"
	"io"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/src/config"
)

type DB struct {
	// Usage. https://entgo.io/ja/docs/crud
	Client *ent.Client
}

// Start DB Connection
//
// Use MySQL.
func NewConnectMySQL() (*DB, error) {
	client, err := ent.Open("mysql", config.Config.DBConfig)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
	}, nil
}

// Connect `empty` table
func NewEmptySQL() (*DB, error) {
	client, err := ent.Open("mysql", "docker:docker@tcp(localhost:3306)/em?parseTime=True")
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
	}, nil
}

// Write SQL Schema
//
// Used:
//
//	f, err := os.Create("schema.sql")
//	if err != nil {
//		log.Fatalf("create migrate file: %v", err)
//	}
//	if err := db.WriteSchema(ctx, f); err != nil {
//		log.Fatalf("failed %v", err)
//	}
func (c *DB) WriteSchema(ctx context.Context, w io.Writer) error {
	return c.Client.Schema.WriteTo(ctx, w)
}

// Close DB Connection
func (c *DB) Close() {
	c.Client.Close()
}
