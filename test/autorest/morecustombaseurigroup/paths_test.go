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

func newPathsClient() *PathsClient {
	// dnsSuffix string, subscriptionID string
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPathsClient(to.Ptr(":3000"), "test12", pl)
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	// vault string, secret string, keyName string, options *PathsGetEmptyOptions
	result, err := client.GetEmpty(context.Background(), "http://localhost", "", "key1", &PathsClientGetEmptyOptions{
		KeyVersion: to.Ptr("v1"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}
