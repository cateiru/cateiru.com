package handler_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
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

		m, err := mock.NewGet("", "/user/me")
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

		e := m.Echo()

		err = h.MeHandler(e)
		require.NoError(t, err)

		m.Ok(t)
	})

	t.Run("no login", func(t *testing.T) {
		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		m, err := mock.NewGet("", "/user/me")
		require.NoError(t, err)

		e := m.Echo()

		err = h.MeHandler(e)
		require.Error(t, err)
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

		form := contents.NewMultipart()
		form.Insert("family_name", "naaaaa")
		form.Insert("given_name", "givennnnn")
		form.Insert("family_name_ja", "いいい")
		form.Insert("given_name_ja", "ああああ")

		b := time.Now()
		b = b.AddDate(-10, 0, 0)
		form.Insert("birth_date", b.Format("2006-01-02T15:04:05-0700"))

		form.Insert("location", "locattttttt")
		form.Insert("location_ja", "ロケーション")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

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

		form := contents.NewMultipart()
		form.Insert("location", "locattttttt")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

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

		form := contents.NewMultipart()

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)

		u.HandlerSession(ctx, tool.DB, m)

		h, err := tool.Handler()
		require.NoError(t, err)

		e := m.Echo()
		err = h.UpdateUserHandler(e)
		require.Error(t, err)
	})
}
