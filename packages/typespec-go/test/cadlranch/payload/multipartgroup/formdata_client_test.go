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
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().AnonymousModel(context.Background(), jpgFile, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_Basic(t *testing.T) {
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
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
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	pngFile, err := os.ReadFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png")
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
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
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

func TestFormDataClient_Complex(t *testing.T) {
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	pngFile, err := os.ReadFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png")
	require.NoError(t, err)
	resp, err := client.NewMultiPartFormDataClient().Complex(context.Background(), multipartgroup.ComplexPartsRequest{
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
		PreviousAddresses: []multipartgroup.Address{
			{
				City: to.Ptr("Y"),
			},
			{
				City: to.Ptr("Z"),
			},
		},
		ProfileImage: streaming.MultipartContent{Body: jpgFile},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_JSONArrayParts(t *testing.T) {
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
	resp, err := client.NewMultiPartFormDataClient().JSONArrayParts(context.Background(), multipartgroup.JSONArrayPartsRequest{
		PreviousAddresses: []multipartgroup.Address{
			{
				City: to.Ptr("Y"),
			},
			{
				City: to.Ptr("Z"),
			},
		},
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFormDataClient_JSONPart(t *testing.T) {
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()
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
	client, err := multipartgroup.NewMultiPartClient(nil)
	require.NoError(t, err)
	jpgFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.jpg", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer jpgFile.Close()

	resp, err := client.NewMultiPartFormDataClient().MultiBinaryParts(context.Background(), multipartgroup.MultiBinaryPartsRequest{
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)

	_, err = jpgFile.Seek(0, io.SeekStart)
	require.NoError(t, err)
	pngFile, err := os.OpenFile("../../../../node_modules/@azure-tools/cadl-ranch-specs/assets/image.png", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer pngFile.Close()
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
