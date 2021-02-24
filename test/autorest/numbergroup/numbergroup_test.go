// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newNumberClient() *NumberClient {
	return NewNumberClient(NewDefaultConnection(nil))
}

func TestNumberGetBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimal: %v", err)
	}
	val := 2.5976931e+101
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimalNegativeDecimal: %v", err)
	}
	val := -99999999.99
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalPositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDecimalPositiveDecimal: %v", err)
	}
	val := 99999999.99
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDouble(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDouble: %v", err)
	}
	val := 2.5976931e+101
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoubleNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDoubleNegativeDecimal: %v", err)
	}
	val := -99999999.99
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoublePositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigDoublePositiveDecimal: %v", err)
	}
	val := 99999999.99
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigFloat(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBigFloat: %v", err)
	}
	val := float32(3.402823e+20)
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetInvalidDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDecimal(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestNumberGetInvalidDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDouble(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestNumberGetInvalidFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidFloat(context.Background(), nil)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

func TestNumberGetNull(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, (*float32)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallDecimal: %v", err)
	}
	val := 2.5976931e-101
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDouble(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallDouble: %v", err)
	}
	val := 2.5976931e-101
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallFloat(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetSmallFloat: %v", err)
	}
	val := 3.402823e-20
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberPutBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimal(context.Background(), 2.5976931e+101, nil)
	if err != nil {
		t.Fatalf("PutBigDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDecimalNegativeDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalPositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDecimalPositiveDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDouble(context.Background(), 2.5976931e+101, nil)
	if err != nil {
		t.Fatalf("PutBigDouble: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoubleNegativeDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDoubleNegativeDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoublePositiveDecimal(context.Background(), nil)
	if err != nil {
		t.Fatalf("PutBigDeoublePositiveDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigFloat(context.Background(), 3.402823e+20, nil)
	if err != nil {
		t.Fatalf("PutBigFloat: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDecimal(context.Background(), 2.5976931e-101, nil)
	if err != nil {
		t.Fatalf("PutSmallDecimal: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDouble(context.Background(), 2.5976931e-101, nil)
	if err != nil {
		t.Fatalf("PutSmallDouble: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestNumberPutSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallFloat(context.Background(), 3.402823e-20, nil)
	if err != nil {
		t.Fatalf("PutSmallFloat: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
