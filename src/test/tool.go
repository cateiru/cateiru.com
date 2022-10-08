package test

import (
	"context"

	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/cateiru.com/src/handler"
)

type TestTool struct {
	DB *db.DB

	Users []*TestUser
}

// Create test tool
//
//	defer tool.Close()
func NewTestTool() (*TestTool, error) {
	db, err := db.NewConnectMySQL()
	if err != nil {
		return nil, err
	}

	return &TestTool{
		DB:    db,
		Users: []*TestUser{},
	}, nil
}

func (c *TestTool) Close() {
	c.DB.Close()
}

func (c *TestTool) NewUser(ctx context.Context) (*TestUser, error) {
	user, err := NewUser()
	if err != nil {
		return nil, err
	}

	if err := user.CreateDB(ctx, c.DB); err != nil {
		return nil, err
	}

	c.Users = append(c.Users, user)

	return user, nil
}

// Returns list of user ids
func (c *TestTool) GetUserIds() []uint32 {
	ids := []uint32{}

	for _, user := range c.Users {
		ids = append(ids, user.User.ID)
	}

	return ids
}

func (c *TestTool) Handler() (*handler.Handler, error) {
	return handler.NewHandler(c.DB)
}
