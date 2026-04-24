// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"bytes"
	"context"
	"multipartgroup"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFormDataHTTPPartsClient_JSONArrayAndFileArray(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	pngFile, err := os.ReadFile(pngPath)
	require.NoError(t, err)
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().JSONArrayAndFileArray(context.Background(), multipartgroup.ComplexHTTPPartsModelRequest{
		ID: "123",
		Address: multipartgroup.Address{
			City: to.Ptr("X"),
		},
		ProfileImage: streaming.MultipartContent{
			Body:        jpgFile,
			ContentType: "application/octet-stream",
			Filename:    "hello.jpg",
		},
		PreviousAddresses: []multipartgroup.Address{
			{City: to.Ptr("Y")},
			{City: to.Ptr("Z")},
		},
		Pictures: []streaming.MultipartContent{
			{
				Body:        streaming.NopCloser(bytes.NewReader(pngFile)),
				ContentType: "application/octet-stream",
				Filename:    "hello.png",
			},
			{
				Body:        streaming.NopCloser(bytes.NewReader(pngFile)),
				ContentType: "application/octet-stream",
				Filename:    "hello.png",
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
