package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestLocationHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		// clear locations
		tool.ClearLocation(ctx)

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)
		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateLocationDB(ctx, tool.DB)
		require.NoError(t, err)

		secondU, err := tool.NewUser(ctx)
		require.NoError(t, err)
		secondBio, err := secondU.CreateBio()
		require.NoError(t, err)
		_, err = secondBio.CreateLocationDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.LocationHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		res := new([]ent.Location)
		err = m.Json(res)
		require.NoError(t, err)

		require.Len(t, *res, 2)

		require.EqualValues(t, (*res)[0].Address, bio.Location.Address, secondBio.Location.Address)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.LocationHandler(e)
	})
}

func TestCreateLocationHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("type", "univ")
		form.Insert("name", "cateiru")
		form.Insert("name_ja", "cateiru")
		form.Insert("address", "Saitama")
		form.Insert("address_ja", "埼玉県")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateLocationHandler(e)
		require.NoError(t, err)

		m.Status(t, http.StatusCreated)

		// check
		b := new(ent.Location)
		err = m.Json(b)
		require.NoError(t, err)
		locationInDB, err := tool.DB.Client.Location.
			Query().
			Where(location.ID(b.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, locationInDB.Type, location.Type("univ"))
		require.Equal(t, locationInDB.Name, "cateiru")
		require.Equal(t, locationInDB.NameJa, "cateiru")
		require.Equal(t, locationInDB.Address, "Saitama")
		require.Equal(t, locationInDB.AddressJa, "埼玉県")
	})

	t.Run("failed form", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("type", "univ")
		form.Insert("name", "cateiru")
		form.Insert("name_ja", "cateiru")
		form.Insert("address", "Saitama")
		// form.Insert("address_ja", "埼玉県")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateLocationHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateLocationHandler(e)
	})
}

func TestUpdateLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateLocationHandler(e)
	})
}

func TestDeleteLocationHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteLocationHandler(e)
	})
}
