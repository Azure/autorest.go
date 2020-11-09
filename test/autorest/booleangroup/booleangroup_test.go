// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package booleangroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func newBoolClient() BoolOperations {
	return NewBoolClient(NewDefaultConnection(nil))
}

func TestGetTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetTrue(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetTrue: %v", err)
	}
	val := true
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestGetFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetFalse(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetFalse: %v", err)
	}
	val := false
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &val)
}

func TestGetNull(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*bool)(nil))
}

func TestGetInvalid(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestPutTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutTrue(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutTrue: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPutFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutFalse(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutFalse: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
