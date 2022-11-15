package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/contents"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestBioHandler(t *testing.T) {
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

		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewGet("", "/")
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.BioHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		res := new(handler.BioResponse)
		err = m.Json(res)
		require.NoError(t, err)

		require.Equal(t, res.Biography.ID, bio.Biography.ID)
		require.Equal(t, res.Location.Address, bio.Location.Address)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.BioHandler(e)
	})
}

func TestCreateBioHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create a new location
		bio, err := u.CreateBio()
		require.NoError(t, err)
		loc, err := bio.CreateLocationDB(ctx, tool.DB)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("is_public", "true")
		form.Insert("location_id", strconv.Itoa(int(loc.ID)))
		form.Insert("position", "web application engineer")
		form.Insert("position_ja", "Webアプリケーションエンジニア")
		form.Insert("join_date", "2022-11-12T07:00:29.967Z")
		form.Insert("leave_date", "2120-01-31T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateBioHandler(e)
		require.NoError(t, err)

		m.Status(t, http.StatusCreated)

		// check
		b := new(ent.Biography)
		err = m.Json(b)
		require.NoError(t, err)
		bioInDB, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(b.ID)).
			First(ctx)
		require.NoError(t, err)

		require.True(t, bioInDB.IsPublic)
		require.Equal(t, bioInDB.LocationID, loc.ID)
		require.Equal(t, bioInDB.Position, "web application engineer")
		require.Equal(t, bioInDB.PositionJa, "Webアプリケーションエンジニア")
		require.Equal(t, bioInDB.Join.Year(), 2022)
		require.Equal(t, bioInDB.Join.Day(), 12)
		require.Equal(t, bioInDB.Leave.Year(), 2120)
		require.Equal(t, bioInDB.Leave.Day(), 31)
	})

	t.Run("no leave and success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create a new location
		bio, err := u.CreateBio()
		require.NoError(t, err)
		loc, err := bio.CreateLocationDB(ctx, tool.DB)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("is_public", "true")
		form.Insert("location_id", strconv.Itoa(int(loc.ID)))
		form.Insert("position", "web application engineer")
		form.Insert("position_ja", "Webアプリケーションエンジニア")
		form.Insert("join_date", "2022-11-12T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateBioHandler(e)
		require.NoError(t, err)

		m.Status(t, http.StatusCreated)

		// check
		b := new(ent.Biography)
		err = m.Json(b)
		require.NoError(t, err)
		bioInDB, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(b.ID)).
			First(ctx)
		require.NoError(t, err)

		require.True(t, bioInDB.IsPublic)
		require.Equal(t, bioInDB.LocationID, loc.ID)
		require.Equal(t, bioInDB.Position, "web application engineer")
		require.Equal(t, bioInDB.PositionJa, "Webアプリケーションエンジニア")
		require.Equal(t, bioInDB.Join.Year(), 2022)
		require.Equal(t, bioInDB.Join.Day(), 12)

		// IsZero reports whether t represents the zero time instant,
		// January 1, year 1, 00:00:00 UTC.
		require.True(t, time.Time.IsZero(bioInDB.Leave))
	})

	t.Run("no exists location", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("is_public", "true")
		form.Insert("location_id", "123456")
		form.Insert("position", "web application engineer")
		form.Insert("position_ja", "Webアプリケーションエンジニア")
		form.Insert("join_date", "2022-11-15T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPost)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.CreateBioHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateBioHandler(e)
	})
}

func TestUpdateBioHandler(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create a new bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("bio_id", strconv.Itoa(int(bio.Biography.ID)))
		form.Insert("position", "web application engineer")
		form.Insert("position_ja", "Webアプリケーションエンジニア")

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateBioHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		// check
		bioInDB, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(bio.Biography.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, bioInDB.Position, "web application engineer")
		require.Equal(t, bioInDB.PositionJa, "Webアプリケーションエンジニア")
	})

	t.Run("all changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create a new bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		bioUsedLoc, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bioUsedLoc.CreateLocationDB(ctx, tool.DB)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("bio_id", strconv.Itoa(int(bio.Biography.ID)))
		form.Insert("is_public", "true")
		form.Insert("location_id", strconv.Itoa(int(bioUsedLoc.LocationId)))
		form.Insert("position", "web application engineer")
		form.Insert("position_ja", "Webアプリケーションエンジニア")
		form.Insert("join_date", "2022-11-15T07:00:29.967Z")
		form.Insert("leave_date", "2120-11-15T07:00:29.967Z")

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateBioHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		// check
		bioInDB, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(bio.Biography.ID)).
			First(ctx)
		require.NoError(t, err)

		require.True(t, bioInDB.IsPublic)
		require.Equal(t, bioInDB.Position, "web application engineer")
		require.Equal(t, bioInDB.PositionJa, "Webアプリケーションエンジニア")
		require.Equal(t, bioInDB.LocationID, bioUsedLoc.LocationId)
		require.Equal(t, bioInDB.Join.Year(), 2022)
		require.Equal(t, bioInDB.Leave.Year(), 2120)
	})

	t.Run("no changes", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()
		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		// create a new bio
		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		// create post form
		form := contents.NewMultipart()
		form.Insert("bio_id", strconv.Itoa(int(bio.Biography.ID)))

		m, err := mock.NewFormData("/", form, http.MethodPut)
		require.NoError(t, err)
		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)
		e := m.Echo()
		h, err := tool.Handler()
		require.NoError(t, err)

		err = h.UpdateBioHandler(e)
		require.Error(t, err)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateBioHandler(e)
	})
}

func TestDeleteBioHandler(t *testing.T) {
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

		bio, err := u.CreateBio()
		require.NoError(t, err)
		_, err = bio.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodDelete, fmt.Sprintf("/?bio_id=%v", bio.Biography.ID))
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.DeleteBioHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		exist, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(bio.Biography.ID)).
			Exist(ctx)
		require.NoError(t, err)

		require.False(t, exist)
	})

	t.Run("delete other user", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		u2, err := tool.NewUser(ctx)
		require.NoError(t, err)

		bio2, err := u2.CreateBio()
		require.NoError(t, err)
		_, err = bio2.CreateDB(ctx, tool.DB)
		require.NoError(t, err)

		m, err := mock.NewMock("", http.MethodDelete, fmt.Sprintf("/?bio_id=%v", bio2.Biography.ID))
		require.NoError(t, err)

		err = u.HandlerSession(ctx, tool.DB, m)
		require.NoError(t, err)

		e := m.Echo()

		err = h.DeleteBioHandler(e)
		require.NoError(t, err)

		m.Ok(t)

		exist, err := tool.DB.Client.Biography.
			Query().
			Where(biography.ID(bio2.Biography.ID)).
			Exist(ctx)
		require.NoError(t, err)

		require.True(t, exist)
	})

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteBioHandler(e)
	})
}
