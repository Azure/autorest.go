package clientopgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstClient_One(t *testing.T) {
	client, err := NewFirstClient(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.One(context.Background(), &FirstClientOneOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
