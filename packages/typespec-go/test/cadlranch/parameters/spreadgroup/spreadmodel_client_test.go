// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package spreadgroup_test

import (
	"context"
	"spreadgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestSpreadModelClient_SpreadAsRequestBody(t *testing.T) {
	client, err := spreadgroup.NewSpreadClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadModelClient().SpreadAsRequestBody(context.Background(), spreadgroup.BodyParameter{
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadModelClient_SpreadCompositeRequest(t *testing.T) {
	client, err := spreadgroup.NewSpreadClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadModelClient().SpreadCompositeRequest(context.Background(), "foo", "bar", spreadgroup.BodyParameter{
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadModelClient_SpreadCompositeRequestMix(t *testing.T) {
	client, err := spreadgroup.NewSpreadClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadModelClient().SpreadCompositeRequestMix(context.Background(), "foo", "bar", spreadgroup.CompositeRequestMix{
		Prop: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadModelClient_SpreadCompositeRequestOnlyWithBody(t *testing.T) {
	client, err := spreadgroup.NewSpreadClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadModelClient().SpreadCompositeRequestOnlyWithBody(context.Background(), spreadgroup.BodyParameter{
		Name: to.Ptr("foo"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadModelClient_SpreadCompositeRequestWithoutBody(t *testing.T) {
	client, err := spreadgroup.NewSpreadClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadModelClient().SpreadCompositeRequestWithoutBody(context.Background(), "foo", "bar", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
