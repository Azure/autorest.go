// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newHTTPFailureClient(t *testing.T) *HTTPFailureClient {
	client, err := NewHTTPFailureClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewHTTPFailureClient(options *azcore.ClientOptions) (*HTTPFailureClient, error) {
	client, err := azcore.NewClient("httpinfrastructuregroup.HTTPFailureClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPFailureClient{internal: client}, nil
}

func TestHTTPFailureGetEmptyError(t *testing.T) {
	client := newHTTPFailureClient(t)
	result, err := client.GetEmptyError(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPFailureGetNoModelEmpty(t *testing.T) {
	client := newHTTPFailureClient(t)
	result, err := client.GetNoModelEmpty(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPFailureGetNoModelError(t *testing.T) {
	client := newHTTPFailureClient(t)
	result, err := client.GetNoModelError(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}
