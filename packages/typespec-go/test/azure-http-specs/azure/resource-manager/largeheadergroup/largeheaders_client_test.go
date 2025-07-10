package largeheadergroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func TestNewLargeHeadersClient_BeginTwo6K(t *testing.T) {
	client := clientFactory.NewLargeHeadersClient()
	poller, err := client.BeginTwo6K(context.Background(), resourceGroupExpected, "header1", nil)
	require.NoError(t, err)
	require.NotNil(t, poller)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Succeeded)
	require.True(t, *resp.Succeeded)
}
