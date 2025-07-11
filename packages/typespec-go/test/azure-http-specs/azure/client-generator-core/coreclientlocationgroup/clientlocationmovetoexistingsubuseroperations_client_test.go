package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToExistingSubUserOperationsClient_GetUser(t *testing.T) {
	factory, err := NewClientLocationMoveToExistingSubClient(nil)
	require.NoError(t, err)
	client := factory.NewClientLocationMoveToExistingSubUserOperationsClient()
	require.NotNil(t, client)
	resp, err := client.GetUser(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
