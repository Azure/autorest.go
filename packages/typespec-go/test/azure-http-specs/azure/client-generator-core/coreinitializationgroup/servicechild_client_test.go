//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceChildClient_DeleteStandalone(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()
	resp, err := client.DeleteStandalone(context.Background(), blobName, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceChildClientDeleteStandaloneResponse{}, resp)
	_, err = client.DeleteStandalone(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
}

func TestServiceChildClient_GetStandalone(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()
	resp, err := client.GetStandalone(context.Background(), blobName, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceChildClientGetStandaloneResponse{}, resp)
	_, err = client.GetStandalone(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
}

func TestServiceChildClient_WithQuery(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()
	resp, err := client.WithQuery(context.Background(), blobName, nil)
	require.NoError(t, err)
	require.Equal(t, ServiceChildClientWithQueryResponse{}, resp)
	_, err = client.WithQuery(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
}
