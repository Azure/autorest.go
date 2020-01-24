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

func getBasicClient(t *testing.T) *complexgroup.BasicClient {
	client, err := complexgroup.NewBasicClient(complexgroup.DefaultEndpoint, nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client
}

func getPrimitiveClient(t *testing.T) *complexgroup.PrimitiveClient {
	client, err := complexgroup.NewPrimitiveClient(complexgroup.DefaultEndpoint, nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

func TestGetValid(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	var v complexgroup.ColorType
	colors := complexgroup.PossibleColorValues()
	for _, c := range colors {
		if string(c) == "YELLOW" {
			v = c
			break
		}
	}
	i, s := int(2), "abc"
	expected := &complexgroup.GetValidResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{ID: &i, Name: &s, Color: &v},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutValid(t *testing.T) {
	client := getBasicClient(t)
	var v complexgroup.ColorType
	colors := complexgroup.PossibleColorValues()
	for _, c := range colors {
		if string(c) == "Magenta" {
			v = c
			break
		}
	}
	i, s := int(2), "abc"
	result, err := client.PutValid(context.Background(), complexgroup.Basic{ID: &i, Name: &s, Color: &v})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	expected := &complexgroup.PutValidResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO check this
func TestGetInvalid(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatalf("GetInvalid expected an error")
	}
	var expected *complexgroup.GetInvalidResponse
	expected = nil
	deepEqualOrFatal(t, result, expected)
}

func TestGetEmpty(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	expected := &complexgroup.GetEmptyResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNull(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	expected := &complexgroup.GetNullResponse{
		StatusCode: http.StatusOK,
		Basic:      &complexgroup.Basic{},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetNotProvided(t *testing.T) {
	client := getBasicClient(t)
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	expected := &complexgroup.GetNotProvidedResponse{
		StatusCode: http.StatusOK,
		Basic:      nil,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetInt(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetInt(context.Background())
	if err != nil {
		t.Fatalf("GetInt: %v", err)
	}
	a, b := int32(-1), int32(2)
	expected := &complexgroup.GetIntResponse{
		StatusCode: http.StatusOK,
		IntWrapper: &complexgroup.IntWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutInt(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), complexgroup.IntWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutInt: %v", err)
	}
	expected := &complexgroup.PutIntResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetLong(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetLong(context.Background())
	if err != nil {
		t.Fatalf("GetLong: %v", err)
	}
	a, b := int64(1099511627775), int64(-999511627788)
	expected := &complexgroup.GetLongResponse{
		StatusCode:  http.StatusOK,
		LongWrapper: &complexgroup.LongWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutLong(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := int64(1099511627775), int64(-999511627788)
	result, err := client.PutLong(context.Background(), complexgroup.LongWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutLong: %v", err)
	}
	expected := &complexgroup.PutLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetFloat(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetFloat(context.Background())
	if err != nil {
		t.Fatalf("GetFloat: %v", err)
	}
	a, b := float32(1.05), float32(-0.003)
	expected := &complexgroup.GetFloatResponse{
		StatusCode:   http.StatusOK,
		FloatWrapper: &complexgroup.FloatWrapper{Field1: &a, Field2: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutFloat(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := float32(1.05), float32(-0.003)
	result, err := client.PutFloat(context.Background(), complexgroup.FloatWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutFloat: %v", err)
	}
	expected := &complexgroup.PutFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetDouble(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetDouble(context.Background())
	if err != nil {
		t.Fatalf("GetDouble: %v", err)
	}
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	expected := &complexgroup.GetDoubleResponse{
		StatusCode:    http.StatusOK,
		DoubleWrapper: &complexgroup.DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutDouble(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	result, err := client.PutDouble(context.Background(), complexgroup.DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b})
	if err != nil {
		t.Fatalf("PutDouble: %v", err)
	}
	expected := &complexgroup.PutDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetBool(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetBool(context.Background())
	if err != nil {
		t.Fatalf("GetBool: %v", err)
	}
	a, b := true, false
	expected := &complexgroup.GetBoolResponse{
		StatusCode:     http.StatusOK,
		BooleanWrapper: &complexgroup.BooleanWrapper{FieldTrue: &a, FieldFalse: &b},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutBool(t *testing.T) {
	client := getPrimitiveClient(t)
	a, b := true, false
	result, err := client.PutBool(context.Background(), complexgroup.BooleanWrapper{FieldTrue: &a, FieldFalse: &b})
	if err != nil {
		t.Fatalf("PutBool: %v", err)
	}
	expected := &complexgroup.PutBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestGetString(t *testing.T) {
	client := getPrimitiveClient(t)
	result, err := client.GetString(context.Background())
	if err != nil {
		t.Fatalf("GetString: %v", err)
	}
	var c *string
	a, b, c := "goodrequest", "", nil
	expected := &complexgroup.GetStringResponse{
		StatusCode:    http.StatusOK,
		StringWrapper: &complexgroup.StringWrapper{Field: &a, Empty: &b, Null: c},
	}
	deepEqualOrFatal(t, result, expected)
}

func TestPutString(t *testing.T) {
	client := getPrimitiveClient(t)
	var c *string
	a, b, c := "goodrequest", "", nil
	result, err := client.PutString(context.Background(), complexgroup.StringWrapper{Field: &a, Empty: &b, Null: c})
	if err != nil {
		t.Fatalf("PutString: %v", err)
	}
	expected := &complexgroup.PutStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
