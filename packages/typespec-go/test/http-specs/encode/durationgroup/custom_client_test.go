// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package durationgroup_test

import (
	"context"
	"durationgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestHeaderClientDefault(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Default(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloat64Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Float64Milliseconds(context.Background(), 35625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Float64Seconds(context.Background(), 35.625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloatMilliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().FloatMilliseconds(context.Background(), 35625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloatMillisecondsLargerUnit(t *testing.T) {
	t.Skip("https://github.com/microsoft/typespec/issues/8987")
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().FloatMillisecondsLargerUnit(context.Background(), 210000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().FloatSeconds(context.Background(), 35.625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloatSecondsLargerUnit(t *testing.T) {
	t.Skip("https://github.com/microsoft/typespec/issues/8987")
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().FloatSecondsLargerUnit(context.Background(), 150.0, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32Milliseconds(context.Background(), 36000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32MillisecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32MillisecondsArray(context.Background(), []int32{36000, 47000}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32MillisecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32MillisecondsLargerUnit(context.Background(), 180000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32Seconds(context.Background(), 36, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32SecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32SecondsLargerUnit(context.Background(), 120, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().ISO8601(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientISO8601Array(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().ISO8601Array(context.Background(), []string{"P40D", "P50D"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClientDefault(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Default(context.Background(), durationgroup.DefaultDurationProperty{
		Value: to.Ptr("P40D"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "P40D", *resp.Value)
}

func TestPropertyClientFloat64Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Float64Milliseconds(context.Background(), durationgroup.Float64MillisecondsDurationProperty{
		Value: to.Ptr(35625.0),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float64(35625), *resp.Value)
}

func TestPropertyClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Float64Seconds(context.Background(), durationgroup.Float64SecondsDurationProperty{
		Value: to.Ptr(35.625),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float64(35.625), *resp.Value)
}

func TestPropertyClientFloatMilliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatMilliseconds(context.Background(), durationgroup.FloatMillisecondsDurationProperty{
		Value: to.Ptr[float32](35625),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float32(35625), *resp.Value)
}

func TestPropertyClientFloatMillisecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatMillisecondsArray(context.Background(), durationgroup.FloatMillisecondsDurationArrayProperty{
		Value: []*float32{
			to.Ptr[float32](35625),
			to.Ptr[float32](46750),
		},
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []*float32{
		to.Ptr[float32](35625),
		to.Ptr[float32](46750),
	}, resp.Value)
}

func TestPropertyClientFloatMillisecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatMillisecondsLargerUnit(context.Background(), durationgroup.FloatMillisecondsLargerUnitDurationProperty{
		Value: to.Ptr[float32](210000),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float32(210000), *resp.Value)
}

func TestPropertyClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatSeconds(context.Background(), durationgroup.FloatSecondsDurationProperty{
		Value: to.Ptr[float32](35.625),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float32(35.625), *resp.Value)
}

func TestPropertyClientFloatSecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatSecondsArray(context.Background(), durationgroup.FloatSecondsDurationArrayProperty{
		Value: []*float32{
			to.Ptr[float32](35.625),
			to.Ptr[float32](46.75),
		},
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []*float32{
		to.Ptr[float32](35.625),
		to.Ptr[float32](46.75),
	}, resp.Value)
}

func TestPropertyClientFloatSecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatSecondsLargerUnit(context.Background(), durationgroup.FloatSecondsLargerUnitDurationProperty{
		Value: to.Ptr[float32](150.0),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float32(150.0), *resp.Value)
}

func TestPropertyClientInt32Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Int32Milliseconds(context.Background(), durationgroup.Int32MillisecondsDurationProperty{
		Value: to.Ptr[int32](36000),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, int32(36000), *resp.Value)
}

func TestPropertyClientInt32MillisecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Int32MillisecondsLargerUnit(context.Background(), durationgroup.Int32MillisecondsLargerUnitDurationProperty{
		Value: to.Ptr[int32](180000),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, int32(180000), *resp.Value)
}

func TestPropertyClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Int32Seconds(context.Background(), durationgroup.Int32SecondsDurationProperty{
		Value: to.Ptr[int32](36),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, int32(36), *resp.Value)
}

func TestPropertyClientInt32SecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Int32SecondsLargerUnit(context.Background(), durationgroup.Int32SecondsLargerUnitDurationProperty{
		Value: to.Ptr[int32](120),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, int32(120), *resp.Value)
}

func TestPropertyClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().ISO8601(context.Background(), durationgroup.ISO8601DurationProperty{
		Value: to.Ptr("P40D"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "P40D", *resp.Value)
}

func TestQueryClientDefault(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Default(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloat64Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Float64Milliseconds(context.Background(), 35625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Float64Seconds(context.Background(), 35.625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloatMilliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().FloatMilliseconds(context.Background(), 35625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloatMillisecondsLargerUnit(t *testing.T) {
	t.Skip("https://github.com/microsoft/typespec/issues/8987")
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().FloatMillisecondsLargerUnit(context.Background(), 210000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().FloatSeconds(context.Background(), 35.625, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloatSecondsLargerUnit(t *testing.T) {
	t.Skip("https://github.com/microsoft/typespec/issues/8987")
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().FloatSecondsLargerUnit(context.Background(), 150.0, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32Milliseconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32Milliseconds(context.Background(), 36000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32MillisecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32MillisecondsArray(context.Background(), []int32{36000, 47000}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32MillisecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32MillisecondsLargerUnit(context.Background(), 180000, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32Seconds(context.Background(), 36, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32SecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32SecondsArray(context.Background(), []int32{36, 47}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32SecondsLargerUnit(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32SecondsLargerUnit(context.Background(), 120, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().ISO8601(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
