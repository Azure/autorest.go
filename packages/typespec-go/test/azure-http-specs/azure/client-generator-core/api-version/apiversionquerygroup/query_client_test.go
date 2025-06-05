//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionquerygroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryClient_QueryAPIVersion(t *testing.T) {
	client,err := NewQueryClient(nil)
	require.NoError(t, err)
	version := "2025-01-01"
	resp, err := client.QueryAPIVersion(context.Background(), version, nil)
	require.NoError(t, err)
	require.Equal(t, QueryClientQueryAPIVersionResponse{}, resp)
}