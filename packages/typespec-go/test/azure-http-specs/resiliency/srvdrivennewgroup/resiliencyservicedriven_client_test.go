// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package srvdrivennewgroup_test

import (
	"context"
	"srvdrivennewgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestResiliencyServiceDrivenClientv2_AddOperation(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.AddOperation(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv1_FromNone(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromNoneOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestResiliencyServiceDrivenClientv1_FromOneOptional(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		NewParameter: to.Ptr("new"),
		Parameter:    to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv1_FromOneRequired(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v1", nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneRequiredOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv2_FromNone(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromNone(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromNoneOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestResiliencyServiceDrivenClientv2_FromOneOptional(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromOneOptional(context.Background(), &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneOptionalOptions{
		NewParameter: to.Ptr("new"),
		Parameter:    to.Ptr("optional"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestResiliencyServiceDrivenClientv2_FromOneRequired(t *testing.T) {
	client, err := srvdrivennewgroup.NewResiliencyServiceDrivenClient("v2", nil)
	require.NoError(t, err)
	resp, err := client.FromOneRequired(context.Background(), "required", &srvdrivennewgroup.ResiliencyServiceDrivenClientFromOneRequiredOptions{
		NewParameter: to.Ptr("new"),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
