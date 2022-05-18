// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergroup

import (
	"context"
	"generatortests"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newHeaderClient() *HeaderClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewHeaderClient(pl)
}

func TestHeaderCustomRequestID(t *testing.T) {
	client := newHeaderClient()
	header := http.Header{}
	header.Set("x-ms-client-request-id", "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	result, err := client.CustomRequestID(runtime.WithHTTPHeader(context.Background(), header), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamBool(context.Background(), "true", true, nil)
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamBool(context.Background(), "false", false, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamByte(context.Background(), "valid", []byte("啊齄丂狛狜隣郎隣兀﨩"), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamDate(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02", "2010-01-01")
	require.NoError(t, err)
	result, err := client.ParamDate(context.Background(), "valid", val, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamDatetime(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02T15:04:05Z", "2010-01-01T12:34:56Z")
	require.NoError(t, err)
	result, err := client.ParamDatetime(context.Background(), "valid", val, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	require.NoError(t, err)
	result, err := client.ParamDatetimeRFC1123(context.Background(), "valid", &HeaderClientParamDatetimeRFC1123Options{Value: &val})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDouble(context.Background(), "positive", 7e120, nil)
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamDouble(context.Background(), "negative", -3.0, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDuration(context.Background(), "valid", "P123DT22H14M12.011S", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamEnum(t *testing.T) {
	client := newHeaderClient()
	val := GreyscaleColorsGREY
	result, err := client.ParamEnum(context.Background(), "valid", &HeaderClientParamEnumOptions{Value: &val})
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamEnum(context.Background(), "null", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// func TestHeaderParamExistingKey(t *testing.T) {
// 	client := newHeaderClient()
// 	result, err := client.ParamExistingKey(context.Background(), "overwrite")
// 	if err != nil {
// 		t.Fatalf("ParamExistingKey: %v", err)
// 	}
// 	if !reflect.ValueOf(result).IsZero() {
//		t.Fatal("expected zero-value result")
//	}
// }

func TestHeaderParamFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamFloat(context.Background(), "positive", 0.07, nil)
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamFloat(context.Background(), "negative", -3.0, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamInteger(context.Background(), "positive", 1, nil)
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamInteger(context.Background(), "negative", -2, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamLong(context.Background(), "positive", 105, nil)
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamLong(context.Background(), "negative", -2, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamProtectedKey(context.Background(), "text/html", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderParamString(t *testing.T) {
	client := newHeaderClient()
	val := "The quick brown fox jumps over the lazy dog"
	result, err := client.ParamString(context.Background(), "valid", &HeaderClientParamStringOptions{Value: &val})
	require.NoError(t, err)
	require.Zero(t, result)

	result, err = client.ParamString(context.Background(), "null", nil)
	require.NoError(t, err)
	require.Zero(t, result)

	val = ""
	result, err = client.ParamString(context.Background(), "empty", &HeaderClientParamStringOptions{Value: &val})
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderResponseBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseBool(context.Background(), "true", nil)
	require.NoError(t, err)
	val := true
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseBool(context.Background(), "false", nil)
	require.NoError(t, err)
	val = false
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseByte(context.Background(), "valid", nil)
	require.NoError(t, err)
	val := []byte("啊齄丂狛狜隣郎隣兀﨩")
	if r := cmp.Diff(result.Value, val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDate(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDate(context.Background(), "valid", nil)
	require.NoError(t, err)
	val, err := time.Parse("2006-01-02", "2010-01-01")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDate(context.Background(), "min", nil)
	require.NoError(t, err)
	val, err = time.Parse("2006-01-02", "0001-01-01")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDatetime(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetime(context.Background(), "valid", nil)
	require.NoError(t, err)
	val, err := time.Parse(time.RFC3339, "2010-01-01T12:34:56Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDatetime(context.Background(), "min", nil)
	require.NoError(t, err)
	val, err = time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetimeRFC1123(context.Background(), "valid", nil)
	require.NoError(t, err)
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseDatetimeRFC1123(context.Background(), "min", nil)
	require.NoError(t, err)
	val, err = time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDouble(context.Background(), "positive", nil)
	require.NoError(t, err)
	if *result.Value != 7e120 {
		t.Fatalf("unexpected value %f", *result.Value)
	}

	result, err = client.ResponseDouble(context.Background(), "negative", nil)
	require.NoError(t, err)
	if *result.Value != -3 {
		t.Fatalf("unexpected value %f", *result.Value)
	}
}

func TestHeaderResponseDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDuration(context.Background(), "valid", nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr("P123DT22H14M12.011S")); r != "" {
		t.Fatal(r)
	}
}

func TestHeaderResponseEnum(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseEnum(context.Background(), "valid", nil)
	require.NoError(t, err)
	val := GreyscaleColors("GREY")
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
	result, err = client.ResponseEnum(context.Background(), "null", nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestHeaderResponseExistingKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseExistingKey(context.Background(), nil)
	require.NoError(t, err)
	if *result.UserAgent != "overwrite" {
		t.Fatalf("unexpected value %s", *result.UserAgent)
	}
}

func TestHeaderResponseFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseFloat(context.Background(), "positive", nil)
	require.NoError(t, err)
	if *result.Value != 0.07 {
		t.Fatalf("unexpected value %f", *result.Value)
	}

	result, err = client.ResponseFloat(context.Background(), "negative", nil)
	require.NoError(t, err)
	if *result.Value != -3 {
		t.Fatalf("unexpected value %f", *result.Value)
	}
}

func TestHeaderResponseInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseInteger(context.Background(), "positive", nil)
	require.NoError(t, err)
	if *result.Value != 1 {
		t.Fatalf("unexpected value %d", *result.Value)
	}

	result, err = client.ResponseInteger(context.Background(), "negative", nil)
	require.NoError(t, err)
	if *result.Value != -2 {
		t.Fatalf("unexpected value %d", *result.Value)
	}
}

func TestHeaderResponseLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseLong(context.Background(), "positive", nil)
	require.NoError(t, err)
	if *result.Value != 105 {
		t.Fatalf("unexpected value %d", *result.Value)
	}

	result, err = client.ResponseLong(context.Background(), "negative", nil)
	require.NoError(t, err)
	if *result.Value != -2 {
		t.Fatalf("unexpected value %d", *result.Value)
	}
}

func TestHeaderResponseProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseProtectedKey(context.Background(), nil)
	require.NoError(t, err)
	if *result.ContentType != "text/html; charset=utf-8" {
		t.Fatalf("unexpected value %s", *result.ContentType)
	}
}

func TestHeaderResponseString(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseString(context.Background(), "valid", nil)
	require.NoError(t, err)
	if *result.Value != "The quick brown fox jumps over the lazy dog" {
		t.Fatalf("unexpected value %s", *result.Value)
	}

	result, err = client.ResponseString(context.Background(), "null", nil)
	require.NoError(t, err)
	if *result.Value != "null" {
		t.Fatalf("unexpected value %s", *result.Value)
	}

	result, err = client.ResponseString(context.Background(), "empty", nil)
	require.NoError(t, err)
	if result.Value != nil {
		t.Fatal("expected nil value")
	}
}
