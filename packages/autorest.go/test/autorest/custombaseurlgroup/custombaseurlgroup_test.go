// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

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
	client, err := NewPathsClient(to.Ptr(":3000"), &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewPathsClient(host *string, options *azcore.ClientOptions) (*PathsClient, error) {
	client, err := azcore.NewClient("custombaseurlgroup.PathsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	if host == nil {
		host = to.Ptr("host")
	}
	return &PathsClient{internal: client, host: *host}, nil
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetEmpty(context.Background(), "localhost", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
