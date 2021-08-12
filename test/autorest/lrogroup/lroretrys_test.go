// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newLRORetrysClient() *LRORetrysClient {
	options := ConnectionOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewLRORetrysClient(NewDefaultConnection(&options))
}

func TestLRORetrysBeginDelete202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDelete202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysDelete202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLRORetrysBeginDeleteAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysDeleteAsyncRelativeRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLRORetrysBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysDeleteProvisioning202Accepted200SucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(res.Product, Product{
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

func TestLRORetrysBeginPost202Retry200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysPost202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLRORetrysBeginPostAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysPostAsyncRelativeRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLRORetrysBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysPut201CreatingSucceeded200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(res.Product, Product{
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

func TestLRORetrysBeginPutAsyncRelativeRetrySucceeded(t *testing.T) {
	op := newLRORetrysClient()
	resp, err := op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LRORetrysPutAsyncRelativeRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(res.Product, Product{
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
