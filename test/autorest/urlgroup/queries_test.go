// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"net/http"
	"testing"
)

func newQueriesClient() *QueriesClient {
	return NewQueriesClient(nil)
}

// ArrayStringCSVEmpty - Get an empty array [] of string using the csv-array format
func TestArrayStringCSVEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVEmpty(context.Background(), &QueriesArrayStringCSVEmptyOptions{
		ArrayQuery: []string{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ArrayStringCSVNull - Get a null array of string using the csv-array format
func TestArrayStringCSVNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ArrayStringCSVValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
func TestArrayStringCsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVValid(context.Background(), &QueriesArrayStringCSVValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ArrayStringPipesValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format
func TestArrayStringPipesValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringPipesValid(context.Background(), &QueriesArrayStringPipesValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ArrayStringSsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format
func TestArrayStringSsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringSsvValid(context.Background(), &QueriesArrayStringSsvValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ArrayStringTsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format
func TestArrayStringTsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringTsvValid(context.Background(), &QueriesArrayStringTsvValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ByteEmpty - Get '' as byte array
func TestByteEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
func TestByteMultiByte(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteMultiByte(context.Background(), &QueriesByteMultiByteOptions{
		ByteQuery: []byte("啊齄丂狛狜隣郎隣兀﨩"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// ByteNull - Get null as byte array (no query parameters in uri)
func TestByteNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DateNull - Get null as date - this should result in no query parameters in uri
func TestDateNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DateTimeNull - Get null as date-time, should result in no query parameters in uri
func TestDateTimeNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateTimeNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time
func TestDateTimeValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateTimeValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DateValid - Get '2012-01-01' as date
func TestDateValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DoubleDecimalNegative - Get '-9999999.999' numeric value
func TestDoubleDecimalNegative(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleDecimalNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DoubleDecimalPositive - Get '9999999.999' numeric value
func TestDoubleDecimalPositive(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleDecimalPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// DoubleNull - Get null numeric value (no query parameter)
func TestDoubleNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// EnumNull - Get null (no query parameter in url)
func TestEnumNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.EnumNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// EnumValid - Get using uri with query parameter 'green color'
func TestEnumValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.EnumValid(context.Background(), &QueriesEnumValidOptions{
		EnumQuery: URIColorGreenColor.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// FloatNull - Get null numeric value (no query parameter)
func TestFloatNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// FloatScientificNegative - Get '-1.034E-20' numeric value
func TestFloatScientificNegative(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatScientificNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// FloatScientificPositive - Get '1.034E+20' numeric value
func TestFloatScientificPositive(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatScientificPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetBooleanFalse - Get false Boolean value on path
func TestGetBooleanFalse(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanFalse(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetBooleanNull - Get null Boolean value on query (query string should be absent)
func TestGetBooleanNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetBooleanTrue - Get true Boolean value on path
func TestGetBooleanTrue(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanTrue(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetIntNegativeOneMillion - Get '-1000000' integer value
func TestGetIntNegativeOneMillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntNegativeOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetIntNull - Get null integer value (no query parameter)
func TestGetIntNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetIntOneMillion - Get '1000000' integer value
func TestGetIntOneMillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetLongNull - Get 'null 64 bit integer value (no query param in uri)
func TestGetLongNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetLongNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value
func TestGetNegativeTenBillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetNegativeTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetTenBillion - Get '10000000000' 64 bit integer value
func TestGetTenBillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// StringEmpty - Get ''
func TestStringEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// StringNull - Get null (no query parameter in url)
func TestStringNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end
func TestStringURLEncoded(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringURLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value
func TestStringUnicode(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringUnicode(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
