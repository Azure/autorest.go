// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package custombaseurlgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newPathsClient() *PathsClient {
	pl := runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPathsClient(to.Ptr(":3000"), pl)
}

func TestGetEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetEmpty(context.Background(), "localhost", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
