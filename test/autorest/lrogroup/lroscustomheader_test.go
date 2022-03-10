// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newLrOSCustomHeaderClient() *LROsCustomHeaderClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = time.Second
	options.Transport = httpClientWithCookieJar()
	return NewLROsCustomHeaderClient(&options)
}

func ctxWithHTTPHeader() context.Context {
	header := http.Header{}
	header.Add("x-ms-client-request-id", "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	return runtime.WithHTTPHeader(context.Background(), header)
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func TestBeginPost202Retry200(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPost202Retry200(ctxWithHTTPHeader(), nil)
	require.NoError(t, err)
	tk, err := env.Poller.ResumeToken()
	require.NoError(t, err)
	env = LROsCustomHeaderClientPost202Retry200PollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		require.NoError(t, err)
		if env.Poller.Done() {
			break
		}
	}
	_, err = env.Poller.FinalResponse(context.Background())
	require.NoError(t, err)
}

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func TestBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPostAsyncRetrySucceeded(ctxWithHTTPHeader(), nil)
	require.NoError(t, err)
	tk, err := env.Poller.ResumeToken()
	require.NoError(t, err)
	env = LROsCustomHeaderClientPostAsyncRetrySucceededPollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		require.NoError(t, err)
		if env.Poller.Done() {
			break
		}
	}
	_, err = env.Poller.FinalResponse(context.Background())
	require.NoError(t, err)
}

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func TestBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPut201CreatingSucceeded200(ctxWithHTTPHeader(), nil)
	require.NoError(t, err)
	tk, err := env.Poller.ResumeToken()
	require.NoError(t, err)
	env = LROsCustomHeaderClientPut201CreatingSucceeded200PollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		require.NoError(t, err)
		if env.Poller.Done() {
			break
		}
	}
	pr, err := env.Poller.FinalResponse(ctxWithHTTPHeader())
	require.NoError(t, err)
	if r := cmp.Diff(pr.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func TestBeginPutAsyncRetrySucceeded(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPutAsyncRetrySucceeded(ctxWithHTTPHeader(), nil)
	require.NoError(t, err)
	tk, err := env.Poller.ResumeToken()
	require.NoError(t, err)
	env = LROsCustomHeaderClientPutAsyncRetrySucceededPollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		require.NoError(t, err)
		if env.Poller.Done() {
			break
		}
	}
	pr, err := env.Poller.FinalResponse(ctxWithHTTPHeader())
	require.NoError(t, err)
	if r := cmp.Diff(pr.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
