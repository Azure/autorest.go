// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dictionarygroup

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newDictionaryClient() *DictionaryClient {
	return NewDictionaryClient(NewDefaultConnection(nil))
}

// GetArrayEmpty - Get an empty dictionary {}
func TestGetArrayEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(len(resp.Value), 0); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemEmpty - Get an array of array of strings [{"0": ["1", "2", "3"], "1": [], "2": ["7", "8", "9"]}
func TestGetArrayItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string][]string{
		"0": {"1", "2", "3"},
		"1": {},
		"2": {"7", "8", "9"},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayItemNull - Get an dictionary of array of strings {"0": ["1", "2", "3"], "1": null, "2": ["7", "8", "9"]}
func TestGetArrayItemNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayItemNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// TODO: this should technically fail since there's no x-nullable
	if r := cmp.Diff(resp.Value, map[string][]string{
		"0": {"1", "2", "3"},
		"1": nil,
		"2": {"7", "8", "9"},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetArrayNull - Get a null array
func TestGetArrayNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Value != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetArrayValid - Get an array of array of strings {"0": ["1", "2", "3"], "1": ["4", "5", "6"], "2": ["7", "8", "9"]}
func TestGetArrayValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetArrayValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string][]string{
		"0": {"1", "2", "3"},
		"1": {"4", "5", "6"},
		"2": {"7", "8", "9"},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetBase64URL - Get base64url dictionary value {"0": "a string that gets encoded with base64url", "1": "test string", "2": "Lorem ipsum"}
func TestGetBase64URL(t *testing.T) {
	t.Skip("unmarshalling fails")
	client := newDictionaryClient()
	resp, err := client.GetBase64URL(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
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
	t.Skip("no x-nullable, should fail")
	client := newDictionaryClient()
	resp, err := client.GetBooleanInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetBooleanInvalidString - Get boolean dictionary value '{"0": true, "1": "boolean", "2": false}'
func TestGetBooleanInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetBooleanInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetBooleanTfft - Get boolean dictionary value {"0": true, "1": false, "2": false, "3": true }
func TestGetBooleanTfft(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetBooleanTfft(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]bool{
		"0": true,
		"1": false,
		"2": false,
		"3": true,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetByteInvalidNull - Get byte dictionary value {"0": hex(FF FF FF FA), "1": null} with the first item base64 encoded
func TestGetByteInvalidNull(t *testing.T) {
	t.Skip("no x-nullable, should fail")
	client := newDictionaryClient()
	resp, err := client.GetByteInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetByteValid - Get byte dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each item encoded in base64
func TestGetByteValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetByteValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]Widget{}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemEmpty - Get dictionary of complex type with empty item {"0": {"integer": 1, "string": "2"}, "1:" {}, "2": {"integer": 5, "string": "6"}}
func TestGetComplexItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]Widget{
		"0": {Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		"1": {},
		"2": {Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetComplexItemNull - Get dictionary of complex type with null item {"0": {"integer": 1, "string": "2"}, "1": null, "2": {"integer": 5, "string": "6"}}
func TestGetComplexItemNull(t *testing.T) {
	t.Skip("test is invalid, expects nil value but missing x-nullable")
	/*client := newDictionaryClient()
	resp, err := client.GetComplexItemNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]Widget{
		"0": Widget{Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		"1": nil,
		"2": Widget{Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	})*/
}

// GetComplexNull - Get dictionary of complex type null value
func TestGetComplexNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Value != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetComplexValid - Get dictionary of complex type with {"0": {"integer": 1, "string": "2"}, "1": {"integer": 3, "string": "4"}, "2": {"integer": 5, "string": "6"}}
func TestGetComplexValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetComplexValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]Widget{
		"0": {Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		"1": {Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		"2": {Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateInvalidChars - Get date dictionary value {"0": "2011-03-22", "1": "date"}
func TestGetDateInvalidChars(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateInvalidChars(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDateInvalidNull - Get date dictionary value {"0": "2012-01-01", "1": null, "2": "1776-07-04"}
func TestGetDateInvalidNull(t *testing.T) {
	t.Skip("x-nullable")
}

// GetDateTimeInvalidChars - Get date dictionary value {"0": "2000-12-01t00:00:01z", "1": "date-time"}
func TestGetDateTimeInvalidChars(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeInvalidChars(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDateTimeInvalidNull - Get date dictionary value {"0": "2000-12-01t00:00:01z", "1": null}
func TestGetDateTimeInvalidNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDateTimeRFC1123Valid - Get date-time-rfc1123 dictionary value {"0": "Fri, 01 Dec 2000 00:00:01 GMT", "1": "Wed, 02 Jan 1980 00:11:35 GMT", "2": "Wed, 12 Oct 1492 10:15:01 GMT"}
func TestGetDateTimeRFC1123Valid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeRFC1123Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	dt2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	dt3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	if r := cmp.Diff(resp.Value, map[string]time.Time{
		"0": dt1,
		"1": dt2,
		"2": dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateTimeValid - Get date-time dictionary value {"0": "2000-12-01t00:00:01z", "1": "1980-01-02T00:11:35+01:00", "2": "1492-10-12T10:15:01-08:00"}
func TestGetDateTimeValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateTimeValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	dt2, _ := time.Parse(time.RFC3339, "1980-01-02T00:11:35+01:00")
	dt3, _ := time.Parse(time.RFC3339, "1492-10-12T10:15:01-08:00")
	if r := cmp.Diff(resp.Value, map[string]time.Time{
		"0": dt1,
		"1": dt2,
		"2": dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDateValid - Get integer dictionary value {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
func TestGetDateValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDateValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	dt1 := time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC)
	dt3 := time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, map[string]time.Time{
		"0": dt1,
		"1": dt2,
		"2": dt3,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryEmpty - Get an dictionaries of dictionaries of type <string, string> with value {}
func TestGetDictionaryEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]map[string]string{}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemEmpty - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryItemEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryItemEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]map[string]string{
		"0": {
			"1": "one",
			"2": "two",
			"3": "three",
		},
		"1": {},
		"2": {
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDictionaryItemNull - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": null, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryItemNull(t *testing.T) {
	t.Skip("x-nullable")
}

// GetDictionaryNull - Get an dictionaries of dictionaries with value null
func TestGetDictionaryNull(t *testing.T) {
	t.Skip("x-nullable")
}

// GetDictionaryValid - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {"4": "four", "5": "five", "6": "six"}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestGetDictionaryValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDictionaryValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]map[string]string{
		"0": {
			"1": "one",
			"2": "two",
			"3": "three",
		},
		"1": {
			"4": "four",
			"5": "five",
			"6": "six",
		},
		"2": {
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDoubleInvalidNull - Get float dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
func TestGetDoubleInvalidNull(t *testing.T) {
	t.Skip("should fail as mising x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetDoubleInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDoubleInvalidString - Get boolean dictionary value {"0": 1.0, "1": "number", "2": 0.0}
func TestGetDoubleInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDoubleInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetDoubleValid - Get float dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestGetDoubleValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDoubleValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]float64{
		"0": 0,
		"1": -0.01,
		"2": -1.2e20,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetDurationValid - Get duration dictionary value {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
func TestGetDurationValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetDurationValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]string{
		"0": "P123DT22H14M12.011S",
		"1": "P5DT1H",
	}); r != "" {
		t.Fatal(r)
	}
}

// GetEmpty - Get empty dictionary value {}
func TestGetEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Value) != 0 {
		t.Fatal("expected empty dictionary")
	}
}

// GetEmptyStringKey - Get Dictionary with key as empty string
func TestGetEmptyStringKey(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetEmptyStringKey(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]string{"": "val1"}); r != "" {
		t.Fatal(r)
	}
}

// GetFloatInvalidNull - Get float dictionary value {"0": 0.0, "1": null, "2": 1.2e20}
func TestGetFloatInvalidNull(t *testing.T) {
	t.Skip("should fail, nil but no x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetFloatInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetFloatInvalidString - Get boolean dictionary value {"0": 1.0, "1": "number", "2": 0.0}
func TestGetFloatInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetFloatInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetFloatValid - Get float dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestGetFloatValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetFloatValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]float32{
		"0": 0,
		"1": -0.01,
		"2": -1.2e20,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetIntInvalidNull - Get integer dictionary value {"0": 1, "1": null, "2": 0}
func TestGetIntInvalidNull(t *testing.T) {
	t.Skip("should fail, nil but no x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetIntInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetIntInvalidString - Get integer dictionary value {"0": 1, "1": "integer", "2": 0}
func TestGetIntInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetIntInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetIntegerValid - Get integer dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
func TestGetIntegerValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetIntegerValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]int32{
		"0": 1,
		"1": -1,
		"2": 3,
		"3": 300,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetInvalid - Get invalid Dictionary value
func TestGetInvalid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetLongInvalidNull - Get long dictionary value {"0": 1, "1": null, "2": 0}
func TestGetLongInvalidNull(t *testing.T) {
	t.Skip("should fail, nil but no x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetLongInvalidNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetLongInvalidString - Get long dictionary value {"0": 1, "1": "integer", "2": 0}
func TestGetLongInvalidString(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetLongInvalidString(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetLongValid - Get integer dictionary value {"0": 1, "1": -1, "2": 3, "3": 300}
func TestGetLongValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetLongValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]int64{
		"0": 1,
		"1": -1,
		"2": 3,
		"3": 300,
	}); r != "" {
		t.Fatal(r)
	}
}

// GetNull - Get null dictionary value
func TestGetNull(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Value != nil {
		t.Fatal("expected nil map")
	}
}

// GetNullKey - Get Dictionary with null key
func TestGetNullKey(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetNullKey(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetNullValue - Get Dictionary with null value
func TestGetNullValue(t *testing.T) {
	t.Skip("should fail, nil but no x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetNullValue(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Value != nil {
		t.Fatal("expected nil dictionary")
	}
}

// GetStringValid - Get string dictionary value {"0": "foo1", "1": "foo2", "2": "foo3"}
func TestGetStringValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetStringValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(resp.Value, map[string]string{
		"0": "foo1",
		"1": "foo2",
		"2": "foo3",
	}); r != "" {
		t.Fatal(r)
	}
}

// GetStringWithInvalid - Get string dictionary value {"0": "foo", "1": 123, "2": "foo2"}
func TestGetStringWithInvalid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.GetStringWithInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// GetStringWithNull - Get string dictionary value {"0": "foo", "1": null, "2": "foo2"}
func TestGetStringWithNull(t *testing.T) {
	t.Skip("should fail, nil but no x-nullable")
	client := newDictionaryClient()
	resp, err := client.GetStringWithNull(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(resp).IsZero() {
		t.Fatal("expected empty response")
	}
}

// PutArrayValid - Put An array of array of strings {"0": ["1", "2", "3"], "1": ["4", "5", "6"], "2": ["7", "8", "9"]}
func TestPutArrayValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutArrayValid(context.Background(), map[string][]string{
		"0": {"1", "2", "3"},
		"1": {"4", "5", "6"},
		"2": {"7", "8", "9"},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutBooleanTfft - Set dictionary value empty {"0": true, "1": false, "2": false, "3": true }
func TestPutBooleanTfft(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutBooleanTfft(context.Background(), map[string]bool{
		"0": true,
		"1": false,
		"2": false,
		"3": true,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutByteValid - Put the dictionary value {"0": hex(FF FF FF FA), "1": hex(01 02 03), "2": hex (25, 29, 43)} with each elementencoded in base 64
func TestPutByteValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutByteValid(context.Background(), map[string][]byte{
		"0": {255, 255, 255, 250},
		"1": {1, 2, 3},
		"2": {37, 41, 67},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutComplexValid - Put an dictionary of complex type with values {"0": {"integer": 1, "string": "2"}, "1": {"integer": 3, "string": "4"}, "2": {"integer": 5, "string": "6"}}
func TestPutComplexValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutComplexValid(context.Background(), map[string]Widget{
		"0": {Integer: to.Int32Ptr(1), String: to.StringPtr("2")},
		"1": {Integer: to.Int32Ptr(3), String: to.StringPtr("4")},
		"2": {Integer: to.Int32Ptr(5), String: to.StringPtr("6")},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateTimeRFC1123Valid - Set dictionary value empty {"0": "Fri, 01 Dec 2000 00:00:01 GMT", "1": "Wed, 02 Jan 1980 00:11:35 GMT", "2": "Wed, 12 Oct 1492 10:15:01 GMT"}
func TestPutDateTimeRFC1123Valid(t *testing.T) {
	client := newDictionaryClient()
	dt1, _ := time.Parse(time.RFC1123, "Fri, 01 Dec 2000 00:00:01 GMT")
	dt2, _ := time.Parse(time.RFC1123, "Wed, 02 Jan 1980 00:11:35 GMT")
	dt3, _ := time.Parse(time.RFC1123, "Wed, 12 Oct 1492 10:15:01 GMT")
	resp, err := client.PutDateTimeRFC1123Valid(context.Background(), map[string]time.Time{
		"0": dt1,
		"1": dt2,
		"2": dt3,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateTimeValid - Set dictionary value  {"0": "2000-12-01t00:00:01z", "1": "1980-01-02T00:11:35+01:00", "2": "1492-10-12T10:15:01-08:00"}
func TestPutDateTimeValid(t *testing.T) {
	client := newDictionaryClient()
	dt1, _ := time.Parse(time.RFC3339, "2000-12-01T00:00:01Z")
	dt2, _ := time.Parse(time.RFC3339, "1980-01-01T23:11:35Z")
	dt3, _ := time.Parse(time.RFC3339, "1492-10-12T18:15:01Z")
	resp, err := client.PutDateTimeValid(context.Background(), map[string]time.Time{
		"0": dt1,
		"1": dt2,
		"2": dt3,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDateValid - Set dictionary value  {"0": "2000-12-01", "1": "1980-01-02", "2": "1492-10-12"}
func TestPutDateValid(t *testing.T) {
	client := newDictionaryClient()
	d1 := time.Date(2000, 12, 01, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(1980, 01, 02, 0, 0, 0, 0, time.UTC)
	d3 := time.Date(1492, 10, 12, 0, 0, 0, 0, time.UTC)
	resp, err := client.PutDateValid(context.Background(), map[string]time.Time{
		"0": d1,
		"1": d2,
		"2": d3,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDictionaryValid - Get an dictionaries of dictionaries of type <string, string> with value {"0": {"1": "one", "2": "two", "3": "three"}, "1": {"4": "four", "5": "five", "6": "six"}, "2": {"7": "seven", "8": "eight", "9": "nine"}}
func TestPutDictionaryValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDictionaryValid(context.Background(), map[string]map[string]string{
		"0": {
			"1": "one",
			"2": "two",
			"3": "three",
		},
		"1": {
			"4": "four",
			"5": "five",
			"6": "six",
		},
		"2": {
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDoubleValid - Set dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestPutDoubleValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDoubleValid(context.Background(), map[string]float64{
		"0": 0,
		"1": -0.01,
		"2": -1.2e20,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutDurationValid - Set dictionary value  {"0": "P123DT22H14M12.011S", "1": "P5DT1H0M0S"}
func TestPutDurationValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutDurationValid(context.Background(), map[string]string{
		"0": "P123DT22H14M12.011S",
		"1": "P5DT1H",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutEmpty - Set dictionary value empty {}
func TestPutEmpty(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutEmpty(context.Background(), map[string]string{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutFloatValid - Set dictionary value {"0": 0, "1": -0.01, "2": 1.2e20}
func TestPutFloatValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutFloatValid(context.Background(), map[string]float32{
		"0": 0,
		"1": -0.01,
		"2": -1.2e20,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutIntegerValid - Set dictionary value empty {"0": 1, "1": -1, "2": 3, "3": 300}
func TestPutIntegerValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutIntegerValid(context.Background(), map[string]int32{
		"0": 1,
		"1": -1,
		"2": 3,
		"3": 300,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutLongValid - Set dictionary value empty {"0": 1, "1": -1, "2": 3, "3": 300}
func TestPutLongValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutLongValid(context.Background(), map[string]int64{
		"0": 1,
		"1": -1,
		"2": 3,
		"3": 300,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutStringValid - Set dictionary value {"0": "foo1", "1": "foo2", "2": "foo3"}
func TestPutStringValid(t *testing.T) {
	client := newDictionaryClient()
	resp, err := client.PutStringValid(context.Background(), map[string]string{
		"0": "foo1",
		"1": "foo2",
		"2": "foo3",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := resp.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
