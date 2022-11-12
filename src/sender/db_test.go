package sender_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru.com/ent/contact"
	"github.com/cateiru/cateiru.com/src/test"
	"github.com/stretchr/testify/require"
)

func TestInsertDB(t *testing.T) {
	test.Init()

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		tool, err := test.NewTestTool()
		require.NoError(t, err)
		defer tool.Close()

		err = tool.ClearUser(ctx)
		require.NoError(t, err)

		u, err := tool.NewUser(ctx)
		require.NoError(t, err)

		err = u.SelectStatus(ctx, true)
		require.NoError(t, err)

		_, err = u.CreateNotice()
		require.NoError(t, err)

		form, err := tool.NewForm()
		require.NoError(t, err)

		_, err = form.InsertDB(ctx, tool.DB, u.User.ID)
		require.NoError(t, err)

		dbForm, err := tool.DB.Client.Contact.Query().
			Where(contact.ToUserID(u.User.ID)).
			First(ctx)
		require.NoError(t, err)

		require.Equal(t, dbForm.Name, form.Name)
	})
}
