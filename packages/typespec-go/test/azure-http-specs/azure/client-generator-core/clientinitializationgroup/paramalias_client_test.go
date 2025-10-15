// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package clientinitializationgroup_test

import (
    "context"
    "testing"

    "clientinitializationgroup"
    "github.com/stretchr/testify/require"
)

func TestParamAliasClient_WithAliasedName(t *testing.T) {
    client, err := clientinitializationgroup.NewParamAliasClientWithNoCredential("sample-blob", "http://localhost:3000", nil)
    require.NoError(t, err)
    require.NotNil(t, client)

    resp, err := client.WithAliasedName(context.Background(), nil)
    require.NoError(t, err)
    require.NotNil(t, resp)
}

func TestParamAliasClient_WithOriginalName(t *testing.T) {
    client, err := clientinitializationgroup.NewParamAliasClientWithNoCredential("sample-blob", "http://localhost:3000", nil)
    require.NoError(t, err)
    require.NotNil(t, client)

    resp, err := client.WithOriginalName(context.Background(), nil)
    require.NoError(t, err)
    require.NotNil(t, resp)
}
