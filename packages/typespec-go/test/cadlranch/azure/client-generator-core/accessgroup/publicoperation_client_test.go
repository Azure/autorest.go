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
	client, err := NewAccessClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPublicOperationClient().NoDecoratorInPublic(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Equal(t, NoDecoratorModelInPublic{
		Name: to.Ptr("foo"),
	}, resp.NoDecoratorModelInPublic)
}

func TestPublicOperationClient_PublicDecoratorInPublic(t *testing.T) {
	client, err := NewAccessClient(nil)
	require.NoError(t, err)
	resp, err := client.NewPublicOperationClient().PublicDecoratorInPublic(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Equal(t, PublicDecoratorModelInPublic{
		Name: to.Ptr("foo"),
	}, resp.PublicDecoratorModelInPublic)
}
