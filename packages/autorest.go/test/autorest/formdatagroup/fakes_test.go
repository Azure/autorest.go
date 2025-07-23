// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package formdatagroup_test

import (
	"context"
	"generatortests"
	"generatortests/formdatagroup"
	"generatortests/formdatagroup/fake"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func TestFakeUploadFile(t *testing.T) {
	server := fake.FormdataServer{
		UploadFile: func(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *formdatagroup.FormdataClientUploadFileOptions) (resp azfake.Responder[formdatagroup.FormdataClientUploadFileResponse], err azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, formdatagroup.FormdataClientUploadFileResponse{Body: fileContent}, nil)
			return
		},
	}
	client, err := formdatagroup.NewFormdataClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewFormdataServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.UploadFile(context.Background(), streaming.NopCloser(strings.NewReader("the data")), "sample", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.EqualValues(t, "the data", string(data))
}

func TestFakeUploadFiles(t *testing.T) {
	server := fake.FormdataServer{
		UploadFiles: func(ctx context.Context, fileContent []io.ReadSeekCloser, options *formdatagroup.FormdataClientUploadFilesOptions) (resp azfake.Responder[formdatagroup.FormdataClientUploadFilesResponse], errResp azfake.ErrorResponder) {
			var rawBody []byte
			for _, content := range fileContent {
				data, err := io.ReadAll(content)
				if err != nil {
					errResp.SetError(err)
					return
				}
				rawBody = append(rawBody, data...)
			}
			resp.SetResponse(http.StatusOK, formdatagroup.FormdataClientUploadFilesResponse{Body: io.NopCloser(strings.NewReader(string(rawBody)))}, nil)
			return
		},
	}
	client, err := formdatagroup.NewFormdataClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewFormdataServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.UploadFiles(context.Background(), []io.ReadSeekCloser{
		streaming.NopCloser(strings.NewReader("the data")),
		streaming.NopCloser(strings.NewReader(" to be uploaded")),
	}, nil)
	require.NoError(t, err)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.EqualValues(t, "the data to be uploaded", string(data))
}
