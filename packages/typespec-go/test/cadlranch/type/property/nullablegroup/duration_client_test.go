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

func TestDurationClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.DurationProperty{
		NullableProperty: to.Ptr("P123DT22H14M12.011S"),
		RequiredProperty: to.Ptr("foo"),
	}, resp.DurationProperty)
}

func TestDurationClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.DurationProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.DurationProperty)
}

func TestDurationClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.PatchNonNull(context.Background(), nullablegroup.DurationProperty{
		NullableProperty: to.Ptr("P123DT22H14M12.011S"),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDurationClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewDurationClient(nil)
	require.NoError(t, err)
	resp, err := client.PatchNull(context.Background(), nullablegroup.DurationProperty{
		NullableProperty: azcore.NullValue[*string](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
