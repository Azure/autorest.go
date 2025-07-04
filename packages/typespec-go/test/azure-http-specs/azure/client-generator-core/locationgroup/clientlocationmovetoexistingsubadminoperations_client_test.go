package locationgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToExistingSubClient_ArchiveProduct(t *testing.T) {
	factory := ClientLocationMoveToExistingSubClient{internal: &azcore.Client{}}
	client := factory.NewClientLocationMoveToExistingSubAdminOperationsClient()
	require.NotNil(t, client)
	resp, err := client.DeleteUser(context.Background(), nil) // Use appropriate context and options
	require.NoError(t, err)
	require.NotNil(t, resp)
}
