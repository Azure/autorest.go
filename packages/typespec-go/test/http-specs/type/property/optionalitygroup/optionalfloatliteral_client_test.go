// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalitygroup_test

import (
	"context"
	"optionalitygroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOptionalFloatLiteralClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalFloatLiteralClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, float32(1.25), *resp.Property)
}

func TestOptionalFloatLiteralClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalFloatLiteralClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalFloatLiteralClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalFloatLiteralClient().PutAll(context.Background(), optionalitygroup.FloatLiteralProperty{
		Property: to.Ptr[float32](1.25),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalFloatLiteralClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalFloatLiteralClient().PutDefault(context.Background(), optionalitygroup.FloatLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
