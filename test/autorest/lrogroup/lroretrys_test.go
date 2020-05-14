// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogrouptest

import (
	"generatortests/autorest/generated/lrogroup"
	"testing"
	"time"
)

func getLRORetrysOperations(t *testing.T) lrogroup.LroRetrysOperations {
	options := lrogroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	client, err := lrogroup.NewDefaultClient(&options)
	if err != nil {
		t.Fatalf("failed to create lro client: %v", err)
	}
	return client.LroRetrysOperations()
}

// func TestLRORetrysBeginDelete202Retry200(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginDelete202Retry200(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysDelete202Retry200Poller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginDeleteAsyncRelativeRetrySucceeded(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginDeleteAsyncRelativeRetrySucceeded(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysDeleteAsyncRelativeRetrySucceededPoller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysDeleteProvisioning202Accepted200SucceededPoller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginPost202Retry200(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginPost202Retry200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysPost202Retry200Poller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginPostAsyncRelativeRetrySucceeded(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginPostAsyncRelativeRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysPostAsyncRelativeRetrySucceededPoller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginPut201CreatingSucceeded200(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysPut201CreatingSucceeded200Poller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLRORetrysBeginPutAsyncRelativeRetrySucceeded(t *testing.T) {
// 	op := getLRORetrysOperations(t)
// 	poller, err := op.BeginPutAsyncRelativeRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLroRetrysPutAsyncRelativeRetrySucceededPoller(rt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for poller.Poll(context.Background()) {
// 		time.Sleep(200 * time.Millisecond)
// 	}
// 	resp, err := poller.Response()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }
