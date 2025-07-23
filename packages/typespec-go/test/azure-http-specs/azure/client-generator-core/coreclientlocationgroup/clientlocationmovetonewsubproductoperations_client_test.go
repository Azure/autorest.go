package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationMoveToNewSubProductOperationsClient_ListProducts(t *testing.T) {
	client, err := NewClientLocationMoveToNewSubProductOperationsClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	resp, err := client.ListProducts(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
