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

func TestCollectionsModelClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsModelClient().GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsModelProperty{
		NullableProperty: []*nullablegroup.InnerModel{
			{
				Property: to.Ptr("hello"),
			},
			{
				Property: to.Ptr("world"),
			},
		},
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsModelProperty)
}

func TestCollectionsModelClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsModelClient().GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.CollectionsModelProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.CollectionsModelProperty)
}

func TestCollectionsModelClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsModelClient().PatchNonNull(context.Background(), nullablegroup.CollectionsModelProperty{
		NullableProperty: []*nullablegroup.InnerModel{
			{
				Property: to.Ptr("hello"),
			},
			{
				Property: to.Ptr("world"),
			},
		},
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestCollectionsModelClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNullableCollectionsModelClient().PatchNull(context.Background(), nullablegroup.CollectionsModelProperty{
		NullableProperty: azcore.NullValue[[]*nullablegroup.InnerModel](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
