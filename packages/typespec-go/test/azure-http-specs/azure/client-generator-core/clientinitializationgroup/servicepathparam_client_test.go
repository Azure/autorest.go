// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServicePathParamClient_DeleteStandalone(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.DeleteStandalone(context.Background(), "blob1", nil)
	require.Contains(t, err.Error(), "Not Found")
	_, err = client.DeleteStandalone(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
	_, err = client.DeleteStandalone(context.Background(), "blob1", nil)
	assert.Error(t, err)
	_, err = client.DeleteStandalone(context.Background(), "sample-blob", nil)
	assert.NoError(t, err)
}

func TestServicePathParamClient_GetStandalone(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.GetStandalone(context.Background(), "blob1", nil)
	require.Contains(t, err.Error(), "Not Found")
	_, err = client.GetStandalone(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
	_, err = client.GetStandalone(context.Background(), "sample-blob", nil)
	assert.NoError(t, err)
}

func TestServicePathParamClient_WithQuery(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.WithQuery(context.Background(), "blob1", &ServicePathParamClientWithQueryOptions{Format: to.Ptr("text")})
	require.Contains(t, err.Error(), "Not Found")
	_, err = client.WithQuery(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
	_, err = client.WithQuery(context.Background(), "sample-blob", &ServicePathParamClientWithQueryOptions{Format: to.Ptr("text")})
	assert.NoError(t, err)
}
