// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dategroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func newDateClient() DateClient {
	return NewDateClient(NewDefaultConnection(nil))
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
	helpers.DeepEqualOrFatal(t, resp.Value, &dt)
}

func TestGetMinDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetMinDate(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	helpers.DeepEqualOrFatal(t, resp.Value, &dt)
}

func TestGetNull(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	var dt *time.Time
	helpers.DeepEqualOrFatal(t, resp.Value, dt)
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
	helpers.DeepEqualOrFatal(t, resp.StatusCode, http.StatusOK)
}

func TestPutMinDate(t *testing.T) {
	client := newDateClient()
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	resp, err := client.PutMinDate(context.Background(), dt, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StatusCode, http.StatusOK)
}
