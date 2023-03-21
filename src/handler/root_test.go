package handler_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestRootHandler(t *testing.T) {
	test.Init()

	t.Run("can access", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		h, err := tool.Handler()
		require.NoError(t, err)

		c := m.Echo()
		err = h.RootHandler(c)
		require.NoError(t, err)

		m.Status(t, http.StatusMovedPermanently)

		location, err := m.W.Result().Location()
		require.NoError(t, err)

		require.Equal(t, *location, config.Config.PageDomain)
	})
}
