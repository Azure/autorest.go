// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package basicgroup_test

import (
	"basicgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBasicClient_CreateOrReplace(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.CreateOrReplace(context.Background(), 1, basicgroup.User{
		Name: to.Ptr("Madge"),
	}, nil)
	require.NoError(t, err)
	require.EqualValues(t, basicgroup.User{
		Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Madge"),
	}, resp.User)
}

func TestBasicClient_CreateOrUpdate(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.CreateOrUpdate(context.Background(), 1, basicgroup.User{
		Name: to.Ptr("Madge"),
	}, nil)
	require.NoError(t, err)
	require.EqualValues(t, basicgroup.User{
		Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Madge"),
	}, resp.User)
}

func TestBasicClient_Delete(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Delete(context.Background(), 1, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestBasicClient_Export(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Export(context.Background(), 1, "json", nil)
	require.NoError(t, err)
	require.EqualValues(t, basicgroup.User{
		Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Madge"),
	}, resp.User)
}

func TestBasicClient_ExportAllUsers(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.ExportAllUsers(context.Background(), "json", nil)
	require.NoError(t, err)
	require.EqualValues(t, []*basicgroup.User{
		{
			Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
			ID:   to.Ptr[int32](1),
			Name: to.Ptr("Madge"),
		},
		{
			Etag: to.Ptr[azcore.ETag]("22bdc430-65e8-45ad-81d9-8ffa60d55b59"),
			ID:   to.Ptr[int32](2),
			Name: to.Ptr("John"),
		},
	}, resp.Users)
}

func TestBasicClient_Get(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), 1, nil)
	require.NoError(t, err)
	require.EqualValues(t, basicgroup.User{
		Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
		ID:   to.Ptr[int32](1),
		Name: to.Ptr("Madge"),
	}, resp.User)
}

func TestBasicClient_NewListPager(t *testing.T) {
	client, err := basicgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	pager := client.NewListPager(&basicgroup.BasicClientListOptions{
		Top:     to.Ptr[int32](5),
		Skip:    to.Ptr[int32](10),
		Orderby: []string{"id"},
		Filter:  to.Ptr("id lt 10"),
		Select:  []string{"id", "orders", "etag"},
		Expand:  []string{"orders"},
	})
	pages := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.EqualValues(t, []*basicgroup.User{
			{
				ID:   to.Ptr[int32](1),
				Name: to.Ptr("Madge"),
				Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b59"),
				Orders: []*basicgroup.UserOrder{
					{
						ID:     to.Ptr[int32](1),
						UserID: to.Ptr[int32](1),
						Detail: to.Ptr("a recorder"),
					},
				},
			},
			{
				ID:   to.Ptr[int32](2),
				Name: to.Ptr("John"),
				Etag: to.Ptr[azcore.ETag]("11bdc430-65e8-45ad-81d9-8ffa60d55b5a"),
				Orders: []*basicgroup.UserOrder{
					{
						ID:     to.Ptr[int32](2),
						UserID: to.Ptr[int32](2),
						Detail: to.Ptr("a TV"),
					},
				},
			},
		}, page.Value)
		pages++
	}
	require.EqualValues(t, 1, pages)
}
