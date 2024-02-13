//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nullablegroup_test

import (
	"context"
	"nullablegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBytesClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBytesClient().GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.BytesProperty{
		NullableProperty: []byte("hello, world!"),
		RequiredProperty: to.Ptr("foo"),
	}, resp.BytesProperty)
}

func TestBytesClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBytesClient().GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.BytesProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.BytesProperty)
}

func TestBytesClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBytesClient().PatchNonNull(context.Background(), nullablegroup.BytesProperty{
		NullableProperty: []byte("hello, world!"),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBytesClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBytesClient().PatchNull(context.Background(), nullablegroup.BytesProperty{
		NullableProperty: azcore.NullValue[[]byte](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
