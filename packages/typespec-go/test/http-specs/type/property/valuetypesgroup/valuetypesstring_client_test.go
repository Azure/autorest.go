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

func TestValueTypesStringClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.EqualValues(t, "hello", *resp.Property)
}

func TestValueTypesStringClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesStringClient().Put(context.Background(), valuetypesgroup.StringProperty{
		Property: to.Ptr("hello"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
