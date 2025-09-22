// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesCollectionsStringClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []string{"hello", "world"}, resp.Property)
}

func TestValueTypesCollectionsStringClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesCollectionsStringClient().Put(context.Background(), valuetypesgroup.CollectionsStringProperty{
		Property: []string{"hello", "world"},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
