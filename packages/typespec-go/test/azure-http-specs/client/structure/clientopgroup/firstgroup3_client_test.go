package clientopgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstGroup3Client_Three(t *testing.T) {
	client, err := NewFirstGroup3Client(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Three(context.Background(), &FirstGroup3ClientThreeOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}

func TestFirstGroup3Client_Two(t *testing.T) {
	client, err := NewFirstGroup3Client(nil)
	require.Nil(t, err)
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Two(context.Background(), &FirstGroup3ClientTwoOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
