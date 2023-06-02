// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package integergroup_test

import (
	"context"
	"generatortests/integergroup"
	"generatortests/integergroup/fake"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeGetInvalid(t *testing.T) {
	value := int32(123)
	server := fake.IntServer{
		GetInvalid: func(ctx context.Context, options *integergroup.IntClientGetInvalidOptions) (resp azfake.Responder[integergroup.IntClientGetInvalidResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, integergroup.IntClientGetInvalidResponse{
				Value: &value,
			}, nil)
			return
		},
	}
	client, err := integergroup.NewIntClient(&azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetInvalid(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, value, *resp.Value)
}

func TestFakeGetUnixTime(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	server := fake.IntServer{
		GetUnixTime: func(ctx context.Context, options *integergroup.IntClientGetUnixTimeOptions) (resp azfake.Responder[integergroup.IntClientGetUnixTimeResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, integergroup.IntClientGetUnixTimeResponse{
				Value: &now,
			}, nil)
			return
		},
	}
	client, err := integergroup.NewIntClient(&azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetUnixTime(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, now, *resp.Value)
}

func TestFakePutMax64(t *testing.T) {
	value := int64(99887766554433221)
	server := fake.IntServer{
		PutMax64: func(ctx context.Context, intBody int64, options *integergroup.IntClientPutMax64Options) (resp azfake.Responder[integergroup.IntClientPutMax64Response], errResp azfake.ErrorResponder) {
			require.EqualValues(t, value, intBody)
			resp.SetResponse(http.StatusOK, integergroup.IntClientPutMax64Response{}, nil)
			return
		},
	}
	client, err := integergroup.NewIntClient(&azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutMax64(context.TODO(), value, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakePutUnixTimeDate(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	server := fake.IntServer{
		PutUnixTimeDate: func(ctx context.Context, intBody time.Time, options *integergroup.IntClientPutUnixTimeDateOptions) (resp azfake.Responder[integergroup.IntClientPutUnixTimeDateResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, now, intBody)
			resp.SetResponse(http.StatusOK, integergroup.IntClientPutUnixTimeDateResponse{}, nil)
			return
		},
	}
	client, err := integergroup.NewIntClient(&azcore.ClientOptions{
		Transport: fake.NewIntServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutUnixTimeDate(context.TODO(), now, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
