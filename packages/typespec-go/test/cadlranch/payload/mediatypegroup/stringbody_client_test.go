// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypegroup_test

import (
	"context"
	"mediatypegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringBodyClient_GetAsJSON(t *testing.T) {
	client, err := mediatypegroup.NewMediaTypeClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringBodyClient().GetAsJSON(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, "foo", *resp.Value)
}

func TestStringBodyClient_GetAsText(t *testing.T) {
	client, err := mediatypegroup.NewMediaTypeClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringBodyClient().GetAsText(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.Equal(t, "{cat}", *resp.Value)
}

func TestStringBodyClient_SendAsJSON(t *testing.T) {
	client, err := mediatypegroup.NewMediaTypeClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringBodyClient().SendAsJSON(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestStringBodyClient_SendAsText(t *testing.T) {
	client, err := mediatypegroup.NewMediaTypeClient(nil)
	require.NoError(t, err)
	resp, err := client.NewStringBodyClient().SendAsText(context.Background(), "{cat}", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
