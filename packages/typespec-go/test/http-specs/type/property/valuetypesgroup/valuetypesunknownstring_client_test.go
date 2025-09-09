// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package valuetypesgroup_test

import (
	"context"
	"testing"
	"valuetypesgroup"

	"github.com/stretchr/testify/require"
)

func TestValueTypesUnknownStringClient_Get(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownStringClient().Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, "hello", resp.Property)
}

func TestValueTypesUnknownStringClient_Put(t *testing.T) {
	client, err := valuetypesgroup.NewValueTypesClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewValueTypesUnknownStringClient().Put(context.Background(), valuetypesgroup.UnknownStringProperty{
		Property: "hello",
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
