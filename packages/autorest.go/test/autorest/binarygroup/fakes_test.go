// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package binarygroup_test

import (
	"bytes"
	"context"
	"generatortests"
	"generatortests/binarygroup"
	"generatortests/binarygroup/fake"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/stretchr/testify/require"
)

func TestFakeBinary(t *testing.T) {
	server := fake.UploadServer{
		Binary: func(ctx context.Context, fileParam io.ReadSeekCloser, options *binarygroup.UploadClientBinaryOptions) (resp azfake.Responder[binarygroup.UploadClientBinaryResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			body, err := io.ReadAll(fileParam)
			require.NoError(t, err)
			require.EqualValues(t, []byte{0xff, 0xfe, 0xfd}, body)
			resp.SetResponse(http.StatusOK, binarygroup.UploadClientBinaryResponse{}, nil)
			return
		},
	}
	client, err := binarygroup.NewUploadClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewUploadServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Binary(context.Background(), streaming.NopCloser(bytes.NewReader([]byte{0xff, 0xfe, 0xfd})), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
