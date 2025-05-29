// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup_test

import (
	"apikeygroup"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiKeyClient_Invalid(t *testing.T) {
	client, err := apikeygroup.NewApiKeyClient(nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), &apikeygroup.APIKeyClientInvalidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestApiKeyClient_OutputToInputOutput(t *testing.T) {
	client, err := apikeygroup.NewApiKeyClient(nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), &apikeygroup.APIKeyClientValidOptions{})
	require.NoError(t, err)
	require.Zero(t, resp)
}
