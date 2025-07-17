//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package usagegroup_test

import (
	"context"
	"testing"
	"usagegroup"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestUsageClientInput(t *testing.T) {
	client, err := usagegroup.NewUsageClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Input(context.Background(), usagegroup.InputRecord{
		RequiredProp: to.Ptr("example-value"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUsageClientInputAndOutput(t *testing.T) {
	client, err := usagegroup.NewUsageClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.InputAndOutput(context.Background(), usagegroup.InputOutputRecord{
		RequiredProp: to.Ptr("example-value"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.RequiredProp)
	require.EqualValues(t, "example-value", *resp.RequiredProp)
}

func TestUsageClientOutput(t *testing.T) {
	client, err := usagegroup.NewUsageClient("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.Output(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.RequiredProp)
	require.EqualValues(t, "example-value", *resp.RequiredProp)
}
