// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package booleangroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newBoolClient() *BoolClient {
	return NewBoolClient(nil)
}

func TestGetTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetTrue(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetTrue: %v", err)
	}
	if r := cmp.Diff(result.Value, to.BoolPtr(true)); r != "" {
		t.Fatal(r)
	}
}

func TestGetFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetFalse(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetFalse: %v", err)
	}
	if r := cmp.Diff(result.Value, to.BoolPtr(false)); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if r := cmp.Diff(result.Value, (*bool)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestGetInvalid(t *testing.T) {
	client := newBoolClient()
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestPutTrue(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutTrue(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutTrue: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPutFalse(t *testing.T) {
	client := newBoolClient()
	result, err := client.PutFalse(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutFalse: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
