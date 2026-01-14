// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipleservicegroup_test

import (
	"context"
	"multipleservicegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinedClient(t *testing.T) {
	client, err := multipleservicegroup.NewCombinedClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.NewCombinedFooClient().Test(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	resp2, err := client.NewCombinedBarClient().Test(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp2)
}
