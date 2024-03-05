//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package arraygroup_test

import (
	"arraygroup"
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBooleanValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBooleanValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, []bool{true, false}, resp.Value)
}

func TestBooleanValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBooleanValueClient().Put(context.Background(), []bool{true, false}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDatetimeValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDatetimeValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []time.Time{time.Date(2022, time.August, 26, 18, 38, 0, 0, time.UTC)}, resp.Value)
}

func TestDatetimeValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDatetimeValueClient().Put(context.Background(), []time.Time{time.Date(2022, time.August, 26, 18, 38, 0, 0, time.UTC)}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDurationValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []string{"P123DT22H14M12.011S"}, resp.Value)
}

func TestDurationValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewDurationValueClient().Put(context.Background(), []string{"P123DT22H14M12.011S"}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFloat32ValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewFloat32ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []float32{43.125}, resp.Value)
}

func TestFloat32ValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewFloat32ValueClient().Put(context.Background(), []float32{43.125}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestInt32ValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewInt32ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []int32{1, 2}, resp.Value)
}

func TestInt32ValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewInt32ValueClient().Put(context.Background(), []int32{1, 2}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestInt64ValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewInt64ValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []int64{9007199254740991, -9007199254740991}, resp.Value)
}

func TestInt64ValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewInt64ValueClient().Put(context.Background(), []int64{9007199254740991, -9007199254740991}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewModelValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []arraygroup.InnerModel{
		{
			Property: to.Ptr("hello"),
		},
		{
			Property: to.Ptr("world"),
		},
	}, resp.Value)
}

func TestModelValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewModelValueClient().Put(context.Background(), []arraygroup.InnerModel{
		{
			Property: to.Ptr("hello"),
		},
		{
			Property: to.Ptr("world"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestNullableFloatValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableFloatValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []*float32{to.Ptr[float32](1.25), nil, to.Ptr[float32](3)}, resp.Value)
}

func TestNullableFloatValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableFloatValueClient().Put(context.Background(), []*float32{to.Ptr[float32](1.25), nil, to.Ptr[float32](3)}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestStringValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []string{"hello", ""}, resp.Value)
}

func TestStringValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringValueClient().Put(context.Background(), []string{"hello", ""}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUnknownValueClientGet(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnknownValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []any{float64(1), "hello", nil}, resp.Value)
}

func TestUnknownValueClientPut(t *testing.T) {
	client, err := arraygroup.NewArrayClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUnknownValueClient().Put(context.Background(), []any{float64(1), "hello", nil}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
