// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"testing"
)

func newHTTPFailureClient() HTTPFailureOperations {
	return NewHTTPFailureClient(NewDefaultClient(nil))
}

func TestHTTPFailureGetEmptyError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetEmptyError(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPFailureGetNoModelEmpty(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelEmpty(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPFailureGetNoModelError(t *testing.T) {
	client := newHTTPFailureClient()
	result, err := client.GetNoModelError(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}
