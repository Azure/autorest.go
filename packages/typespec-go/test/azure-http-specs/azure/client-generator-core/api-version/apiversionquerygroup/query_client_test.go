// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionquerygroup_test

import (
	"context"
	"testing"

	"apiversionquerygroup"
	"github.com/stretchr/testify/require"
)

func TestQueryClient_QueryAPIVersion(t *testing.T) {
	client, err := apiversionquerygroup.NewQueryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.QueryAPIVersion(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
