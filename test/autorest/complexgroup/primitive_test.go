// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getPrimitiveOperations(t *testing.T) complexgroup.PrimitiveOperations {
	client, err := complexgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create complex client: %v", err)
	}
	return client.PrimitiveOperations()
}

func TestPrimitiveGetInt(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetInt(context.Background())
	if err != nil {
		t.Fatalf("GetInt: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.IntWrapper, &complexgroup.IntWrapper{Field1: to.Int32Ptr(-1), Field2: to.Int32Ptr(2)})
}

func TestPrimitivePutInt(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), complexgroup.IntWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutInt: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetLong(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetLong(context.Background())
	if err != nil {
		t.Fatalf("GetLong: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.LongWrapper, &complexgroup.LongWrapper{
		Field1: to.Int64Ptr(1099511627775),
		Field2: to.Int64Ptr(-999511627788),
	})
}

func TestPrimitivePutLong(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := int64(1099511627775), int64(-999511627788)
	result, err := client.PutLong(context.Background(), complexgroup.LongWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutLong: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetFloat(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetFloat(context.Background())
	if err != nil {
		t.Fatalf("GetFloat: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.FloatWrapper, &complexgroup.FloatWrapper{
		Field1: to.Float32Ptr(1.05),
		Field2: to.Float32Ptr(-0.003),
	})
}

func TestPrimitivePutFloat(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := float32(1.05), float32(-0.003)
	result, err := client.PutFloat(context.Background(), complexgroup.FloatWrapper{Field1: &a, Field2: &b})
	if err != nil {
		t.Fatalf("PutFloat: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetDouble(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetDouble(context.Background())
	if err != nil {
		t.Fatalf("GetDouble: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.DoubleWrapper, &complexgroup.DoubleWrapper{
		Field1: to.Float64Ptr(3e-100),
		Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: to.Float64Ptr(-0.000000000000000000000000000000000000000000000000000000005),
	})
}

func TestPrimitivePutDouble(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	result, err := client.PutDouble(context.Background(), complexgroup.DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b})
	if err != nil {
		t.Fatalf("PutDouble: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetBool(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetBool(context.Background())
	if err != nil {
		t.Fatalf("GetBool: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.BooleanWrapper, &complexgroup.BooleanWrapper{
		FieldFalse: to.BoolPtr(false),
		FieldTrue:  to.BoolPtr(true),
	})
}

func TestPrimitiveGetByte(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetByte(context.Background())
	if err != nil {
		t.Fatalf("GetByte: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.ByteWrapper, &complexgroup.ByteWrapper{Field: &[]byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}})
}

func TestPrimitivePutBool(t *testing.T) {
	client := getPrimitiveOperations(t)
	a, b := true, false
	result, err := client.PutBool(context.Background(), complexgroup.BooleanWrapper{FieldTrue: &a, FieldFalse: &b})
	if err != nil {
		t.Fatalf("PutBool: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitivePutByte(t *testing.T) {
	client := getPrimitiveOperations(t)
	val := []byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}
	result, err := client.PutByte(context.Background(), complexgroup.ByteWrapper{Field: &val})
	if err != nil {
		t.Fatalf("PutByte: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetString(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetString(context.Background())
	if err != nil {
		t.Fatalf("GetString: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.StringWrapper, &complexgroup.StringWrapper{
		Empty: to.StringPtr(""),
		Field: to.StringPtr("goodrequest"),
	})
}

func TestPrimitivePutString(t *testing.T) {
	client := getPrimitiveOperations(t)
	var c *string
	a, b, c := "goodrequest", "", nil
	result, err := client.PutString(context.Background(), complexgroup.StringWrapper{Field: &a, Empty: &b, Null: c})
	if err != nil {
		t.Fatalf("PutString: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// func TestPrimitiveGetDate(t *testing.T) {
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

// func TestPrimitivePutDate(t *testing.T) {
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

func TestPrimitiveGetDuration(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetDuration(context.Background())
	if err != nil {
		t.Fatalf("GetDuration: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.DurationWrapper, &complexgroup.DurationWrapper{
		Field: to.StringPtr("P123DT22H14M12.011S"),
	})
}

func TestPrimitivePutDuration(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.PutDuration(context.Background(), complexgroup.DurationWrapper{Field: to.StringPtr("P123DT22H14M12.011S")})
	if err != nil {
		t.Fatalf("PutDuration: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitiveGetDateTime(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetDateTime(context.Background())
	if err != nil {
		t.Fatalf("GetDateTime: %v", err)
	}
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	helpers.DeepEqualOrFatal(t, result.DatetimeWrapper, &complexgroup.DatetimeWrapper{
		Field: &f,
		Now:   &n,
	})
}

func TestPrimitiveGetDateTimeRFC1123(t *testing.T) {
	client := getPrimitiveOperations(t)
	result, err := client.GetDateTimeRFC1123(context.Background())
	if err != nil {
		t.Fatalf("GetDateTimeRFC1123: %v", err)
	}
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	helpers.DeepEqualOrFatal(t, result.Datetimerfc1123Wrapper, &complexgroup.Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	})
}

func TestPrimitivePutDateTime(t *testing.T) {
	client := getPrimitiveOperations(t)
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	result, err := client.PutDateTime(context.Background(), complexgroup.DatetimeWrapper{
		Field: &f,
		Now:   &n,
	})
	if err != nil {
		t.Fatalf("PutDateTime: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPrimitivePutDateTimeRFC1123(t *testing.T) {
	client := getPrimitiveOperations(t)
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	result, err := client.PutDateTimeRFC1123(context.Background(), complexgroup.Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	})
	if err != nil {
		t.Fatalf("PutDateTimeRFC1123: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
