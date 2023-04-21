// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newPathsClient(t *testing.T) *PathsClient {
	client, err := NewPathsClient(nil)
	require.NoError(t, err)
	return client
}

func NewPathsClient(options *azcore.ClientOptions) (*PathsClient, error) {
	client, err := azcore.NewClient("urlgroup.PathsClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &PathsClient{internal: client}, nil
}

func TestArrayCSVInPath(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.ArrayCSVInPath(context.Background(), []string{"ArrayPath1", "begin!*'();:@ &=+$,/?#[]end", "", ""}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsBase64URL(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.Base64URL(context.Background(), []byte("lorem"), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsByteEmpty(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.ByteEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsByteMultiByte(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.ByteMultiByte(context.Background(), []byte("啊齄丂狛狜隣郎隣兀﨩"), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// TODO: check
func TestPathsByteNull(t *testing.T) {
	client := newPathsClient(t)
	_, err := client.ByteNull(context.Background(), nil, nil)
	require.Error(t, err)
}

func TestPathsDateNull(t *testing.T) {
	client := newPathsClient(t)
	var time time.Time
	result, err := client.DateNull(context.Background(), time, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsDateTimeNull(t *testing.T) {
	client := newPathsClient(t)
	var time time.Time
	result, err := client.DateTimeNull(context.Background(), time, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsDateTimeValid(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.DateTimeValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsDateValid(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.DateValid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsDoubleDecimalNegative(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.DoubleDecimalNegative(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsDoubleDecimalPositive(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.DoubleDecimalPositive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsEnumNull(t *testing.T) {
	client := newPathsClient(t)
	var color URIColor
	_, err := client.EnumNull(context.Background(), color, nil)
	require.Error(t, err)
}

func TestPathsEnumValid(t *testing.T) {
	client := newPathsClient(t)
	color := URIColorGreenColor
	result, err := client.EnumValid(context.Background(), color, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsFloatScientificNegative(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.FloatScientificNegative(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsFloatScientificPositive(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.FloatScientificPositive(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetBooleanFalse(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetBooleanFalse(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetBooleanTrue(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetBooleanTrue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetIntNegativeOneMillion(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetIntNegativeOneMillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetIntOneMillion(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetIntOneMillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetNegativeTenBillion(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetNegativeTenBillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsGetTenBillion(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.GetTenBillion(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsStringEmpty(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.StringEmpty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsStringNull(t *testing.T) {
	client := newPathsClient(t)
	var s string
	_, err := client.StringNull(context.Background(), s, nil)
	require.Error(t, err)
}

func TestPathsStringURLEncoded(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.StringURLEncoded(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsStringURLNonEncoded(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.StringURLNonEncoded(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsStringUnicode(t *testing.T) {
	client := newPathsClient(t)
	result, err := client.StringUnicode(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPathsUnixTimeURL(t *testing.T) {
	client := newPathsClient(t)
	d, err := time.Parse("2006-01-02", "2016-04-13")
	require.NoError(t, err)
	result, err := client.UnixTimeURL(context.Background(), d, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
