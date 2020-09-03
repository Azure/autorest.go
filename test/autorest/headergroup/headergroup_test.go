// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergrouptest

import (
	"context"
	"generatortests/autorest/generated/headergroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newHeaderClient() headergroup.HeaderOperations {
	return headergroup.NewHeaderClient(headergroup.NewDefaultClient(nil))
}

// func TestHeaderCustomRequestID(t *testing.T) {
// 	client := newHeaderClient()
// 	result, err := client.CustomRequestID(context.Background())
// 	if err != nil {
// 		t.Fatalf("CustomRequestID: %v", err)
// 	}
// 	expected := &headergroup.HeaderCustomRequestIDResponse{
// 		StatusCode: http.StatusOK,
// 	}
// 	helpers.DeepEqualOrFatal(t, result, expected)
// }

func TestHeaderParamBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamBool(context.Background(), "true", true)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamBool(context.Background(), "false", false)
	if err != nil {
		t.Fatalf("ParamBool: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// TODO this required a change in the generated code so that it outputs base64.STDEncoding.EncodeToString()
func TestHeaderParamByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamByte(context.Background(), "valid", []byte("啊齄丂狛狜隣郎隣兀﨩"))
	if err != nil {
		t.Fatalf("ParamByte: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamDate(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02", "2010-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDate(context.Background(), "valid", val)
	if err != nil {
		t.Fatalf("ParamDate: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamDatetime(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse("2006-01-02T15:04:05Z", "2010-01-01T12:34:56Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDatetime(context.Background(), "valid", val)
	if err != nil {
		t.Fatalf("ParamDatetime: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	result, err := client.ParamDatetimeRFC1123(context.Background(), "valid", &headergroup.HeaderParamDatetimeRFC1123Options{Value: &val})
	if err != nil {
		t.Fatalf("ParamDatetimeRFC1123: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDouble(context.Background(), "positive", 7e120)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamDouble(context.Background(), "negative", -3.0)
	if err != nil {
		t.Fatalf("ParamDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamDuration(context.Background(), "valid", "P123DT22H14M12.011S")
	if err != nil {
		t.Fatalf("ParamDuration: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamEnum(t *testing.T) {
	client := newHeaderClient()
	val := headergroup.GreyscaleColorsGrey
	result, err := client.ParamEnum(context.Background(), "valid", &headergroup.HeaderParamEnumOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamEnum: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamEnum(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ParamEnum: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// func TestHeaderParamExistingKey(t *testing.T) {
// 	client := newHeaderClient()
// 	result, err := client.ParamExistingKey(context.Background(), "overwrite")
// 	if err != nil {
// 		t.Fatalf("ParamExistingKey: %v", err)
// 	}
// 	helpers.VerifyStatusCode(t, result, http.StatusOK)
// }

func TestHeaderParamFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamFloat(context.Background(), "positive", 0.07)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamFloat(context.Background(), "negative", -3.0)
	if err != nil {
		t.Fatalf("ParamFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamInteger(context.Background(), "positive", 1)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamInteger(context.Background(), "negative", -2)
	if err != nil {
		t.Fatalf("ParamInteger: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamLong(context.Background(), "positive", 105)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamLong(context.Background(), "negative", -2)
	if err != nil {
		t.Fatalf("ParamLong: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ParamProtectedKey(context.Background(), "text/html")
	if err != nil {
		t.Fatalf("ParamProtectedKey: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestHeaderParamString(t *testing.T) {
	client := newHeaderClient()
	val := "The quick brown fox jumps over the lazy dog"
	result, err := client.ParamString(context.Background(), "valid", &headergroup.HeaderParamStringOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	result, err = client.ParamString(context.Background(), "null", nil)
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)

	val = ""
	result, err = client.ParamString(context.Background(), "empty", &headergroup.HeaderParamStringOptions{Value: &val})
	if err != nil {
		t.Fatalf("ParamString: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// TODO check why we dont check for the value returned in all of the tests below this comment
func TestHeaderResponseBool(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseBool(context.Background(), "true")
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := true
	expected := headergroup.HeaderResponseBoolResponse{RawResponse: result.RawResponse, Value: &val}
	helpers.DeepEqualOrFatal(t, result, &expected)
	result, err = client.ResponseBool(context.Background(), "false")
	if err != nil {
		t.Fatalf("ResponseBool: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val = false
	expected = headergroup.HeaderResponseBoolResponse{RawResponse: result.RawResponse, Value: &val}
	helpers.DeepEqualOrFatal(t, result, &expected)

}

func TestHeaderResponseByte(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseByte(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseByte: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := []byte("啊齄丂狛狜隣郎隣兀﨩")
	helpers.DeepEqualOrFatal(t, result, &headergroup.HeaderResponseByteResponse{RawResponse: result.RawResponse, Value: &val})
}

func TestHeaderResponseDate(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDate(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err := time.Parse("2006-01-02", "2010-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDateResponse{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
	result, err = client.ResponseDate(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDate: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err = time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDateResponse{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
}

func TestHeaderResponseDatetime(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetime(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err := time.Parse(time.RFC3339, "2010-01-01T12:34:56Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDatetimeResponse{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
	result, err = client.ResponseDatetime(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDatetime: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err = time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDatetimeResponse{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
}

func TestHeaderResponseDatetimeRFC1123(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDatetimeRFC1123(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err := time.Parse(time.RFC1123, "Wed, 01 Jan 2010 12:34:56 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDatetimeRFC1123Response{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
	result, err = client.ResponseDatetimeRFC1123(context.Background(), "min")
	if err != nil {
		t.Fatalf("ResponseDatetimeRFC1123: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val, err = time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatalf("Unable to parse time: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result,
		&headergroup.HeaderResponseDatetimeRFC1123Response{
			RawResponse: result.RawResponse,
			Value:       &val,
		})
}

func TestHeaderResponseDouble(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDouble(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseDouble(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseDuration(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseDuration(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseDuration: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, to.StringPtr("P123DT22H14M12.011S"))
}

func TestHeaderResponseEnum(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseEnum(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	val := headergroup.GreyscaleColors("GREY")
	helpers.DeepEqualOrFatal(t, result, &headergroup.HeaderResponseEnumResponse{RawResponse: result.RawResponse, Value: &val})
	result, err = client.ResponseEnum(context.Background(), "null")
	if err != nil {
		t.Fatalf("ResponseEnum: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseExistingKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseExistingKey(context.Background())
	if err != nil {
		t.Fatalf("ResponseExistingKey: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseFloat(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseFloat(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseFloat(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseInteger(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseInteger(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseInteger(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseInteger: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseLong(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseLong(context.Background(), "positive")
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseLong(context.Background(), "negative")
	if err != nil {
		t.Fatalf("ResponseLong: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseProtectedKey(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseProtectedKey(context.Background())
	if err != nil {
		t.Fatalf("ResponseProtectedKey: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestHeaderResponseString(t *testing.T) {
	client := newHeaderClient()
	result, err := client.ResponseString(context.Background(), "valid")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseString(context.Background(), "null")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)

	result, err = client.ResponseString(context.Background(), "empty")
	if err != nil {
		t.Fatalf("ResponseString: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
