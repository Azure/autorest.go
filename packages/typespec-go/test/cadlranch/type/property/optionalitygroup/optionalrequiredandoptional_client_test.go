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

func TestOptionalRequiredAndOptionalClient_GetAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalRequiredAndOptionalClient().GetAll(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, optionalitygroup.RequiredAndOptionalProperty{
		RequiredProperty: to.Ptr[int32](42),
		OptionalProperty: to.Ptr("hello"),
	}, resp.RequiredAndOptionalProperty)
}

func TestOptionalRequiredAndOptionalClient_GetRequiredOnly(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalRequiredAndOptionalClient().GetRequiredOnly(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, optionalitygroup.RequiredAndOptionalProperty{
		RequiredProperty: to.Ptr[int32](42),
	}, resp.RequiredAndOptionalProperty)
}

func TestOptionalRequiredAndOptionalClient_PutAll(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalRequiredAndOptionalClient().PutAll(context.Background(), optionalitygroup.RequiredAndOptionalProperty{
		RequiredProperty: to.Ptr[int32](42),
		OptionalProperty: to.Ptr("hello"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOptionalRequiredAndOptionalClient_PutRequiredOnly(t *testing.T) {
	client, err := optionalitygroup.NewOptionalClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOptionalRequiredAndOptionalClient().PutRequiredOnly(context.Background(), optionalitygroup.RequiredAndOptionalProperty{
		RequiredProperty: to.Ptr[int32](42),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
