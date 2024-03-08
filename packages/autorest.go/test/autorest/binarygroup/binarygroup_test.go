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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func newDownloadClient(t *testing.T) *DownloadClient {
	client, err := NewDownloadClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func newUploadClient(t *testing.T) *UploadClient {
	client, err := NewUploadClient(&azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
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

func TestErrorStream(t *testing.T) {
	client := newDownloadClient(t)
	resp, err := client.ErrorStream(context.Background(), nil)
	var respError *azcore.ResponseError
	require.ErrorAs(t, err, &respError)
	const want = `GET http://localhost:3000/binary/error
--------------------------------------------------------------------------------
RESPONSE 400: 400 Bad Request
ERROR CODE UNAVAILABLE
--------------------------------------------------------------------------------
{
  "status": 400,
  "message": "I failed on purpose"
}
--------------------------------------------------------------------------------
`
	require.EqualValues(t, want, respError.Error())
	require.Zero(t, resp)
}
