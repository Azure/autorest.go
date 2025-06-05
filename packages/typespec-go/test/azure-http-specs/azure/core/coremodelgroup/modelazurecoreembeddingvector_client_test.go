// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coremodelgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModelAzureCoreEmbeddingVectorClient_Get(t *testing.T) {
	expected := []*int32{toInt32(0), toInt32(1), toInt32(2), toInt32(3), toInt32(4)}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)

	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, expected, resp.Int32Array)
}

func TestModelAzureCoreEmbeddingVectorClient_Post(t *testing.T) {
	input := AzureEmbeddingModel{Embedding: []*int32{toInt32(0), toInt32(1), toInt32(2), toInt32(3), toInt32(4)}}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)
	resp, err := client.Post(context.Background(), input, nil)
	require.NoError(t, err)
	require.Len(t, resp.AzureEmbeddingModel.Embedding, 5)
}

func TestModelAzureCoreEmbeddingVectorClient_Put(t *testing.T) {
	expected := []*int32{toInt32(0), toInt32(1), toInt32(2), toInt32(3), toInt32(4)}
	client, err := NewModelAzureCoreEmbeddingVectorClient(nil)
	require.NoError(t, err)
	_, err = client.Put(context.Background(), expected, nil)
	require.NoError(t, err)
}

func toInt32(v int32) *int32 {
	return &v
}
