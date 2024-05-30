// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package customgroup_test

import (
	"context"
	"customgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/require"
)

func TestCustomClient_Invalid(t *testing.T) {
	client, err := customgroup.NewCustomClientWithKeyCredential(azcore.NewKeyCredential("invalid-key"), nil)
	require.NoError(t, err)
	resp, err := client.Invalid(context.Background(), nil)
	require.Error(t, err)
	require.Zero(t, resp)
}

func TestCustomClient_Valid(t *testing.T) {
	client, err := customgroup.NewCustomClientWithKeyCredential(azcore.NewKeyCredential("valid-key"), nil)
	require.NoError(t, err)
	resp, err := client.Valid(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
