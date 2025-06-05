// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup_test

import (
	"context"
	"customgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomClient_Invalid_SharedAccessKey(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Expected SharedAccessKey invalid-key")
	require.Zero(t, resp)
}

func TestCustomClient_Valid_WithValidKey(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Expected SharedAccessKey valid-key but got undefined")
	require.Zero(t, resp)
}

func TestCustomClient_OutputToInputOutput(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), &customgroup.CustomClientValidOptions{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Expected SharedAccessKey valid-key but got undefined")
	require.Zero(t, resp)
}
