// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package dategroup

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newDateClient() *DateClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewDateClient(pl)
}

func TestGetInvalidDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetInvalidDate(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestGetMaxDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetMaxDate(context.Background(), nil)
	require.NoError(t, err)
	dt := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, &dt); r != "" {
		t.Fatal(r)
	}
}

func TestGetMinDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetMinDate(context.Background(), nil)
	require.NoError(t, err)
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	if r := cmp.Diff(resp.Value, &dt); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if resp.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestGetOverflowDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetOverflowDate(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestGetUnderflowDate(t *testing.T) {
	client := newDateClient()
	resp, err := client.GetUnderflowDate(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestPutMaxDate(t *testing.T) {
	client := newDateClient()
	dt := time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
	result, err := client.PutMaxDate(context.Background(), dt, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutMinDate(t *testing.T) {
	client := newDateClient()
	dt := time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)
	result, err := client.PutMinDate(context.Background(), dt, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
