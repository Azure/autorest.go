// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesUnknownIntClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownIntClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, float64(42), resp.Property)
}

func TestValueTypesUnknownIntClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownIntClient().Put(context.Background(), valuetypesgroup.UnknownIntProperty{
		Property: float64(42),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
