// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"testing"
)

func TestHTTPServerFailureDelete505(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPServerFailureOperations()
	result, err := client.Delete505(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPServerFailureGet501(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPServerFailureOperations()
	result, err := client.Get501(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPServerFailureHead501(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPServerFailureOperations()
	result, err := client.Head501(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPServerFailurePost505(t *testing.T) {
	client := httpinfrastructuregroup.NewDefaultClient(nil).HTTPServerFailureOperations()
	result, err := client.Post505(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}
