// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivenoldgroup_test

import (
	"context"
	"srvdrivenoldgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestResiliencyServiceDrivenClient_FromNone(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneOptional(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivenoldgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		Parameter: to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneRequired(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
