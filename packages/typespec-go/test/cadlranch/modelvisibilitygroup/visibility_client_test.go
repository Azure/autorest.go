//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package modelvisibilitygroup_test

import (
	"context"
	"modelvisibilitygroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestVisibilityClientDeleteModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.DeleteModel(context.Background(), modelvisibilitygroup.VisibilityModel{
		DeleteProp: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientGetModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.GetModel(context.Background(), modelvisibilitygroup.VisibilityModel{
		QueryProp: to.Ptr[int32](123),
	}, nil)
	require.NoError(t, err)
	require.EqualValues(t, modelvisibilitygroup.VisibilityModel{
		ReadProp: to.Ptr("abc"),
	}, resp.VisibilityModel)
}

func TestVisibilityClientHeadModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.HeadModel(context.Background(), modelvisibilitygroup.VisibilityModel{
		QueryProp: to.Ptr[int32](123),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientPatchModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PatchModel(context.Background(), modelvisibilitygroup.VisibilityModel{
		UpdateProp: []*int32{
			to.Ptr[int32](1),
			to.Ptr[int32](2),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientPostModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PostModel(context.Background(), modelvisibilitygroup.VisibilityModel{
		CreateProp: []*string{
			to.Ptr("foo"),
			to.Ptr("bar"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestVisibilityClientPutModel(t *testing.T) {
	client, err := modelvisibilitygroup.NewVisibilityClient(nil)
	require.NoError(t, err)
	resp, err := client.PutModel(context.Background(), modelvisibilitygroup.VisibilityModel{
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
