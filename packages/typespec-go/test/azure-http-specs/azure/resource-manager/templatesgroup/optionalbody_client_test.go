package templatesgroup

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"
)

func TestOptionalBodyClient_CheckGlobal(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)
	subscriptionIdExpected := "00000000-0000-0000-0000-000000000000"
	widgetName := "widget1"
	resourceGroup := "test-rg"
	clinet, err := NewOptionalBodyClient(subscriptionIdExpected, cred, nil)
	require.NoError(t, err)

	resp, err := clinet.Get(ctx, resourceGroup, widgetName, &OptionalBodyClientGetOptions{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
