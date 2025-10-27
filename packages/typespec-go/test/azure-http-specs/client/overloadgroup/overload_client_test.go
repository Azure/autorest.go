// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package overloadgroup_test

import (
	"context"
	"testing"

	"overloadgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOverloadClient_List(t *testing.T) {
	client, err := overloadgroup.NewOverloadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	expectedResources := []*overloadgroup.Resource{
		{
			ID:    to.Ptr("1"),
			Name:  to.Ptr("foo"),
			Scope: to.Ptr("car"),
		},
		{
			ID:    to.Ptr("2"),
			Name:  to.Ptr("bar"),
			Scope: to.Ptr("bike"),
		},
	}
	resp, err := client.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.ResourceArray)
	require.Equal(t, expectedResources, resp.ResourceArray)

}

func TestOverloadClient_ListByScope(t *testing.T) {
	client, err := overloadgroup.NewOverloadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	expectedResources := []*overloadgroup.Resource{
		{
			ID:    to.Ptr("1"),
			Name:  to.Ptr("foo"),
			Scope: to.Ptr("car"),
		},
	}
	scope := "car"
	resp, err := client.ListByScope(context.Background(), scope, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.ResourceArray)
	require.Equal(t, expectedResources, resp.ResourceArray)
}
