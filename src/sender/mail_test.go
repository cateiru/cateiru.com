package sender_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src/sender"
	"github.com/stretchr/testify/require"
)

type TestTemplateValues struct {
	Title string
	Body  string
}

func TestTemplate(t *testing.T) {
	text := TestTemplateValues{
		Title: "sample title",
		Body:  "sample body",
	}

	mail := sender.NewMail("example@example.com", "test title")

	_, err := mail.AddContentsFromTemplate("test", text)
	require.NoError(t, err)

	require.NotEqual(t, len(mail.HTMLText), 0)
	require.NotEqual(t, len(mail.PlainText), 0)
}
