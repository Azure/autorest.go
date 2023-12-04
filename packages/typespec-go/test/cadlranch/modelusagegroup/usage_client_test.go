//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package modelusagegroup_test

import (
	"context"
	"modelusagegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestUsageClientInput(t *testing.T) {
	client, err := modelusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.Input(context.Background(), modelusagegroup.InputRecord{
		RequiredProp: to.Ptr("example-value"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestUsageClientInputAndOutput(t *testing.T) {
	client, err := modelusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.InputAndOutput(context.Background(), modelusagegroup.InputOutputRecord{
		RequiredProp: to.Ptr("example-value"),
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.RequiredProp)
	require.EqualValues(t, "example-value", *resp.RequiredProp)
}

func TestUsageClientOutput(t *testing.T) {
	client, err := modelusagegroup.NewUsageClient(nil)
	require.NoError(t, err)
	resp, err := client.Output(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.RequiredProp)
	require.EqualValues(t, "example-value", *resp.RequiredProp)
}
