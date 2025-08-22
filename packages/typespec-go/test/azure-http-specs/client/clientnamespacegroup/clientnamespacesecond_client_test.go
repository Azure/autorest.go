// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientnamespacegroup_test

import (
	"clientnamespacegroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientNamespaceSecondClient_GetSecond(t *testing.T) {
	client, err := clientnamespacegroup.NewClientNamespaceClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	secondClient := client.NewClientNamespaceSecondClient()
	require.NotNil(t, secondClient)
	resp, err := secondClient.GetSecond(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.SecondClientResult)
	require.NotNil(t, resp.SecondClientResult.Type)
	require.Equal(t, clientnamespacegroup.SecondClientEnumTypeSecond, *resp.SecondClientResult.Type)
}
