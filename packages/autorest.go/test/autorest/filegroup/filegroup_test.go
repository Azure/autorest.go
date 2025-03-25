// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filegroup

import (
	"context"
	"generatortests"
	"io"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newFilesClient(t *testing.T) *FilesClient {
	client, err := NewFilesClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewFilesClient(options *azcore.ClientOptions) (*FilesClient, error) {
	client, err := azcore.NewClient("filegroup.FilesClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &FilesClient{internal: client}, nil
}

func TestGetEmptyFile(t *testing.T) {
	client := newFilesClient(t)
	result, err := client.GetEmptyFile(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	require.NoError(t, result.Body.Close())
}

func TestGetFile(t *testing.T) {
	client := newFilesClient(t)
	result, err := client.GetFile(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	require.NoError(t, result.Body.Close())
	if l := len(b); l != 8725 {
		t.Fatalf("unexpected byte count: want 8725, got %d", l)
	}
}

func TestGetFileLarge(t *testing.T) {
	t.Skip("test is unreliable, can fail when running on a machine with low memory")
	client := newFilesClient(t)
	result, err := client.GetFileLarge(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	require.NoError(t, result.Body.Close())
	const size = 3000 * 1024 * 1024
	if l := len(b); l != size {
		t.Fatalf("unexpected byte count: want %d, got %d", size, l)
	}
}
