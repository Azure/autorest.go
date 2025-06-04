// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup_test

import (
	"context"
	"customgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomClient_Invalid(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), &customgroup.CustomClientInvalidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestCustomClient_OutputToInputOutput(t *testing.T) {
	client, err := customgroup.NewCustomClient(nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), &customgroup.CustomClientValidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
