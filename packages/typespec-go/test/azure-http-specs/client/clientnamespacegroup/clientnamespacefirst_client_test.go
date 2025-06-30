package clientnamespacegroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientNamespaceFirstClient_GetFirst(t *testing.T) {
	factory := &ClientNamespaceClient{}
	client := factory.NewClientNamespaceFirstClient() 
	require.NotNil(t, client)
	resp, err := client.GetFirst(context.Background(), nil) // Use appropriate context and options
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "first", resp.FirstClientResult.Name)
}
