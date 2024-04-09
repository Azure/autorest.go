// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivengroup_test

import (
	"context"
	"srvdrivengroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestResiliencyServiceDrivenClient_AddOperation(t *testing.T) {
	client, err := srvdrivengroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.AddOperation(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromNone(t *testing.T) {
	client, err := srvdrivengroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), &srvdrivengroup.ResiliencyServiceDrivenClientFromNoneOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneOptional(t *testing.T) {
	client, err := srvdrivengroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivengroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		NewParameter: to.Ptr("new"),
		Parameter:    to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClient_FromOneRequired(t *testing.T) {
	client, err := srvdrivengroup.NewResiliencyServiceDrivenClient(nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", &srvdrivengroup.ResiliencyServiceDrivenClientFromOneRequiredOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
