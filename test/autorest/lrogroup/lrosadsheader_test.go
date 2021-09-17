// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func newLrosaDsClient() *LROSADsClient {
	options := ConnectionOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewLROSADsClient(NewDefaultConnection(&options))
}

func TestLROSADSBeginDelete202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDelete202NonRetry400(context.Background(), nil)
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

func TestLROSADSBeginDelete202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDelete202RetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDelete204Succeeded(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDelete204Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if rt != "" {
		t.Fatal("expected an empty resume token")
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetry400(context.Background(), nil)
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

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
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

func TestLROSADSBeginDeleteAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background(), nil)
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

func TestLROSADSBeginDeleteNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NoLocation(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202NoLocation(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPost202NonRetry400(context.Background(), nil)
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

func TestLROSADSBeginPost202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPost202RetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPostAsyncRelativeRetry400(context.Background(), nil)
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

func TestLROSADSBeginPostAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPostAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
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

func TestLROSADSBeginPostAsyncRelativeRetryNoPayload(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPostAsyncRelativeRetryNoPayload(context.Background(), nil)
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

func TestLROSADSBeginPostNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPostNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPut200InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPut200InvalidJSON(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutAsyncRelativeRetry400(context.Background(), nil)
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

func TestLROSADSBeginPutAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
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

func TestLROSADSBeginPutAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutAsyncRelativeRetryNoStatus(context.Background(), nil)
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

func TestLROSADSBeginPutAsyncRelativeRetryNoStatusPayload(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutAsyncRelativeRetryNoStatusPayload(context.Background(), nil)
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

func TestLROSADSBeginPutError201NoProvisioningStatePayload(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutError201NoProvisioningStatePayload(context.Background(), nil)
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

func TestLROSADSBeginPutNonRetry201Creating400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutNonRetry201Creating400(context.Background(), nil)
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

func TestLROSADSBeginPutNonRetry201Creating400InvalidJSON(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginPutNonRetry201Creating400InvalidJSON(context.Background(), nil)
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

func TestLROSADSBeginPutNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}
