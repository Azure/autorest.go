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

func TestHeaderParamProtectedKey(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamProtectedKey(context.Background(), "text/html")
	if err != nil {
		t.Fatalf("ParamProtectedKey: %v", err)
	}
	expected := &headergroup.HeaderParamProtectedKeyResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderParamString(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ParamString(context.Background(), "valid", "The quick brown fox jumps over the lazy dog")
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	expected := &headergroup.HeaderParamStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	// result, err = client.ParamString(context.Background(), "null", "")
	// if err != nil {
	// 	t.Fatalf("ParamString: %v", err)
	// }
	// expected = &headergroup.HeaderParamStringResponse{
	// 	StatusCode: http.StatusOK,
	// }
	// deepEqualOrFatal(t, result, expected)

	result, err = client.ParamString(context.Background(), "empty", "")
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	expected = &headergroup.HeaderParamStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

// TODO check why we dont check for the value returned in all of the tests below this comment
func TestHeaderResponseBool(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseBool(context.Background(), "true")
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	expected := &headergroup.HeaderResponseBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseBool(context.Background(), "false")
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	expected = &headergroup.HeaderResponseBoolResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseByte(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseByte(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseByte: %v", err)
	}
	expected := &headergroup.HeaderResponseByteResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseDate(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseDate(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	expected := &headergroup.HeaderResponseDateResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseDate(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	expected = &headergroup.HeaderResponseDateResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseDatetime(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseDatetime(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	expected := &headergroup.HeaderResponseDatetimeResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseDatetime(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	expected = &headergroup.HeaderResponseDatetimeResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseDatetimeRFC1123(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseDatetimeRFC1123(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	expected := &headergroup.HeaderResponseDatetimeRFC1123Response{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseDatetimeRFC1123(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	expected = &headergroup.HeaderResponseDatetimeRFC1123Response{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseDouble(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseDouble(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	expected := &headergroup.HeaderResponseDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseDouble(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	expected = &headergroup.HeaderResponseDoubleResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseDuration(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseDuration(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDuration: %v", err)
	}
	expected := &headergroup.HeaderResponseDurationResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseEnum(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseEnum(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	expected := &headergroup.HeaderResponseEnumResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseEnum(context.Background(), "null")
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	expected = &headergroup.HeaderResponseEnumResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseExistingKey(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseExistingKey(context.Background())
	if err != nil {
		t.Fatalf("ResponseExistingKey: %v", err)
	}
	expected := &headergroup.HeaderResponseExistingKeyResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseFloat(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseFloat(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	expected := &headergroup.HeaderResponseFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseFloat(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	expected = &headergroup.HeaderResponseFloatResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseInteger(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseInteger(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	expected := &headergroup.HeaderResponseIntegerResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseInteger(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	expected = &headergroup.HeaderResponseIntegerResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseLong(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseLong(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	expected := &headergroup.HeaderResponseLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseLong(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	expected = &headergroup.HeaderResponseLongResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseProtectedKey(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseProtectedKey(context.Background())
	if err != nil {
		t.Fatalf("ResponseProtectedKey: %v", err)
	}
	expected := &headergroup.HeaderResponseProtectedKeyResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}

func TestHeaderResponseString(t *testing.T) {
	client := getHeaderClient(t)
	result, err := client.ResponseString(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	expected := &headergroup.HeaderResponseStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseString(context.Background(), "null")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	expected = &headergroup.HeaderResponseStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)

	result, err = client.ResponseString(context.Background(), "empty")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	expected = &headergroup.HeaderResponseStringResponse{
		StatusCode: http.StatusOK,
	}
	deepEqualOrFatal(t, result, expected)
}
