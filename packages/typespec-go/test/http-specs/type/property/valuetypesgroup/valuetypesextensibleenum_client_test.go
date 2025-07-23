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

func TestValueTypesExtensibleEnumClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesExtensibleEnumClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, valuetypesgroup.InnerEnum("UnknownValue"), *resp.Property)
}

func TestValueTypesExtensibleEnumClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesExtensibleEnumClient().Put(context.Background(), valuetypesgroup.ExtensibleEnumProperty{
		Property: to.Ptr(valuetypesgroup.InnerEnum("UnknownValue")),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
