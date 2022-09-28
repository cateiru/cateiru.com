package handler_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

func TestRootHandler(t *testing.T) {
	test.Init()

	t.Run("can access", func(t *testing.T) {
		m, err := mock.NewMock("", http.MethodGet, "/")
		require.NoError(t, err)

		c := m.Echo()
		err = handler.RootHandler(c)
		require.NoError(t, err)

		m.Status(t, http.StatusMovedPermanently)

		location, err := m.W.Result().Location()
		require.NoError(t, err)

		require.Equal(t, *location, config.Config.PageDomain)
	})
}
