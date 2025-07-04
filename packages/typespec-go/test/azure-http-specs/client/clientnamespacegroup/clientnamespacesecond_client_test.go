package clientnamespacegroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestClientNamespaceSecondClient_GetSecond(t *testing.T) {
	factory := &ClientNamespaceClient{internal: &azcore.Client{}}
	resp, err := factory.NewClientNamespaceSecondClient().GetSecond(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "second", resp.SecondClientResult.Type)
}
