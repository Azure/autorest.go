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

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDelete202NoRetry204(context.Background())
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
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp.RawResponse, 204)
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("did not receive an error but was expecting one")
	}
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDelete202Retry200(context.Background())
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
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("did not receive an error but was expecting one")
	}
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDelete204Succeeded(context.Background())
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
	helpers.VerifyStatusCode(t, resp, 204)
	resp, err = poller.Wait(context.Background(), time.Duration(1)*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, 204)
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("did not receive an error but was expecting one")
	}
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background())
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
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("did not receive an error but was expecting one")
	}
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background())
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
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("did not receive an error but was expecting one")
	}
}

// func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
// 	op := getLROSOperations(t)
// 	poller, err := op.BeginDeleteAsyncRetryFailed(context.Background())
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
// 	_, err = poller.ResumeToken()
// 	if err == nil {
// 		t.Fatal("did not receive an error but was expecting one")
// 	}
// }
