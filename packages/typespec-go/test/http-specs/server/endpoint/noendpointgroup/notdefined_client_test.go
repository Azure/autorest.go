// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package noendpointgroup_test

import (
	"context"
	"noendpointgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotDefinedClient_Valid(t *testing.T) {
	client, err := noendpointgroup.NewNotDefinedClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, resp.Success)
}
