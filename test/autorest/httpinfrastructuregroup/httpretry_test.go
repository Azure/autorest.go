// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newHTTPRetryClient() *HTTPRetryClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.Transport = httpClientWithCookieJar()
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &options)
	return NewHTTPRetryClient(pl)
}

func httpClientWithCookieJar() policy.Transporter {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return http.DefaultClient
}

func TestHTTPRetryDelete503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Delete503(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryGet502(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Get502(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryHead408(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Head408(context.Background(), nil)
	require.NoError(t, err)
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRetryOptions502(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := newHTTPRetryClient()
	result, err := client.Options502(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryPatch500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch500(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryPatch504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch504(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryPost503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Post503(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryPut500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put500(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHTTPRetryPut504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put504(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
