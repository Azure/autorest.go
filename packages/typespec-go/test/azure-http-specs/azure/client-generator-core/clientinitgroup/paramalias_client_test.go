// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitgroup_test

import (
	"clientinitgroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamAliasClient_WithAliasedName(t *testing.T) {
	client, err := clientinitgroup.NewParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.WithAliasedName(context.Background(), nil)
	require.NoError(t, err)
}

func TestParamAliasClient_WithOriginalName(t *testing.T) {
	client, err := clientinitgroup.NewParamAliasClientWithNoCredential("http://localhost:3000", "sample-blob", nil)
	require.NoError(t, err)
	_, err = client.WithOriginalName(context.Background(), nil)
	require.NoError(t, err)
}
