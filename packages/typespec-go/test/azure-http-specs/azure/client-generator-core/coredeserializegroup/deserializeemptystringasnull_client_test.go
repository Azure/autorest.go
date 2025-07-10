package coredeserializegroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeserializeEmptyStringAsNullClient_Get(t *testing.T) {
	client, err := NewDeserializeEmptyStringAsNullClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
