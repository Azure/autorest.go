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

func TestModelInOperationClient_InputToInputOutput(t *testing.T) {
	client, err := coreusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUsageModelInOperationClient().InputToInputOutput(context.Background(), coreusagegroup.InputModel{
		Name: to.Ptr("Madge"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelInOperationClient_OutputToInputOutput(t *testing.T) {
	client, err := coreusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUsageModelInOperationClient().OutputToInputOutput(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, coreusagegroup.OutputModel{
		Name: to.Ptr("Madge"),
	}, resp.OutputModel)
}

func TestModelInOperationClient_ModelInReadOnlyProperty(t *testing.T) {
	client, err := coreusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.NewUsageModelInOperationClient().ModelInReadOnlyProperty(context.Background(), coreusagegroup.RoundTripModel{}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "Madge", *resp.Result.Name)
}
