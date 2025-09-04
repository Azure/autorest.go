// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesUnknownDictClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownDictClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, map[string]any{
		"k1": "hello",
		"k2": float64(42),
	}, resp.Property)
}

func TestValueTypesUnknownDictClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownDictClient().Put(context.Background(), valuetypesgroup.UnknownDictProperty{
		Property: map[string]any{
			"k1": "hello",
			"k2": float64(42),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
