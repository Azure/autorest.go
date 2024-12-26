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

func TestValueTypesUnionEnumValueClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionEnumValueClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.ExtendedEnumEnumValue2, *resp.Property)
}

func TestValueTypesUnionEnumValueClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnionEnumValueClient().Put(context.Background(), valuetypesgroup.UnionEnumValueProperty{
		Property: to.Ptr(valuetypesgroup.ExtendedEnumEnumValue2),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
