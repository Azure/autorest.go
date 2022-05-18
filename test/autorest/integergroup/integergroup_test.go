// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package integergroup

import (
	"context"
	"generatortests"
	"math"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newIntClient() *IntClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewIntClient(pl)
}

func TestIntGetInvalid(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetInvalidUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetInvalidUnixTime(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetNull(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, (*int32)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestIntGetNullUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetNullUnixTime(context.Background(), nil)
	require.NoError(t, err)
	if result.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestIntGetOverflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt32(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetOverflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetOverflowInt64(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetUnderflowInt32(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt32(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetUnderflowInt64(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnderflowInt64(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestIntGetUnixTime(t *testing.T) {
	client := newIntClient()
	result, err := client.GetUnixTime(context.Background(), nil)
	require.NoError(t, err)
	t1 := time.Unix(1460505600, 0)
	if r := cmp.Diff(result.Value, &t1); r != "" {
		t.Fatal(r)
	}
}

func TestIntPutMax32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax32(context.Background(), math.MaxInt32, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestIntPutMax64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMax64(context.Background(), math.MaxInt64, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestIntPutMin32(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin32(context.Background(), math.MinInt32, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestIntPutMin64(t *testing.T) {
	client := newIntClient()
	result, err := client.PutMin64(context.Background(), math.MinInt64, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestIntPutUnixTimeDate(t *testing.T) {
	client := newIntClient()
	t1 := time.Unix(1460505600, 0)
	result, err := client.PutUnixTimeDate(context.Background(), t1, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
