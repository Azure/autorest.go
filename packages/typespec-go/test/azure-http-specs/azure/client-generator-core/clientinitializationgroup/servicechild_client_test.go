//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup

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
	require.Error(t, err)
	require.Equal(t, ServiceChildClientDeleteStandaloneResponse{}, resp)
	_, err = client.DeleteStandalone(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
	_, err = client.DeleteStandalone(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}

func TestServiceChildClient_GetStandalone(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()
	resp, err := client.GetStandalone(context.Background(), blobName, nil)
	require.Contains(t, err.Error(), "Not Found")
	require.Equal(t, ServiceChildClientGetStandaloneResponse{}, resp)
	_, err = client.GetStandalone(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
	_, err = client.GetStandalone(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}

func TestServiceChildClient_WithQuery(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()
	resp, err := client.WithQuery(context.Background(), blobName, nil)
	require.Contains(t, err.Error(), "Not Found")
	require.Equal(t, ServiceChildClientWithQueryResponse{}, resp)
	_, err = client.WithQuery(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
	_, err = client.WithQuery(context.Background(), "sample-blob", nil)
	require.Error(t, err)
}
