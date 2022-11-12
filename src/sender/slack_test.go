package sender_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru.com/src/test"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const MOCKED_SUCCESS_SLACK_WEBHOOK = "https://example.com/slack/webhook"
const MOCKED_FAILED_SLACK_WEBHOOK = "https://example.com/slack/webhook/failed"

func TestSlackSender(t *testing.T) {
	test.Init()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", MOCKED_SUCCESS_SLACK_WEBHOOK,
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	httpmock.RegisterResponder("POST", MOCKED_FAILED_SLACK_WEBHOOK,
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	tool, err := test.NewTestTool()
	require.NoError(t, err)
	defer tool.Close()

	t.Run("success", func(t *testing.T) {
		form, err := tool.NewForm()
		require.NoError(t, err)

		err = form.SlackSender(MOCKED_SUCCESS_SLACK_WEBHOOK)
		require.NoError(t, err)
	})

	t.Run("invalid url", func(t *testing.T) {
		form, err := tool.NewForm()
		require.NoError(t, err)

		err = form.SlackSender(MOCKED_FAILED_SLACK_WEBHOOK)
		require.Error(t, err)
	})

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	require.Equal(t, info[fmt.Sprintf("POST %s", MOCKED_SUCCESS_SLACK_WEBHOOK)], 1)
	require.Equal(t, info[fmt.Sprintf("POST %s", MOCKED_FAILED_SLACK_WEBHOOK)], 1)
}
