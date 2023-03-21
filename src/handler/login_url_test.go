package handler_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestLoginURLHandler(t *testing.T) {
	test.Init()

	tool, err := test.NewTestTool()
	require.NoError(t, err)

	m, err := easy.NewMock("/", http.MethodGet, "")
	require.NoError(t, err)

	h, err := tool.Handler()
	require.NoError(t, err)

	e := m.Echo()

	err = h.LoginURLHandler(e)
	require.NoError(t, err)

	m.Status(t, http.StatusMovedPermanently)

	location, err := m.W.Result().Location()
	require.NoError(t, err)
	require.Equal(t, location.String(), fmt.Sprintf("https://sso.cateiru.com/sso/login?scope=openid&response_type=code&client_id=12345&redirect_uri=%s&prompt=consent", url.QueryEscape("http://localhost:8080/login")))
}
