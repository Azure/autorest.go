// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package naminggroup_test

import (
	"context"
	"naminggroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestModelClient_Client(t *testing.T) {
	client, err := naminggroup.NewNamingClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNamingClientModelClient().Client(context.Background(), naminggroup.ClientModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelClient_Language(t *testing.T) {
	client, err := naminggroup.NewNamingClient(nil)
	require.NoError(t, err)
	resp, err := client.NewNamingClientModelClient().Language(context.Background(), naminggroup.GoModel{
		DefaultName: to.Ptr(true),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
