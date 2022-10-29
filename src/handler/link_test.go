package handler_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestLinkHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.LinkHandler(e)
	})
}

func TestCreateLinkHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.CreateLinkHandler(e)
	})
}

func TestUpdateLinkHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.UpdateLinkHandler(e)
	})
}

func TestDeleteLinkHandler(t *testing.T) {
	test.Init()

	test.LoginTestGet(t, func(h *handler.Handler, e echo.Context) error {
		return h.DeleteLinkHandler(e)
	})
}

func TestGetFavicon(t *testing.T) {
	test.Init()

	t.Run("get favicon from HTML", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/html/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)

		fav, err := handler.GetFavicon("https://example.com")
		require.NoError(t, err)

		require.Equal(t, fav, "https://example.com/html/favicon.ico")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 0)
		require.Equal(t, info["GET https://example.com/html/favicon.ico"], 1)
	})

	t.Run("no defined favicon in HTML", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon_no_defined.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/html/favicon.ico",
			httpmock.NewStringResponder(http.StatusOK, ""),
		)

		fav, err := handler.GetFavicon("https://example.com")
		require.NoError(t, err)

		require.Equal(t, fav, "https://example.com/favicon.ico")

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 1)
		require.Equal(t, info["GET https://example.com/html/favicon.ico"], 0)
	})

	t.Run("favicon not found", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		f, err := os.ReadFile("./favicon_no_defined.html")
		require.NoError(t, err)

		httpmock.RegisterResponder(
			"GET",
			"https://example.com",
			httpmock.NewBytesResponder(http.StatusOK, f),
		)
		httpmock.RegisterResponder(
			"GET",
			"https://example.com/favicon.ico",
			httpmock.NewStringResponder(http.StatusNotFound, ""),
		)

		_, err = handler.GetFavicon("https://example.com")
		require.Error(t, err)

		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		require.Equal(t, info["GET https://example.com"], 1)
		require.Equal(t, info["GET https://example.com/favicon.ico"], 1)
	})
}
