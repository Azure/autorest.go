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

func newSubscriptionInCredentialsClient() *SubscriptionInCredentialsClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewSubscriptionInCredentialsClient("1234-5678-9012-3456", pl)
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostMethodGlobalNotProvidedValid(t *testing.T) {
	client := newSubscriptionInCredentialsClient()
	result, err := client.PostMethodGlobalNotProvidedValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to null, and client-side validation should prevent you from making this call
func TestPostMethodGlobalNull(t *testing.T) {
	t.Skip("invalid test, subscription ID is not x-nullable")
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostMethodGlobalValid(t *testing.T) {
	client := newSubscriptionInCredentialsClient()
	result, err := client.PostMethodGlobalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostPathGlobalValid(t *testing.T) {
	client := newSubscriptionInCredentialsClient()
	result, err := client.PostPathGlobalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials.  Set the credential subscriptionId to '1234-5678-9012-3456' to succeed
func TestPostSwaggerGlobalValid(t *testing.T) {
	client := newSubscriptionInCredentialsClient()
	result, err := client.PostSwaggerGlobalValid(runtime.WithHTTPHeader(context.Background(), http.Header{
		"x-ms-client-request-id": []string{"9C4D50EE-2D56-4CD3-8152-34347DC9F2B0"},
	}), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
