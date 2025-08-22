// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestValueTypesUnionFloatLiteralClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionFloatLiteralClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.UnionFloatLiteralPropertyProperty46875, *resp.Property)
}

func TestValueTypesUnionFloatLiteralClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionFloatLiteralClient().Put(context.Background(), valuetypesgroup.UnionFloatLiteralProperty{
		Property: to.Ptr(valuetypesgroup.UnionFloatLiteralPropertyProperty46875),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
