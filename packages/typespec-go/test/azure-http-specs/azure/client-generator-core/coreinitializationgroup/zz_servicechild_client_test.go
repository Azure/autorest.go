//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"errors"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestServiceChildClient_DeleteStandalone_Success(t *testing.T) {
	// Arrange
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()

	// Act
	resp, err := client.DeleteStandalone(context.Background(), blobName, nil)

	// Assert
	require.NoError(t, err)
	require.Equal(t, ServiceChildClientDeleteStandaloneResponse{}, resp)
}

func TestServiceChildClient_DeleteStandalone_EmptyBlobName(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()

	_, err = client.DeleteStandalone(context.Background(), "", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parameter blobName cannot be empty")
}

func TestServiceChildClient_DeleteStandalone_ErrorFromPipeline(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()

	_, err = client.DeleteStandalone(context.Background(), blobName, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "pipeline error")
}

func TestServiceChildClient_DeleteStandalone_Non204Status(t *testing.T) {
	blobName := "testblob"
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceChildClient()

	_, err = client.DeleteStandalone(context.Background(), blobName, nil)
	require.Error(t, err)
	var respErr *azcore.ResponseError
	require.True(t, errors.As(err, &respErr))
}
