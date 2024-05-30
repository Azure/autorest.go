// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apikeygroup_test

import (
	"apikeygroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestAPIKeyClient_Invalid(t *testing.T) {
	client, err := apikeygroup.NewAPIKeyClientWithKeyCredential(azcore.NewKeyCredential("invalid-key"), nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestAPIKeyClient_Valid(t *testing.T) {
	client, err := apikeygroup.NewAPIKeyClientWithKeyCredential(azcore.NewKeyCredential("valid-key"), nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
