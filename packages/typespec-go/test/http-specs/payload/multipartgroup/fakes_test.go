// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"bytes"
	"context"
	"multipartgroup"
	"multipartgroup/fake"
	"net/http"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newFakeClient(t *testing.T, transport policy.Transporter) *multipartgroup.MultiPartClient {
	t.Helper()
	client, err := multipartgroup.NewMultiPartClientWithNoCredential("http://localhost:3000", &multipartgroup.MultiPartClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: transport,
		},
	})
	require.NoError(t, err)
	return client
}

func TestFakeFormDataClient_Basic(t *testing.T) {
	calledBasic := false
	srv := fake.MultiPartServer{
		MultiPartFormDataServer: fake.MultiPartFormDataServer{
			Basic: func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.MultiPartFormDataClientBasicOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientBasicResponse], errResp azfake.ErrorResponder) {
				require.EqualValues(t, "123", body.ID)
				require.NotNil(t, body.ProfileImage.Body)
				calledBasic = true
				resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataClientBasicResponse{}, nil)
				return
			},
		},
	}
	client := newFakeClient(t, fake.NewMultiPartServerTransport(&srv))
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	_, err = client.NewMultiPartFormDataClient().Basic(context.Background(), multipartgroup.MultiPartRequest{
		ID: "123",
		ProfileImage: streaming.MultipartContent{
			Body: jpgFile,
		},
	}, nil)
	require.NoError(t, err)
	require.True(t, calledBasic)
}

func TestFakeFormDataFileClient_UploadFileRequiredFilename(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.go/issues/1937")
	calledUpload := false
	srv := fake.MultiPartServer{
		MultiPartFormDataServer: fake.MultiPartFormDataServer{
			MultiPartFormDataFileServer: fake.MultiPartFormDataFileServer{
				UploadFileRequiredFilename: func(ctx context.Context, file streaming.MultipartContent, options *multipartgroup.MultiPartFormDataFileClientUploadFileRequiredFilenameOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataFileClientUploadFileRequiredFilenameResponse], errResp azfake.ErrorResponder) {
					require.NotNil(t, file.Body)
					calledUpload = true
					resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataFileClientUploadFileRequiredFilenameResponse{}, nil)
					return
				},
			},
		},
	}
	client := newFakeClient(t, fake.NewMultiPartServerTransport(&srv))
	pngFile, err := os.OpenFile(pngPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = pngFile.Close() }()
	_, err = client.NewMultiPartFormDataClient().NewMultiPartFormDataFileClient().UploadFileRequiredFilename(context.Background(), streaming.MultipartContent{
		Body:     pngFile,
		Filename: "image.png",
	}, nil)
	require.NoError(t, err)
	require.True(t, calledUpload)
}

func TestFakeFormDataHTTPPartsClient_JSONArrayAndFileArray(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.go/issues/1937")
	calledJSONArrayAndFileArray := false
	srv := fake.MultiPartServer{
		MultiPartFormDataServer: fake.MultiPartFormDataServer{
			MultiPartFormDataHTTPPartsServer: fake.MultiPartFormDataHTTPPartsServer{
				JSONArrayAndFileArray: func(ctx context.Context, body multipartgroup.ComplexHTTPPartsModelRequest, options *multipartgroup.MultiPartFormDataHTTPPartsClientJSONArrayAndFileArrayOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataHTTPPartsClientJSONArrayAndFileArrayResponse], errResp azfake.ErrorResponder) {
					require.EqualValues(t, "123", body.ID)
					require.NotNil(t, body.Address.City)
					require.EqualValues(t, "X", *body.Address.City)
					require.Len(t, body.PreviousAddresses, 2)
					require.Len(t, body.Pictures, 2)
					require.NotNil(t, body.ProfileImage.Body)
					calledJSONArrayAndFileArray = true
					resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataHTTPPartsClientJSONArrayAndFileArrayResponse{}, nil)
					return
				},
			},
		},
	}
	client := newFakeClient(t, fake.NewMultiPartServerTransport(&srv))
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	pngFile, err := os.ReadFile(pngPath)
	require.NoError(t, err)
	_, err = client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().JSONArrayAndFileArray(context.Background(), multipartgroup.ComplexHTTPPartsModelRequest{
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
	require.True(t, calledJSONArrayAndFileArray)
}

func TestFakeFormDataHTTPPartsContentTypeClient_RequiredContentType(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.go/issues/1937")
	calledRequiredContentType := false
	srv := fake.MultiPartServer{
		MultiPartFormDataServer: fake.MultiPartFormDataServer{
			MultiPartFormDataHTTPPartsServer: fake.MultiPartFormDataHTTPPartsServer{
				MultiPartFormDataHTTPPartsContentTypeServer: fake.MultiPartFormDataHTTPPartsContentTypeServer{
					RequiredContentType: func(ctx context.Context, body multipartgroup.FileWithHTTPPartRequiredContentTypeRequest, options *multipartgroup.MultiPartFormDataHTTPPartsContentTypeClientRequiredContentTypeOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataHTTPPartsContentTypeClientRequiredContentTypeResponse], errResp azfake.ErrorResponder) {
						require.NotNil(t, body.ProfileImage.Body)
						calledRequiredContentType = true
						resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataHTTPPartsContentTypeClientRequiredContentTypeResponse{}, nil)
						return
					},
				},
			},
		},
	}
	client := newFakeClient(t, fake.NewMultiPartServerTransport(&srv))
	jpgFile, err := os.OpenFile(jpgPath, os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = jpgFile.Close() }()
	_, err = client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsContentTypeClient().RequiredContentType(context.Background(), multipartgroup.FileWithHTTPPartRequiredContentTypeRequest{
		ProfileImage: streaming.MultipartContent{
			Body:        jpgFile,
			ContentType: "application/octet-stream",
			Filename:    "hello.jpg",
		},
	}, nil)
	require.NoError(t, err)
	require.True(t, calledRequiredContentType)
}

func TestFakeFormDataHTTPPartsNonStringClient_Float(t *testing.T) {
	t.Skip("https://github.com/Azure/autorest.go/issues/1937")
	calledFloat := false
	srv := fake.MultiPartServer{
		MultiPartFormDataServer: fake.MultiPartFormDataServer{
			MultiPartFormDataHTTPPartsServer: fake.MultiPartFormDataHTTPPartsServer{
				MultiPartFormDataHTTPPartsNonStringServer: fake.MultiPartFormDataHTTPPartsNonStringServer{
					Float: func(ctx context.Context, temperature float64, options *multipartgroup.MultiPartFormDataHTTPPartsNonStringClientFloatOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataHTTPPartsNonStringClientFloatResponse], errResp azfake.ErrorResponder) {
						require.EqualValues(t, 0.5, temperature)
						calledFloat = true
						resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataHTTPPartsNonStringClientFloatResponse{}, nil)
						return
					},
				},
			},
		},
	}
	client := newFakeClient(t, fake.NewMultiPartServerTransport(&srv))
	_, err := client.NewMultiPartFormDataClient().NewMultiPartFormDataHTTPPartsClient().NewMultiPartFormDataHTTPPartsNonStringClient().Float(context.Background(), 0.5, nil)
	require.NoError(t, err)
	require.True(t, calledFloat)
}
