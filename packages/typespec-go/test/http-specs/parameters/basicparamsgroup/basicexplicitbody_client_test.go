// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package basicparamsgroup_test

import (
	"basicparamsgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestBasicExplicitBodyClient_Simple(t *testing.T) {
	client, err := basicparamsgroup.NewBasicClient(nil)
	require.NoError(t, err)
	resp, err := client.NewBasicExplicitBodyClient().Simple(context.Background(), basicparamsgroup.User{
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
