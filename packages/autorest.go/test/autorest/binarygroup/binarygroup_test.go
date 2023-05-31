// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binarygroup

import (
	"bytes"
	"context"
	"generatortests"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func newUploadClient(t *testing.T) *UploadClient {
	client, err := NewUploadClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func NewUploadClient(options *azcore.ClientOptions) (*UploadClient, error) {
	client, err := azcore.NewClient("binarygroup.UploadClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &UploadClient{internal: client}, nil
}

func TestBinary(t *testing.T) {
	client := newUploadClient(t)
	resp, err := client.Binary(context.Background(), streaming.NopCloser(bytes.NewReader([]byte{0xff, 0xfe, 0xfd})), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFile(t *testing.T) {
	client := newUploadClient(t)
	jsonFile := strings.NewReader(`{ "more": "cowbell" }`)
	resp, err := client.File(context.Background(), streaming.NopCloser(jsonFile), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
