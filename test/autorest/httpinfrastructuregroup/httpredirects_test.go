// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"reflect"
	"testing"
)

func newHTTPRedirectsClient() *HTTPRedirectsClient {
	return NewHTTPRedirectsClient(nil)
}

func TestHTTPRedirectsDelete307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Delete307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsGet300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient()
	result, err := client.Get300(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsGet301(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsGet302(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsGet307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Get307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsHead300(t *testing.T) {
	t.Skip("does not automatically redirect")
	client := newHTTPRedirectsClient()
	result, err := client.Head300(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsHead301(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !result.Success {
		t.Fatal("expected Success")
	}
}

func TestHTTPRedirectsHead302(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRedirectsHead307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Head307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !result.Success {
		t.Fatal("unexpected Success")
	}
}

func TestHTTPRedirectsOptions307(t *testing.T) {
	t.Skip("receive a status code of 204 which is not expected")
	client := newHTTPRedirectsClient()
	result, err := client.Options307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPatch302(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient()
	result, err := client.Patch302(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPatch307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Patch307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPost303(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Post303(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPost307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Post307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPut301(t *testing.T) {
	t.Skip("HTTP client automatically redirects, test server doesn't expect it")
	client := newHTTPRedirectsClient()
	result, err := client.Put301(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPRedirectsPut307(t *testing.T) {
	client := newHTTPRedirectsClient()
	result, err := client.Put307(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
