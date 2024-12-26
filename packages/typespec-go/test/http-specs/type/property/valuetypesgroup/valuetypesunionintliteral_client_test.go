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

func TestValueTypesUnionIntLiteralClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionIntLiteralClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.UnionIntLiteralPropertyProperty42, *resp.Property)
}

func TestValueTypesUnionIntLiteralClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionIntLiteralClient().Put(context.Background(), valuetypesgroup.UnionIntLiteralProperty{
		Property: to.Ptr(valuetypesgroup.UnionIntLiteralPropertyProperty42),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
