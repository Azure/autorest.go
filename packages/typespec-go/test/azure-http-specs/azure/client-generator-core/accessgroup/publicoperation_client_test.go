// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package accessgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestPublicOperationClient_NoDecoratorInPublic(t *testing.T) {
	client, err := NewAccessClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessPublicOperationClient().NoDecoratorInPublic(context.Background(), "sample", nil)
	require.NoError(t, err)
	require.Equal(t, NoDecoratorModelInPublic{
		Name: to.Ptr("sample"),
	}, resp.NoDecoratorModelInPublic)
}

func TestPublicOperationClient_PublicDecoratorInPublic(t *testing.T) {
	client, err := NewAccessClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewAccessPublicOperationClient().PublicDecoratorInPublic(context.Background(), "sample", nil)
	require.NoError(t, err)
	require.Equal(t, PublicDecoratorModelInPublic{
		Name: to.Ptr("sample"),
	}, resp.PublicDecoratorModelInPublic)
}
