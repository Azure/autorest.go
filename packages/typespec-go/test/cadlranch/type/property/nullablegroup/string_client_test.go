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

func TestStringClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringClient().GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.StringProperty{
		NullableProperty: to.Ptr("hello"),
		RequiredProperty: to.Ptr("foo"),
	}, resp.StringProperty)
}

func TestStringClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringClient().GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.StringProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.StringProperty)
}

func TestStringClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringClient().PatchNonNull(context.Background(), nullablegroup.StringProperty{
		NullableProperty: to.Ptr("hello"),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestStringClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewNullableClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringClient().PatchNull(context.Background(), nullablegroup.StringProperty{
		NullableProperty: azcore.NullValue[*string](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
