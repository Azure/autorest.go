// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"time"
	"valuetypesgroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestValueTypesDatetimeClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDatetimeClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.EqualValues(t, time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC), *resp.Property)
}

func TestValueTypesDatetimeClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesDatetimeClient().Put(context.Background(), valuetypesgroup.DatetimeProperty{
		Property: to.Ptr(time.Date(2022, 8, 26, 18, 38, 0, 0, time.UTC)),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
