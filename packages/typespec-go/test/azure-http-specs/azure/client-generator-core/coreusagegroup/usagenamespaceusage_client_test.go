// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coreusagegroup_test

import (
	"context"
	"coreusagegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestUsageNamespaceUsageClient_NamespaceModelSerializable(t *testing.T) {
	client, err := coreusagegroup.NewUsageClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewUsageNamespaceUsageClient().NamespaceModelSerializable(
		context.Background(),
		coreusagegroup.NamespaceModel{
			Name: to.Ptr("test"),
		},
		nil,
	)
	require.NoError(t, err)
	require.Zero(t, resp)
}
