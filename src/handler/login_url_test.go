package handler_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/cateiru/go-http-easy-test/handler/mock"
	"github.com/stretchr/testify/require"
)

func TestLoginURLHandler(t *testing.T) {
	test.Init()

	m, err := mock.NewMock("", http.MethodGet, "/")
	require.NoError(t, err)

	e := m.Echo()

	err = handler.LoginURLHandler(e)
	require.NoError(t, err)

	m.Status(t, http.StatusMovedPermanently)

	location, err := m.W.Result().Location()
	require.NoError(t, err)
	require.Equal(t, location.String(), fmt.Sprintf("https://sso.cateiru.com/sso/login?scope=openid&response_type=code&client_id=12345&redirect_uri=%s&prompt=consent", url.QueryEscape("http://localhost:8080/login")))
}
