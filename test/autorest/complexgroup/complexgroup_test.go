// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"net/http"
	"reflect"
	"testing"
)

func getBasicOperations(t *testing.T) complexgroup.BasicOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.BasicOperations()
}

func getPrimitiveOperations(t *testing.T) complexgroup.PrimitiveOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.PrimitiveOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetValid(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	var v complexgroup.CMYKColors
	colors := complexgroup.PossibleCMYKColorsValues()
	for _, c := range colors {
		if string(c) == "YELLOW" {
			v = c
			break
		}
	}
	i, s := int32(2), "abc"
	expected := &complexgroup.BasicGetValidResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{ID: &i, Name: &s, Color: &v},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutValid(t *testing.T) {
	client := getBasicOperations(t)
	var v complexgroup.CMYKColors
	colors := complexgroup.PossibleCMYKColorsValues()
	for _, c := range colors {
		if string(c) == "Magenta" {
			v = c
			break
		}
	}
	i, s := int32(2), "abc"
	result, err := client.PutValid(context.Background(), complexgroup.Basic{ID: &i, Name: &s, Color: &v})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.BasicPutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO check this
func TestGetInvalid(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatalf("GetInvalid expected an error")
	}
	var expected *complexgroup.BasicGetInvalidResponse
	expected = nil
	deepEqualOrFatal(t, result, expected)
}

func TestGetEmpty(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &complexgroup.BasicGetEmptyResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNull(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &complexgroup.BasicGetNullResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNotProvided(t *testing.T) {
	client := getBasicOperations(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	expected := &complexgroup.BasicGetNotProvidedResponse{
		StatusCode: http.StatusOK,
		Basic:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetInt(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetInt(context.Background())
	if err != nil {
		t.Fatalf("GetInt: %v", err)
	}
	a, b := int32(-1), int32(2)
	expected := &complexgroup.PrimitiveGetIntResponse{
		StatusCode: http.StatusOK,
		IntWrapper: &complexgroup.IntWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutInt(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), complexgroup.IntWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutInt: %v", err)
	}
	expected := &complexgroup.PrimitivePutIntResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetLong(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetLong(context.Background())
	if err != nil {
		t.Fatalf("GetLong: %v", err)
	}
	a, b := int64(1099511627775), int64(-999511627788)
	expected := &complexgroup.PrimitiveGetLongResponse{
		StatusCode:  http.StatusOK,
		LongWrapper: &complexgroup.LongWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutLong(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := int64(1099511627775), int64(-999511627788)
	result, err := client.PutLong(context.Background(), complexgroup.LongWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutLong: %v", err)
	}
	expected := &complexgroup.PrimitivePutLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetFloat(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetFloat(context.Background())
	if err != nil {
		t.Fatalf("GetFloat: %v", err)
	}
	a, b := float32(1.05), float32(-0.003)
	expected := &complexgroup.PrimitiveGetFloatResponse{
		StatusCode:   http.StatusOK,
		FloatWrapper: &complexgroup.FloatWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutFloat(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := float32(1.05), float32(-0.003)
	result, err := client.PutFloat(context.Background(), complexgroup.FloatWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutFloat: %v", err)
	}
	expected := &complexgroup.PrimitivePutFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetDouble(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetDouble(context.Background())
	if err != nil {
		t.Fatalf("GetDouble: %v", err)
	}
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	expected := &complexgroup.PrimitiveGetDoubleResponse{
		StatusCode:    http.StatusOK,
		DoubleWrapper: &complexgroup.DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutDouble(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	result, err := client.PutDouble(context.Background(), complexgroup.DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b})
	if err != nil {
		t.Fatalf("PutDouble: %v", err)
	}
	expected := &complexgroup.PrimitivePutDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetBool(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetBool(context.Background())
	if err != nil {
		t.Fatalf("GetBool: %v", err)
	}
	a, b := true, false
	expected := &complexgroup.PrimitiveGetBoolResponse{
		StatusCode:     http.StatusOK,
		BooleanWrapper: &complexgroup.BooleanWrapper{FieldTrue: &a, FieldFalse: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutBool(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := true, false
	result, err := client.PutBool(context.Background(), complexgroup.BooleanWrapper{FieldTrue: &a, FieldFalse: &b})
	if err != nil {
		t.Fatalf("PutBool: %v", err)
	}
	expected := &complexgroup.PrimitivePutBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetString(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetString(context.Background())
	if err != nil {
		t.Fatalf("GetString: %v", err)
	}
	var c *string
	a, b, c := "goodrequest", "", nil
	expected := &complexgroup.PrimitiveGetStringResponse{
		StatusCode:    http.StatusOK,
		StringWrapper: &complexgroup.StringWrapper{Field: &a, Empty: &b, Null: c},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutString(t *testing.T) {
	client := getPrimitiveOperations(t)
	var c *string
	a, b, c := "goodrequest", "", nil
	result, err := client.PutString(context.Background(), complexgroup.StringWrapper{Field: &a, Empty: &b, Null: c})
	if err != nil {
		t.Fatalf("PutString: %v", err)
	}
	expected := &complexgroup.PrimitivePutStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestGetDate(t *testing.T) {
// 	client := getPrimitiveOperations(t)
// 	result, err := client.GetDate(context.Background())
// 	if err != nil {
// 		t.Fatalf("GetDate: %v", err)
// 	}
// 	a, err := time.Parse("2006-01-02", "0001-01-01")
// 	if err != nil {
// 		t.Fatalf("Unable to parse date string: %v", err)
// 	}
// 	b, err := time.Parse("2006-01-02", "2016-02-29")
// 	if err != nil {
// 		t.Fatalf("Unable to parse leap year date string: %v", err)
// 	}
// 	expected := &complexgroup.PrimitiveGetDateResponse{
// 		StatusCode:  http.StatusOK,
// 		DateWrapper: &complexgroup.DateWrapper{Field: &a, Leap: &b},
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

// func TestPutDate(t *testing.T) {
// 	client := getPrimitiveOperations(t)
// 	a, err := time.Parse("2006-01-02", "0001-01-01")
// 	if err != nil {
// 		t.Fatalf("Unable to parse date string: %v", err)
// 	}
// 	b, err := time.Parse("2006-01-02", "2016-02-29")
// 	if err != nil {
// 		t.Fatalf("Unable to parse leap year date string: %v", err)
// 	}
// 	result, err := client.PutDate(context.Background(), complexgroup.DateWrapper{Field: &a, Leap: &b})
// 	if err != nil {
// 		t.Fatalf("PutDate: %v", err)
// 	}
// 	expected := &complexgroup.PrimitivePutDateResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }
