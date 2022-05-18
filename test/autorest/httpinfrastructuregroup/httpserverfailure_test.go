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

func newHTTPServerFailureClient() *HTTPServerFailureClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewHTTPServerFailureClient(pl)
}

func TestHTTPServerFailureDelete505(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Delete505(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPServerFailureGet501(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Get501(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestHTTPServerFailureHead501(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Head501(context.Background(), nil)
	require.Error(t, err)
	require.False(t, result.Success)
}

func TestHTTPServerFailurePost505(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Post505(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}
