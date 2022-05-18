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

func newAPIVersionDefaultClient() *APIVersionDefaultClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewAPIVersionDefaultClient(pl)
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalNotProvidedValid(t *testing.T) {
	client := newAPIVersionDefaultClient()
	result, err := client.GetMethodGlobalNotProvidedValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func TestGetMethodGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient()
	result, err := client.GetMethodGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func TestGetPathGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient()
	result, err := client.GetPathGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func TestGetSwaggerGlobalValid(t *testing.T) {
	client := newAPIVersionDefaultClient()
	result, err := client.GetSwaggerGlobalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
