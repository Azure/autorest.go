// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func getHTTPRetryOperations(t *testing.T) httpinfrastructuregroup.HTTPRetryOperations {
	client, err := httpinfrastructuregroup.NewDefaultClient(&httpinfrastructuregroup.ClientOptions{Retry: azcore.RetryOptions{MaxRetries: 5, RetryDelay: 10 * time.Millisecond, MaxRetryDelay: 15 * time.Millisecond, TryTimeout: 10 * time.Millisecond}})
	if err != nil {
		t.Fatalf("failed to create HTTPRetry client: %v", err)
	}
	return client.HTTPRetryOperations()
}

// TODO check should this be 200? or 503?
func TestHTTPRetryDelete503(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Delete503(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Did not expect a response")
	}
}

func TestHTTPRetryGet502(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Get502(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusBadGateway)
}

func TestHTTPRetryHead408(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Head408(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryOptions502(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Options502(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRetryPatch500(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Patch500(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPatch504(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Patch504(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPost503(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Post503(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPut500(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Put500(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRetryPut504(t *testing.T) {
	client := getHTTPRetryOperations(t)
	result, err := client.Put504(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
