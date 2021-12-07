// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newLROSClient() *LROsClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.Transport = httpClientWithCookieJar()
	return NewLROsClient(&options)
}

func httpClientWithCookieJar() policy.Transporter {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return http.DefaultClient
}

func TestLROResumeWrongPoller(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp2 := LROsDelete202Retry200PollerResponse{}
	if err = resp2.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not find receive one")
	}
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDelete202NoRetry204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDelete202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete204Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteAsyncNoHeaderInRetryPollerResponse{}
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

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteAsyncNoRetrySucceededPollerResponse{}
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

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteAsyncRetrySucceededPollerResponse{}
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

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteAsyncRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsDeleteProvisioning202Accepted200SucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost200WithPayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPost200WithPayloadPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPost202ListPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.ProductArray, []*Product{
		{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPost202NoRetry204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPost202Retry200PollerResponse{}
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

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostAsyncNoRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostAsyncRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostAsyncRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostDoubleHeadersFinalAzureHeaderGetPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID: to.StringPtr("100"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostDoubleHeadersFinalAzureHeaderGetDefaultPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPostDoubleHeadersFinalLocationGetPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200SucceededNoState(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPut200UpdatingSucceeded204PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	if err = resp.Resume(context.Background(), op, rt); err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), &LROsBeginPut201CreatingSucceeded200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPut201CreatingSucceeded200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut202Retry200(context.Background(), &LROsBeginPut202Retry200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPut202Retry200PollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), &LROsBeginPutAsyncNoHeaderInRetryOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), &LROsBeginPutAsyncNoRetrySucceededOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncNoRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncNoRetrycanceledPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncNonResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncNonResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncRetryFailedPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if !reflect.ValueOf(res).IsZero() {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncRetrySucceededPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("foo"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutAsyncSubResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SubProduct, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutNoHeaderInRetryPollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, Product{
		ID: to.StringPtr("100"),
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNonResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutNonResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutSubResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	resp = LROsPutSubResourcePollerResponse{}
	if err = resp.Resume(context.Background(), op, rt); err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SubProduct, SubProduct{
		ID: to.StringPtr("100"),
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
