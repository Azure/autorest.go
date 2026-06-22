// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package querygroup_test

import (
	"context"
	"querygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuerySpecialCharClient_DollarSign(t *testing.T) {
	t.Skip("waiting for fix https://github.com/microsoft/typespec/pull/10962")
	client, err := querygroup.NewQueryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewQuerySpecialCharClient().DollarSign(context.Background(), "status eq 'active'", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
