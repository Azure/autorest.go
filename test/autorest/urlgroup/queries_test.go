// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newQueriesClient() *QueriesClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewQueriesClient(pl)
}

// ArrayStringCSVEmpty - Get an empty array [] of string using the csv-array format
func TestArrayStringCSVEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVEmpty(context.Background(), &QueriesClientArrayStringCSVEmptyOptions{
		ArrayQuery: []string{},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ArrayStringCSVNull - Get a null array of string using the csv-array format
func TestArrayStringCSVNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// ArrayStringCSVValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the csv-array format
func TestArrayStringCsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringCSVValid(context.Background(), &QueriesClientArrayStringCSVValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ArrayStringPipesValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the pipes-array format
func TestArrayStringPipesValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringPipesValid(context.Background(), &QueriesClientArrayStringPipesValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ArrayStringSsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the ssv-array format
func TestArrayStringSsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringSsvValid(context.Background(), &QueriesClientArrayStringSsvValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ArrayStringTsvValid - Get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the tsv-array format
func TestArrayStringTsvValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ArrayStringTsvValid(context.Background(), &QueriesClientArrayStringTsvValidOptions{
		ArrayQuery: []string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ByteEmpty - Get '' as byte array
func TestByteEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// ByteMultiByte - Get '啊齄丂狛狜隣郎隣兀﨩' multibyte value as utf-8 encoded byte array
func TestByteMultiByte(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteMultiByte(context.Background(), &QueriesClientByteMultiByteOptions{
		ByteQuery: []byte("啊齄丂狛狜隣郎隣兀﨩"),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// ByteNull - Get null as byte array (no query parameters in uri)
func TestByteNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.ByteNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DateNull - Get null as date - this should result in no query parameters in uri
func TestDateNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DateTimeNull - Get null as date-time, should result in no query parameters in uri
func TestDateTimeNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateTimeNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DateTimeValid - Get '2012-01-01T01:01:01Z' as date-time
func TestDateTimeValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateTimeValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DateValid - Get '2012-01-01' as date
func TestDateValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DateValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DoubleDecimalNegative - Get '-9999999.999' numeric value
func TestDoubleDecimalNegative(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleDecimalNegative(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DoubleDecimalPositive - Get '9999999.999' numeric value
func TestDoubleDecimalPositive(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleDecimalPositive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// DoubleNull - Get null numeric value (no query parameter)
func TestDoubleNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.DoubleNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// EnumNull - Get null (no query parameter in url)
func TestEnumNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.EnumNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// EnumValid - Get using uri with query parameter 'green color'
func TestEnumValid(t *testing.T) {
	client := newQueriesClient()
	result, err := client.EnumValid(context.Background(), &QueriesClientEnumValidOptions{
		EnumQuery: to.Ptr(URIColorGreenColor),
	})
	require.NoError(t, err)
	require.Zero(t, result)
}

// FloatNull - Get null numeric value (no query parameter)
func TestFloatNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// FloatScientificNegative - Get '-1.034E-20' numeric value
func TestFloatScientificNegative(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatScientificNegative(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// FloatScientificPositive - Get '1.034E+20' numeric value
func TestFloatScientificPositive(t *testing.T) {
	client := newQueriesClient()
	result, err := client.FloatScientificPositive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetBooleanFalse - Get false Boolean value on path
func TestGetBooleanFalse(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanFalse(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetBooleanNull - Get null Boolean value on query (query string should be absent)
func TestGetBooleanNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetBooleanTrue - Get true Boolean value on path
func TestGetBooleanTrue(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetBooleanTrue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetIntNegativeOneMillion - Get '-1000000' integer value
func TestGetIntNegativeOneMillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntNegativeOneMillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetIntNull - Get null integer value (no query parameter)
func TestGetIntNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetIntOneMillion - Get '1000000' integer value
func TestGetIntOneMillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetIntOneMillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetLongNull - Get 'null 64 bit integer value (no query param in uri)
func TestGetLongNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetLongNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetNegativeTenBillion - Get '-10000000000' 64 bit integer value
func TestGetNegativeTenBillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetNegativeTenBillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// GetTenBillion - Get '10000000000' 64 bit integer value
func TestGetTenBillion(t *testing.T) {
	client := newQueriesClient()
	result, err := client.GetTenBillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// StringEmpty - Get ''
func TestStringEmpty(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// StringNull - Get null (no query parameter in url)
func TestStringNull(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringNull(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// StringURLEncoded - Get 'begin!*'();:@ &=+$,/?#[]end
func TestStringURLEncoded(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringURLEncoded(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// StringUnicode - Get '啊齄丂狛狜隣郎隣兀﨩' multi-byte string value
func TestStringUnicode(t *testing.T) {
	client := newQueriesClient()
	result, err := client.StringUnicode(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
