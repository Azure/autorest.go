// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dictionarygroup

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

func newDictionaryClient() *DictionaryClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewDictionaryClient(pl)
}

// GetArrayEmpty - Get an empty dictionary {}
func TestGetArrayEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(len(resp.Value), 0); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemEmpty - Get an array of array of strings [{"0": ["1", "2", "3"], "1": [], "2": ["7", "8", "9"]}
func TestGetArrayItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string][]*string{
		"0": to.SliceOfPtrs("1", "2", "3"),
		"1": {},
		"2": to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemNull - Get an dictionary of array of strings {"0": ["1", "2", "3"], "1": null, "2": ["7", "8", "9"]}
func TestGetArrayItemNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayItemNull(context.Background(), nil)
	require.NoError(t, err)
	// TODO: this should technically fail since there's no x-nullable
	if r := cmp.Diff(resp.Value, map[string][]*string{
		"0": to.SliceOfPtrs("1", "2", "3"),
		"1": nil,
		"2": to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayNull - Get a null array
func TestGetArrayNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayNull(context.Background(), nil)
	require.NoError(t, err)
	if resp.Value != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetArrayValid - Get an array of array of strings {"0": ["1", "2", "3"], "1": ["4", "5", "6"], "2": ["7", "8", "9"]}
func TestGetArrayValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string][]*string{
		"0": to.SliceOfPtrs("1", "2", "3"),
		"1": to.SliceOfPtrs("4", "5", "6"),
		"2": to.SliceOfPtrs("7", "8", "9"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBase64URL - Get base64url dictionary value {"0": "a string that gets encoded with base64url", "1": "test string", "2": "Lorem ipsum"}
func TestGetBase64URL(t *testing.T) {
	t.Skip("unmarshalling fails")
	client := newDictionaryClient()
	resp, err := client.GetBase64URL(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string][]byte{
		"0": {},
		"1": {},
		"2": {},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBooleanInvalidNull - Get boolean dictionary value {"0": true, "1": null, "2": false }
func TestGetBooleanInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetBooleanInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*bool{
		"0": to.Ptr(true),
		"1": nil,
		"2": to.Ptr(false),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBooleanInvalidString - Get boolean dictionary value '{"0": true, "1": "boolean", "2": false}'
func TestGetBooleanInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetBooleanInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetBooleanTfft - Get boolean dictionary value {"0": true, "1": false, "2": false, "3": true }
func TestGetBooleanTfft(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetBooleanTfft(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*bool{
		"0": to.Ptr(true),
		"1": to.Ptr(false),
		"2": to.Ptr(false),
		"3": to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetByteInvalidNull - Get byte dictionary value {"0": hex(FF FF FF FA), "1": null} with the first item base64 encoded
func TestGetByteInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetByteInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string][]byte{
		"0": {0xab, 0xac, 0xad},
		"1": nil,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetByteValid - Get byte dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each item encoded in base64
func TestGetByteValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetByteValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string][]byte{
		"0": {255, 255, 255, 250},
		"1": {1, 2, 3},
		"2": {37, 41, 67},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexEmpty - Get empty dictionary of complex type {}
func TestGetComplexEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*Widget{}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemEmpty - Get dictionary of complex type with empty item {"0": {"integer": 1, "string": "2"}, "1:" {}, "2": {"integer": 5, "string": "6"}}
func TestGetComplexItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*Widget{
		"0": {Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		"1": {},
		"2": {Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemNull - Get dictionary of complex type with null item {"0": {"integer": 1, "string": "2"}, "1": null, "2": {"integer": 5, "string": "6"}}
func TestGetComplexItemNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexItemNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*Widget{
		"0": {Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		"1": nil,
		"2": {Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexNull - Get dictionary of complex type null value
func TestGetComplexNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexNull(context.Background(), nil)
	require.NoError(t, err)
	if resp.Value != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetComplexValid - Get dictionary of complex type with {"0": {"integer": 1, "string": "2"}, "1": {"integer": 3, "string": "4"}, "2": {"integer": 5, "string": "6"}}
func TestGetComplexValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*Widget{
		"0": {Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		"1": {Integer: to.Ptr[int32](3), String: to.Ptr("4")},
		"2": {Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateInvalidChars - Get date dictionary value {"0": "2011-03-22", "1": "date"}
func TestGetDateInvalidChars(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateInvalidChars(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDateInvalidNull - Get date dictionary value {"0": "2012-01-01", "1": null, "2": "1776-07-04"}
func TestGetDateInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	v1 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	v3 := time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, map[string]*time.Time{
		"0": &v1,
		"1": nil,
		"2": &v3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeInvalidChars - Get date dictionary value {"0": "2000-12-01t00:00:01z", "1": "date-time"}
func TestGetDateTimeInvalidChars(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeInvalidChars(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDateTimeInvalidNull - Get date dictionary value {"0": "2000-12-01t00:00:01z", "1": null}
func TestGetDateTimeInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	dt1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	if r := cmp.Diff(resp.Value, map[string]*time.Time{
		"0": &dt1,
		"1": nil,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeRFC1123Valid - Get date-time-rfc1123 dictionary value {"0": "Fri, 01 Dec 2000 00:00:01 GMT", "1": "Wed, 02 Jan 1980 00:11:35 GMT", "2": "Wed, 12 Oct 1492 10:15:01 GMT"}
func TestGetDateTimeRFC1123Valid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeRFC1123Valid(context.Background(), nil)
	require.NoError(t, err)
	dt1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	dt2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	dt3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	if r := cmp.Diff(resp.Value, map[string]*time.Time{
		"0": &dt1,
		"1": &dt2,
		"2": &dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeValid - Get date-time dictionary value {"0": "2000-12-01t00:00:01z", "1": "1980-01-02T00:11:35+01:00", "2": "1492-10-12T10:15:01-08:00"}
func TestGetDateTimeValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeValid(context.Background(), nil)
	require.NoError(t, err)
	dt1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	dt2, _ := time.Parse(time.RFC3339, "1980-01-02T00:11:35+01:00")
	dt3, _ := time.Parse(time.RFC3339, "1492-10-12T10:15:01-08:00")
	if r := cmp.Diff(resp.Value, map[string]*time.Time{
		"0": &dt1,
		"1": &dt2,
		"2": &dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateValid - Get integer dictionary value {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
func TestGetDateValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateValid(context.Background(), nil)
	require.NoError(t, err)
	dt1 := time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC)
	dt3 := time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, map[string]*time.Time{
		"0": &dt1,
		"1": &dt2,
		"2": &dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryEmpty - Get an dictionaries of dictionaries of type <string, string> with value {}
func TestGetDictionaryEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]map[string]*string{}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemEmpty - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryItemEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]map[string]*string{
		"0": {
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		"1": {},
		"2": {
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemNull - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": null, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryItemNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryItemNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]map[string]*string{
		"0": {
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		"1": nil,
		"2": {
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryNull - Get an dictionaries of dictionaries with value null
func TestGetDictionaryNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryNull(context.Background(), nil)
	require.NoError(t, err)
	if resp.Value != nil {
		t.Fatal("expected nil value")
	}
}

// GetDictionaryValid - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {"4": "four", "5": "five", "6": "six"}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]map[string]*string{
		"0": {
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		"1": {
			"4": to.Ptr("four"),
			"5": to.Ptr("five"),
			"6": to.Ptr("six"),
		},
		"2": {
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidNull - Get float dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
func TestGetDoubleInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDoubleInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*float64{
		"0": to.Ptr[float64](0),
		"1": nil,
		"2": to.Ptr[float64](-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidString - Get boolean dictionary value {"0": 1.0, "1": "number", "2": 0.0}
func TestGetDoubleInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDoubleInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetDoubleValid - Get float dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestGetDoubleValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDoubleValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*float64{
		"0": to.Ptr[float64](0),
		"1": to.Ptr[float64](-0.01),
		"2": to.Ptr[float64](-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDurationValid - Get duration dictionary value {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
func TestGetDurationValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDurationValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*string{
		"0": to.Ptr("P123DT22H14M12.011S"),
		"1": to.Ptr("P5DT1H"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetEmpty - Get empty dictionary value {}
func TestGetEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if len(resp.Value) != 0 {
		t.Fatal("expected empty dictionary")
	}
}

// GetEmptyStringKey - Get Dictionary with key as empty string
func TestGetEmptyStringKey(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetEmptyStringKey(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*string{"": to.Ptr("val1")}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidNull - Get float dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
func TestGetFloatInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetFloatInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*float32{
		"0": to.Ptr[float32](0),
		"1": nil,
		"2": to.Ptr[float32](-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidString - Get boolean dictionary value {"0": 1.0, "1": "number", "2": 0.0}
func TestGetFloatInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetFloatInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetFloatValid - Get float dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestGetFloatValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetFloatValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*float32{
		"0": to.Ptr[float32](0),
		"1": to.Ptr[float32](-0.01),
		"2": to.Ptr[float32](-1.2e20),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidNull - Get integer dictionary value {"0": 1, "1": null, "2": 0}
func TestGetIntInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetIntInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*int32{
		"0": to.Ptr[int32](1),
		"1": nil,
		"2": to.Ptr[int32](0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidString - Get integer dictionary value {"0": 1, "1": "integer", "2": 0}
func TestGetIntInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetIntInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetIntegerValid - Get integer dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
func TestGetIntegerValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetIntegerValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*int32{
		"0": to.Ptr[int32](1),
		"1": to.Ptr[int32](-1),
		"2": to.Ptr[int32](3),
		"3": to.Ptr[int32](300),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetInvalid - Get invalid Dictionary value
func TestGetInvalid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetLongInvalidNull - Get long dictionary value {"0": 1, "1": null, "2": 0}
func TestGetLongInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetLongInvalidNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*int64{
		"0": to.Ptr[int64](1),
		"1": nil,
		"2": to.Ptr[int64](0),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetLongInvalidString - Get long dictionary value {"0": 1, "1": "integer", "2": 0}
func TestGetLongInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetLongInvalidString(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetLongValid - Get integer dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
func TestGetLongValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetLongValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*int64{
		"0": to.Ptr[int64](1),
		"1": to.Ptr[int64](-1),
		"2": to.Ptr[int64](3),
		"3": to.Ptr[int64](300),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetNull - Get null dictionary value
func TestGetNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if resp.Value != nil {
		t.Fatal("expected nil map")
	}
}

// GetNullKey - Get Dictionary with null key
func TestGetNullKey(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetNullKey(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetNullValue - Get Dictionary with null value
func TestGetNullValue(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetNullValue(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*string{
		"key1": nil,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetStringValid - Get string dictionary value {"0": "foo1", "1": "foo2", "2": "foo3"}
func TestGetStringValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetStringValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*string{
		"0": to.Ptr("foo1"),
		"1": to.Ptr("foo2"),
		"2": to.Ptr("foo3"),
	}); r != "" {
		t.Fatal(r)
	}
}

// GetStringWithInvalid - Get string dictionary value {"0": "foo", "1": 123, "2": "foo2"}
func TestGetStringWithInvalid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetStringWithInvalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

// GetStringWithNull - Get string dictionary value {"0": "foo", "1": null, "2": "foo2"}
func TestGetStringWithNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetStringWithNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(resp.Value, map[string]*string{
		"0": to.Ptr("foo"),
		"1": nil,
		"2": to.Ptr("foo2"),
	}); r != "" {
		t.Fatal(r)
	}
}

// PutArrayValid - Put An array of array of strings {"0": ["1", "2", "3"], "1": ["4", "5", "6"], "2": ["7", "8", "9"]}
func TestPutArrayValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutArrayValid(context.Background(), map[string][]*string{
		"0": to.SliceOfPtrs("1", "2", "3"),
		"1": to.SliceOfPtrs("4", "5", "6"),
		"2": to.SliceOfPtrs("7", "8", "9"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutBooleanTfft - Set dictionary value empty {"0": true, "1": false, "2": false, "3": true }
func TestPutBooleanTfft(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutBooleanTfft(context.Background(), map[string]*bool{
		"0": to.Ptr(true),
		"1": to.Ptr(false),
		"2": to.Ptr(false),
		"3": to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutByteValid - Put the dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each elementencoded in base 64
func TestPutByteValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutByteValid(context.Background(), map[string][]byte{
		"0": {255, 255, 255, 250},
		"1": {1, 2, 3},
		"2": {37, 41, 67},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutComplexValid - Put an dictionary of complex type with values {"0": {"integer": 1, "string": "2"}, "1": {"integer": 3, "string": "4"}, "2": {"integer": 5, "string": "6"}}
func TestPutComplexValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutComplexValid(context.Background(), map[string]*Widget{
		"0": {Integer: to.Ptr[int32](1), String: to.Ptr("2")},
		"1": {Integer: to.Ptr[int32](3), String: to.Ptr("4")},
		"2": {Integer: to.Ptr[int32](5), String: to.Ptr("6")},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateTimeRFC1123Valid - Set dictionary value empty {"0": "Fri, 01 Dec 2000 00:00:01 GMT", "1": "Wed, 02 Jan 1980 00:11:35 GMT", "2": "Wed, 12 Oct 1492 10:15:01 GMT"}
func TestPutDateTimeRFC1123Valid(t *testing.T) {
	client := newDictionaryClient()
	dt1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	dt2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	dt3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	resp, err := client.PutDateTimeRFC1123Valid(context.Background(), map[string]*time.Time{
		"0": &dt1,
		"1": &dt2,
		"2": &dt3,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateTimeValid - Set dictionary value  {"0": "2000-12-01t00:00:01z", "1": "1980-01-02T00:11:35+01:00", "2": "1492-10-12T10:15:01-08:00"}
func TestPutDateTimeValid(t *testing.T) {
	client := newDictionaryClient()
	dt1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	dt2, _ := time.Parse(time.RFC3339, "1980-01-01T23:11:35Z")
	dt3, _ := time.Parse(time.RFC3339, "1492-10-12T18:15:01Z")
	resp, err := client.PutDateTimeValid(context.Background(), map[string]*time.Time{
		"0": &dt1,
		"1": &dt2,
		"2": &dt3,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDateValid - Set dictionary value  {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
func TestPutDateValid(t *testing.T) {
	client := newDictionaryClient()
	d1 := time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC)
	d3 := time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)
	resp, err := client.PutDateValid(context.Background(), map[string]*time.Time{
		"0": &d1,
		"1": &d2,
		"2": &d3,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDictionaryValid - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {"4": "four", "5": "five", "6": "six"}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestPutDictionaryValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDictionaryValid(context.Background(), map[string]map[string]*string{
		"0": {
			"1": to.Ptr("one"),
			"2": to.Ptr("two"),
			"3": to.Ptr("three"),
		},
		"1": {
			"4": to.Ptr("four"),
			"5": to.Ptr("five"),
			"6": to.Ptr("six"),
		},
		"2": {
			"7": to.Ptr("seven"),
			"8": to.Ptr("eight"),
			"9": to.Ptr("nine"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDoubleValid - Set dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestPutDoubleValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDoubleValid(context.Background(), map[string]*float64{
		"0": to.Ptr[float64](0),
		"1": to.Ptr[float64](-0.01),
		"2": to.Ptr[float64](-1.2e20),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutDurationValid - Set dictionary value  {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
func TestPutDurationValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDurationValid(context.Background(), map[string]*string{
		"0": to.Ptr("P123DT22H14M12.011S"),
		"1": to.Ptr("P5DT1H"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutEmpty - Set dictionary value empty {}
func TestPutEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutEmpty(context.Background(), map[string]*string{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutFloatValid - Set dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestPutFloatValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutFloatValid(context.Background(), map[string]*float32{
		"0": to.Ptr[float32](0),
		"1": to.Ptr[float32](-0.01),
		"2": to.Ptr[float32](-1.2e20),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutIntegerValid - Set dictionary value empty {"0": 1, "1": -1, "2": 3, "3": 300}
func TestPutIntegerValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutIntegerValid(context.Background(), map[string]*int32{
		"0": to.Ptr[int32](1),
		"1": to.Ptr[int32](-1),
		"2": to.Ptr[int32](3),
		"3": to.Ptr[int32](300),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutLongValid - Set dictionary value empty {"0": 1, "1": -1, "2": 3, "3": 300}
func TestPutLongValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutLongValid(context.Background(), map[string]*int64{
		"0": to.Ptr[int64](1),
		"1": to.Ptr[int64](-1),
		"2": to.Ptr[int64](3),
		"3": to.Ptr[int64](300),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

// PutStringValid - Set dictionary value {"0": "foo1", "1": "foo2", "2": "foo3"}
func TestPutStringValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutStringValid(context.Background(), map[string]*string{
		"0": to.Ptr("foo1"),
		"1": to.Ptr("foo2"),
		"2": to.Ptr("foo3"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
