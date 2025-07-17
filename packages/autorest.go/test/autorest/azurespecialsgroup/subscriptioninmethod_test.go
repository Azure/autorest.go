// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"generatortests"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newSubscriptionInMethodClient(t *testing.T) *SubscriptionInMethodClient {
	client, err := NewSubscriptionInMethodClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewSubscriptionInMethodClient(endpoint string, options *azcore.ClientOptions) (*SubscriptionInMethodClient, error) {
	client, err := azcore.NewClient("azurespecialsgroup.SubscriptionInMethodClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &SubscriptionInMethodClient{internal: client, endpoint: endpoint}, nil
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method.  pass in subscription id = null, client-side validation should prevent you from making this call
func TestPostMethodLocalNull(t *testing.T) {
	t.Skip("invalid test, missing x-nullable")
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostMethodLocalValid(t *testing.T) {
	client := newSubscriptionInMethodClient(t)
	result, err := client.PostMethodLocalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), "1234-5678-9012-3456", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostPathLocalValid(t *testing.T) {
	client := newSubscriptionInMethodClient(t)
	result, err := client.PostPathLocalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), "1234-5678-9012-3456", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func TestPostSwaggerLocalValid(t *testing.T) {
	client := newSubscriptionInMethodClient(t)
	result, err := client.PostSwaggerLocalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), "1234-5678-9012-3456", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
