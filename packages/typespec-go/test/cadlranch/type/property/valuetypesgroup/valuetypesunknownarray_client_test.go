// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesUnknownArrayClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownArrayClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, []any{"hello", "world"}, resp.Property)
}

func TestValueTypesUnknownArrayClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownArrayClient().Put(context.Background(), valuetypesgroup.UnknownArrayProperty{
		Property: []any{"hello", "world"},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
