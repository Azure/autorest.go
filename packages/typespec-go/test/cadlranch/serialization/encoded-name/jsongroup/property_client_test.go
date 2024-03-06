// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package jsongroup_test

import (
	"context"
	"jsongroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPropertyClient_Get(t *testing.T) {
	client, err := jsongroup.NewJsonClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, jsongroup.JSONEncodedNameModel{
		DefaultName: to.Ptr(true),
	}, resp.JSONEncodedNameModel)
}

func TestPropertyClient_Send(t *testing.T) {
	client, err := jsongroup.NewJsonClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPropertyClient().Send(context.Background(), jsongroup.JSONEncodedNameModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
