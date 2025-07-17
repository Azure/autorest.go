// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalitygroup_test

import (
	"context"
	"optionalitygroup"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestOptionalDatetimeClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDatetimeClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.WithinDuration(t, time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC), *resp.Property, 0)
}

func TestOptionalDatetimeClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDatetimeClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalDatetimeClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDatetimeClient().PutAll(context.Background(), optionalitygroup.DatetimeProperty{
		Property: to.Ptr(time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC)),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalDatetimeClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDatetimeClient().PutDefault(context.Background(), optionalitygroup.DatetimeProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
