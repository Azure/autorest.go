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

func newHTTPSuccessClient() *HTTPSuccessClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewHTTPSuccessClient(pl)
}

func TestHTTPSuccessDelete200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete200(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessDelete202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete202(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessDelete204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete204(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessGet200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Get200(context.Background(), nil)
	require.NoError(t, err)
	if !*result.Value {
		t.Fatal("expected Success")
	}
}

func TestHTTPSuccessHead200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head200(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessHead204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head204(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessHead404(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head404(context.Background(), nil)
	require.NoError(t, err)
	if result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessOptions200(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := newHTTPSuccessClient()
	result, err := client.Options200(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPatch200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch200(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPatch202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch202(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPatch204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch204(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPost200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post200(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPost201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post201(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPost202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post202(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPost204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post204(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPut200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put200(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPut201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put201(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPut202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put202(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPSuccessPut204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put204(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
