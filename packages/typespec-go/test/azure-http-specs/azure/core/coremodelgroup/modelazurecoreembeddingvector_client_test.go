// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coremodelgroup

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestModelAzureCoreEmbeddingVectorClient_Get(t *testing.T) {
	input := []*int32{to.Ptr(int32(0)), to.Ptr(int32(1)), to.Ptr(int32(2)), to.Ptr(int32(3)), to.Ptr(int32(4))}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)

	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, input, resp.Int32Array)
}

func TestModelAzureCoreEmbeddingVectorClient_Post(t *testing.T) {
	input := AzureEmbeddingModel{Embedding: []*int32{to.Ptr(int32(0)), to.Ptr(int32(1)), to.Ptr(int32(2)), to.Ptr(int32(3)), to.Ptr(int32(4))}}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)
	resp, err := client.Post(context.Background(), input, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	expected := AzureEmbeddingModel{Embedding: []*int32{to.Ptr(int32(5)), to.Ptr(int32(6)), to.Ptr(int32(7)), to.Ptr(int32(8)), to.Ptr(int32(9))}}
	require.Equal(t, expected, resp.AzureEmbeddingModel)
}

func TestModelAzureCoreEmbeddingVectorClient_Put(t *testing.T) {
	input := []*int32{to.Ptr(int32(0)), to.Ptr(int32(1)), to.Ptr(int32(2)), to.Ptr(int32(3)), to.Ptr(int32(4))}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), input, nil)
	require.NoError(t, err)
	require.Equal(t, ModelAzureCoreEmbeddingVectorClientPutResponse{}, resp)
}
