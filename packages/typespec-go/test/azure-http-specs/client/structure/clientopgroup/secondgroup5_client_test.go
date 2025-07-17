package clientopgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecondGroup5Client_Six(t *testing.T) {
	client, err := NewSecondGroup5Client(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Six(context.Background(), &SecondGroup5ClientSixOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
