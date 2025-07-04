package clientopgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestFirstGroup3Client_Two(t *testing.T) {
	factory := &FirstClient{internal: &azcore.Client{}}
	client := factory.NewFirstGroup3Client()
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Two(context.Background(), &FirstGroup3ClientTwoOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}

func TestFirstGroup3Client_Three(t *testing.T) {
	factory := &FirstClient{internal: &azcore.Client{}}
	client := factory.NewFirstGroup3Client()
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Three(context.Background(), &FirstGroup3ClientThreeOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}

func TestFirstGroup4Client_Two(t *testing.T) {
	factory := &FirstClient{internal: &azcore.Client{}}
	client := factory.NewFirstGroup4Client()
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Four(context.Background(), &FirstGroup4ClientFourOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}

func TestFirstGroup5Client_Three(t *testing.T) {
	factory := &SecondClient{internal: &azcore.Client{}}
	client := factory.NewSecondGroup5Client()
	require.NotNil(t, client)
	client.endpoint = "http://localhost:3000"
	client.client = ClientType(ClientTypeClientOperationGroup)
	resp, err := client.Six(context.Background(), &SecondGroup5ClientSixOptions{})
	require.Nil(t, err)
	require.NotNil(t, resp)
}
