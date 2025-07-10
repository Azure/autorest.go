package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToExistingSubAdminOperationsClient_DeleteUser(t *testing.T) {
	client, err := NewClientLocationMoveToExistingSubAdminOperationsClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.DeleteUser(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientLocationMoveToExistingSubAdminOperationsClient_GetAdminInfo(t *testing.T) {
	client, err := NewClientLocationMoveToExistingSubAdminOperationsClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.GetAdminInfo(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
