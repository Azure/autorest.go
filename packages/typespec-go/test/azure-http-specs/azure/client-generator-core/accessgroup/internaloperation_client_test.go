// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package accessgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestInternalOperationClient_internalDecoratorInInternal(t *testing.T) {
	client, err := NewAccessClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessInternalOperationClient().internalDecoratorInInternal(context.Background(), "sample", nil)
	require.NoError(t, err)
	require.Equal(t, internalDecoratorModelInInternal{
		Name: to.Ptr("sample"),
	}, resp.internalDecoratorModelInInternal)
}

func TestInternalOperationClient_noDecoratorInInternal(t *testing.T) {
	client, err := NewAccessClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessInternalOperationClient().noDecoratorInInternal(context.Background(), "sample", nil)
	require.NoError(t, err)
	require.Equal(t, noDecoratorModelInInternal{
		Name: to.Ptr("sample"),
	}, resp.noDecoratorModelInInternal)
}

func TestInternalOperationClient_publicDecoratorInInternal(t *testing.T) {
	client, err := NewAccessClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessInternalOperationClient().publicDecoratorInInternal(context.Background(), "sample", nil)
	require.NoError(t, err)
	require.Equal(t, PublicDecoratorModelInInternal{
		Name: to.Ptr("sample"),
	}, resp.PublicDecoratorModelInInternal)
}
