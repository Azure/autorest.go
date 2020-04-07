// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"testing"
)

func getHTTPFailureOperations(t *testing.T) httpinfrastructuregroup.HTTPFailureOperations {
	client, err := httpinfrastructuregroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create HTTPFailure client: %v", err)
	}
	return client.HTTPFailureOperations()
}

func TestHTTPFailureGetEmptyError(t *testing.T) {
	client := getHTTPFailureOperations(t)
	result, err := client.GetEmptyError(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPFailureGetNoModelEmpty(t *testing.T) {
	client := getHTTPFailureOperations(t)
	result, err := client.GetNoModelEmpty(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPFailureGetNoModelError(t *testing.T) {
	client := getHTTPFailureOperations(t)
	result, err := client.GetNoModelError(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}
