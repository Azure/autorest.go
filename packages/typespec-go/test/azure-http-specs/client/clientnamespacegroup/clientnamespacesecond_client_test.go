package clientnamespacegroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientNamespaceSecondClient_GetSecond(t *testing.T) {
	factory := &ClientNamespaceClient{}
	resp, err := factory.NewClientNamespaceSecondClient().GetSecond(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "second", resp.SecondClientResult.Type)
}
