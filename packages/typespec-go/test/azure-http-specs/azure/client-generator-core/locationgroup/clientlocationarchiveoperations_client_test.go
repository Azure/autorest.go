package locationgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestClientLocationArchiveOperationsClient_ArchiveProduct(t *testing.T) {
	factory := ClientLocationClient{internal: &azcore.Client{}}
	client := factory.NewClientLocationArchiveOperationsClient()
	require.NotNil(t, client)
	resp, err := client.ArchiveProduct(context.Background(), nil) // Use appropriate context and options
	require.NoError(t, err)
	require.NotNil(t, resp)
}
