package coreclientlocationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientLocationClient_GetHealthStatus(t *testing.T) {
	client, err := NewClientLocationClient(nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	resp, err := client.GetHealthStatus(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
