// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"bytes"
	"context"
	"io"
	"multipartgroup"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFormDataClient_AnonymousModel(t *testing.T) {
	t.Skip("waiting on https://github.com/Azure/typespec-azure/pull/4317")
	/*client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().AnonymousModel(context.Background(), jpgFile, nil)
	require.NoError(t, err)
	require.Zero(t, resp)*/
}

func TestFormDataClient_Basic(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	resp, err := client.NewMultiPartFormDataClient().Basic(context.Background(), multipartgroup.MultiPartRequest{
		ID: "123",
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_BinaryArrayParts(t *testing.T) {
	client := newClient(t)
	pngFile, err := os.ReadFile(pngPath)
	require.NoError(t, err)
	resp, err := client.NewMultiPartFormDataClient().BinaryArrayParts(context.Background(), multipartgroup.BinaryArrayPartsRequest{
		ID: "123",
		Pictures: []streaming.MultipartContent{
			{
				Body: streaming.NopCloser(bytes.NewReader(pngFile)),
			},
			{
				Body: streaming.NopCloser(bytes.NewReader(pngFile)),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_CheckFileNameAndContentType(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	resp, err := client.NewMultiPartFormDataClient().CheckFileNameAndContentType(context.Background(), multipartgroup.MultiPartRequest{
		ID: "123",
		ProfileImage: streaming.MultipartContent{
			Body:        jpgFile,
			ContentType: "image/jpg",
			Filename:    "hello.jpg",
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_FileArrayAndBasic(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	pngFile, err := os.ReadFile(pngPath)
	require.NoError(t, err)
	resp, err := client.NewMultiPartFormDataClient().FileArrayAndBasic(context.Background(), multipartgroup.ComplexPartsRequest{
		Address: multipartgroup.Address{
			City: to.Ptr("X"),
		},
		ID: "123",
		Pictures: []streaming.MultipartContent{
			{
				Body: streaming.NopCloser(bytes.NewReader(pngFile)),
			},
			{
				Body: streaming.NopCloser(bytes.NewReader(pngFile)),
			},
		},
		ProfileImage: streaming.MultipartContent{Body: jpgFile},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_JSONPart(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	resp, err := client.NewMultiPartFormDataClient().JSONPart(context.Background(), multipartgroup.JSONPartRequest{
		Address: multipartgroup.Address{
			City: to.Ptr("X"),
		},
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_MultiBinaryParts(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()

	// without optional picture
	resp, err := client.NewMultiPartFormDataClient().MultiBinaryParts(context.Background(), multipartgroup.MultiBinaryPartsRequest{
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)

	// with optional picture
	_, err = jpgFile.Seek(0, io.SeekStart)
	require.NoError(t, err)
	pngFile, err := os.OpenFile(pngPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = pngFile.Close() }()
	resp, err = client.NewMultiPartFormDataClient().MultiBinaryParts(context.Background(), multipartgroup.MultiBinaryPartsRequest{
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
		Picture: &streaming.MultipartContent{
			Body: pngFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_OptionalParts(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	resp, err := client.NewMultiPartFormDataClient().OptionalParts(context.Background(), multipartgroup.MultiPartOptionalRequest{
		ID: to.Ptr("123"),
		ProfileImage: &streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_WithWireName(t *testing.T) {
	client := newClient(t)
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	resp, err := client.NewMultiPartFormDataClient().WithWireName(context.Background(), multipartgroup.MultiPartRequestWithWireName{
		Identifier: "123",
		Image: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
