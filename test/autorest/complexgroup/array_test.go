// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newArrayClient() *ArrayClient {
	return NewArrayClient(nil)
}

func TestArrayGetEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{
		Array: []*string{},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayGetNotProvided(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayGetValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{
		Array: []*string{
			to.StringPtr("1, 2, 3, 4"),
			to.StringPtr(""),
			nil,
			to.StringPtr("&S#$(*Y"),
			to.StringPtr("The quick brown fox jumps over the lazy dog"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayPutEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutEmpty(context.Background(), ArrayWrapper{Array: []*string{}}, nil)
	if err != nil {
		t.Fatalf("PutEmpty: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestArrayPutValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutValid(context.Background(), ArrayWrapper{Array: []*string{
		to.StringPtr("1, 2, 3, 4"),
		to.StringPtr(""),
		nil,
		to.StringPtr("&S#$(*Y"),
		to.StringPtr("The quick brown fox jumps over the lazy dog"),
	}}, nil)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
