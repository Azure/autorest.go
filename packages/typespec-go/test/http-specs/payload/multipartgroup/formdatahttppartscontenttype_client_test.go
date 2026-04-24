// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"context"
	"multipartgroup"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func TestFormDataHTTPPartsContentTypeClient_ImageJPEGContentType(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsContentTypeClient().ImageJPEGContentType(context.Background(), multipartgroup.FileWithHTTPPartSpecificContentTypeRequest{
		ProfileImage: streaming.MultipartContent{
			Body:     jpgFile,
			Filename: "hello.jpg",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataHTTPPartsContentTypeClient_OptionalContentType(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsContentTypeClient().OptionalContentType(context.Background(), multipartgroup.FileWithHTTPPartOptionalContentTypeRequest{
		ProfileImage: streaming.MultipartContent{
			Body:     jpgFile,
			Filename: "hello.jpg",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataHTTPPartsContentTypeClient_RequiredContentType(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsContentTypeClient().RequiredContentType(context.Background(), multipartgroup.FileWithHTTPPartRequiredContentTypeRequest{
		ProfileImage: streaming.MultipartContent{
			Body:        jpgFile,
			ContentType: "application/octet-stream",
			Filename:    "hello.jpg",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
