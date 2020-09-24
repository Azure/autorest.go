// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newHTTPRedirectsClient() HTTPRedirectsOperations {
	return NewHTTPRedirectsClient(NewDefaultClient(nil))
}

func TestHTTPRedirectsDelete307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Delete307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsGet300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient()
	result, err := client.Get300(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	switch x := result.(type) {
	case *HTTPRedirectsGet300Response:
		helpers.VerifyStatusCode(t, x.RawResponse, http.StatusOK)
	case *StringArrayResponse:
		helpers.VerifyStatusCode(t, x.RawResponse, http.StatusMultipleChoices)
	default:
		t.Fatalf("unhandled response type %v", x)
	}
}

func TestHTTPRedirectsGet301(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsGet302(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsGet307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsHead300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient()
	result, err := client.Head300(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsHead301(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsHead302(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsHead307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsOptions307(t *testing.T) {
	t.Skip("receive a status code of 204 which is not expected")
	client := newHTTPRedirectsClient()
	result, err := client.Options307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsPatch302(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient()
	result, err := client.Patch302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPatch307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Patch307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsPost303(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Post303(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPost307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Post307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHTTPRedirectsPut301(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient()
	result, err := client.Put301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHTTPRedirectsPut307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Put307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
