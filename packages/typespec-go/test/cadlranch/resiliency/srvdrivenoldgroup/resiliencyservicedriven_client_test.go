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

func TestResiliencyServiceDrivenClientv1_FromNone(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv1_FromOneOptional(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivenoldgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		Parameter: to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv1_FromOneRequired(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv2_FromNone(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv2_FromOneOptional(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivenoldgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		Parameter: to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv2_FromOneRequired(t *testing.T) {
	client, err := srvdrivenoldgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
