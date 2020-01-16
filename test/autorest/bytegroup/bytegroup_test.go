// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegrouptest

import (
	"context"
	"generatortests/autorest/generated/bytegroup"
	"net/http"
	"reflect"
	"testing"
)

func getByteClient(t *testing.T) *bytegroup.ByteClient {
	client, err := bytegroup.NewByteClient(nil)
	if err != nil {
		t.Fatalf("failed to create byte client: %v", err)
	}
	return client
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetEmpty(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &bytegroup.GetEmptyResponse{
		StatusCode: http.StatusOK,
		Value:      []byte{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetInvalid(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetInvalid(context.Background())
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetNonASCII(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetNonASCII(context.Background())
	if err != nil {
		t.Fatalf("GetNonASCII: %v", err)
	}
	expected := &bytegroup.GetNonASCIIResponse{
		StatusCode: http.StatusOK,
		Value:      []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNull(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &bytegroup.GetNullResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutNonASCII(t *testing.T) {
	client := getByteClient(t)
	result, err := client.PutNonASCII(context.Background(), []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6})
	if err != nil {
		t.Fatalf("PutNonASCII: %v", err)
	}
	expected := &bytegroup.PutNonASCIIResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
