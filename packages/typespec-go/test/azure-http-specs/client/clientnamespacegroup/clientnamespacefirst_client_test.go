// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientnamespacegroup_test

import (
	"clientnamespacegroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientNamespaceFirstClient_GetFirst(t *testing.T) {
	client, err := clientnamespacegroup.NewClientNamespaceFirstClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.GetFirst(context.Background(), nil) // Use appropriate context and options
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.FirstClientResult)
	require.NotNil(t, resp.FirstClientResult.Name)
	require.Equal(t, "first", *resp.FirstClientResult.Name)
}
