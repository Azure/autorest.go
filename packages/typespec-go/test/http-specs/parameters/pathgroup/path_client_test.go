// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package pathgroup_test

import (
	"context"
	"pathgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestNormal(t *testing.T) {
	client, err := pathgroup.NewPathClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Normal(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptional_none(t *testing.T) {
	client, err := pathgroup.NewPathClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Optional(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptional_some(t *testing.T) {
	client, err := pathgroup.NewPathClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Optional(context.Background(), &pathgroup.PathClientOptionalOptions{
		Name: to.Ptr("foo"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
