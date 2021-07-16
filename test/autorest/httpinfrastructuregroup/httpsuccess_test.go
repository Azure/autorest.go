// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"net/http"
	"testing"
)

func newHTTPSuccessClient() *HTTPSuccessClient {
	return NewHTTPSuccessClient(NewDefaultConnection(nil))
}

func TestHTTPSuccessDelete200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessDelete202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessDelete204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessGet200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Get200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessHead200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessHead204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessHead404(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Head404(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPSuccessOptions200(t *testing.T) {
	t.Skip("options method not enabled by test server")
	client := newHTTPSuccessClient()
	result, err := client.Options200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPatch200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPatch202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPatch204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPost200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPost201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post201(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusCreated {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPost202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPost204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPut200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPut201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put201(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusCreated {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPut202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHTTPSuccessPut204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}
