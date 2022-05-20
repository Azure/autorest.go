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

func newHTTPFailureClient() *HTTPFailureClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewHTTPFailureClient(pl)
}

func TestHTTPFailureGetEmptyError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetEmptyError(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPFailureGetNoModelEmpty(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelEmpty(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPFailureGetNoModelError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelError(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}
