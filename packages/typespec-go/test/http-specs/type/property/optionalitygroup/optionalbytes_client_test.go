// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalitygroup_test

import (
	"context"
	"optionalitygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionalBytesClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalBytesClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, []byte("hello, world!"), resp.Property)
}

func TestOptionalBytesClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalBytesClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalBytesClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalBytesClient().PutAll(context.Background(), optionalitygroup.BytesProperty{
		Property: []byte("hello, world!"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalBytesClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalBytesClient().PutDefault(context.Background(), optionalitygroup.BytesProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
