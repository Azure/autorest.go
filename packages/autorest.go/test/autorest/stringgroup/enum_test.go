// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newEnumClient(t *testing.T) *EnumClient {
	client, err := NewEnumClient(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewEnumClient(endpoint string, options *azcore.ClientOptions) (*EnumClient, error) {
	client, err := azcore.NewClient("stringgroup.EnumClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &EnumClient{internal: client, endpoint: endpoint}, nil
}

func TestEnumGetNotExpandable(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.GetNotExpandable(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(ColorsRedColor)); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferenced(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.GetReferenced(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, to.Ptr(ColorsRedColor)); r != "" {
		t.Fatal(r)
	}
}

func TestEnumGetReferencedConstant(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.GetReferencedConstant(context.Background(), nil)
	require.NoError(t, err)
	val := "Sample String"
	if r := cmp.Diff(result.RefColorConstant, RefColorConstant{Field1: &val}); r != "" {
		t.Fatal(r)
	}
}

func TestEnumPutNotExpandable(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.PutNotExpandable(context.Background(), ColorsRedColor, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestEnumPutReferenced(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.PutReferenced(context.Background(), ColorsRedColor, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestEnumPutReferencedConstant(t *testing.T) {
	client := newEnumClient(t)
	result, err := client.PutReferencedConstant(context.Background(), RefColorConstant{}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
