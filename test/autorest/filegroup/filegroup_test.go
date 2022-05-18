// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filegroup

import (
	"context"
	"generatortests"
	"io/ioutil"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newFilesClient() *FilesClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewFilesClient(pl)
}

func TestGetEmptyFile(t *testing.T) {
	client := newFilesClient()
	result, err := client.GetEmptyFile(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	result.Body.Close()
}

func TestGetFile(t *testing.T) {
	client := newFilesClient()
	result, err := client.GetFile(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.Body)
	require.NoError(t, err)
	result.Body.Close()
	if l := len(b); l != 8725 {
		t.Fatalf("unexpected byte count: want 8725, got %d", l)
	}
}

func TestGetFileLarge(t *testing.T) {
	t.Skip("test is unreliable, can fail when running on a machine with low memory")
	client := newFilesClient()
	result, err := client.GetFileLarge(context.Background(), nil)
	require.NoError(t, err)
	if result.Body == nil {
		t.Fatal("unexpected nil response body")
	}
	b, err := ioutil.ReadAll(result.Body)
	require.NoError(t, err)
	result.Body.Close()
	const size = 3000 * 1024 * 1024
	if l := len(b); l != size {
		t.Fatalf("unexpected byte count: want %d, got %d", size, l)
	}
}
