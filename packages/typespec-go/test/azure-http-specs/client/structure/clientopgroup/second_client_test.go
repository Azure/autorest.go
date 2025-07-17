package clientopgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecondClientClient_Five(t *testing.T) {
	client, err := NewSecondClient(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Five(context.Background(), &SecondClientFiveOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
