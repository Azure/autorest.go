// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package corescalargroup_test

import (
	"context"
	"corescalargroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestAzureLocationScalarClient_Get(t *testing.T) {
	client, err := corescalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarAzureLocationScalarClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "eastus", *resp.Value)
}

func TestAzureLocationScalarClient_Header(t *testing.T) {
	client, err := corescalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarAzureLocationScalarClient().Header(context.Background(), "eastus", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestAzureLocationScalarClient_Post(t *testing.T) {
	client, err := corescalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarAzureLocationScalarClient().Post(context.Background(), corescalargroup.AzureLocationModel{
		Location: to.Ptr("eastus"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Location)
	require.EqualValues(t, "eastus", *resp.Location)
}

func TestAzureLocationScalarClient_Put(t *testing.T) {
	client, err := corescalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarAzureLocationScalarClient().Put(context.Background(), "eastus", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestAzureLocationScalarClient_Query(t *testing.T) {
	client, err := corescalargroup.NewScalarClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewScalarAzureLocationScalarClient().Query(context.Background(), "eastus", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
