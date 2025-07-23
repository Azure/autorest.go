package clientnamespacegroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientNamespaceSecondClient_GetSecond(t *testing.T) {
	client, err := NewClientNamespaceClient(nil)
	require.NoError(t, err)
	client.endpoint = "http://localhost:3000"
	secondClient := client.NewClientNamespaceSecondClient()
	require.NotNil(t, secondClient)
	resp, err := secondClient.GetSecond(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.SecondClientResult)
	require.NotNil(t, resp.SecondClientResult.Type)
	require.Equal(t, SecondClientEnumTypeSecond, *resp.SecondClientResult.Type)
}
