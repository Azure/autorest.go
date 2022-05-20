// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimegroup

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

func newDatetimeClient() *DatetimeClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewDatetimeClient(pl)
}

func TestGetInvalid(t *testing.T) {
	client := newDatetimeClient()
	_, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
}

func TestGetLocalNegativeOffsetLowercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalNegativeOffsetLowercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalNegativeOffsetMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalNegativeOffsetMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00-14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalNegativeOffsetUppercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalNegativeOffsetUppercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalPositiveOffsetLowercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalPositiveOffsetLowercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalPositiveOffsetMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalPositiveOffsetMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00+14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalNoOffsetMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalNoOffsetMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetLocalPositiveOffsetUppercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetLocalPositiveOffsetUppercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if result.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestGetOverflow(t *testing.T) {
	t.Skip("API doesn't actually overflow")
	client := newDatetimeClient()
	_, err := client.GetOverflow(context.Background(), nil)
	require.Error(t, err)
}

func TestGetUnderflow(t *testing.T) {
	client := newDatetimeClient()
	_, err := client.GetUnderflow(context.Background(), nil)
	require.Error(t, err)
}

func TestGetUTCLowercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetUTCLowercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetUTCMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetUTCMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetUTCUppercaseMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetUTCUppercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetUTCUppercaseMaxDateTime7Digits(t *testing.T) {
	client := newDatetimeClient()
	result, err := client.GetUTCUppercaseMaxDateTime7Digits(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.9999999Z")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestPutLocalNegativeOffsetMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	require.NoError(t, err)
	result, err := client.PutLocalNegativeOffsetMaxDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutLocalNegativeOffsetMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00-14:00")
	require.NoError(t, err)
	result, err := client.PutLocalNegativeOffsetMinDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutLocalPositiveOffsetMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	require.NoError(t, err)
	result, err := client.PutLocalPositiveOffsetMaxDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutLocalPositiveOffsetMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00+14:00")
	require.NoError(t, err)
	result, err := client.PutLocalPositiveOffsetMinDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutUTCMaxDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	require.NoError(t, err)
	result, err := client.PutUTCMaxDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutUTCMaxDateTime7Digits(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.9999999Z")
	require.NoError(t, err)
	result, err := client.PutUTCMaxDateTime7Digits(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutUTCMinDateTime(t *testing.T) {
	client := newDatetimeClient()
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	require.NoError(t, err)
	result, err := client.PutUTCMinDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
