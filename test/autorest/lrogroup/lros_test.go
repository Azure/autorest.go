// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogrouptest

import (
	"context"
	"generatortests/autorest/generated/lrogroup"
	"generatortests/helpers"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func getLROSOperations(t *testing.T) lrogroup.LrOSOperations {
	options := lrogroup.DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	client, err := lrogroup.NewDefaultClient(&options)
	if err != nil {
		t.Fatalf("failed to create lro client: %v", err)
	}
	return client.LrOSOperations()
}

func httpClientWithCookieJar() azcore.Transport {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return azcore.TransportFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req.WithContext(ctx))
	})
}

// func TestLROResumeWrongPoller(t *testing.T) {
// 	op := getLROSOperations(t)
// 	resp, err := op.Delete202NoRetry204(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller := resp.GetPoller()
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = op.ResumeProductPoller(rt)
// 	if err == nil {
// 		t.Fatal("expected an error but did not find receive one")
// 	}
// }

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDelete202NoRetry204(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDelete202Retry200(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDelete204Succeeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	t.Skip("pending on CloudError field fix")
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	t.Skip("CloudError unmarshalling is failing")
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteAsyncRetryFailed(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
}

func TestLROBeginDeleteAsyncRetrySucceeded(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteAsyncRetrySucceeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	t.Skip("CloudError unmarshalling is failing")
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteAsyncRetrycanceled(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROBeginDeleteNoHeaderInRetry(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteNoHeaderInRetry(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Accepted200Succeeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202DeletingFailed200(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Deletingcanceled200(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginPost200WithPayload(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost200WithPayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, resp.Sku, &lrogroup.Sku{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	})
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginPost202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := getLROSOperations(t)
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = resp.PollUntilDone(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
}

// func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostAsyncNoRetrySucceededPoller(rt)
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

// func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
// 	t.Skip("CloudError unmarshalling fails")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostAsyncRetryFailedPoller(rt)
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

// func TestLROBeginPostAsyncRetrySucceeded(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostAsyncRetrySucceededPoller(rt)
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

// func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
// 	t.Skip("CloudError unmarshalling failed")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostAsyncRetrycanceledPoller(rt)
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

// func TestLROBeginPostDoubleHeadersFinalAzureHeaderGet(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostDoubleHeadersFinalAzureHeaderGetPoller(rt)
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

// func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostDoubleHeadersFinalAzureHeaderGetDefaultPoller(rt)
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

// func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPostDoubleHeadersFinalLocationGetPoller(rt)
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

// func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
// 	t.Skip("missing error info returned for error")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPut200Acceptedcanceled200Poller(rt)
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

// func TestLROBeginPut200Succeeded(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut200Succeeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = poller.ResumeToken()
// 	if err == nil {
// 		t.Fatal("expected an error but did not receive one")
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

// func TestLROBeginPut200SucceededNoState(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut200SucceededNoState(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = poller.ResumeToken()
// 	if err == nil {
// 		t.Fatal("expected an error but did not receive one")
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

// func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPut200UpdatingSucceeded204Poller(rt)
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

// func TestLROBeginPut201CreatingFailed200(t *testing.T) {
// 	t.Skip("missing error info message returned for error")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPut201CreatingFailed200Poller(rt)
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

// func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut201CreatingSucceeded200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPut201CreatingSucceeded200Poller(rt)
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

// func TestLROBeginPut202Retry200(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPut202Retry200(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPut202Retry200Poller(rt)
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

// func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncNoHeaderInRetryPoller(rt)
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

// func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncNoRetrySucceededPoller(rt)
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

// func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
// 	t.Skip("CloudError unmarshalling failed")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncNoRetrycanceledPoller(rt)
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

// func TestLROBeginPutAsyncNonResource(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncNonResource(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncNonResourcePoller(rt)
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

// func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
// 	t.Skip("CloudError unmarshalling failed")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncRetryFailedPoller(rt)
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

// func TestLROBeginPutAsyncRetrySucceeded(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncRetrySucceededPoller(rt)
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

// func TestLROBeginPutAsyncSubResource(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutAsyncSubResource(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutAsyncSubResourcePoller(rt)
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

// func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
// 	t.Skip("The test needs to fix some underlying problems with the poller returning an error")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutNoHeaderInRetryPoller(rt)
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
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLROBeginPutNonResource(t *testing.T) {
// 	t.Skip("The test needs to fix some underlying problems with the poller returning an error")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutNonResource(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutNonResourcePoller(rt)
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
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }

// func TestLROBeginPutSubResource(t *testing.T) {
// 	t.Skip("The test needs to fix some underlying problems with the poller returning an error")
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginPutSubResource(context.Background(), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rt, err := poller.ResumeToken()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	poller, err = op.ResumeLrOSPutSubResourcePoller(rt)
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
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 202)
// 	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	helpers.VerifyStatusCode(t, resp.RawResponse, 200)
// }
