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

func TestValueTypesFloatClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesFloatClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, float32(43.125), *resp.Property)
}

func TestValueTypesFloatClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesFloatClient().Put(context.Background(), valuetypesgroup.FloatProperty{
		Property: to.Ptr[float32](43.125),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
