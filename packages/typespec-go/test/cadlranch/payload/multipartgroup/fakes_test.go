// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package multipartgroup_test

import (
	"bytes"
	"context"
	"io"
	"multipartgroup"
	"multipartgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeFormDataClientBasic(t *testing.T) {
	id := "12345"
	bodyContent := []byte{1, 2, 3, 4, 5}
	contentType := "binary"
	filename := "data.bin"
	server := fake.MultiPartFormDataServer{
		Basic: func(ctx context.Context, body multipartgroup.MultiPartRequest, options *multipartgroup.MultiPartFormDataClientBasicOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientBasicResponse], errResp azfake.ErrorResponder) {
			require.Equal(t, id, body.ID)
			b, err := io.ReadAll(body.ProfileImage.Body)
			require.NoError(t, err)
			require.Equal(t, bodyContent, b)
			require.Equal(t, contentType, body.ProfileImage.ContentType)
			require.Equal(t, filename, body.ProfileImage.Filename)
			resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataClientBasicResponse{}, nil)
			return
		},
	}
	client, err := multipartgroup.NewMultiPartClient(&azcore.ClientOptions{
		Transport: fake.NewMultiPartFormDataServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.NewMultiPartFormDataClient().Basic(context.Background(), multipartgroup.MultiPartRequest{
		ID: id,
		ProfileImage: streaming.MultipartContent{
			Body:        streaming.NopCloser(bytes.NewReader(bodyContent)),
			ContentType: contentType,
			Filename:    filename,
		},
	}, nil)
	require.NoError(t, err)
}

func TestFakeFormDataClientBinaryArrayParts(t *testing.T) {
	id := "abc123"
	bodyWithDefaultsContent := []byte{1, 2, 3, 4, 5}
	dataDotBinContent := []byte{201, 202, 203}
	stuffContent := []byte{79, 81, 83, 85}
	server := fake.MultiPartFormDataServer{
		BinaryArrayParts: func(ctx context.Context, body multipartgroup.BinaryArrayPartsRequest, options *multipartgroup.MultiPartFormDataClientBinaryArrayPartsOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientBinaryArrayPartsResponse], errResp azfake.ErrorResponder) {
			require.Equal(t, id, body.ID)
			require.Len(t, body.Pictures, 3)
			// entries should be in the same order
			entries := []struct {
				Body        []byte
				ContentType string
				Filename    string
			}{
				{
					Body:        bodyWithDefaultsContent,
					ContentType: "application/octet-stream",
					Filename:    "pictures",
				},
				{
					Body:        dataDotBinContent,
					ContentType: "binary",
					Filename:    "data.bin",
				},
				{
					Body:        stuffContent,
					ContentType: "bits",
					Filename:    "stuff",
				},
			}
			for i, entry := range entries {
				b, err := io.ReadAll(body.Pictures[i].Body)
				require.NoError(t, err)
				require.Equal(t, entry.Body, b)
				require.EqualValues(t, entry.ContentType, body.Pictures[i].ContentType)
				require.EqualValues(t, entry.Filename, body.Pictures[i].Filename)
			}
			resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataClientBinaryArrayPartsResponse{}, nil)
			return
		},
	}
	client, err := multipartgroup.NewMultiPartClient(&azcore.ClientOptions{
		Transport: fake.NewMultiPartFormDataServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.NewMultiPartFormDataClient().BinaryArrayParts(context.Background(), multipartgroup.BinaryArrayPartsRequest{
		ID: id,
		Pictures: []streaming.MultipartContent{
			{
				Body: streaming.NopCloser(bytes.NewReader(bodyWithDefaultsContent)),
			},
			{
				Body:        streaming.NopCloser(bytes.NewReader(dataDotBinContent)),
				ContentType: "binary",
				Filename:    "data.bin",
			},
			{
				Body:        streaming.NopCloser(bytes.NewReader(stuffContent)),
				ContentType: "bits",
				Filename:    "stuff",
			},
		},
	}, nil)
	require.NoError(t, err)
}

func TestFakeFormDataClientJSONPart(t *testing.T) {
	city := "Someplace"
	bodyContent := []byte{1, 2, 3, 4, 5}
	contentType := "binary"
	filename := "data.bin"
	server := fake.MultiPartFormDataServer{
		JSONPart: func(ctx context.Context, body multipartgroup.JSONPartRequest, options *multipartgroup.MultiPartFormDataClientJSONPartOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientJSONPartResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, body.Address.City)
			require.Equal(t, city, *body.Address.City)
			b, err := io.ReadAll(body.ProfileImage.Body)
			require.NoError(t, err)
			require.Equal(t, bodyContent, b)
			require.Equal(t, contentType, body.ProfileImage.ContentType)
			require.Equal(t, filename, body.ProfileImage.Filename)
			resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataClientJSONPartResponse{}, nil)
			return
		},
	}
	client, err := multipartgroup.NewMultiPartClient(&azcore.ClientOptions{
		Transport: fake.NewMultiPartFormDataServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.NewMultiPartFormDataClient().JSONPart(context.Background(), multipartgroup.JSONPartRequest{
		Address: multipartgroup.Address{
			City: &city,
		},
		ProfileImage: streaming.MultipartContent{
			Body:        streaming.NopCloser(bytes.NewReader(bodyContent)),
			ContentType: contentType,
			Filename:    filename,
		},
	}, nil)
	require.NoError(t, err)
}

func TestFakeFormDataClientJSONArrayParts(t *testing.T) {
	previous := []multipartgroup.Address{
		{
			City: to.Ptr("City1"),
		},
		{
			City: to.Ptr("CitwTwo"),
		},
	}
	bodyContent := []byte{1, 2, 3, 4, 5}
	contentType := "binary"
	filename := "data.bin"
	server := fake.MultiPartFormDataServer{
		JSONArrayParts: func(ctx context.Context, body multipartgroup.JSONArrayPartsRequest, options *multipartgroup.MultiPartFormDataClientJSONArrayPartsOptions) (resp azfake.Responder[multipartgroup.MultiPartFormDataClientJSONArrayPartsResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, previous, body.PreviousAddresses)
			b, err := io.ReadAll(body.ProfileImage.Body)
			require.NoError(t, err)
			require.Equal(t, bodyContent, b)
			require.Equal(t, contentType, body.ProfileImage.ContentType)
			require.Equal(t, filename, body.ProfileImage.Filename)
			resp.SetResponse(http.StatusNoContent, multipartgroup.MultiPartFormDataClientJSONArrayPartsResponse{}, nil)
			return
		},
	}
	client, err := multipartgroup.NewMultiPartClient(&azcore.ClientOptions{
		Transport: fake.NewMultiPartFormDataServerTransport(&server),
	})
	require.NoError(t, err)
	_, err = client.NewMultiPartFormDataClient().JSONArrayParts(context.Background(), multipartgroup.JSONArrayPartsRequest{
		PreviousAddresses: previous,
		ProfileImage: streaming.MultipartContent{
			Body:        streaming.NopCloser(bytes.NewReader(bodyContent)),
			ContentType: contentType,
			Filename:    filename,
		},
	}, nil)
	require.NoError(t, err)
}
