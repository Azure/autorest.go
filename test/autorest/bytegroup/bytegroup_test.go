// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newByteClient() *ByteClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewByteClient(pl)
}

func TestGetEmpty(t *testing.T) {
	client := newByteClient()
	result, err := client.GetEmpty(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, []byte{}); r != "" {
		t.Fatal(r)
	}
}

func TestGetInvalid(t *testing.T) {
	client := newByteClient()
	result, err := client.GetInvalid(context.Background(), nil)
	// TODO: verify error response is clear and actionable
	require.Error(t, err)
	require.Zero(t, result)
}

func TestGetNonASCII(t *testing.T) {
	client := newByteClient()
	result, err := client.GetNonASCII(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6}); r != "" {
		t.Fatal(r)
	}
}

func TestGetNull(t *testing.T) {
	client := newByteClient()
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, ([]byte)(nil)); r != "" {
		t.Fatal(r)
	}
}

func TestPutNonASCII(t *testing.T) {
	client := newByteClient()
	result, err := client.PutNonASCII(context.Background(), []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
