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

func TestOptionalDurationClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDurationClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, "P123DT22H14M12.011S", *resp.Property)
}

func TestOptionalDurationClient_GetDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDurationClient().GetDefault(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalDurationClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDurationClient().PutAll(context.Background(), optionalitygroup.DurationProperty{
		Property: to.Ptr("P123DT22H14M12.011S"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalDurationClient_PutDefault(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalDurationClient().PutDefault(context.Background(), optionalitygroup.DurationProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
