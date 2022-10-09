package test

import (
	"context"
	"crypto/rand"
	"errors"

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

func (c *TestTool) ClearUser(ctx context.Context) error {
	if _, err := c.DB.Client.User.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (c *TestTool) ClearBio(ctx context.Context) error {
	if _, err := c.DB.Client.Biography.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := c.DB.Client.Location.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
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

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
