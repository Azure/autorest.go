//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package recursivegroup_test

import (
	"context"
	"recursivegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestRecursiveClientGet(t *testing.T) {
	client, err := recursivegroup.NewRecursiveClient(nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, resp.Extension, recursivegroup.Extension{
		Extension: []recursivegroup.Extension{
			{
				Extension: []recursivegroup.Extension{
					{
						Level: to.Ptr[int32](2),
					},
				},
				Level: to.Ptr[int32](1),
			},
			{
				Level: to.Ptr[int32](1),
			},
		},
	})
}

func TestRecursiveClientPut(t *testing.T) {
	client, err := recursivegroup.NewRecursiveClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), recursivegroup.Extension{
		Extension: []recursivegroup.Extension{
			{
				Extension: []recursivegroup.Extension{
					{
						Level: to.Ptr[int32](2),
					},
				},
				Level: to.Ptr[int32](1),
			},
			{
				Level: to.Ptr[int32](1),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
