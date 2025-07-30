package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationArchiveOperationsClient_ArchiveProduct(t *testing.T) {
	factory, err := NewClientLocationClient("http://localhost:3000", nil)
	require.NoError(t, err)
	client := factory.NewClientLocationArchiveOperationsClient()
	require.NotNil(t, client)
	resp, err := client.ArchiveProduct(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientLocationMoveToRootResourceOperationsClient_GetResource(t *testing.T) {
	factory, err := NewClientLocationMoveToRootClient("http://localhost:3000", nil)
	require.NoError(t, err)
	client := factory.NewClientLocationMoveToRootResourceOperationsClient()
	require.NotNil(t, client)
	resp, err := client.GetResource(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
