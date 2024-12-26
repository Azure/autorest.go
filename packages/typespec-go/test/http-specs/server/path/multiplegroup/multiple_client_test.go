// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multiplegroup_test

import (
	"context"
	"multiplegroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipleClient_NoOperationParams(t *testing.T) {
	client, err := multiplegroup.NewMultipleClient(nil)
	require.NoError(t, err)
	resp, err := client.NoOperationParams(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestMultipleClient_WithOperationPathParam(t *testing.T) {
	client, err := multiplegroup.NewMultipleClient(nil)
	require.NoError(t, err)
	resp, err := client.WithOperationPathParam(context.Background(), "test", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
