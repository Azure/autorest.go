//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package apiversionpathgroup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathClient_PathAPIVersion(t *testing.T) {
	client,err := NewPathClient(nil)
	require.NoError(t, err)
	version := "2024-06-01"
	resp, err := client.PathAPIVersion(context.Background(), version, nil)
	require.NoError(t, err)
	require.Equal(t, PathClientPathAPIVersionResponse{}, resp)

	version = ""
	_, err = client.PathAPIVersion(context.Background(), version, nil)
	require.Error(t, err)
	require.Equal(t, "parameter version cannot be empty", err.Error())
}