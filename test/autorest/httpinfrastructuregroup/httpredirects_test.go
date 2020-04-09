// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getHTTPRedirectsOperations(t *testing.T) httpinfrastructuregroup.HTTPRedirectsOperations {
	client, err := httpinfrastructuregroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create HTTPRedirects client: %v", err)
	}
	return client.HTTPRedirectsOperations()
}

func TestHTTPRedirectsDelete307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Delete307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsGet300(t *testing.T) {
	t.Skip()
	client := getHTTPRedirectsOperations(t)
	result, err := client.Get300(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsGet301(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Get301(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsGet302(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Get302(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsGet307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Get307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsHead300(t *testing.T) {
	t.Skip()
	client := getHTTPRedirectsOperations(t)
	result, err := client.Head300(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsHead301(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Head301(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsHead302(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Head302(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsHead307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Head307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsOptions307(t *testing.T) {
	t.Skip()
	client := getHTTPRedirectsOperations(t)
	result, err := client.Options307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPatch302(t *testing.T) {
	t.Skip()
	client := getHTTPRedirectsOperations(t)
	result, err := client.Patch302(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPatch307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Patch307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPost303(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Post303(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPost307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Post307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPut301(t *testing.T) {
	t.Skip()
	client := getHTTPRedirectsOperations(t)
	result, err := client.Put301(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPut307(t *testing.T) {
	client := getHTTPRedirectsOperations(t)
	result, err := client.Put307(context.Background())
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
