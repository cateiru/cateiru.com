package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/user"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

// Response login user data
func TestMeHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := easy.NewMock("/user/me", http.MethodGet, "")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.MeHandler(e)
		require.NoError(t, err)

		m.Ok(t)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.MeHandler(e)
	})
}

func TestUpdateUserHandler(t *testing.T) {
	test.Init()

	t.Run("all changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("family_name", "naaaaa")
		form.Insert("given_name", "givennnnn")
		form.Insert("family_name_ja", "いいい")
		form.Insert("given_name_ja", "ああああ")

		b := time.Now()
		b = b.AddDate(-10, 0, 0)
		form.Insert("birth_date", b.Format(time.RFC3339))

		form.Insert("location", "locattttttt")
		form.Insert("location_ja", "ロケーション")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateUserHandler(e)
		require.NoError(t, err)

		updatedUser, err := tool.DB.Client.User.Get(ctx, u.User.ID)
		require.NoError(t, err)

		require.Equal(t, updatedUser.FamilyName, "naaaaa")
		require.Equal(t, updatedUser.GivenName, "givennnnn")
		require.Equal(t, updatedUser.FamilyNameJa, "いいい")
		require.Equal(t, updatedUser.GivenNameJa, "ああああ")

		require.Equal(t, updatedUser.BirthDate.Year(), b.Year())

		require.Equal(t, updatedUser.Location, "locattttttt")
		require.Equal(t, updatedUser.LocationJa, "ロケーション")
	})

	t.Run("one change", func(t *testing.T) {
		ctx := context.Background()
		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()
		form.Insert("location", "locattttttt")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateUserHandler(e)
		require.NoError(t, err)

		updatedUser, err := tool.DB.Client.User.Get(ctx, u.User.ID)
		require.NoError(t, err)

		require.Equal(t, updatedUser.FamilyName, u.FamilyName)
		require.Equal(t, updatedUser.GivenName, u.GivenName)
		require.Equal(t, updatedUser.FamilyNameJa, u.FamilyNameJa)
		require.Equal(t, updatedUser.GivenNameJa, u.GivenNameJa)

		require.Equal(t, updatedUser.BirthDate.Year(), u.BirthDate.Year())

		require.Equal(t, updatedUser.Location, "locattttttt")
		require.Equal(t, updatedUser.LocationJa, u.LocationJa)
	})

	t.Run("no changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		form := easy.NewMultipart()

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateUserHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateUserHandler(e)
	})
}

func TestAllUsersHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearUser(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		_, err = tool.NewUser(ctx)
		require.NoError(t, err)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.AllUsersHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		body := new([]ent.User)

		err = m.Json(body)
		require.NoError(t, err)

		require.Len(t, *body, 2)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.AllUsersHandler(e)
	})
}

func TestChangeSelectHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearUser(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		u2, err := tool.NewUser(ctx)
		require.NoError(t, err)

		tool.DB.Client.User.Update().Where(user.ID(u2.User.ID)).SetSelected(true).SaveX(ctx)

		u2Changed, err := tool.DB.Client.User.Query().Where(user.ID(u2.User.ID)).First(ctx)
		require.NoError(t, err)
		require.True(t, u2Changed.Selected)

		form := easy.NewMultipart()
		form.Insert("id", fmt.Sprintf("%d", u.User.ID))

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		h, err := tool.Handler()
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()
		err = h.ChangeSelect(e)
		require.NoError(t, err)

		u2ChangedSecond, err := tool.DB.Client.User.Query().Where(user.ID(u2.User.ID)).First(ctx)
		require.NoError(t, err)
		require.False(t, u2ChangedSecond.Selected)

		uChanged, err := tool.DB.Client.User.Query().Where(user.ID(u.User.ID)).First(ctx)
		require.NoError(t, err)
		require.True(t, uChanged.Selected)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.ChangeSelect(e)
	})
}
