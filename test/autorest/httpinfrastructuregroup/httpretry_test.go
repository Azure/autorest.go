// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func newHTTPRetryClient() HTTPRetryOperations {
	options := DefaultClientOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewHTTPRetryClient(NewDefaultClient(&options))
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

func TestHTTPRetryDelete503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Delete503(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryGet502(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Get502(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryHead408(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Head408(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryOptions502(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := newHTTPRetryClient()
	result, err := client.Options502(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRetryPatch500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch500(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPatch504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch504(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPost503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Post503(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPut500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put500(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPut504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put504(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
