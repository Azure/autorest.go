// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package mediatypesgroup_test

import (
	"context"
	"generatortests"
	"generatortests/mediatypesgroup"
	"generatortests/mediatypesgroup/fake"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeBinaryBodyWithThreeContentTypes(t *testing.T) {
	server := fake.MediaTypesServer{
		BinaryBodyWithThreeContentTypes: func(ctx context.Context, contentType mediatypesgroup.ContentType2, message io.ReadSeekCloser, options *mediatypesgroup.MediaTypesClientBinaryBodyWithThreeContentTypesOptions) (resp azfake.Responder[mediatypesgroup.MediaTypesClientBinaryBodyWithThreeContentTypesResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, mediatypesgroup.ContentType2TextPlain, contentType)
			body, err := io.ReadAll(message)
			require.NoError(t, err)
			require.EqualValues(t, "this is a test", string(body))
			resp.SetResponse(http.StatusOK, mediatypesgroup.MediaTypesClientBinaryBodyWithThreeContentTypesResponse{
				Value: to.Ptr("success"),
			}, nil)
			return
		},
	}
	client, err := mediatypesgroup.NewMediaTypesClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewMediaTypesServerTransport(&server),
	})
	require.NoError(t, err)
	message := streaming.NopCloser(strings.NewReader("this is a test"))
	resp, err := client.BinaryBodyWithThreeContentTypes(context.Background(), mediatypesgroup.ContentType2TextPlain, message, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, "success", *resp.Value)
}
