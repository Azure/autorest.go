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

// ORIGINAL method according to current guidelines
// TODO fix poll func to check for error in poller.go
func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDelete202NoRetry204(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for {
		w, err := poller.Poll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if poller.Done() {
			helpers.VerifyStatusCode(t, w.RawResponse, 204)
			break
		}
		time.Sleep(1 * time.Second)
	}

}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := getLROSOperations(t)
	poller, err := op.BeginDelete202Retry200(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for {
		w, err := poller.Poll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if poller.Done() {
			helpers.VerifyStatusCode(t, w.RawResponse, 200)
			break
		}
		time.Sleep(1 * time.Second)
	}

}
