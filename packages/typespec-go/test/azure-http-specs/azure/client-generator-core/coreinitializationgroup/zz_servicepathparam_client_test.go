// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServicePathParamClient_DeleteStandalone_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.DeleteStandalone(context.Background(), "blob1", nil)
	assert.NoError(t, err)
}

func TestServicePathParamClient_DeleteStandalone_EmptyBlobName(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.DeleteStandalone(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
}

func TestServicePathParamClient_DeleteStandalone_ErrorFromPipeline(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.DeleteStandalone(context.Background(), "blob1", nil)
	assert.Error(t, err)
	assert.Equal(t, "pipeline error", err.Error())
}

func TestServicePathParamClient_DeleteStandalone_Non204Status(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.DeleteStandalone(context.Background(), "blob1", nil)
	assert.Error(t, err)
}

func TestServicePathParamClient_GetStandalone_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	// Mocking a successful response
	_, err = client.GetStandalone(context.Background(), "blob1", nil)
	assert.NoError(t, err)
}

func TestServicePathParamClient_GetStandalone_EmptyBlobName(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.GetStandalone(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
}

func TestServicePathParamClient_GetStandalone_ErrorFromPipeline(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.GetStandalone(context.Background(), "blob1", nil)
	assert.Error(t, err)
	assert.Equal(t, "pipeline error", err.Error())
}

func TestServicePathParamClient_GetStandalone_Non200Status(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.GetStandalone(context.Background(), "blob1", nil)
	assert.Error(t, err)
}

func TestServicePathParamClient_WithQuery_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.WithQuery(context.Background(), "blob1", &ServicePathParamClientWithQueryOptions{FormatParam: nil})
	assert.NoError(t, err)
}

func TestServicePathParamClient_WithQuery_EmptyBlobName(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.WithQuery(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Equal(t, "parameter blobName cannot be empty", err.Error())
}

func TestServicePathParamClient_WithQuery_ErrorFromPipeline(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.WithQuery(context.Background(), "blob1", nil)
	assert.Error(t, err)
	assert.Equal(t, "pipeline error", err.Error())
}

func TestServicePathParamClient_WithQuery_Non204Status(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServicePathParamClient()
	_, err = client.WithQuery(context.Background(), "blob1", nil)
	assert.Error(t, err)
}
