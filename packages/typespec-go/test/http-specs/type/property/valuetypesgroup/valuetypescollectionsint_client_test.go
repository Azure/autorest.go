// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesCollectionsIntClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsIntClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []int32{1, 2}, resp.Property)
}

func TestValueTypesCollectionsIntClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsIntClient().Put(context.Background(), valuetypesgroup.CollectionsIntProperty{
		Property: []int32{1, 2},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
