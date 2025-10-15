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

func TestOptionalPlainDateClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalPlainDateClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.WithinDuration(t, time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC), *resp.Property, 0)
}

func TestOptionalPlainDateClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalPlainDateClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp.Property)
}

func TestOptionalPlainDateClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewOptionalPlainDateClient().PutAll(context.Background(), optionalitygroup.PlainDateProperty{Property: to.Ptr(time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC))}, nil)
	require.NoError(t, err)
}

func TestOptionalPlainDateClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewOptionalPlainDateClient().PutDefault(context.Background(), optionalitygroup.PlainDateProperty{}, nil)
	require.NoError(t, err)
}
