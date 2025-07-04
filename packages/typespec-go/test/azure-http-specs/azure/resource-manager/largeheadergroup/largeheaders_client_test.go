package largeheadergroup

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"
)

func TestNewLargeHeadersClient_BeginTwo6K(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	client, err := NewLargeHeadersClient(subscriptionID, cred, nil)
	require.NoError(t, err)
	resourceGroupName := "test-rg"
	resourceName := "header1"
	poller, err := client.BeginTwo6K(context.Background(), resourceGroupName, resourceName, &LargeHeadersClientBeginTwo6KOptions{})
	require.NoError(t, err)
	require.NotNil(t, poller)
	resp, err := poller.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{Frequency: time.Second})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
