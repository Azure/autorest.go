// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newLrOSCustomHeaderClient() *LROsCustomHeaderClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.Transport = httpClientWithCookieJar()
	return NewLROsCustomHeaderClient(NewDefaultConnection(&options))
}

func ctxWithHTTPHeader() context.Context {
	header := http.Header{}
	header.Add("x-ms-client-request-id", "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	return policy.WithHTTPHeader(context.Background(), header)
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func TestBeginPost202Retry200(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPost202Retry200(ctxWithHTTPHeader(), nil)
	if err != nil {
		t.Fatal(err)
	}
	tk, err := env.Poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	env = LROsCustomHeaderPost202Retry200PollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		if err != nil {
			t.Fatal(err)
		}
		if env.Poller.Done() {
			break
		}
	}
	resp, err := env.Poller.FinalResponse(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func TestBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPostAsyncRetrySucceeded(ctxWithHTTPHeader(), nil)
	if err != nil {
		t.Fatal(err)
	}
	tk, err := env.Poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	env = LROsCustomHeaderPostAsyncRetrySucceededPollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		if err != nil {
			t.Fatal(err)
		}
		if env.Poller.Done() {
			break
		}
	}
	resp, err := env.Poller.FinalResponse(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func TestBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLrOSCustomHeaderClient()
	env, err := op.BeginPut201CreatingSucceeded200(ctxWithHTTPHeader(), nil)
	if err != nil {
		t.Fatal(err)
	}
	tk, err := env.Poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	env = LROsCustomHeaderPut201CreatingSucceeded200PollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		if err != nil {
			t.Fatal(err)
		}
		if env.Poller.Done() {
			break
		}
	}
	pr, err := env.Poller.FinalResponse(ctxWithHTTPHeader())
	if err != nil {
		t.Fatal(err)
	}
	if s := pr.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pr.Product, Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
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
	if err != nil {
		t.Fatal(err)
	}
	tk, err := env.Poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	env = LROsCustomHeaderPutAsyncRetrySucceededPollerResponse{}
	if err = env.Resume(ctxWithHTTPHeader(), op, tk); err != nil {
		t.Fatal(err)
	}
	for {
		_, err = env.Poller.Poll(ctxWithHTTPHeader())
		if err != nil {
			t.Fatal(err)
		}
		if env.Poller.Done() {
			break
		}
	}
	pr, err := env.Poller.FinalResponse(ctxWithHTTPHeader())
	if err != nil {
		t.Fatal(err)
	}
	if s := pr.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pr.Product, Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
