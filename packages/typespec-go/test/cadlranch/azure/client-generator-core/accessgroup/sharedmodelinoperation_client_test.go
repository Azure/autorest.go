// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package accessgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestSharedModelInOperationClient_internalMethod(t *testing.T) {
	client, err := NewAccessClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSharedModelInOperationClient().internalMethod(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Equal(t, SharedModel{
		Name: to.Ptr("foo"),
	}, resp.SharedModel)
}

func TestSharedModelInOperationClient_Public(t *testing.T) {
	client, err := NewAccessClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSharedModelInOperationClient().Public(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Equal(t, SharedModel{
		Name: to.Ptr("foo"),
	}, resp.SharedModel)
}
