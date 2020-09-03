// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogrouptest

import (
	"context"
	"generatortests/autorest/generated/lrogroup"
	"generatortests/helpers"
	"testing"
	"time"
)

func newLrosaDsClient() lrogroup.LrosaDsOperations {
	options := lrogroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return lrogroup.NewLrosaDsClient(lrogroup.NewDefaultClient(&options))
}

func TestLROSADSBeginDelete202NonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDelete202NonRetry400(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202NonRetry400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
	}
}

func TestLROSADSBeginDelete202RetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDelete202RetryInvalidHeader(context.Background())
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDelete204Succeeded(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDelete204Succeeded(context.Background())
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
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 204)
}

func TestLROSADSBeginDeleteAsyncRelativeRetry400(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetry400(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRelativeRetry400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteAsyncRelativeRetryInvalidHeader(context.Background())
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRelativeRetryInvalidJSONPolling(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryNoStatus(t *testing.T) {
	op := newLrosaDsClient()
	resp, err := op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRelativeRetryNoStatus(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
	}
}

func TestLROSADSBeginDeleteNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginDeleteNonRetry400(context.Background())
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
	poller, err = op.ResumePost202NonRetry400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePostAsyncRelativeRetry400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePostAsyncRelativeRetryInvalidJSONPolling(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePostAsyncRelativeRetryNoPayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutAsyncRelativeRetry400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutAsyncRelativeRetryInvalidJSONPolling(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutAsyncRelativeRetryNoStatus(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutAsyncRelativeRetryNoStatusPayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutError201NoProvisioningStatePayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutNonRetry201Creating400(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
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
	poller, err = op.ResumePutNonRetry201Creating400InvalidJSON(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response with the error")
	}
}

func TestLROSADSBeginPutNonRetry400(t *testing.T) {
	op := newLrosaDsClient()
	_, err := op.BeginPutNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}
