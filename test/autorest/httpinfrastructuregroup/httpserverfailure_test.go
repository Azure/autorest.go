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

func newHTTPServerFailureClient(t *testing.T) *HTTPServerFailureClient {
	client, err := NewHTTPServerFailureClient(nil)
	require.NoError(t, err)
	return client
}

func NewHTTPServerFailureClient(options *azcore.ClientOptions) (*HTTPServerFailureClient, error) {
	client, err := azcore.NewClient("httpinfrastructuregroup.HTTPServerFailureClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPServerFailureClient{internal: client}, nil
}

func TestHTTPServerFailureDelete505(t *testing.T) {
	client := newHTTPServerFailureClient(t)
	result, err := client.Delete505(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPServerFailureGet501(t *testing.T) {
	client := newHTTPServerFailureClient(t)
	result, err := client.Get501(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPServerFailureHead501(t *testing.T) {
	client := newHTTPServerFailureClient(t)
	result, err := client.Head501(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPServerFailurePost505(t *testing.T) {
	client := newHTTPServerFailureClient(t)
	result, err := client.Post505(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}
