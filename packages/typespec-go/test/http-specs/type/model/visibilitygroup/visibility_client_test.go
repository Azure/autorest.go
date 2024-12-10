//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package visibilitygroup_test

import (
	"context"
	"testing"
	"visibilitygroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestVisibilityClientDeleteModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.DeleteModel(context.Background(), visibilitygroup.VisibilityModel{
		DeleteProp: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientGetModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.GetModel(context.Background(), visibilitygroup.VisibilityModel{
		QueryProp: to.Ptr[int32](123),
	}, nil)
	require.NoError(t, err)
	require.EqualValues(t, visibilitygroup.VisibilityModel{
		ReadProp: to.Ptr("abc"),
	}, resp.VisibilityModel)
}

func TestVisibilityClientHeadModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.HeadModel(context.Background(), visibilitygroup.VisibilityModel{
		QueryProp: to.Ptr[int32](123),
	}, nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestVisibilityClientPatchModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PatchModel(context.Background(), visibilitygroup.VisibilityModel{
		UpdateProp: []*int32{
			to.Ptr[int32](1),
			to.Ptr[int32](2),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientPostModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PostModel(context.Background(), visibilitygroup.VisibilityModel{
		CreateProp: []*string{
			to.Ptr("foo"),
			to.Ptr("bar"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientPutModel(t *testing.T) {
	client, err := visibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PutModel(context.Background(), visibilitygroup.VisibilityModel{
		CreateProp: []*string{
			to.Ptr("foo"),
			to.Ptr("bar"),
		},
		UpdateProp: []*int32{
			to.Ptr[int32](1),
			to.Ptr[int32](2),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
