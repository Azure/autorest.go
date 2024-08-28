// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package flattengroup_test

import (
	"context"
	"flattengroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFlattenPropertyClient_PutFlattenModel(t *testing.T) {
	client, err := flattengroup.NewFlattenPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.PutFlattenModel(context.Background(), flattengroup.FlattenModel{
		Name: to.Ptr("foo"),
		Properties: &flattengroup.ChildModel{
			Description: to.Ptr("bar"),
			Age:         to.Ptr[int32](10),
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, flattengroup.FlattenModel{
		Name: to.Ptr("test"),
		Properties: &flattengroup.ChildModel{
			Description: to.Ptr("test"),
			Age:         to.Ptr[int32](1),
		},
	}, resp.FlattenModel)
}

func TestFlattenPropertyClient_PutNestedFlattenModel(t *testing.T) {
	client, err := flattengroup.NewFlattenPropertyClient(nil)
	require.NoError(t, err)
	resp, err := client.PutNestedFlattenModel(context.Background(), flattengroup.NestedFlattenModel{
		Name: to.Ptr("foo"),
		Properties: &flattengroup.ChildFlattenModel{
			Summary: to.Ptr("bar"),
			Properties: &flattengroup.ChildModel{
				Description: to.Ptr("test"),
				Age:         to.Ptr[int32](10),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Equal(t, flattengroup.NestedFlattenModel{
		Name: to.Ptr("test"),
		Properties: &flattengroup.ChildFlattenModel{
			Summary: to.Ptr("test"),
			Properties: &flattengroup.ChildModel{
				Description: to.Ptr("foo"),
				Age:         to.Ptr[int32](1),
			},
		},
	}, resp.NestedFlattenModel)
}
