// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceParamAliasClient_WithAliasedName_Success(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceParamAliasClient()
	resp, err := client.WithAliasedName(context.Background(), "blobValue", nil)
	require.Contains(t, err.Error(), "Not Found")
	require.Equal(t, ServiceParamAliasClientWithAliasedNameResponse{}, resp)
	_, err = client.WithAliasedName(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}

func TestServiceParamAliasClient_WithAliasedName_EmptyBlob(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceParamAliasClient()
	_, err = client.WithAliasedName(context.Background(), "", nil)
	require.Error(t, err)
	require.Equal(t, "parameter blob cannot be empty", err.Error())
	require.Error(t, err)
	_, err = client.WithAliasedName(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}

func TestServiceParamAliasClient_WithOriginalName_EmptyBlobName(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceParamAliasClient()
	_, err = client.WithOriginalName(context.Background(), "", nil)
	require.Error(t, err)
	require.Equal(t, "parameter blobName cannot be empty", err.Error())
	_, err = client.WithOriginalName(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}

func TestServiceParamAliasClient_WithOriginalName_StatusCodeError(t *testing.T) {
	serviceClient, err := NewServiceClient(nil)
	require.NoError(t, err)
	client := serviceClient.NewServiceParamAliasClient()
	_, err = client.WithOriginalName(context.Background(), "blobNameValue", nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Not Found")
	_, err = client.WithOriginalName(context.Background(), "sample-blob", nil)
	require.NoError(t, err)
}
