// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func TestFormDataFileClient_UploadFileArray(t *testing.T) {
	client := newClient(t)
	pngFile, err := os.ReadFile(pngPath)
	require.NoError(t, err)
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataFileClient().UploadFileArray(context.Background(), []streaming.MultipartContent{
		{
			Body:     streaming.NopCloser(bytes.NewReader(pngFile)),
			Filename: "image1.png",
		},
		{
			Body:     streaming.NopCloser(bytes.NewReader(pngFile)),
			Filename: "image2.png",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataFileClient_UploadFileRequiredFilename(t *testing.T) {
	client := newClient(t)
	pngFile, err := os.OpenFile(pngPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer pngFile.Close()
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataFileClient().UploadFileRequiredFilename(context.Background(), streaming.MultipartContent{
		Body:     pngFile,
		Filename: "image.png",
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataFileClient_UploadFileSpecificContentType(t *testing.T) {
	client := newClient(t)
	pngFile, err := os.OpenFile(pngPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer pngFile.Close()
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataFileClient().UploadFileSpecificContentType(context.Background(), streaming.MultipartContent{
		Body:     pngFile,
		Filename: "image.png",
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
