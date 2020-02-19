// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergrouptest

import (
	"context"
	"generatortests/autorest/generated/headergroup"
	"net/http"
	"reflect"
	"testing"
)

func getHeaderClient(t *testing.T) headergroup.HeaderOperations {
	client, err := headergroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create header client: %v", err)
	}
	return client.HeaderOperations()
}

func deepEqualOrFatal(t *testing.T, result interface{}, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("got %+v, want %+v", result, expected)
	}
}

// func TestHeaderCustomRequestID(t *testing.T) {
// 	client := getHeaderClient(t)
// 	result, err := client.CustomRequestID(context.Background())
// 	if err != nil {
// 		t.Fatalf("CustomRequestID: %v", err)
// 	}
// 	expected := &headergroup.HeaderCustomRequestIDResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamBool(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamBool(context.Background(), "true", true)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	expected := &headergroup.HeaderParamBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ParamBool(context.Background(), "false", false)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	expected = &headergroup.HeaderParamBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO this required a change in the generated code so that it outputs base64.STDEncoding.EncodeToString()
func TestHeaderParamByte(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamByte(context.Background(), "valid", []byte("啊齄丂狛狜隣郎隣兀﨩"))
	if err != nil {
		t.Fatalf("ParamByte: %v", err)
	}
	expected := &headergroup.HeaderParamByteResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestHeaderParamDate(t *testing.T) {
// 	client := getHeaderClient(t)
// 	val, err := time.Parse("2006-01-02", "2010-01-01")
// 	if err != nil {
// 		t.Fatalf("Unable to parse time: %v", err)
// 	}
// 	result, err := client.ParamDate(context.Background(), "valid", date.Date{time.Date(2010, time.January, 01, 0, 0, 0, 0, time.UTC)})
// 	if err != nil {
// 		t.Fatalf("ParamDate: %v", err)
// 	}
// 	expected := &headergroup.HeaderParamDateResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

// func TestHeaderParamDatetime(t *testing.T) {
// 	client := getHeaderClient(t)
// 	val, err := time.Parse("2006-01-02T15:04:05Z", "2010-01-01T12:34:56Z")
// 	if err != nil {
// 		t.Fatalf("Unable to parse time: %v", err)
// 	}
// 	result, err := client.ParamDatetime(context.Background(), "valid", val)
// 	if err != nil {
// 		t.Fatalf("ParamDatetime: %v", err)
// 	}
// 	expected := &headergroup.HeaderParamDatetimeResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

// func TestHeaderParamDatetimeRFC1123(t *testing.T) {
// 	client := getHeaderClient(t)
// 	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
// 	if err != nil {
// 		t.Fatalf("Unable to parse time: %v", err)
// 	}
// 	result, err := client.ParamDatetimeRFC1123(context.Background(), "valid", val)
// 	if err != nil {
// 		t.Fatalf("ParamDatetimeRFC1123: %v", err)
// 	}
// 	expected := &headergroup.HeaderParamDatetimeRFC1123Response{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamDouble(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamDouble(context.Background(), "positive", 7e120)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	expected := &headergroup.HeaderParamDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ParamDouble(context.Background(), "negative", -3.0)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	expected = &headergroup.HeaderParamDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// func TestHeaderParamDuration(t *testing.T) {
// 	client := getHeaderClient(t)
// 	val, err := time.ParseDuration("P123DT22H14M12.011S")
// 	if err != nil {
// 		t.Fatalf("Unable to parse duration: %v", err)
// 	}
// 	result, err := client.ParamDuration(context.Background(), "valid", val)
// 	if err != nil {
// 		t.Fatalf("ParamDuration: %v", err)
// 	}
// 	expected := &headergroup.HeaderParamDurationResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamEnum(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamEnum(context.Background(), "valid", headergroup.GreyscaleColorsGrey)
	if err != nil {
		t.Fatalf("ParamEnum: %v", err)
	}
	expected := &headergroup.HeaderParamEnumResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	// var color headergroup.GreyscaleColors
	// result, err = client.ParamEnum(context.Background(), "null", color)
	// if err != nil {
	// 	t.Fatalf("ParamEnum: %v", err)
	// }
	// expected = &headergroup.HeaderParamEnumResponse{
	// 	StatusCode: http.StatusOK,
	// }
	// deepEqualOrFatal(t, result, expected)
}

// func TestHeaderParamExistingKey(t *testing.T) {
// 	client := getHeaderClient(t)
// 	result, err := client.ParamExistingKey(context.Background(), "overwrite")
// 	if err != nil {
// 		t.Fatalf("ParamExistingKey: %v", err)
// 	}
// 	expected := &headergroup.HeaderParamExistingKeyResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	deepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamFloat(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamFloat(context.Background(), "positive", 0.07)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	expected := &headergroup.HeaderParamFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ParamFloat(context.Background(), "negative", -3.0)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	expected = &headergroup.HeaderParamFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderParamInteger(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamInteger(context.Background(), "positive", 1)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	expected := &headergroup.HeaderParamIntegerResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ParamInteger(context.Background(), "negative", -2)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	expected = &headergroup.HeaderParamIntegerResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderParamLong(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamLong(context.Background(), "positive", 105)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	expected := &headergroup.HeaderParamLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ParamLong(context.Background(), "negative", -2)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	expected = &headergroup.HeaderParamLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
