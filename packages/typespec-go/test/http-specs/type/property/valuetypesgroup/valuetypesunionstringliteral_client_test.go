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

func TestValueTypesUnionStringLiteralClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionStringLiteralClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.UnionStringLiteralPropertyPropertyWorld, *resp.Property)
}

func TestValueTypesUnionStringLiteralClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionStringLiteralClient().Put(context.Background(), valuetypesgroup.UnionStringLiteralProperty{
		Property: to.Ptr(valuetypesgroup.UnionStringLiteralPropertyPropertyWorld),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
