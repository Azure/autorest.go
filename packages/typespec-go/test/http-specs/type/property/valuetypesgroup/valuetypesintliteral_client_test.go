// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesIntLiteralClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesIntLiteralClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Property)
	require.Equal(t, int32(42), *resp.Property)
}

func TestValueTypesIntLiteralClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient(nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesIntLiteralClient().Put(context.Background(), valuetypesgroup.IntLiteralProperty{}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
