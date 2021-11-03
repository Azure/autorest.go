// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package integergroup

import (
	"context"
	"math"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func newIntClient() *IntClient {
	return NewIntClient(nil)
}

func TestIntGetInvalid(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetInvalidUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalidUnixTime(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetNull(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, (*int32)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestIntGetNullUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNullUnixTime(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNullUnixTime: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if result.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestIntGetOverflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt32(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetOverflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt64(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetUnderflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt32(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetUnderflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt64(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestIntGetUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnixTime(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetUnixTime: %v", err)
	}
	t1 := time.Unix(1460505600, 0)
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &t1); r != "" {
		t.Fatal(r)
	}
}

func TestIntPutMax32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax32(context.Background(), math.MaxInt32, nil)
	if err != nil {
		t.Fatalf("PutMax32: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestIntPutMax64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax64(context.Background(), math.MaxInt64, nil)
	if err != nil {
		t.Fatalf("PutMax64: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestIntPutMin32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin32(context.Background(), math.MinInt32, nil)
	if err != nil {
		t.Fatalf("PutMin32: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestIntPutMin64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin64(context.Background(), math.MinInt64, nil)
	if err != nil {
		t.Fatalf("PutMin64: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestIntPutUnixTimeDate(t *testing.T) {
	client := newIntClient()
	t1 := time.Unix(1460505600, 0)
	result, err := client.PutUnixTimeDate(context.Background(), t1, nil)
	if err != nil {
		t.Fatalf("PutUnixTimeDate: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
