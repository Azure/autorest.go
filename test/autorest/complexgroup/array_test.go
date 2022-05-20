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

func newArrayClient() *ArrayClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewArrayClient(pl)
}

func TestArrayGetEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{
		Array: []*string{},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayGetNotProvided(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayGetValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ArrayWrapper, ArrayWrapper{
		Array: []*string{
			to.Ptr("1, 2, 3, 4"),
			to.Ptr(""),
			nil,
			to.Ptr("&S#$(*Y"),
			to.Ptr("The quick brown fox jumps over the lazy dog"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestArrayPutEmpty(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutEmpty(context.Background(), ArrayWrapper{Array: []*string{}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestArrayPutValid(t *testing.T) {
	client := newArrayClient()
	result, err := client.PutValid(context.Background(), ArrayWrapper{Array: []*string{
		to.Ptr("1, 2, 3, 4"),
		to.Ptr(""),
		nil,
		to.Ptr("&S#$(*Y"),
		to.Ptr("The quick brown fox jumps over the lazy dog"),
	}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
