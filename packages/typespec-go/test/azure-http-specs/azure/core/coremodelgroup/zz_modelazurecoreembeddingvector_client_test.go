// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package coremodelgroup

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModelAzureCoreEmbeddingVectorClient_postCreateRequest(t *testing.T) {
	client := &ModelAzureCoreEmbeddingVectorClient{}
	ctx := context.Background()

	// Prepare a sample AzureEmbeddingModel
	body := AzureEmbeddingModel{
		// Fill with sample data as per the model definition
		Embedding: []*int32{toInt32Ptr(1), toInt32Ptr(2), toInt32Ptr(3)},
	}

	req, err := client.postCreateRequest(ctx, body, nil)
	require.NoError(t, err)
	require.NotNil(t, req)

	// Check method and URL
	rawReq := req.Raw()
	require.Equal(t, http.MethodPost, rawReq.Method)
	require.Contains(t, rawReq.URL.Path, "/azure/core/model/embeddingVector")

	// Check headers
	require.Equal(t, "application/json", rawReq.Header.Get("Accept"))
	require.Equal(t, "application/json", rawReq.Header.Get("Content-Type"))

	// Check body
	b, err := io.ReadAll(rawReq.Body)
	require.NoError(t, err)
	defer rawReq.Body.Close()

	var got AzureEmbeddingModel
	err = json.Unmarshal(b, &got)
	require.NoError(t, err)
	require.Equal(t, body, got)
}

// Helper to get *int32
func toInt32Ptr(v int32) *int32 {
	return &v
}
