// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package querygroup_test

import (
	"context"
	"querygroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryConstantClient_Post(t *testing.T) {
	client, err := querygroup.NewQueryClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewQueryConstantClient().Post(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
