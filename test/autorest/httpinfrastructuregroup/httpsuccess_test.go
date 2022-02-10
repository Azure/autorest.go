// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"reflect"
	"testing"
)

func newHTTPSuccessClient() *HTTPSuccessClient {
	return NewHTTPSuccessClient(nil)
}

func TestHTTPSuccessDelete200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessDelete202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessDelete204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Delete204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessGet200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Get200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !*result.Value {
		t.Fatal("expected Success")
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPatch200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPatch202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPatch204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Patch204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPost200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPost201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post201(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPost202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPost204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Post204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPut200(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put200(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPut201(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put201(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPut202(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put202(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestHTTPSuccessPut204(t *testing.T) {
	client := newHTTPSuccessClient()
	result, err := client.Put204(context.Background(), nil)
	if err != nil {
		t.Fatalf("Did not expect an error, but received: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
