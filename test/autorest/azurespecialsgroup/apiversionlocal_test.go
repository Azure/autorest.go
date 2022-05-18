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

func newAPIVersionLocalClient() *APIVersionLocalClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewAPIVersionLocalClient(pl)
}

// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
func TestGetMethodLocalNull(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetMethodLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetPathLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetPathLocalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetSwaggerLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetSwaggerLocalValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
