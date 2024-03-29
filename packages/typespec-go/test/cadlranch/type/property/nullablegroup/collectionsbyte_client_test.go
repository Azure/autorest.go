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

func TestCollectionsByteClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsByteClient().GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsByteProperty{
		NullableProperty: [][]byte{
			[]byte("hello, world!"),
			[]byte("hello, world!"),
		},
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsByteProperty)
}

func TestCollectionsByteClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsByteClient().GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsByteProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsByteProperty)
}

func TestCollectionsByteClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsByteClient().PatchNonNull(context.Background(), nullablegroup.CollectionsByteProperty{
		NullableProperty: [][]byte{
			[]byte("hello, world!"),
			[]byte("hello, world!"),
		},
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestCollectionsByteClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsByteClient().PatchNull(context.Background(), nullablegroup.CollectionsByteProperty{
		NullableProperty: azcore.NullValue[[][]byte](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
