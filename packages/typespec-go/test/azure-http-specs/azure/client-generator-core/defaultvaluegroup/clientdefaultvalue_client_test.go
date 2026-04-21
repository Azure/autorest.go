// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package defaultvaluegroup_test

import (
	"context"
	"defaultvaluegroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestClientDefaultValueClient_GetHeaderParameter(t *testing.T) {
	client, err := defaultvaluegroup.NewClientDefaultValueClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetHeaderParameter(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientDefaultValueClient_GetOperationParameter(t *testing.T) {
	client, err := defaultvaluegroup.NewClientDefaultValueClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetOperationParameter(context.Background(), "test", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientDefaultValueClient_GetPathParameter(t *testing.T) {
	client, err := defaultvaluegroup.NewClientDefaultValueClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.GetPathParameter(context.Background(), "segment2", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClientDefaultValueClient_PutModelProperty(t *testing.T) {
	t.Skip("https://github.com/Azure/typespec-azure/issues/4295")
	client, err := defaultvaluegroup.NewClientDefaultValueClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.PutModelProperty(context.Background(), defaultvaluegroup.ModelWithDefaultValues{
		Name: to.Ptr("test"),
	}, nil)
	require.NoError(t, err)
	require.Equal(t, defaultvaluegroup.ModelWithDefaultValues{
		Name:    to.Ptr("test"),
		Timeout: to.Ptr(int32(30)),
		Tier:    to.Ptr("standard"),
		Retry:   to.Ptr(true),
	}, resp.ModelWithDefaultValues)
}
