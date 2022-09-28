package test

import (
	"context"

	"github.com/cateiru/cateiru.com/src/db"
)

type TestTool struct {
	DB *db.DB

	Users []*TestUser
}

func NewTestTool(db *db.DB) *TestTool {
	return &TestTool{
		DB: db,
	}
}

// Create test tool
//
// Usage:
//
//	tool, err := test.NewTestToolDB()
//	require.NoError(t, err)
//	defer tool.Close()
func NewTestToolDB() (*TestTool, error) {
	dbOp, err := db.NewConnectMySQL()
	if err != nil {
		return nil, err
	}

	return NewTestTool(dbOp), nil
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