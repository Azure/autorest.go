// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestValueTypesCollectionsModelClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsModelClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []valuetypesgroup.InnerModel{
		{
			Property: to.Ptr("hello"),
		},
		{
			Property: to.Ptr("world"),
		},
	}, resp.Property)
}

func TestValueTypesCollectionsModelClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsModelClient().Put(context.Background(), valuetypesgroup.CollectionsModelProperty{
		Property: []valuetypesgroup.InnerModel{
			{
				Property: to.Ptr("hello"),
			},
			{
				Property: to.Ptr("world"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
