// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"reflect"
	"testing"
)

func newHTTPServerFailureClient() *HTTPServerFailureClient {
	return NewHTTPServerFailureClient(NewDefaultConnection(nil))
}

func TestHTTPServerFailureDelete505(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Delete505(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPServerFailureGet501(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Get501(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPServerFailureHead501(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Head501(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Success {
		t.Fatal("unexpected success")
	}
}

func TestHTTPServerFailurePost505(t *testing.T) {
	client := newHTTPServerFailureClient()
	result, err := client.Post505(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatalf("Expected a nil result")
	}
}
