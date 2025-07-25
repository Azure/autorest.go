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

func newAPIVersionDefaultClient(t *testing.T) *APIVersionDefaultClient {
	client, err := NewAPIVersionDefaultClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewAPIVersionDefaultClient(endpoint string, options *azcore.ClientOptions) (*APIVersionDefaultClient, error) {
	client, err := azcore.NewClient("azurespecialsgroup.APIVersionDefaultClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &APIVersionDefaultClient{internal: client, endpoint: endpoint}, nil
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalNotProvidedValid(t *testing.T) {
	client := newAPIVersionDefaultClient(t)
	result, err := client.GetMethodGlobalNotProvidedValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient(t)
	result, err := client.GetMethodGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func TestGetPathGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient(t)
	result, err := client.GetPathGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func TestGetSwaggerGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient(t)
	result, err := client.GetSwaggerGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
