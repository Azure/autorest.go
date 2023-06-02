// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package morecustombaseurigroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newPathsClient(t *testing.T) *PathsClient {
	client, err := NewPathsClient("test12", to.Ptr(":3000"), &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewPathsClient(subscriptionID string, dnsSuffix *string, options *azcore.ClientOptions) (*PathsClient, error) {
	client, err := azcore.NewClient("morecustombaseurigroup.PathsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	if dnsSuffix == nil {
		dnsSuffix = to.Ptr("host")
	}
	return &PathsClient{
		internal:       client,
		dnsSuffix:      *dnsSuffix,
		subscriptionID: subscriptionID,
	}, nil
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient(t)
	// vault string, secret string, keyName string, options *PathsGetEmptyOptions
	result, err := client.GetEmpty(context.Background(), "http://localhost", "", "key1", &PathsClientGetEmptyOptions{
		KeyVersion: to.Ptr("v1"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}
