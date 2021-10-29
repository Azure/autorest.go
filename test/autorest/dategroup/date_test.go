// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dategroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func newDateClient() *DateClient {
	return NewDateClient(nil)
}

func TestGetInvalidDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetInvalidDate(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetMaxDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetMaxDate(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, &dt); r != "" {
		t.Fatal(r)
	}
}

func TestGetMinDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetMinDate(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, &dt); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestGetOverflowDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetOverflowDate(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestGetUnderflowDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetUnderflowDate(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestPutMaxDate(t *testing.T) {
	client := newDateClient()
	dt := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	resp, err := client.PutMaxDate(context.Background(), dt, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.RawResponse.StatusCode, http.StatusOK); r != "" {
		t.Fatal(r)
	}
}

func TestPutMinDate(t *testing.T) {
	client := newDateClient()
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	resp, err := client.PutMinDate(context.Background(), dt, nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.RawResponse.StatusCode, http.StatusOK); r != "" {
		t.Fatal(r)
	}
}
