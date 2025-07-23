// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup_test

import (
	"context"
	"generatortests"
	"generatortests/nonstringenumgroup"
	"generatortests/nonstringenumgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeGet(t *testing.T) {
	value := nonstringenumgroup.IntEnumFourHundredSix
	server := fake.IntServer{
		Get: func(ctx context.Context, options *nonstringenumgroup.IntClientGetOptions) (resp azfake.Responder[nonstringenumgroup.IntClientGetResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, nonstringenumgroup.IntClientGetResponse{
				Value: &value,
			}, nil)
			return
		},
	}
	client, err := nonstringenumgroup.NewIntClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, value, *resp.Value)
}

func TestFakePut(t *testing.T) {
	value := nonstringenumgroup.IntEnumFourHundredSix
	result := "hooray"
	server := fake.IntServer{
		Put: func(ctx context.Context, input nonstringenumgroup.IntEnum, options *nonstringenumgroup.IntClientPutOptions) (resp azfake.Responder[nonstringenumgroup.IntClientPutResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, value, input)
			resp.SetResponse(http.StatusOK, nonstringenumgroup.IntClientPutResponse{
				Value: &result,
			}, nil)
			return
		},
	}
	client, err := nonstringenumgroup.NewIntClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.Put(context.Background(), value, nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, result, *resp.Value)
}
