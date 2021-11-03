// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

func newHTTPRetryClient() *HTTPRetryClient {
	options := azcore.ClientOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.Transport = httpClientWithCookieJar()
	return NewHTTPRetryClient(&options)
}

func httpClientWithCookieJar() policy.Transporter {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return http.DefaultClient
}

func TestHTTPRetryDelete503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Delete503(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryGet502(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Get502(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryHead408(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Head408(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRetryOptions502(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := newHTTPRetryClient()
	result, err := client.Options502(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryPatch500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch500(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryPatch504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Patch504(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryPost503(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Post503(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryPut500(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put500(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPRetryPut504(t *testing.T) {
	client := newHTTPRetryClient()
	result, err := client.Put504(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
