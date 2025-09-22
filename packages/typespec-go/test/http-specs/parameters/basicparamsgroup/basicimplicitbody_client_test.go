// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package basicparamsgroup_test

import (
	"basicparamsgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicImplicitBodyClient_Simple(t *testing.T) {
	client, err := basicparamsgroup.NewBasicClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewBasicImplicitBodyClient().Simple(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
