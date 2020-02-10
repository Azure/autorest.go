// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package booleangrouptest

import (
	"context"
	"generatortests/autorest/generated/booleangroup"
	"net/http"
	"reflect"
	"testing"
)

func getBoolClient(t *testing.T) booleangroup.BoolOperations {
	client, err := booleangroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create bool client: %v", err)
	}
	return client.BoolOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetTrue(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.GetTrue(context.Background())
	if err != nil {
		t.Fatalf("GetTrue: %v", err)
	}
	val := true
	expected := &booleangroup.BoolGetTrueResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetFalse(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.GetFalse(context.Background())
	if err != nil {
		t.Fatalf("GetFalse: %v", err)
	}
	val := false
	expected := &booleangroup.BoolGetFalseResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNull(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &booleangroup.BoolGetNullResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetInvalid(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.GetInvalid(context.Background())
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestPutTrue(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.PutTrue(context.Background())
	if err != nil {
		t.Fatalf("PutTrue: %v", err)
	}
	expected := &booleangroup.BoolPutTrueResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutFalse(t *testing.T) {
	client := getBoolClient(t)
	result, err := client.PutFalse(context.Background())
	if err != nil {
		t.Fatalf("PutFalse: %v", err)
	}
	expected := &booleangroup.BoolPutFalseResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
