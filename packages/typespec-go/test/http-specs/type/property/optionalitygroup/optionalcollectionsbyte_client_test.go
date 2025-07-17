// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalitygroup_test

import (
	"context"
	"optionalitygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionalCollectionsByteClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsByteClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, [][]byte{
		[]byte("hello, world!"),
		[]byte("hello, world!"),
	}, resp.Property)
}

func TestOptionalCollectionsByteClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsByteClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalCollectionsByteClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsByteClient().PutAll(context.Background(), optionalitygroup.CollectionsByteProperty{
		Property: [][]byte{
			[]byte("hello, world!"),
			[]byte("hello, world!"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalCollectionsByteClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalCollectionsByteClient().PutDefault(context.Background(), optionalitygroup.CollectionsByteProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
