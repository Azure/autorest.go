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

func TestSpreadAliasClient_SpreadAsRequestBody(t *testing.T) {
	client, err := spreadgroup.NewSpreadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadAliasClient().SpreadAsRequestBody(context.Background(), "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadAliasClient_SpreadAsRequestParameter(t *testing.T) {
	client, err := spreadgroup.NewSpreadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadAliasClient().SpreadAsRequestParameter(context.Background(), "1", "bar", "foo", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadAliasClient_SpreadWithMultipleParameters(t *testing.T) {
	client, err := spreadgroup.NewSpreadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpreadAliasClient().SpreadWithMultipleParameters(context.Background(), "1", "bar", "foo", []int32{1, 2}, &spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions{
		OptionalInt:        to.Ptr[int32](1),
		OptionalStringList: []string{"foo", "bar"},
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestSpreadAliasClient_SpreadWithInnerAlias(t *testing.T) {
	client, err := spreadgroup.NewSpreadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewSpreadAliasClient().SpreadParameterWithInnerAlias(context.Background(), "1", "foo", 1, "bar", nil)
	require.NoError(t, err)
}

func TestSpreadAliasClient_SpreadWithInnerModel(t *testing.T) {
	client, err := spreadgroup.NewSpreadClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	_, err = client.NewSpreadAliasClient().SpreadParameterWithInnerModel(context.Background(), "1", "foo", "bar", nil)
	require.NoError(t, err)
}
