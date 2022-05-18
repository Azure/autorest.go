// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package numbergroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newNumberClient() *NumberClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewNumberClient(pl)
}

func TestNumberGetBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := 2.5976931e+101
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalNegativeDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := -99999999.99
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDecimalPositiveDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := 99999999.99
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDouble(context.Background(), nil)
	require.NoError(t, err)
	val := 2.5976931e+101
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoubleNegativeDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := -99999999.99
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigDoublePositiveDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := 99999999.99
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetBigFloat(context.Background(), nil)
	require.NoError(t, err)
	val := float32(3.402823e+20)
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetInvalidDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDecimal(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestNumberGetInvalidDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidDouble(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestNumberGetInvalidFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetInvalidFloat(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, result)
}

func TestNumberGetNull(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, (*float32)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDecimal(context.Background(), nil)
	require.NoError(t, err)
	val := 2.5976931e-101
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallDouble(context.Background(), nil)
	require.NoError(t, err)
	val := 2.5976931e-101
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberGetSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.GetSmallFloat(context.Background(), nil)
	require.NoError(t, err)
	val := 3.402823e-20
	if r := cmp.Diff(result.Value, &val); r != "" {
		t.Fatal(r)
	}
}

func TestNumberPutBigDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimal(context.Background(), 2.5976931e+101, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigDecimalNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalNegativeDecimal(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigDecimalPositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDecimalPositiveDecimal(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDouble(context.Background(), 2.5976931e+101, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigDoubleNegativeDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoubleNegativeDecimal(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigDoublePositiveDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigDoublePositiveDecimal(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutBigFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutBigFloat(context.Background(), 3.402823e+20, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutSmallDecimal(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDecimal(context.Background(), 2.5976931e-101, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutSmallDouble(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallDouble(context.Background(), 2.5976931e-101, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestNumberPutSmallFloat(t *testing.T) {
	client := newNumberClient()
	result, err := client.PutSmallFloat(context.Background(), 3.402823e-20, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
