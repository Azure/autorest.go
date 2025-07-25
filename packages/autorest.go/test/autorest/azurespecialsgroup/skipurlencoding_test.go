// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newSkipURLEncodingClient(t *testing.T) *SkipURLEncodingClient {
	client, err := NewSkipURLEncodingClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewSkipURLEncodingClient(endpoint string, options *azcore.ClientOptions) (*SkipURLEncodingClient, error) {
	client, err := azcore.NewClient("azurespecialsgroup.SkipURLEncodingClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SkipURLEncodingClient{internal: client, endpoint: endpoint}, nil
}

// GetMethodPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetMethodPathValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetMethodPathValid(context.Background(), "path1/path2/path3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodQueryNull - Get method with unencoded query parameter with value null
func TestGetMethodQueryNull(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetMethodQueryNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetMethodQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetMethodQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetPathQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetPathQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetPathValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetPathValid(context.Background(), "path1/path2/path3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetSwaggerPathValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetSwaggerPathValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetSwaggerQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient(t)
	result, err := client.GetSwaggerQueryValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
