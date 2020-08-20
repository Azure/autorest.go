// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygrouptest

import (
	"context"
	"generatortests/autorest/generated/arraygroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

// GetArrayEmpty - Get an empty array []
func TestGetArrayEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetArrayEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.StringArrayArray == nil {
		t.Fatal("unexpected nil array")
	}
	if l := len(*resp.StringArrayArray); l != 0 {
		helpers.DeepEqualOrFatal(t, l, 0)
	}
}

// GetArrayItemEmpty - Get an array of array of strings [['1', '2', '3'], [], ['7', '8', '9']]
func TestGetArrayItemEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetArrayItemEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArrayArray, &[][]string{
		[]string{"1", "2", "3"},
		[]string{},
		[]string{"7", "8", "9"},
	})
}

// GetArrayItemNull - Get an array of array of strings [['1', '2', '3'], null, ['7', '8', '9']]
func TestGetArrayItemNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetArrayItemNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArrayArray, &[][]string{
		[]string{"1", "2", "3"},
		nil,
		[]string{"7", "8", "9"},
	})
}

// GetArrayNull - Get a null array
func TestGetArrayNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetArrayNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.StringArrayArray != nil {
		t.Fatal("expected nil array")
	}
}

// GetArrayValid - Get an array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestGetArrayValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetArrayValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArrayArray, &[][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	})
}

// GetBase64URL - Get array value ['a string that gets encoded with base64url', 'test string' 'Lorem ipsum'] with the items base64url encoded
func TestGetBase64URL(t *testing.T) {
	t.Skip("decoding fails")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetBase64URL(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.ByteArrayArray, &[][]byte{
		[]byte{0},
		[]byte{0},
		[]byte{0},
	})
}

// GetBooleanInvalidNull - Get boolean array value [true, null, false]
func TestGetBooleanInvalidNull(t *testing.T) {
	t.Skip("unmarshalling succeeds")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetBooleanInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetBooleanInvalidString - Get boolean array value [true, 'boolean', false]
func TestGetBooleanInvalidString(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetBooleanInvalidString(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetBooleanTfft - Get boolean array value [true, false, false, true]
func TestGetBooleanTfft(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetBooleanTfft(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.BoolArray, &[]bool{true, false, false, true})
}

// GetByteInvalidNull - Get byte array value [hex(AB, AC, AD), null] with the first item base64 encoded
func TestGetByteInvalidNull(t *testing.T) {
	t.Skip("needs investigation")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetByteInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetByteValid - Get byte array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each item encoded in base64
func TestGetByteValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetByteValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.ByteArrayArray, &[][]byte{
		{255, 255, 255, 250},
		{1, 2, 3},
		{37, 41, 67},
	})
}

// GetComplexEmpty - Get empty array of complex type []
func TestGetComplexEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetComplexEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.ProductArray == nil {
		t.Fatal("unexpected nil array")
	}
	if l := len(*resp.ProductArray); l != 0 {
		helpers.DeepEqualOrFatal(t, l, 0)
	}
}

// GetComplexItemEmpty - Get array of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
func TestGetComplexItemEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetComplexItemEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.ProductArray, &[]arraygroup.Product{
		arraygroup.Product{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		arraygroup.Product{},
		arraygroup.Product{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	})
}

// GetComplexItemNull - Get array of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
func TestGetComplexItemNull(t *testing.T) {
	t.Skip("arrays with nil elements")
}

// GetComplexNull - Get array of complex type null value
func TestGetComplexNull(t *testing.T) {
	t.Skip("arrays with nil elements")
}

// GetComplexValid - Get array of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestGetComplexValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetComplexValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.ProductArray, &[]arraygroup.Product{
		arraygroup.Product{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		arraygroup.Product{Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		arraygroup.Product{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	})
}

// GetDateInvalidChars - Get date array value ['2011-03-22', 'date']
func TestGetDateInvalidChars(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateInvalidChars(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetDateInvalidNull - Get date array value ['2012-01-01', null, '1776-07-04']
func TestGetDateInvalidNull(t *testing.T) {
	t.Skip("arrays with nil elements")
}

// GetDateTimeInvalidChars - Get date array value ['2000-12-01t00:00:01z', 'date-time']
func TestGetDateTimeInvalidChars(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateTimeInvalidChars(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetDateTimeInvalidNull - Get date array value ['2000-12-01t00:00:01z', null]
func TestGetDateTimeInvalidNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateTimeInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetDateTimeRFC1123Valid - Get date-time array value ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestGetDateTimeRFC1123Valid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateTimeRFC1123Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	helpers.DeepEqualOrFatal(t, resp.TimeArray, &[]time.Time{
		v1,
		v2,
		v3,
	})
}

// GetDateTimeValid - Get date-time array value ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestGetDateTimeValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateTimeValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T01:11:35+01:00")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T02:15:01-08:00")
	helpers.DeepEqualOrFatal(t, resp.TimeArray, &[]time.Time{
		v1,
		v2,
		v3,
	})
}

// GetDateValid - Get integer array value ['2000-12-01', '1980-01-02', '1492-10-12']
func TestGetDateValid(t *testing.T) {
	t.Skip("needs codegen fix")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDateValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.TimeArray, &[]time.Time{
		time.Date(2000, time.December, 01, 0, 0, 0, 0, time.UTC),
		time.Date(1980, time.January, 02, 0, 0, 0, 0, time.UTC),
		time.Date(1492, time.October, 12, 0, 0, 0, 0, time.UTC),
	})
}

// GetDictionaryEmpty - Get an array of Dictionaries of type <string, string> with value []
func TestGetDictionaryEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDictionaryEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.MapOfStringArray, &[]map[string]string{})
}

// GetDictionaryItemEmpty - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDictionaryItemEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.MapOfStringArray, &[]map[string]string{
		map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
		},
		map[string]string{},
		map[string]string{
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	})
}

// GetDictionaryItemNull - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, null, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDictionaryItemNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.MapOfStringArray, &[]map[string]string{
		map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
		},
		nil,
		map[string]string{
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	})
}

// GetDictionaryNull - Get an array of Dictionaries with value null
func TestGetDictionaryNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDictionaryNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.MapOfStringArray != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDictionaryValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.MapOfStringArray, &[]map[string]string{
		map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
		},
		map[string]string{
			"4": "four",
			"5": "five",
			"6": "six",
		},
		map[string]string{
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	})
}

// GetDoubleInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetDoubleInvalidNull(t *testing.T) {
	t.Skip("arrays with nil elements")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDoubleInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetDoubleInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetDoubleInvalidString(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDoubleInvalidString(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetDoubleValid - Get float array value [0, -0.01, 1.2e20]
func TestGetDoubleValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDoubleValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.Float64Array, &[]float64{0, -0.01, -1.2e20})
}

// GetDurationValid - Get duration array value ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestGetDurationValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetDurationValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArray, &[]string{"P123DT22H14M12.011S", "P5DT1H0M0S"})
}

// GetEmpty - Get empty array value []
func TestGetEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Int32Array == nil {
		t.Fatal("unexpected nil array")
	}
	if l := len(*resp.Int32Array); l != 0 {
		helpers.DeepEqualOrFatal(t, l, 0)
	}
}

// GetEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetEnumValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetEnumValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.FooEnumArray, &[]arraygroup.FooEnum{arraygroup.FooEnumFoo1, arraygroup.FooEnumFoo2, arraygroup.FooEnumFoo3})
}

// GetFloatInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetFloatInvalidNull(t *testing.T) {
	t.Skip("arrays with nil elements")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetFloatInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetFloatInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetFloatInvalidString(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetFloatInvalidString(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetFloatValid - Get float array value [0, -0.01, 1.2e20]
func TestGetFloatValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetFloatValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.Float32Array, &[]float32{0, -0.01, -1.2e20})
}

// GetIntInvalidNull - Get integer array value [1, null, 0]
func TestGetIntInvalidNull(t *testing.T) {
	t.Skip("arrays with nil elements")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetIntInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetIntInvalidString - Get integer array value [1, 'integer', 0]
func TestGetIntInvalidString(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetIntInvalidString(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetIntegerValid - Get integer array value [1, -1, 3, 300]
func TestGetIntegerValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetIntegerValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.Int32Array, &[]int32{1, -1, 3, 300})
}

// GetInvalid - Get invalid array [1, 2, 3
func TestGetInvalid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetLongInvalidNull - Get long array value [1, null, 0]
func TestGetLongInvalidNull(t *testing.T) {
	t.Skip("arrays with nil elements")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetLongInvalidNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetLongInvalidString - Get long array value [1, 'integer', 0]
func TestGetLongInvalidString(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetLongInvalidString(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetLongValid - Get integer array value [1, -1, 3, 300]
func TestGetLongValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetLongValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.Int64Array, &[]int64{1, -1, 3, 300})
}

// GetNull - Get null array value
func TestGetNull(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Int32Array != nil {
		t.Fatal("expected nil array")
	}
}

// GetStringEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetStringEnumValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetStringEnumValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.Enum0Array, &[]arraygroup.Enum0{arraygroup.Enum0Foo1, arraygroup.Enum0Foo2, arraygroup.Enum0Foo3})
}

// GetStringValid - Get string array value ['foo1', 'foo2', 'foo3']
func TestGetStringValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetStringValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArray, &[]string{"foo1", "foo2", "foo3"})
}

// GetStringWithInvalid - Get string array value ['foo', 123, 'foo2']
func TestGetStringWithInvalid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetStringWithInvalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetStringWithNull - Get string array value ['foo', null, 'foo2']
func TestGetStringWithNull(t *testing.T) {
	t.Skip("arrays with nil elements")
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetStringWithNull(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

// GetUUIDInvalidChars - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'foo']
func TestGetUUIDInvalidChars(t *testing.T) {
	t.Skip("no strongly-typed UUID")
}

// GetUUIDValid - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestGetUUIDValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.GetUUIDValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, resp.StringArray, &[]string{"6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205"})
}

// PutArrayValid - Put An array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestPutArrayValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutArrayValid(context.Background(), [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutBooleanTfft - Set array value empty [true, false, false, true]
func TestPutBooleanTfft(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutBooleanTfft(context.Background(), []bool{true, false, false, true})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutByteValid - Put the array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each elementencoded in base 64
func TestPutByteValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutByteValid(context.Background(), [][]byte{
		[]byte{0xFF, 0xFF, 0xFF, 0xFA},
		[]byte{0x01, 0x02, 0x03},
		[]byte{0x25, 0x29, 0x43},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutComplexValid - Put an array of complex type with values [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestPutComplexValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutComplexValid(context.Background(), []arraygroup.Product{
		arraygroup.Product{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		arraygroup.Product{Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		arraygroup.Product{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutDateTimeRFC1123Valid - Set array value  ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestPutDateTimeRFC1123Valid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	resp, err := client.PutDateTimeRFC1123Valid(context.Background(), []time.Time{v1, v2, v3})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutDateTimeValid - Set array value  ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestPutDateTimeValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T00:11:35Z")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T10:15:01Z")
	resp, err := client.PutDateTimeValid(context.Background(), []time.Time{v1, v2, v3})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutDateValid - Set array value  ['2000-12-01', '1980-01-02', '1492-10-12']
func TestPutDateValid(t *testing.T) {
	t.Skip("needs codegen fix")
}

// PutDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestPutDictionaryValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutDictionaryValid(context.Background(), []map[string]string{
		map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
		},
		map[string]string{
			"4": "four",
			"5": "five",
			"6": "six",
		},
		map[string]string{
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutDoubleValid - Set array value [0, -0.01, 1.2e20]
func TestPutDoubleValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutDoubleValid(context.Background(), []float64{0, -0.01, -1.2e20})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutDurationValid - Set array value  ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestPutDurationValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutDurationValid(context.Background(), []string{"P123DT22H14M12.011S", "P5DT1H"})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutEmpty - Set array value empty []
func TestPutEmpty(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutEmpty(context.Background(), []string{})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutEnumValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutEnumValid(context.Background(), []arraygroup.FooEnum{arraygroup.FooEnumFoo1, arraygroup.FooEnumFoo2, arraygroup.FooEnumFoo3})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutFloatValid - Set array value [0, -0.01, 1.2e20]
func TestPutFloatValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutFloatValid(context.Background(), []float32{0, -0.01, -1.2e20})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutIntegerValid - Set array value empty [1, -1, 3, 300]
func TestPutIntegerValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutIntegerValid(context.Background(), []int32{1, -1, 3, 300})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutLongValid - Set array value empty [1, -1, 3, 300]
func TestPutLongValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutLongValid(context.Background(), []int64{1, -1, 3, 300})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutStringEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringEnumValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutStringEnumValid(context.Background(), []arraygroup.Enum1{arraygroup.Enum1Foo1, arraygroup.Enum1Foo2, arraygroup.Enum1Foo3})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutStringValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutStringValid(context.Background(), []string{"foo1", "foo2", "foo3"})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}

// PutUUIDValid - Set array value  ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestPutUUIDValid(t *testing.T) {
	client := arraygroup.NewDefaultClient(nil).ArrayOperations()
	resp, err := client.PutUUIDValid(context.Background(), []string{"6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205"})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, resp, http.StatusOK)
}
