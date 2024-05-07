//go:build go1.18
// +build go1.18

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
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Default(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Float64Seconds(context.Background(), 35.621, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().FloatSeconds(context.Background(), 35.621, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().Int32Seconds(context.Background(), 36, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().ISO8601(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestHeaderClientISO8601Array(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationHeaderClient().ISO8601Array(context.Background(), []string{"P40D", "P50D"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestPropertyClientDefault(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Default(context.Background(), durationgroup.DefaultDurationProperty{
		Value: to.Ptr("P40D"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "P40D", *resp.Value)
}

func TestPropertyClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Float64Seconds(context.Background(), durationgroup.Float64SecondsDurationProperty{
		Value: to.Ptr(35.621),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float64(35.621), *resp.Value)
}

func TestPropertyClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatSeconds(context.Background(), durationgroup.FloatSecondsDurationProperty{
		Value: to.Ptr[float32](35.621),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, float32(35.621), *resp.Value)
}

func TestPropertyClientFloatSecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().FloatSecondsArray(context.Background(), durationgroup.FloatSecondsDurationArrayProperty{
		Value: []*float32{
			to.Ptr[float32](35.621),
			to.Ptr[float32](46.781),
		},
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, []*float32{
		to.Ptr[float32](35.621),
		to.Ptr[float32](46.781),
	}, resp.Value)
}

func TestPropertyClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().Int32Seconds(context.Background(), durationgroup.Int32SecondsDurationProperty{
		Value: to.Ptr[int32](36),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, int32(36), *resp.Value)
}

func TestPropertyClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationPropertyClient().ISO8601(context.Background(), durationgroup.ISO8601DurationProperty{
		Value: to.Ptr("P40D"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "P40D", *resp.Value)
}

func TestQueryClientDefault(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Default(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloat64Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Float64Seconds(context.Background(), 35.621, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientFloatSeconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().FloatSeconds(context.Background(), 35.621, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32Seconds(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32Seconds(context.Background(), 36, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientInt32SecondsArray(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().Int32SecondsArray(context.Background(), []int32{36, 47}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestQueryClientISO8601(t *testing.T) {
	client, err := durationgroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationQueryClient().ISO8601(context.Background(), "P40D", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
