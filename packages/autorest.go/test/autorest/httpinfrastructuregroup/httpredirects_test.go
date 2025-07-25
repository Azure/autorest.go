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

func newHTTPRedirectsClient(t *testing.T) *HTTPRedirectsClient {
	client, err := NewHTTPRedirectsClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewHTTPRedirectsClient(endpoint string, options *azcore.ClientOptions) (*HTTPRedirectsClient, error) {
	client, err := azcore.NewClient("httpinfrastructuregroup.HTTPRedirectsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPRedirectsClient{internal: client, endpoint: endpoint}, nil
}

func TestHTTPRedirectsDelete307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Delete307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsGet300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient(t)
	result, err := client.Get300(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsGet301(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Get301(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsGet302(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Get302(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsGet307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Get307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsHead300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient(t)
	result, err := client.Head300(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsHead301(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Head301(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("expected Success")
	}
}

func TestHTTPRedirectsHead302(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Head302(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRedirectsHead307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Head307(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRedirectsOptions307(t *testing.T) {
	t.Skip("receive a status code of 204 which is not expected")
	client := newHTTPRedirectsClient(t)
	result, err := client.Options307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPatch302(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient(t)
	result, err := client.Patch302(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPatch307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Patch307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPost303(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Post303(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPost307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Post307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPut301(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient(t)
	result, err := client.Put301(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRedirectsPut307(t *testing.T) {
	client := newHTTPRedirectsClient(t)
	result, err := client.Put307(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
