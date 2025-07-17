package clientopgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstGroup4Client_Two(t *testing.T) {
	client, err := NewFirstGroup4Client(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Four(context.Background(), &FirstGroup4ClientFourOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
