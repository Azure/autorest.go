// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package stringgroup_test

import (
	"context"
	"generatortests/stringgroup"
	"generatortests/stringgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeGetBase64Encoded(t *testing.T) {
	value := []byte("this is the thing")
	server := fake.StringServer{
		GetBase64Encoded: func(ctx context.Context, options *stringgroup.StringClientGetBase64EncodedOptions) (resp azfake.Responder[stringgroup.StringClientGetBase64EncodedResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, stringgroup.StringClientGetBase64EncodedResponse{
				Value: value,
			}, nil)
			return
		},
	}
	client, err := stringgroup.NewStringClient(&azcore.ClientOptions{
		Transport: fake.NewStringServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetBase64Encoded(context.Background(), nil)
	require.NoError(t, err)
	require.EqualValues(t, value, resp.Value)
}

func TestFakePutBase64URLEncoded(t *testing.T) {
	value := []byte("this is the thing")
	server := fake.StringServer{
		PutBase64URLEncoded: func(ctx context.Context, stringBody []byte, options *stringgroup.StringClientPutBase64URLEncodedOptions) (resp azfake.Responder[stringgroup.StringClientPutBase64URLEncodedResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, value, stringBody)
			resp.SetResponse(http.StatusOK, stringgroup.StringClientPutBase64URLEncodedResponse{}, nil)
			return
		},
	}
	client, err := stringgroup.NewStringClient(&azcore.ClientOptions{
		Transport: fake.NewStringServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutBase64URLEncoded(context.Background(), value, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
