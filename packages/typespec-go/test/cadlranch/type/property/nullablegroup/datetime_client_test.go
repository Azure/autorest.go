//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nullablegroup_test

import (
	"context"
	"nullablegroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestDatetimeClientGetNonNull(t *testing.T) {
	client, err := nullablegroup.NewDatetimeClient(nil)
	require.NoError(t, err)
	resp, err := client.GetNonNull(context.Background(), nil)
	require.NoError(t, err)
	timeProp, err := time.Parse(time.RFC3339, "2022-08-26T18:38:00Z")
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.DatetimeProperty{
		NullableProperty: &timeProp,
		RequiredProperty: to.Ptr("foo"),
	}, resp.DatetimeProperty)
}

func TestDatetimeClientGetNull(t *testing.T) {
	client, err := nullablegroup.NewDatetimeClient(nil)
	require.NoError(t, err)
	resp, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, nullablegroup.DatetimeProperty{
		RequiredProperty: to.Ptr("foo"),
	}, resp.DatetimeProperty)
}

func TestDatetimeClientPatchNonNull(t *testing.T) {
	client, err := nullablegroup.NewDatetimeClient(nil)
	require.NoError(t, err)
	timeProp, err := time.Parse(time.RFC3339, "2022-08-26T18:38:00Z")
	require.NoError(t, err)
	resp, err := client.PatchNonNull(context.Background(), nullablegroup.DatetimeProperty{
		NullableProperty: &timeProp,
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestDatetimeClientPatchNull(t *testing.T) {
	client, err := nullablegroup.NewDatetimeClient(nil)
	require.NoError(t, err)
	resp, err := client.PatchNull(context.Background(), nullablegroup.DatetimeProperty{
		NullableProperty: azcore.NullValue[*time.Time](),
		RequiredProperty: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
