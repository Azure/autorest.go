// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimerfc1123group_test

import (
	"context"
	"generatortests/datetimerfc1123group"
	"generatortests/datetimerfc1123group/fake"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeGetUTCMinDateTime(t *testing.T) {
	server := fake.Datetimerfc1123Server{
		GetUTCMinDateTime: func(ctx context.Context, options *datetimerfc1123group.Datetimerfc1123ClientGetUTCMinDateTimeOptions) (resp azfake.Responder[datetimerfc1123group.Datetimerfc1123ClientGetUTCMinDateTimeResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, datetimerfc1123group.Datetimerfc1123ClientGetUTCMinDateTimeResponse{
				Value: to.Ptr(time.Date(2023, 5, 11, 10, 5, 44, 0, time.UTC)),
			}, nil)
			return
		},
	}
	client, err := datetimerfc1123group.NewDatetimerfc1123Client(&azcore.ClientOptions{
		Transport: fake.NewDatetimerfc1123ServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetUTCMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Value)
	require.EqualValues(t, time.Date(2023, 5, 11, 10, 5, 44, 0, time.UTC), *resp.Value)
}

func TestFakePutUTCMaxDateTime(t *testing.T) {
	server := fake.Datetimerfc1123Server{
		PutUTCMaxDateTime: func(ctx context.Context, datetimeBody time.Time, options *datetimerfc1123group.Datetimerfc1123ClientPutUTCMaxDateTimeOptions) (resp azfake.Responder[datetimerfc1123group.Datetimerfc1123ClientPutUTCMaxDateTimeResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, time.Date(2023, 5, 11, 10, 5, 44, 0, time.UTC), datetimeBody)
			resp.SetResponse(http.StatusOK, datetimerfc1123group.Datetimerfc1123ClientPutUTCMaxDateTimeResponse{}, nil)
			return
		},
	}
	client, err := datetimerfc1123group.NewDatetimerfc1123Client(&azcore.ClientOptions{
		Transport: fake.NewDatetimerfc1123ServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutUTCMaxDateTime(context.Background(), time.Date(2023, 5, 11, 10, 5, 44, 0, time.UTC), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
