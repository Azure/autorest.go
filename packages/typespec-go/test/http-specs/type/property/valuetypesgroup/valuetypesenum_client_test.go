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

func TestValueTypesEnumClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesEnumClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.FixedInnerEnumValueOne, *resp.Property)
}

func TestValueTypesEnumClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesEnumClient().Put(context.Background(), valuetypesgroup.EnumProperty{
		Property: to.Ptr(valuetypesgroup.FixedInnerEnumValueOne),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
