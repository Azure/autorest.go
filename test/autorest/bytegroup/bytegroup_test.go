// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newByteClient() *ByteClient {
	return NewByteClient(nil)
}

func TestGetEmpty(t *testing.T) {
	client := newByteClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, []byte{}); r != "" {
		t.Fatal(r)
	}
}

func TestGetInvalid(t *testing.T) {
	client := newByteClient()
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetNonASCII(t *testing.T) {
	client := newByteClient()
	result, err := client.GetNonASCII(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNonASCII: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6}); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newByteClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, ([]byte)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestPutNonASCII(t *testing.T) {
	client := newByteClient()
	result, err := client.PutNonASCII(context.Background(), []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6}, nil)
	if err != nil {
		t.Fatalf("PutNonASCII: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
