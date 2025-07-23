// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesBytesClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesBytesClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, []byte("hello, world!"), resp.Property)
}

func TestValueTypesBytesClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesBytesClient().Put(context.Background(), valuetypesgroup.BytesProperty{
		Property: []byte("hello, world!"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
