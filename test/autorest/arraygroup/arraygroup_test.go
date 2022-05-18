// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newArrayClient() *ArrayClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewArrayClient(pl)
}

// GetArrayEmpty - Get an empty array []
func TestGetArrayEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.StringArrayArray)
	if r := cmp.Diff(len(resp.StringArrayArray), 0); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemEmpty - Get an array of array of strings [['1', '2', '3'], [], ['7', '8', '9']]
func TestGetArrayItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.SliceOfPtrs("1", "2", "3"),
		{},
		to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemNull - Get an array of array of strings [['1', '2', '3'], null, ['7', '8', '9']]
func TestGetArrayItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayItemNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.SliceOfPtrs("1", "2", "3"),
		nil,
		to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayNull - Get a null array
func TestGetArrayNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayNull(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.StringArrayArray)
}

// GetArrayValid - Get an array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestGetArrayValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetArrayValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArrayArray, [][]*string{
		to.SliceOfPtrs("1", "2", "3"),
		to.SliceOfPtrs("4", "5", "6"),
		to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBase64URL - Get array value ['a string that gets encoded with base64url', 'test string' 'Lorem ipsum'] with the items base64url encoded
func TestGetBase64URL(t *testing.T) {
	t.Skip("decoding fails")
	client := newArrayClient()
	resp, err := client.GetBase64URL(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.ByteArrayArray, [][]byte{
		{0},
		{0},
		{0},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBooleanInvalidNull - Get boolean array value [true, null, false]
func TestGetBooleanInvalidNull(t *testing.T) {
	t.Skip("unmarshalling succeeds")
	client := newArrayClient()
	resp, err := client.GetBooleanInvalidNull(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetBooleanInvalidString - Get boolean array value [true, 'boolean', false]
func TestGetBooleanInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetBooleanInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetBooleanTfft - Get boolean array value [true, false, false, true]
func TestGetBooleanTfft(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetBooleanTfft(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.BoolArray, to.SliceOfPtrs(true, false, false, true)); r != "" {
		t.Fatal(r)
	}
}

// GetByteInvalidNull - Get byte array value [hex(AB, AC, AD), null] with the first item base64 encoded
func TestGetByteInvalidNull(t *testing.T) {
	t.Skip("needs investigation")
	client := newArrayClient()
	resp, err := client.GetByteInvalidNull(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetByteValid - Get byte array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each item encoded in base64
func TestGetByteValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetByteValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.ByteArrayArray, [][]byte{
		{255, 255, 255, 250},
		{1, 2, 3},
		{37, 41, 67},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexEmpty - Get empty array of complex type []
func TestGetComplexEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.ProductArray)
	if r := cmp.Diff(len(resp.ProductArray), 0); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemEmpty - Get array of complex type with empty item [{'integer': 1 'string': '2'}, {}, {'integer': 5, 'string': '6'}]
func TestGetComplexItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		{},
		{Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemNull - Get array of complex type with null item [{'integer': 1 'string': '2'}, null, {'integer': 5, 'string': '6'}]
func TestGetComplexItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexItemNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		nil,
		{Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexNull - Get array of complex type null value
func TestGetComplexNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexNull(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.ProductArray)
}

// GetComplexValid - Get array of complex type with [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestGetComplexValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetComplexValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.ProductArray, []*Product{
		{Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		{Integer: to.Ptr[int32](3), String: to.Ptr("4")},
		{Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateInvalidChars - Get date array value ['2011-03-22', 'date']
func TestGetDateInvalidChars(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateInvalidChars(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDateInvalidNull - Get date array value ['2012-01-01', null, '1776-07-04']
func TestGetDateInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	v1 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	v3 := time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		nil,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeInvalidChars - Get date array value ['2000-12-01t00:00:01z', 'date-time']
func TestGetDateTimeInvalidChars(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeInvalidChars(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDateTimeInvalidNull - Get date array value ['2000-12-01t00:00:01z', null]
func TestGetDateTimeInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		nil,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeRFC1123Valid - Get date-time array value ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestGetDateTimeRFC1123Valid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeRFC1123Valid(context.Background(), nil)
	require.NoError(t, err)
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		&v2,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeValid - Get date-time array value ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestGetDateTimeValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateTimeValid(context.Background(), nil)
	require.NoError(t, err)
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T01:11:35+01:00")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T02:15:01-08:00")
	if r := cmp.Diff(resp.TimeArray, []*time.Time{
		&v1,
		&v2,
		&v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateValid - Get integer array value ['2000-12-01', '1980-01-02', '1492-10-12']
func TestGetDateValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDateValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.TimeArray, to.SliceOfPtrs(
		time.Date(2000, time.December, 01, 0, 0, 0, 0, time.UTC),
		time.Date(1980, time.January, 02, 0, 0, 0, 0, time.UTC),
		time.Date(1492, time.October, 12, 0, 0, 0, 0, time.UTC),
	)); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryEmpty - Get an array of Dictionaries of type <string, string> with value []
func TestGetDictionaryEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemEmpty - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		{},
		{
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemNull - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, null, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryItemNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryItemNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		nil,
		{
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryNull - Get an array of Dictionaries with value null
func TestGetDictionaryNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryNull(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.MapOfStringArray)
}

// GetDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestGetDictionaryValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDictionaryValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.MapOfStringArray, []map[string]*string{
		{
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		{
			"4": to.Ptr("four"),
			"5": to.Ptr("five"),
			"6": to.Ptr("six"),
		},
		{
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetDoubleInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Float64Array, []*float64{
		to.Ptr[float64](0),
		nil,
		to.Ptr(-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetDoubleInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDoubleValid - Get float array value [0, -0.01, 1.2e20]
func TestGetDoubleValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDoubleValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Float64Array, to.SliceOfPtrs[float64](0, -0.01, -1.2e20)); r != "" {
		t.Fatal(r)
	}
}

// GetDurationValid - Get duration array value ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestGetDurationValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetDurationValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArray, to.SliceOfPtrs("P123DT22H14M12.011S", "P5DT1H0M0S")); r != "" {
		t.Fatal(r)
	}
}

// GetEmpty - Get empty array value []
func TestGetEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Int32Array)
	if r := cmp.Diff(len(resp.Int32Array), 0); r != "" {
		t.Fatal(r)
	}
}

// GetEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetEnumValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.FooEnumArray, []*FooEnum{
		to.Ptr(FooEnumFoo1), to.Ptr(FooEnumFoo2), to.Ptr(FooEnumFoo3)}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidNull - Get float array value [0.0, null, -1.2e20]
func TestGetFloatInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Float32Array, []*float32{
		to.Ptr[float32](0),
		nil,
		to.Ptr[float32](-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidString - Get boolean array value [1.0, 'number', 0.0]
func TestGetFloatInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetFloatValid - Get float array value [0, -0.01, 1.2e20]
func TestGetFloatValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetFloatValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Float32Array, to.SliceOfPtrs[float32](0, -0.01, -1.2e20)); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidNull - Get integer array value [1, null, 0]
func TestGetIntInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Int32Array, []*int32{
		to.Ptr[int32](1),
		nil,
		to.Ptr[int32](0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidString - Get integer array value [1, 'integer', 0]
func TestGetIntInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetIntegerValid - Get integer array value [1, -1, 3, 300]
func TestGetIntegerValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetIntegerValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Int32Array, to.SliceOfPtrs[int32](1, -1, 3, 300)); r != "" {
		t.Fatal(r)
	}
}

// GetInvalid - Get invalid array [1, 2, 3
func TestGetInvalid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetLongInvalidNull - Get long array value [1, null, 0]
func TestGetLongInvalidNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Int64Array, []*int64{
		to.Ptr[int64](1),
		nil,
		to.Ptr[int64](0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetLongInvalidString - Get long array value [1, 'integer', 0]
func TestGetLongInvalidString(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetLongValid - Get integer array value [1, -1, 3, 300]
func TestGetLongValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetLongValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Int64Array, to.SliceOfPtrs[int64](1, -1, 3, 300)); r != "" {
		t.Fatal(r)
	}
}

// GetNull - Get null array value
func TestGetNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.Int32Array)
}

// GetStringEnumValid - Get enum array value ['foo1', 'foo2', 'foo3']
func TestGetStringEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringEnumValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Enum0Array, []*Enum0{
		to.Ptr(Enum0Foo1), to.Ptr(Enum0Foo2), to.Ptr(Enum0Foo3)}); r != "" {
		t.Fatal(r)
	}
}

// GetStringValid - Get string array value ['foo1', 'foo2', 'foo3']
func TestGetStringValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArray, to.SliceOfPtrs("foo1", "foo2", "foo3")); r != "" {
		t.Fatal(r)
	}
}

// GetStringWithInvalid - Get string array value ['foo', 123, 'foo2']
func TestGetStringWithInvalid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringWithInvalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetStringWithNull - Get string array value ['foo', null, 'foo2']
func TestGetStringWithNull(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetStringWithNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArray, []*string{to.Ptr("foo"), nil, to.Ptr("foo2")}); r != "" {
		t.Fatal(r)
	}
}

// GetUUIDInvalidChars - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'foo']
func TestGetUUIDInvalidChars(t *testing.T) {
	t.Skip("no strongly-typed UUID")
}

// GetUUIDValid - Get uuid array value ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestGetUUIDValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.GetUUIDValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.StringArray, to.SliceOfPtrs("6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205")); r != "" {
		t.Fatal(r)
	}
}

// PutArrayValid - Put An array of array of strings [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']]
func TestPutArrayValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutArrayValid(context.Background(), [][]*string{
		to.SliceOfPtrs("1", "2", "3"),
		to.SliceOfPtrs("4", "5", "6"),
		to.SliceOfPtrs("7", "8", "9"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutBooleanTfft - Set array value empty [true, false, false, true]
func TestPutBooleanTfft(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutBooleanTfft(context.Background(), to.SliceOfPtrs(true, false, false, true), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutByteValid - Put the array value [hex(FF FF FF FA), hex(01 02 03), hex (25, 29, 43)] with each elementencoded in base 64
func TestPutByteValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutByteValid(context.Background(), [][]byte{
		{0xFF, 0xFF, 0xFF, 0xFA},
		{0x01, 0x02, 0x03},
		{0x25, 0x29, 0x43},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutComplexValid - Put an array of complex type with values [{'integer': 1 'string': '2'}, {'integer': 3, 'string': '4'}, {'integer': 5, 'string': '6'}]
func TestPutComplexValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutComplexValid(context.Background(), []*Product{
		{Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		{Integer: to.Ptr[int32](3), String: to.Ptr("4")},
		{Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateTimeRFC1123Valid - Set array value  ['Fri, 01 Dec 2000 00:00:01 GMT', 'Wed, 02 Jan 1980 00:11:35 GMT', 'Wed, 12 Oct 1492 10:15:01 GMT']
func TestPutDateTimeRFC1123Valid(t *testing.T) {
	client := newArrayClient()
	v1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	v2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	v3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	resp, err := client.PutDateTimeRFC1123Valid(context.Background(), []*time.Time{&v1, &v2, &v3}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateTimeValid - Set array value  ['2000-12-01t00:00:01z', '1980-01-02T00:11:35+01:00', '1492-10-12T10:15:01-08:00']
func TestPutDateTimeValid(t *testing.T) {
	client := newArrayClient()
	v1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	v2, _ := time.Parse(time.RFC3339, "1980-01-02T00:11:35Z")
	v3, _ := time.Parse(time.RFC3339, "1492-10-12T10:15:01Z")
	resp, err := client.PutDateTimeValid(context.Background(), []*time.Time{&v1, &v2, &v3}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateValid - Set array value  ['2000-12-01', '1980-01-02', '1492-10-12']
func TestPutDateValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDateValid(context.Background(), to.SliceOfPtrs(time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC),
		time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC),
		time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDictionaryValid - Get an array of Dictionaries of type <string, string> with value [{'1': 'one', '2': 'two', '3': 'three'}, {'4': 'four', '5': 'five', '6': 'six'}, {'7': 'seven', '8': 'eight', '9': 'nine'}]
func TestPutDictionaryValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDictionaryValid(context.Background(), []map[string]*string{
		{
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		{
			"4": to.Ptr("four"),
			"5": to.Ptr("five"),
			"6": to.Ptr("six"),
		},
		{
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDoubleValid - Set array value [0, -0.01, 1.2e20]
func TestPutDoubleValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDoubleValid(context.Background(), to.SliceOfPtrs[float64](0, -0.01, -1.2e20), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDurationValid - Set array value  ['P123DT22H14M12.011S', 'P5DT1H0M0S']
func TestPutDurationValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutDurationValid(context.Background(), to.SliceOfPtrs("P123DT22H14M12.011S", "P5DT1H"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutEmpty - Set array value empty []
func TestPutEmpty(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutEmpty(context.Background(), []*string{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutEnumValid(context.Background(), []*FooEnum{
		to.Ptr(FooEnumFoo1), to.Ptr(FooEnumFoo2), to.Ptr(FooEnumFoo3)}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutFloatValid - Set array value [0, -0.01, 1.2e20]
func TestPutFloatValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutFloatValid(context.Background(), to.SliceOfPtrs[float32](0, -0.01, -1.2e20), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutIntegerValid - Set array value empty [1, -1, 3, 300]
func TestPutIntegerValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutIntegerValid(context.Background(), to.SliceOfPtrs[int32](1, -1, 3, 300), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutLongValid - Set array value empty [1, -1, 3, 300]
func TestPutLongValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutLongValid(context.Background(), to.SliceOfPtrs[int64](1, -1, 3, 300), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutStringEnumValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringEnumValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutStringEnumValid(context.Background(), []*Enum1{
		to.Ptr(Enum1Foo1), to.Ptr(Enum1Foo2), to.Ptr(Enum1Foo3)}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutStringValid - Set array value ['foo1', 'foo2', 'foo3']
func TestPutStringValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutStringValid(context.Background(), to.SliceOfPtrs("foo1", "foo2", "foo3"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutUUIDValid - Set array value  ['6dcc7237-45fe-45c4-8a6b-3a8a3f625652', 'd1399005-30f7-40d6-8da6-dd7c89ad34db', 'f42f6aa1-a5bc-4ddf-907e-5f915de43205']
func TestPutUUIDValid(t *testing.T) {
	client := newArrayClient()
	resp, err := client.PutUUIDValid(context.Background(), to.SliceOfPtrs("6dcc7237-45fe-45c4-8a6b-3a8a3f625652", "d1399005-30f7-40d6-8da6-dd7c89ad34db", "f42f6aa1-a5bc-4ddf-907e-5f915de43205"), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
