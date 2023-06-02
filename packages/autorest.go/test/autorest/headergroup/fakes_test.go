// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package headergroup_test

import (
	"context"
	"generatortests/headergroup"
	"generatortests/headergroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeParamByte(t *testing.T) {
	server := fake.HeaderServer{
		ParamByte: func(ctx context.Context, scenario string, value []byte, options *headergroup.HeaderClientParamByteOptions) (resp azfake.Responder[headergroup.HeaderClientParamByteResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, "scenario", scenario)
			require.EqualValues(t, []byte{0xab, 0xcd, 0xef}, value)
			resp.SetResponse(http.StatusOK, headergroup.HeaderClientParamByteResponse{}, nil)
			return
		},
	}
	client, err := headergroup.NewHeaderClient(&azcore.ClientOptions{
		Transport: fake.NewHeaderServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.ParamByte(context.Background(), "scenario", []byte{0xab, 0xcd, 0xef}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakeResponseEnum(t *testing.T) {
	server := fake.HeaderServer{
		ResponseEnum: func(ctx context.Context, scenario string, options *headergroup.HeaderClientResponseEnumOptions) (resp azfake.Responder[headergroup.HeaderClientResponseEnumResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, "scenario", scenario)
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, headergroup.HeaderClientResponseEnumResponse{
				Value: to.Ptr(headergroup.GreyscaleColorsGREY),
			}, nil)
			return
		},
	}
	client, err := headergroup.NewHeaderClient(&azcore.ClientOptions{
		Transport: fake.NewHeaderServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.ResponseEnum(context.Background(), "scenario", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, headergroup.GreyscaleColorsGREY, *resp.Value)
}
