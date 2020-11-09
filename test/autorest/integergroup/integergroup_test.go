// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package integergroup

import (
	"context"
	"generatortests/helpers"
	"math"
	"net/http"
	"testing"
	"time"
)

func newIntClient() IntOperations {
	return NewIntClient(NewDefaultConnection(nil))
}

func TestIntGetInvalid(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestIntGetInvalidUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalidUnixTime(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestIntGetNull(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*int32)(nil))
}

func TestIntGetNullUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNullUnixTime(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNullUnixTime: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*time.Time)(nil))
}

func TestIntGetOverflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt32(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil response but received one")
	}
}

func TestIntGetOverflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt64(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil response but received one")
	}
}

func TestIntGetUnderflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt32(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil response but received one")
	}
}

func TestIntGetUnderflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt64(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil response but received one")
	}
}

func TestIntGetUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnixTime(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetUnixTime: %v", err)
	}
	t1 := time.Unix(1460505600, 0)
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &t1)
}

func TestIntPutMax32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax32(context.Background(), math.MaxInt32, nil)
	if err != nil {
		t.Fatalf("PutMax32: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestIntPutMax64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax64(context.Background(), math.MaxInt64, nil)
	if err != nil {
		t.Fatalf("PutMax64: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestIntPutMin32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin32(context.Background(), math.MinInt32, nil)
	if err != nil {
		t.Fatalf("PutMin32: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestIntPutMin64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin64(context.Background(), math.MinInt64, nil)
	if err != nil {
		t.Fatalf("PutMin64: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestIntPutUnixTimeDate(t *testing.T) {
	client := newIntClient()
	t1 := time.Unix(1460505600, 0)
	result, err := client.PutUnixTimeDate(context.Background(), t1, nil)
	if err != nil {
		t.Fatalf("PutUnixTimeDate: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
