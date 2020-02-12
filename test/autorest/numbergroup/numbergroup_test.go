// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergrouptest

import (
	"context"
	"generatortests/autorest/generated/numbergroup"
	"net/http"
	"reflect"
	"testing"
)

func getNumberClient(t *testing.T) numbergroup.NumberOperations {
	client, err := numbergroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create number client: %v", err)
	}
	return client.NumberOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestNumberGetBigDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetBigDecimal: %v", err)
	}
	val := 2.5976931e+101
	expected := &numbergroup.NumberGetBigDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigDecimalNegativeDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDecimalNegativeDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetBigDecimalNegativeDecimal: %v", err)
	}
	val := -99999999.99
	expected := &numbergroup.NumberGetBigDecimalNegativeDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigDecimalPositiveDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDecimalPositiveDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetBigDecimalPositiveDecimal: %v", err)
	}
	val := 99999999.99
	expected := &numbergroup.NumberGetBigDecimalPositiveDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigDouble(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDouble(context.Background())
	if err != nil {
		t.Fatalf("GetBigDouble: %v", err)
	}
	val := 2.5976931e+101
	expected := &numbergroup.NumberGetBigDoubleResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigDoubleNegativeDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDoubleNegativeDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetBigDoubleNegativeDecimal: %v", err)
	}
	val := -99999999.99
	expected := &numbergroup.NumberGetBigDoubleNegativeDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigDoublePositiveDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigDoublePositiveDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetBigDoublePositiveDecimal: %v", err)
	}
	val := 99999999.99
	expected := &numbergroup.NumberGetBigDoublePositiveDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetBigFloat(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetBigFloat(context.Background())
	if err != nil {
		t.Fatalf("GetBigFloat: %v", err)
	}
	val := float32(3.402823e+20)
	expected := &numbergroup.NumberGetBigFloatResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetInvalidDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetInvalidDecimal(context.Background())
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetInvalidDouble(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetInvalidDouble(context.Background())
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetInvalidFloat(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetInvalidFloat(context.Background())
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if result != nil {
		t.Fatalf("expected a nil result")
	}
}

func TestNumberGetNull(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &numbergroup.NumberGetNullResponse{
		StatusCode: http.StatusOK,
		Value:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetSmallDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetSmallDecimal(context.Background())
	if err != nil {
		t.Fatalf("GetSmallDecimal: %v", err)
	}
	val := 2.5976931e-101
	expected := &numbergroup.NumberGetSmallDecimalResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetSmallDouble(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetSmallDouble(context.Background())
	if err != nil {
		t.Fatalf("GetSmallDouble: %v", err)
	}
	val := 2.5976931e-101
	expected := &numbergroup.NumberGetSmallDoubleResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberGetSmallFloat(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.GetSmallFloat(context.Background())
	if err != nil {
		t.Fatalf("GetSmallFloat: %v", err)
	}
	val := 3.402823e-20
	expected := &numbergroup.NumberGetSmallFloatResponse{
		StatusCode: http.StatusOK,
		Value:      &val,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDecimal(context.Background(), 2.5976931e+101)
	if err != nil {
		t.Fatalf("PutBigDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutBigDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDecimalNegativeDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDecimalNegativeDecimal(context.Background())
	if err != nil {
		t.Fatalf("PutBigDecimalNegativeDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutBigDecimalNegativeDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDecimalPositiveDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDecimalPositiveDecimal(context.Background())
	if err != nil {
		t.Fatalf("PutBigDecimalPositiveDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutBigDecimalPositiveDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDouble(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDouble(context.Background(), 2.5976931e+101)
	if err != nil {
		t.Fatalf("PutBigDouble: %v", err)
	}
	expected := &numbergroup.NumberPutBigDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDoubleNegativeDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDoubleNegativeDecimal(context.Background())
	if err != nil {
		t.Fatalf("PutBigDoubleNegativeDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutBigDoubleNegativeDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigDoublePositiveDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigDoublePositiveDecimal(context.Background())
	if err != nil {
		t.Fatalf("PutBigDeoublePositiveDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutBigDoublePositiveDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutBigFloat(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutBigFloat(context.Background(), 3.402823e+20)
	if err != nil {
		t.Fatalf("PutBigFloat: %v", err)
	}
	expected := &numbergroup.NumberPutBigFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutSmallDecimal(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutSmallDecimal(context.Background(), 2.5976931e-101)
	if err != nil {
		t.Fatalf("PutSmallDecimal: %v", err)
	}
	expected := &numbergroup.NumberPutSmallDecimalResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutSmallDouble(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutSmallDouble(context.Background(), 2.5976931e-101)
	if err != nil {
		t.Fatalf("PutSmallDouble: %v", err)
	}
	expected := &numbergroup.NumberPutSmallDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestNumberPutSmallFloat(t *testing.T) {
	client := getNumberClient(t)
	result, err := client.PutSmallFloat(context.Background(), 3.402823e-20)
	if err != nil {
		t.Fatalf("PutSmallFloat: %v", err)
	}
	expected := &numbergroup.NumberPutSmallFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
