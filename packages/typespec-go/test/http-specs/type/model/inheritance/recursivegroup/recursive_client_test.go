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
	client, err := recursivegroup.NewRecursiveClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, recursivegroup.Extension{
		Extension: []recursivegroup.Extension{
			{
				Extension: []recursivegroup.Extension{
					{
						Level: to.Ptr[int8](2),
					},
				},
				Level: to.Ptr[int8](1),
			},
			{
				Level: to.Ptr[int8](1),
			},
		},
		Level: to.Ptr[int8](0),
	}, resp.Extension)
}

func TestRecursiveClientPut(t *testing.T) {
	client, err := recursivegroup.NewRecursiveClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), recursivegroup.Extension{
		Extension: []recursivegroup.Extension{
			{
				Extension: []recursivegroup.Extension{
					{
						Level: to.Ptr[int8](2),
					},
				},
				Level: to.Ptr[int8](1),
			},
			{
				Level: to.Ptr[int8](1),
			},
		},
		Level: to.Ptr[int8](0),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
