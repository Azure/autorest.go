// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallyparentgroup_test

import (
	"clientinitindividuallyparentgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndividuallyParentNestedWithParamAliasClient_WithAliasedName(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithAliasedName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyParentNestedWithParamAliasClient_WithOriginalName(t *testing.T) {
	client, err := clientinitindividuallyparentgroup.NewIndividuallyParentIndividuallyParentNestedWithParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithOriginalName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
