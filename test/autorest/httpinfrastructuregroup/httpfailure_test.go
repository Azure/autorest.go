// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"reflect"
	"testing"
)

func newHTTPFailureClient() HTTPFailureClient {
	return NewHTTPFailureClient(NewDefaultConnection(nil))
}

func TestHTTPFailureGetEmptyError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetEmptyError(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestHTTPFailureGetNoModelEmpty(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelEmpty(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestHTTPFailureGetNoModelError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelError(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}
