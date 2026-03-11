// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitindividuallygroup_test

import (
	"clientinitindividuallygroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndividuallyNestedWithParamAliasClient_WithAliasedName(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithAliasedName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestIndividuallyNestedWithParamAliasClient_WithOriginalName(t *testing.T) {
	client, err := clientinitindividuallygroup.NewIndividuallyNestedWithParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	resp, err := client.WithOriginalName(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
