package test

import (
	"context"
	"crypto/rand"
	"errors"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
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
	return c.ClearLocation(ctx)
}

func (c *TestTool) ClearLocation(ctx context.Context) error {
	if _, err := c.DB.Client.Location.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (c *TestTool) ClearProduct(ctx context.Context) error {
	if _, err := c.DB.Client.Product.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (c *TestTool) ClearLink(ctx context.Context) error {
	if _, err := c.DB.Client.Link.Delete().Exec(ctx); err != nil {
		return err
	}
	if _, err := c.DB.Client.Category.Delete().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (c *TestTool) ClearNotice(ctx context.Context) error {
	if _, err := c.DB.Client.Notice.Delete().Exec(ctx); err != nil {
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

func LoginTestGet(t *testing.T, o func(h *handler.Handler, e echo.Context) error) {
	t.Run("no login", func(t *testing.T) {
		tool, err := NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		e := m.Echo()

		err = o(h, e)
		require.Error(t, err)
		CheckHTTPErrorMessage(t, err, "cookie is not found")
	})

	t.Run("invalid cookie: parse failed UUID", func(t *testing.T) {
		tool, err := NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		m.Cookie([]*http.Cookie{
			{
				Name:  config.Config.SessionCookieName,
				Value: "oooooo",
			},
		})

		e := m.Echo()

		err = o(h, e)
		require.Error(t, err)
		CheckHTTPErrorMessage(t, err, "failed parse session token uuid")
	})

	t.Run("invalid session token", func(t *testing.T) {
		tool, err := NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		uuid := uuid.New()

		m.Cookie([]*http.Cookie{
			{
				Name:  config.Config.SessionCookieName,
				Value: uuid.String(),
			},
		})

		e := m.Echo()

		err = o(h, e)
		require.Error(t, err)
		CheckHTTPErrorMessage(t, err, "session token invalid")
	})
}

func CheckHTTPErrorMessage(t *testing.T, err error, message string) {
	httpErr, ok := err.(*echo.HTTPError)
	require.True(t, ok, "err is no echo.HTTPError")

	require.Equal(t, httpErr.Message, message, "invalid error message")
}
