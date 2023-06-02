// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newHTTPClientFailureClient(t *testing.T) *HTTPClientFailureClient {
	options := policy.ClientOptions{
		Retry: policy.RetryOptions{
			MaxRetryDelay: 2 * time.Second,
		},
		TracingProvider: generatortests.NewTracingProvider(t),
	}
	client, err := NewHTTPClientFailureClient(&options)
	require.NoError(t, err)
	return client
}

func NewHTTPClientFailureClient(options *azcore.ClientOptions) (*HTTPClientFailureClient, error) {
	client, err := azcore.NewClient("httpinfrastructuregroup.HTTPClientFailureClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &HTTPClientFailureClient{internal: client}, nil
}

func TestHTTPClientFailureDelete400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Delete400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureDelete407(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Delete407(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureDelete417(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Delete417(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet402(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get402(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet403(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get403(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet411(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get411(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet412(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get412(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureGet416(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Get416(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureHead400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Head400(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPClientFailureHead401(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Head401(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPClientFailureHead410(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Head410(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPClientFailureHead429(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Head429(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPClientFailureOptions400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Options400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureOptions403(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Options403(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailureOptions412(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Options412(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePatch400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Patch400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePatch405(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Patch405(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePatch414(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Patch414(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePost400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Post400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePost406(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Post406(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePost415(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Post415(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePut400(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Put400(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePut404(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Put404(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePut409(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Put409(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPClientFailurePut413(t *testing.T) {
	client := newHTTPClientFailureClient(t)
	result, err := client.Put413(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}
