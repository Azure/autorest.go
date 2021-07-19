// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergroup

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newHeaderClient() *HeaderClient {
	return NewHeaderClient(NewDefaultConnection(nil))
}

func TestHeaderCustomRequestID(t *testing.T) {
	client := newHeaderClient()
	header := http.Header{}
	header.Set("x-ms-client-request-id", "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	result, err := client.CustomRequestID(azcore.WithHTTPHeader(context.Background(), header), nil)
	if err != nil {
		t.Fatalf("CustomRequestID: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamBool(context.Background(), "true", true, nil)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamBool(context.Background(), "false", false, nil)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamByte(context.Background(), "valid", []byte("啊齄丂狛狜隣郎隣兀﨩"), nil)
	if err != nil {
		t.Fatalf("ParamByte: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamDate(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02", "2010-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDate(context.Background(), "valid", val, nil)
	if err != nil {
		t.Fatalf("ParamDate: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamDatetime(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02T15:04:05Z", "2010-01-01T12:34:56Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDatetime(context.Background(), "valid", val, nil)
	if err != nil {
		t.Fatalf("ParamDatetime: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDatetimeRFC1123(context.Background(), "valid", &HeaderParamDatetimeRFC1123Options{Value: &val})
	if err != nil {
		t.Fatalf("ParamDatetimeRFC1123: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDouble(context.Background(), "positive", 7e120, nil)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamDouble(context.Background(), "negative", -3.0, nil)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDuration(context.Background(), "valid", "P123DT22H14M12.011S", nil)
	if err != nil {
		t.Fatalf("ParamDuration: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamEnum(t *testing.T) {
	client := newHeaderClient()
	val := GreyscaleColorsGREY
	result, err := client.ParamEnum(context.Background(), "valid", &HeaderParamEnumOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamEnum: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamEnum(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ParamEnum: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// func TestHeaderParamExistingKey(t *testing.T) {
// 	client := newHeaderClient()
// 	result, err := client.ParamExistingKey(context.Background(), "overwrite")
// 	if err != nil {
// 		t.Fatalf("ParamExistingKey: %v", err)
// 	}
// 	if s := result.RawResponse.StatusCode; s != http.StatusOK {
//		t.Fatalf("unexpected status code %d", s)
//	}
// }

func TestHeaderParamFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamFloat(context.Background(), "positive", 0.07, nil)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamFloat(context.Background(), "negative", -3.0, nil)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamInteger(context.Background(), "positive", 1, nil)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamInteger(context.Background(), "negative", -2, nil)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamLong(context.Background(), "positive", 105, nil)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamLong(context.Background(), "negative", -2, nil)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamProtectedKey(context.Background(), "text/html", nil)
	if err != nil {
		t.Fatalf("ParamProtectedKey: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderParamString(t *testing.T) {
	client := newHeaderClient()
	val := "The quick brown fox jumps over the lazy dog"
	result, err := client.ParamString(context.Background(), "valid", &HeaderParamStringOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ParamString(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	val = ""
	result, err = client.ParamString(context.Background(), "empty", &HeaderParamStringOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseBool(context.Background(), "true", nil)
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val := true
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseBool(context.Background(), "false", nil)
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val = false
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseByte(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseByte: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val := []byte("啊齄丂狛狜隣郎隣兀﨩")
	if r := cmp.Diff(result.Value, val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDate(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDate(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err := time.Parse("2006-01-02", "2010-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDate(context.Background(), "min", nil)
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err = time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDatetime(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetime(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err := time.Parse(time.RFC3339, "2010-01-01T12:34:56Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDatetime(context.Background(), "min", nil)
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err = time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetimeRFC1123(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDatetimeRFC1123(context.Background(), "min", nil)
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val, err = time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDouble(context.Background(), "positive", nil)
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseDouble(context.Background(), "negative", nil)
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDuration(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseDuration: %v", err)
	}
	if r := cmp.Diff(result.Value, to.StringPtr("P123DT22H14M12.011S")); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseEnum(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseEnum(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	val := GreyscaleColors("GREY")
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseEnum(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseExistingKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseExistingKey(context.Background(), nil)
	if err != nil {
		t.Fatalf("ResponseExistingKey: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseFloat(context.Background(), "positive", nil)
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseFloat(context.Background(), "negative", nil)
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseInteger(context.Background(), "positive", nil)
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseInteger(context.Background(), "negative", nil)
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseLong(context.Background(), "positive", nil)
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseLong(context.Background(), "negative", nil)
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseProtectedKey(context.Background(), nil)
	if err != nil {
		t.Fatalf("ResponseProtectedKey: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestHeaderResponseString(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseString(context.Background(), "valid", nil)
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseString(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}

	result, err = client.ResponseString(context.Background(), "empty", nil)
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
