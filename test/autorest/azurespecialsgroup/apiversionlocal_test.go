// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgroup

import (
	"context"
	"net/http"
	"testing"
)

func newAPIVersionLocalClient() *APIVersionLocalClient {
	return NewAPIVersionLocalClient(NewDefaultConnection(nil))
}

// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
func TestGetMethodLocalNull(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetMethodLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetMethodLocalValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetPathLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetPathLocalValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetSwaggerLocalValid(t *testing.T) {
	client := newAPIVersionLocalClient()
	result, err := client.GetSwaggerLocalValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
