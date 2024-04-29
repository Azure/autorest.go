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

func TestValueTypesDecimalClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDecimalClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, 0.33333, *resp.Property)
}

func TestValueTypesDecimalClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDecimalClient().Put(context.Background(), valuetypesgroup.DecimalProperty{
		Property: to.Ptr[float64](0.33333),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
