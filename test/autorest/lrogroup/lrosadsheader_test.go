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

func getLrosaDsOperations(t *testing.T) lrogroup.LrosaDsOperations {
	options := lrogroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	client, err := lrogroup.NewDefaultClient(&options)
	if err != nil {
		t.Fatalf("failed to create lro client: %v", err)
	}
	return client.LrosaDsOperations()
}

func TestLROSADSBeginPost202Retry200(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginDelete202NonRetry400(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsDelete202NonRetry400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginDelete202RetryInvalidHeader(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginDelete202RetryInvalidHeader(context.Background())
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDelete204Succeeded(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginDelete204Succeeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, 204)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, 204)
}

func TestLROSADSBeginDeleteAsyncRelativeRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginDeleteAsyncRelativeRetry400(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsDeleteAsyncRelativeRetry400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginDeleteAsyncRelativeRetryInvalidHeader(context.Background())
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginDeleteAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginDeleteAsyncRelativeRetryInvalidJSONPolling(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsDeleteAsyncRelativeRetryInvalidJsonPollingPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginDeleteAsyncRelativeRetryNoStatus(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginDeleteAsyncRelativeRetryNoStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsDeleteAsyncRelativeRetryNoStatusPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginDeleteNonRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginDeleteNonRetry400(context.Background())
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NoLocation(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPost202NoLocation(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPost202NonRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPost202NonRetry400(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPost202NonRetry400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginPost202RetryInvalidHeader(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPost202RetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPostAsyncRelativeRetry400(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPostAsyncRelativeRetry400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPostAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPostAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPostAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPostAsyncRelativeRetryInvalidJsonPollingPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPostAsyncRelativeRetryNoPayload(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPostAsyncRelativeRetryNoPayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPostAsyncRelativeRetryNoPayloadPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPostNonRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPostNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPut200InvalidJSON(t *testing.T) {
	t.Skip("need to check response handling in the poller")
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPut200InvalidJSON(context.Background(), nil)
	if err != nil {
		t.Fatal("expected an error but did not receive one")
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPut200InvalidJsonPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPutAsyncRelativeRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutAsyncRelativeRetry400(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutAsyncRelativeRetry400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidHeader(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPutAsyncRelativeRetryInvalidHeader(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetryInvalidJSONPolling(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutAsyncRelativeRetryInvalidJSONPolling(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutAsyncRelativeRetryInvalidJsonPollingPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	_, err = poller.Response()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	_, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatus(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutAsyncRelativeRetryNoStatus(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutAsyncRelativeRetryNoStatusPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPutAsyncRelativeRetryNoStatusPayload(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutAsyncRelativeRetryNoStatusPayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutAsyncRelativeRetryNoStatusPayloadPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPutError201NoProvisioningStatePayload(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutError201NoProvisioningStatePayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutError201NoProvisioningStatePayloadPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROSADSBeginPutNonRetry201Creating400(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutNonRetry201Creating400(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutNonRetry201Creating400Poller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	resp, err := poller.Response()
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 400)
}

func TestLROSADSBeginPutNonRetry201Creating400InvalidJSON(t *testing.T) {
	op := getLrosaDsOperations(t)
	poller, err := op.BeginPutNonRetry201Creating400InvalidJSON(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeLrosaDsPutNonRetry201Creating400InvalidJsonPoller(rt)
	if err != nil {
		t.Fatal(err)
	}
	for poller.Poll(context.Background()) {
		time.Sleep(200 * time.Millisecond)
	}
	_, err = poller.Response()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	_, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROSADSBeginPutNonRetry400(t *testing.T) {
	op := getLrosaDsOperations(t)
	_, err := op.BeginPutNonRetry400(context.Background(), nil)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}
