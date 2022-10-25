package handler_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src/handler"
	"github.com/stretchr/testify/require"
)

func TestValidateURL(t *testing.T) {
	urls := map[string]bool{
		"https://example.com":              true,
		"http://example.com":               true,
		"https://cateiru.com/hogehoge":     true,
		"https://sso.cateiru.com/?aaa=bbb": true,

		"http://":            false,
		"hogehoge":           false,
		"https//example.com": false,
	}

	for url, ok := range urls {
		if ok {
			require.NoError(t, handler.ValidateURL(url))
		} else {
			require.Error(t, handler.ValidateURL(url))
		}
	}
}
