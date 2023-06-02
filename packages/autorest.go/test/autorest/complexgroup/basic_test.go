// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

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

func newBasicClient(t *testing.T) *BasicClient {
	client, err := NewBasicClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewBasicClient(options *azcore.ClientOptions) (*BasicClient, error) {
	client, err := azcore.NewClient("complexgroup.BasicClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &BasicClient{internal: client}, nil
}

func TestBasicGetValid(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{ID: to.Ptr[int32](2), Name: to.Ptr("abc"), Color: to.Ptr(CMYKColorsYELLOW)}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicPutValid(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.PutValid(context.Background(), Basic{
		ID:    to.Ptr[int32](2),
		Name:  to.Ptr("abc"),
		Color: to.Ptr(CMYKColorsMagenta),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestBasicGetInvalid(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
	if r := cmp.Diff(result, BasicClientGetInvalidResponse{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetEmpty(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNull(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNotProvided(t *testing.T) {
	client := newBasicClient(t)
	result, err := client.GetNotProvided(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Basic, Basic{}); r != "" {
		t.Fatal(r)
	}
}
