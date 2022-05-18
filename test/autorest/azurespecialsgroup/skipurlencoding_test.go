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

func newSkipURLEncodingClient() *SkipURLEncodingClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewSkipURLEncodingClient(pl)
}

// GetMethodPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetMethodPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodPathValid(context.Background(), "path1/path2/path3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodQueryNull - Get method with unencoded query parameter with value null
func TestGetMethodQueryNull(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodQueryNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetMethodQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetMethodQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetPathQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetPathQueryValid(context.Background(), "value1&q2=value2&q3=value3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetPathValid(context.Background(), "path1/path2/path3", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
func TestGetSwaggerPathValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetSwaggerPathValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
func TestGetSwaggerQueryValid(t *testing.T) {
	client := newSkipURLEncodingClient()
	result, err := client.GetSwaggerQueryValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
