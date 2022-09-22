package src_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	src.Init("test")

	t.Run("connection", func(t *testing.T) {
		db, err := src.NewConnectMySQL()
		require.NoError(t, err)
		require.NotNil(t, db)
	})

	t.Run("connection empty sql", func(t *testing.T) {
		db, err := src.NewEmptySQL()
		require.NoError(t, err)
		require.NotNil(t, db)
	})
}
