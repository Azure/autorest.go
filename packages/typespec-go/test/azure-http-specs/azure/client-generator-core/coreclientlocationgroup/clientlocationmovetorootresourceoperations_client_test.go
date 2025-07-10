package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationArchiveOperationsClient_ArchiveProduct(t *testing.T) {
	factory := ClientLocationMoveToRootClient{}
	client := factory.NewClientLocationMoveToRootResourceOperationsClient()
	require.NotNil(t, client)
	resp, err := client.GetResource(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
